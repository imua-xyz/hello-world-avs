pragma solidity >=0.8.17;
import "./IAVSManager.sol" as avs;
contract AvsServiceContract {
    address public owner;
    // task submitter decides on the criteria for a task to be completed
    // note that this does not mean the task was "correctly" answered (i.e. the number was squared correctly)
    //      this is for the challenge logic to verify
    // task is completed (and contract will accept its TaskResponse)
    // are signed by at least  ThresholdPercentage of the operators
    struct Task {
        string name;
        uint256 taskId;
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
        string msg;
    }
    event TaskCreated(
        uint256 taskId,
        address issuer,
        string name,
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
        uint64 taskResponsePeriod,
        uint64 taskChallengePeriod,
        uint64 thresholdPercentage,
        uint64 taskStatisticalPeriod
    ) public returns (uint64) {
        // create a new task struct
        Task memory newTask;
        newTask.name = name;
        newTask.taskResponsePeriod = taskResponsePeriod;
        newTask.taskChallengePeriod = taskChallengePeriod;
        newTask.taskStatisticalPeriod = taskStatisticalPeriod;
        newTask.thresholdPercentage = thresholdPercentage;

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
        emit TaskCreated(newTask.taskId, msg.sender,newTask.name, newTask.taskResponsePeriod, newTask.taskChallengePeriod,newTask.thresholdPercentage,newTask.taskStatisticalPeriod);
        return taskID;
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
    function getTaskInfo(address taskAddr,uint64 taskID) external view returns (uint64[] memory){
        uint64[] memory  info = avs.AVSMANAGER_CONTRACT.getTaskInfo(
            taskAddr,
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
