pragma solidity >=0.8.17;
import "./IAVSManager.sol" as avs;
contract AvsServiceContract {
    address public owner;
    event DataLogged(bytes data);
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
        string memory avsName,
        uint64 minStakeAmount,
        address taskAddr,
        address slashAddr,
        address rewardAddr,
        string[] memory avsOwnerAddress,
        string[] memory assetIds,
        uint64 avsUnbondingPeriod,
        uint64 minSelfDelegation,
        string memory epochIdentifier,
        uint64[] memory params
    ) public returns (bool) {
        bool success =  avs.AVSMANAGER_CONTRACT.registerAVS(
            msg.sender,
            avsName,
            minStakeAmount,
            taskAddr,
            slashAddr,
            rewardAddr,
            avsOwnerAddress,
            assetIds,
            avsUnbondingPeriod,
            minSelfDelegation,
            epochIdentifier,
            params
        );
        return success;
    }


    function updateAVS(
        string memory avsName,
        uint64 minStakeAmount,
        address taskAddr,
        address slashAddr,
        address rewardAddr,
        string[] memory avsOwnerAddress,
        string[] memory assetIds,
        uint64 avsUnbondingPeriod,
        uint64 minSelfDelegation,
        string memory epochIdentifier,
        uint64[] memory params
    ) public returns (bool) {
        bool success =  avs.AVSMANAGER_CONTRACT.updateAVS(
            msg.sender,
            avsName,
            minStakeAmount,
            taskAddr,
            slashAddr,
            rewardAddr,
            avsOwnerAddress,
            assetIds,
            avsUnbondingPeriod,
            minSelfDelegation,
            epochIdentifier,
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
        string memory name,
        bytes memory pubKey,
        bytes memory pubKeyRegistrationSignature,
        bytes memory pubKeyRegistrationMessageHash
    ) public returns (bool) {
        bool success = avs.AVSMANAGER_CONTRACT.registerBLSPublicKey(
            msg.sender,
            name,
            pubKey,
            pubKeyRegistrationSignature,
            pubKeyRegistrationMessageHash
        );
        return success;
    }

    function getRegisteredPubkey(string memory operator) public returns (bytes memory) {


        bytes memory data = avs.AVSMANAGER_CONTRACT.getRegisteredPubkey(
            operator
        );
        emit DataLogged(data);
        return data;
        //return abi.decode(data, (bytes));
    }

    function getOptInOperators(address avsAddress) public returns (string[] memory) {
        string[] memory data = avs.AVSMANAGER_CONTRACT.getOptInOperators(
            avsAddress
        );


        return data;
    }


    function createNewTask(
        string memory name,
        uint64 taskResponsePeriod,
        uint64 taskChallengePeriod,
        uint64 thresholdPercentage,
        uint64 taskStatisticalPeriod
    ) public returns (bool) {
        // create a new task struct
        Task memory newTask;
        newTask.name = name;
        newTask.taskResponsePeriod = taskResponsePeriod;
        newTask.taskChallengePeriod = taskChallengePeriod;
        newTask.taskStatisticalPeriod = taskStatisticalPeriod;
        newTask.thresholdPercentage = thresholdPercentage;

        bool success = avs.AVSMANAGER_CONTRACT.createTask(
            msg.sender,
            name,
            abi.encodePacked(keccak256(abi.encode(newTask))),
            taskResponsePeriod,
            taskChallengePeriod,
            thresholdPercentage,
            taskStatisticalPeriod
        );

        emit TaskCreated(newTask.taskId, msg.sender,newTask.name, newTask.taskResponsePeriod, newTask.taskChallengePeriod,newTask.thresholdPercentage,newTask.taskStatisticalPeriod);
        return success;
    }

}