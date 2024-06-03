// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "../libraries/EpochUtils.sol";
import "./AVSDirectoryStorage.sol";

contract AVSDirectory is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AVSDirectoryStorage,
    ReentrancyGuardUpgradeable
{
    // @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    // @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /*******************************************************************************
                            INITIALIZING FUNCTIONS
    *******************************************************************************/

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManager, slasher, 
     * and eigenpodManager contracts
     */
    constructor(IDelegationManager _delegation, IStrategyManager _strategyManager) AVSDirectoryStorage(_delegation, _strategyManager) {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     * minWithdrawalDelayBlocks is set only once here
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _transferOwnership(initialOwner);
    }

    /*******************************************************************************
                            Operator<>AVS Registration
    *******************************************************************************/

    /**
     * @notice Called by the AVS's service manager contract to register an oprator to an operator set
     * @param operator The address of the operator to register.
     * @param operatorSetId The ID of the operator set to register the operator to.
     */
    function registerOperatorToOperatorSet(
        address operator,
        bytes4 operatorSetId
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        require(
            operatorSetRegistrations[msg.sender][operator][operatorSetId] != true,
            "AVSDirectory.registerOperatorToOperatorSet: operator already registered to operator set"
        )
        // TODO: check that the operator is allowing registrations for this operator set

        // Update state
        operatorSetRegistrations[msg.sender][operator][operatorSetId] = true;
        operatorAVSOperatorSetCount[msg.sender][operator] += 1;

        // Set the avs as an operator set AVS
        if (!isOperatorSetAVS[msg.sender]) {
            isOperatorSetAVS[msg.sender] = true;
        }
        // Set the operator as registered for the AVS if not already
        if (avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED) {
            avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED, EpochUtils.currentEpoch());
        }

        emit OperatorAddedToOperatorSet(operator, msg.sender, operatorSetId, EpochUtils.currentEpoch());
    }

    /**
     * @notice Called by an avs to deregister an operator from an operator set.
     * @param operator The address of the operator to deregister.
     * @param operatorSetId The ID of the operator set to register the operator to.
     */
    function deregisterOperatorFromOperatorSet(
        address operator,
        bytes4 operatorSetId
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        require(
            avsOperatorStatus[msg.sender][operator] == OperatorAVSRegistrationStatus.REGISTERED,
            "AVSDirectory.deregisterOperatorFromOperatorSet: operator not registered"
        );
        require(
            operatorSetRegistrations[msg.sender][operator][operatorSetId] == true, 
            "AVSDirectory.deregisterOperatorFromOperatorSet: operator not registered to operator set"
        );
        
        operatorAVSOperatorSetCount[msg.sender][operator] -= 1;

        if (operatorAVSOperatorSetCount[msg.sender][operator] == 0) {
            avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED, EpochUtils.currentEpoch() + 2);
        }

        emit OperatorRemovedFromOperatorSet(operator, msg.sender, operatorSetId, EpochUtils.currentEpoch() + 2);
    }


    /**
     * @notice Called by the AVS's service manager contract to register an operator with the avs.
     * @param operator The address of the operator to register.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {

        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToAVS: operator signature expired"
        );
        require(
            avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED,
            "AVSDirectory.registerOperatorToAVS: operator already registered"
        );
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToAVS: salt already spent"
        );
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet");
        require(
            isOperatorSetAVS[msg.sender],
            "AVSDirectory.registerOperatorToAVS: AVS is not a legacy AVS"
        )

        // Calculate the digest hash
        bytes32 operatorRegistrationDigestHash = calculateOperatorAVSRegistrationDigestHash({
            operator: operator,
            avs: msg.sender,
            salt: operatorSignature.salt,
            expiry: operatorSignature.expiry
        });

        // Check that the signature is valid
        EIP1271SignatureUtils.checkSignature_EIP1271(
            operator,
            operatorRegistrationDigestHash,
            operatorSignature.signature
        );

        // Set the operator as registered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;

        // Mark the salt as spent
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        // Legacy registrations will be active from the current epoch
        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED, EpochUtils.currentEpoch());
    }

    /**
     * @notice Called by an avs to deregister an operator with the avs.
     * @param operator The address of the operator to deregister.
     */
    function deregisterOperatorFromAVS(address operator) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        require(
            avsOperatorStatus[msg.sender][operator] == OperatorAVSRegistrationStatus.REGISTERED,
            "AVSDirectory.deregisterOperatorFromAVS: operator not registered"
        );

        // Set the operator as deregistered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED, EpochUtils.currentEpoch() + 2);
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(bytes32 salt) external {
        require(!operatorSaltIsSpent[msg.sender][salt], "AVSDirectory.cancelSalt: cannot cancel spent salt");
        operatorSaltIsSpent[msg.sender][salt] = true;
    }

    /*******************************************************************************
                                AVS Configurations
    *******************************************************************************/

    /**
     * @notice Called by an avs to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an avs
     */
    function updateAVSMetadataURI(string calldata metadataURI) external {
        emit AVSMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     * @notice Called by an avs to add a strategy to its operator set
     * @param operatorSetID the ID of the operator set
     * @param strategies the list strategies to add to the operator set
     */
    function addStrategiesToOperatorSet(
        bytes4 operatorSetID,
        IStrategy[] calldata strategies
    ) {
        uint256 strategiesToAdd = strategies.length;
        for (uint256 i = 0; i < strategiesToAdd; i++) {
            // Require that the strategy is valid 
            IStrategy strategy = strategies[i];
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategy) || strategy == beaconChainETHStrategy,
                "AVSDirectory.addStrategiesToOperatorSet: invalid strategy considered"
            );
            require(
                !avsOperatorSetStrategies[msg.sender][operatorSetID][strategy],
                "AVSDirectory.addStrategiesToOperatorSet: strategy already added"
            )
            emit OperatorSetStrategyAdded(msg.sender, operatorSetID, strategies[i]);
        }
    }

    /**
     * @notice Called by an avs to remove a strategy to its operator set
     * @param operatorSetID the ID of the operator set
     * @param strategies the list strategies to remove from the operator set
     */
    function removeStrategiesFromOperatorSet(
        bytes4 operatorSetID,
        IStrategy[] calldata strategies
    ) {
        uint256 strategiesToRemove = strategies.length;
        for (uint256 i = 0; i < strategiesToRemove; i++) {
            require(
                avsOperatorSetStrategies[msg.sender][operatorSetID][strategies[i]],
                "AVSDirectory.removeStrategiesFromOperatorSet: strategy not added"
            )
            emit OperatorSetStrategyRemoved(msg.sender, operatorSetID, strategies[i]);
        }
    }

    /*******************************************************************************
                            VIEW FUNCTIONS
    *******************************************************************************/

    /**
     * @notice Calculates the digest hash to be signed by an operator to register with an AVS
     * @param operator The account registering as an operator
     * @param avs The address of the service manager contract for the AVS that the operator is registering to
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid
     */
    function calculateOperatorAVSRegistrationDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 structHash = keccak256(
            abi.encode(OPERATOR_AVS_REGISTRATION_TYPEHASH, operator, avs, salt, expiry)
        );
        // calculate the digest hash
        bytes32 digestHash = keccak256(
            abi.encodePacked("\x19\x01", domainSeparator(), structHash)
        );
        return digestHash;
    }

    /**
     * @notice Getter function for the current EIP-712 domain separator for this contract.
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     */
    function domainSeparator() public view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return _calculateDomainSeparator();
        }
    }

    // @notice Internal function for calculating the current domain separator of this contract
    function _calculateDomainSeparator() internal view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
    }
}
