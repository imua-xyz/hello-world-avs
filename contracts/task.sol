pragma solidity ^0.8.0;
import "./IAVSManager.sol" as avs;

contract AVSTask {
    // task submitter decides on the criteria for a task to be completed
    // note that this does not mean the task was "correctly" answered (i.e. the number was squared correctly)
    //      this is for the challenge logic to verify
    // task is completed (and contract will accept its TaskResponse)
    // are signed by at least  ThresholdPercentage of the operators
    struct Task {
        uint256 numberToBeSquared;
        string name;
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
        uint256 numberSquared;
    }

    event TaskChallengedSuccessfully(
        uint64 indexed taskIndex,
        address indexed challenger
    );
    function createNewTask(
        uint256 numberToBeSquared,
        string memory name,
    // bytes calldata hash,
        uint64 taskResponsePeriod,
        uint64 taskChallengePeriod,
        uint64 thresholdPercentage,
        uint64 taskStatisticalPeriod
    ) public returns (bool) {
        // create a new task struct
        Task memory newTask;
        newTask.numberToBeSquared = numberToBeSquared;
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
        return success;
    }


    // NOTE: this function enables a challenger to raise and resolve a challenge.
    // TODO: require challenger to pay a bond for raising a challenge
    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        string memory operatorAddress
    ) public returns (bool) {
        uint64 taskID = taskResponse.taskID;
        uint256 numberToBeSquared = task.numberToBeSquared;

        // logic for checking whether challenge is valid or not
        uint256 actualSquaredOutput = numberToBeSquared * numberToBeSquared;
        bool isResponseCorrect = (actualSquaredOutput ==
            taskResponse.numberSquared);

        // if response was correct, no slashing happens so we return
        if (isResponseCorrect == true) {
            return false;
        }
        // // freeze the operators who signed adversarially

        // the task response has been challenged successfully
        bool success = avs.AVSMANAGER_CONTRACT.challenge(
            msg.sender,
            abi.encodePacked(keccak256(abi.encode(task))),
            taskResponse.taskID,
            abi.encodePacked(keccak256(abi.encode(taskResponse))),
            operatorAddress
        );
        emit TaskChallengedSuccessfully(taskID, msg.sender);
        return success;

    }
}