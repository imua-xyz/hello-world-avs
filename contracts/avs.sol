pragma solidity >=0.8.17;
import "./IAVSManager.sol" as avs;
contract AvsServiceContract {
    address constant BLS_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000809;
    address public owner;
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


    function deregisterAVS(string memory avsName) public returns (bool) {
        bool success = avs.AVSMANAGER_CONTRACT.deregisterAVS(
            msg.sender,
            avsName
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


        bool success; bytes memory data = avs.AVSMANAGER_CONTRACT.getRegisteredPubkey(
        operator
    );

        require(success, "Call failed");
        return abi.decode(data, (bytes));
    }

    function getOptInOperators(address avsAddress) public returns (string[] memory) {
        bool success; string[] memory data = avs.AVSMANAGER_CONTRACT.getOptInOperators(
        avsAddress
    );


        require(success, "Call failed");
        return data;
    }

    function fastAggregateVerify(
        bytes32 msg_,
        bytes calldata signature,
        bytes[] calldata pubkeys
    ) public  returns (bool) {
        (bool success, ) = BLS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("fastAggregateVerify(bytes32,bytes,bytes[])")),
                msg_,signature,pubkeys
            )
        );

        return success;

    }

    function submitProof(
        string memory taskId,
        string memory taskContractAddress,
        string memory aggregator,
        string memory avsAddress,
        bytes memory operatorStatus
    ) public returns (bool) {
        bool success = avs.AVSMANAGER_CONTRACT.submitProof(
            taskId,
            taskContractAddress,
            aggregator,
            avsAddress,
            operatorStatus
        );
        return success;
    }

}