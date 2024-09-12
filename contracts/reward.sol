pragma solidity ^0.8.0;

contract AVSReward {
    struct Operator {
        address operatorAddress;
        uint256 rewardBalance;
    }

    Operator[] public operators;

    modifier onlyOperator() {
        bool isOperator = false;
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i].operatorAddress == msg.sender) {
                isOperator = true;
                break;
            }
        }
        require(isOperator, "Only AVS operator can call this function.");
        _;
    }

    constructor() {
        operators.push(Operator(msg.sender, 0));
    }

    function registerOperator(address _operatorAddress) external onlyOperator {
        operators.push(Operator(_operatorAddress, 0));
    }

    function depositReward() external payable onlyOperator {
        // Deposit reward to be distributed among operators.
        for (uint256 i = 0; i < operators.length; i++) {
            operators[i].rewardBalance += msg.value / operators.length;
        }
    }

    function claimReward() external {
        uint256 rewardAmount = 0;
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i].operatorAddress == msg.sender) {
                rewardAmount = operators[i].rewardBalance;
                operators[i].rewardBalance = 0;
                break;
            }
        }
        require(rewardAmount > 0, "No reward balance available for the operator.");

        // Transfer the reward to the operator.
        (bool success, ) = msg.sender.call{value: rewardAmount}("");
        require(success, "Failed to transfer reward to the operator.");
    }
}