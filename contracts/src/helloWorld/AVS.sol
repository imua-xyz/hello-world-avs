pragma solidity >=0.8.17;
import "./IAVSManager.sol" as avs;
contract AvsServiceContract {
    address public owner;
    mapping(bytes32 => bool) public processedTasks;
    event TaskResolved(uint64 taskId, bool isExpected);

    // task submitter decides on the criteria for a task to be completed
    // note that this does not mean the task was "correctly" answered (i.e. the number was squared correctly)
    //      this is for the challenge logic to verify
    // task is completed (and contract will accept its TaskResponse)
    // are signed by at least  ThresholdPercentage of the operators
    struct Task {
        string name;
        uint256 taskId;
        uint64 numberToBeSquared;
        uint64 taskResponsePeriod;
        uint64 taskChallengePeriod;
        uint64 thresholdPercentage;
        uint64 taskStatisticalPeriod;
    }

    // Task response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct TaskResponse {
        // Can be obtained by the operator from the event NewTaskCreated.
        uint64 taskID;
        // This is just the response that the operator has to compute by itself.
        uint64 numberSquared;
    }
    event TaskCreated(
        uint256 taskId,
        address issuer,
        string name,
        uint64 numberToBeSquared,
        uint64 taskResponsePeriod,
        uint64 taskChallengePeriod,
        uint64 thresholdPercentage,
        uint64 taskStatisticalPeriod
    );

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only contract owner can call this function.");
        _;
    }

    function registerAVS(
        avs.AVSParams calldata params
    ) public returns (bool) {
        bool success =  avs.AVSMANAGER_CONTRACT.registerAVS(
            params
        );
        return success;
    }


    function updateAVS(
        avs.AVSParams calldata params
    ) public returns (bool) {
        bool success =  avs.AVSMANAGER_CONTRACT.updateAVS(
            params
        );
        return success;
    }

    function registerOperatorToAVS() public returns (bool) {

        bool success = avs.AVSMANAGER_CONTRACT.registerOperatorToAVS(
            msg.sender
        );
        return success;
    }

    function deregisterOperatorFromAVS() public returns (bool) {

        bool success = avs.AVSMANAGER_CONTRACT.deregisterOperatorFromAVS(
            msg.sender
        );
        return success;
    }

    function registerBLSPublicKey(
        address  avsAddr,
        bytes memory pubKey,
        bytes memory pubKeyRegistrationSignature,
        bytes memory pubKeyRegistrationMessageHash
    ) public returns (bool) {
        bool success = avs.AVSMANAGER_CONTRACT.registerBLSPublicKey(
            msg.sender,
            avsAddr,
            pubKey,
            pubKeyRegistrationSignature,
            pubKeyRegistrationMessageHash
        );
        return success;
    }

    function operatorSubmitTask(
        uint64 taskID,
        bytes calldata taskResponse,
        bytes calldata blsSignature,
        address taskContractAddress,
        uint8  phase
    ) public returns (bool) {

        bool success  = avs.AVSMANAGER_CONTRACT.operatorSubmitTask(
            msg.sender,
            taskID,
            taskResponse,
            blsSignature,
            taskContractAddress,
            phase
        );
        return success;
    }



    function createNewTask(
        string memory name,
        uint64 numberToBeSquared,
        uint64 taskResponsePeriod,
        uint64 taskChallengePeriod,
        uint8 thresholdPercentage,
        uint64 taskStatisticalPeriod
    ) public returns (uint64) {
        // create a new task struct
        Task memory newTask;
        newTask.name = name;
        newTask.numberToBeSquared = numberToBeSquared;
        newTask.taskResponsePeriod = taskResponsePeriod;
        newTask.taskChallengePeriod = taskChallengePeriod;
        newTask.taskStatisticalPeriod = taskStatisticalPeriod;
        newTask.thresholdPercentage = thresholdPercentage;
        require(
            thresholdPercentage<=100,
            "The threshold cannot be greater than 100"
        );
        uint64 taskID = avs.AVSMANAGER_CONTRACT.createTask(
            msg.sender,
            name,
            abi.encodePacked(keccak256(abi.encode(newTask))),
            taskResponsePeriod,
            taskChallengePeriod,
            thresholdPercentage,
            taskStatisticalPeriod
        );
        newTask.taskId = taskID;
        emit TaskCreated(newTask.taskId, msg.sender,newTask.name,newTask.numberToBeSquared, newTask.taskResponsePeriod, newTask.taskChallengePeriod,newTask.thresholdPercentage,newTask.taskStatisticalPeriod);
        return taskID;
    }

    function raiseAndResolveChallenge(
        address taskAddress,
        uint64 taskID,
        Task memory task
    ) public returns (bool) {
        // some logical checks

        avs.TaskInfo memory info = avs.AVSMANAGER_CONTRACT.getTaskInfo(
            taskAddress,
            taskID
        );
        bytes[] memory taskResponse =avs.AVSMANAGER_CONTRACT.getOperatorTaskResponse(
            taskAddress,
            taskID
        );
        require(compareBytes(info.hash , abi.encodePacked(keccak256(abi.encode(task)))), "Task is not equal.");
        require(info.isExpected == false, "Task already completed.");
        bytes32 key = getCompositeKey(info.taskContractAddress, info.taskId);

        require(!processedTasks[key],  "Task already processed");
        require(info.optInOperators.length > 0, "No opted-in operators");

        uint256 totalApprovedPower;
        uint256 totalPower = stringToUint(info.taskTotalPower);
        address[] memory tempReward = new address[](info.optInOperators.length);
        address[] memory tempSlash = new address[](info.optInOperators.length);
        uint256 rewardCount;
        uint256 slashCount;
        uint64 actualSquaredOutput = task.numberToBeSquared * task.numberToBeSquared;

        for (uint256 i = 0; i < info.optInOperators.length; i++) {
            address operator = info.optInOperators[i];

            TaskResponse memory res = decodeTaskRes(taskResponse[i]);
            bool isValid = (actualSquaredOutput == res.numberSquared);

            uint256 power = findOperatorPower(info.operatorActivePower, operator);

            if (isValid) {
                tempReward[rewardCount++] = info.optInOperators[i];
                totalApprovedPower += power;
            } else {
                tempSlash[slashCount++] = info.optInOperators[i];
            }
        }
        info.eligibleRewardOperators = tempReward;
        info.eligibleSlashOperators = tempSlash;
        uint256 approvalRate = (totalApprovedPower * 100) / totalPower;
        info.isExpected = approvalRate >= info.thresholdPercentage;
        info.thresholdPercentage = uint8(approvalRate);
        processedTasks[key] = true;
        emit TaskResolved(info.taskId, info.isExpected);
        bool success  = avs.AVSMANAGER_CONTRACT.challenge(
            msg.sender,
            info
        );
        return success;
    }

    function decodeTaskRes(bytes memory encodedData) public pure returns (TaskResponse memory) {
        (uint64 numberSquared, uint64 taskId) = abi.decode(encodedData, (uint64, uint64));

        return TaskResponse(numberSquared, taskId);
    }

    function compareBytes(bytes memory a, bytes memory b) public pure returns (bool) {
        if (a.length != b.length) {
            return false;
        }

        for (uint i = 0; i < a.length; i++) {
            if (a[i] != b[i]) {
                return false;
            }
        }

        return true;
    }

    function findOperatorPower(
        avs.OperatorActivePower[] memory powers,
        address operator
    ) internal pure returns (uint256) {
        for (uint256 i = 0; i < powers.length; i++) {
            if (powers[i].operator == operator) {
                return powers[i].power;
            }
        }
        revert("Operator power not found");
    }
    function getCompositeKey(
        address operator,
        uint64 taskId
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(
            bytes1(0xff), // 防碰撞前缀
            operator,
            bytes32(uint256(taskId))
        ));
    }
    function stringToUint(string memory s) internal pure returns (uint256) {
        bytes memory b = bytes(s);
        uint256 result = 0;
        for (uint256 i = 0; i < b.length; i++) {
            uint256 c = uint256(uint8(b[i]));
            if (c >= 48 && c <= 57) {
                result = result * 10 + (c - 48);
            }
        }
        return result;
    }
    //query
    function getOptInOperators(address avsAddress) public view returns (string[] memory) {
        string[] memory data = avs.AVSMANAGER_CONTRACT.getOptInOperators(
            avsAddress
        );

        return data;
    }


    function getRegisteredPubkey(address operator,address avsAddr) public view returns (bytes memory) {


        return avs.AVSMANAGER_CONTRACT.getRegisteredPubkey(
            operator,
            avsAddr
        );
    }

    function getAVSUSDValue(address avsAddr) external view returns (uint256){
        uint256  amount = avs.AVSMANAGER_CONTRACT.getAVSUSDValue(
            avsAddr
        );
        return amount;
    }
    function getOperatorOptedUSDValue(address avsAddr,address operatorAddr) external view returns (uint256){
        uint256  amount = avs.AVSMANAGER_CONTRACT.getOperatorOptedUSDValue(
            avsAddr,
            operatorAddr
        );
        return amount;
    }

    function getAVSEpochIdentifier(address avsAddr) external view returns (string memory){
        string memory epochIdentifier = avs.AVSMANAGER_CONTRACT.getAVSEpochIdentifier(
            avsAddr
        );
        return epochIdentifier;
    }
    function getTaskInfo(address taskAddress,uint64 taskID) external view returns (avs.TaskInfo memory){
        avs.TaskInfo memory info = avs.AVSMANAGER_CONTRACT.getTaskInfo(
            taskAddress,
            taskID
        );
        return info;
    }

    function isOperator(address operator) public view returns (bool) {

        bool flag = avs.AVSMANAGER_CONTRACT.isOperator(
            operator
        );
        return flag;
    }

    function getCurrentEpoch(string memory epochIdentifier) public view returns (int64) {

        int64 currentEpoch = avs.AVSMANAGER_CONTRACT.getCurrentEpoch(
            epochIdentifier
        );
        return currentEpoch;
    }
}
