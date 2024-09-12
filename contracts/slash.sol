pragma solidity ^0.8.0;

contract AVSSlash {
    struct Operator {
        address operatorAddress;
        uint256 stakedTokens;
        bool isSlashed;
    }
    address public owner;
    Operator[] public operators;
    uint256 public slashAmount;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function.");
        _;
    }

    constructor(uint256 _slashAmount) {
        owner = msg.sender;
        slashAmount = _slashAmount;
    }

    function registerOperator(address _operatorAddress, uint256 _stakedTokens) external onlyOwner {
        operators.push(Operator(_operatorAddress, _stakedTokens, false));
    }

    function slashOperator(uint256 _operatorIndex) external onlyOwner {
        require(_operatorIndex < operators.length, "Invalid operator index.");
        require(!operators[_operatorIndex].isSlashed, "Operator is already slashed.");

        // Perform the slashing action here, such as freezing the malicious operator or deducting staked tokens.
        // This is a simplified example, you may need to implement the actual slashing logic based on your requirements.
        operators[_operatorIndex].isSlashed = true;
        operators[_operatorIndex].stakedTokens -= slashAmount;
    }
}