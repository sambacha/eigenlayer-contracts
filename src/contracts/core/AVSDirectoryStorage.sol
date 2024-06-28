// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../libraries/LibBit.sol";

abstract contract AVSDirectoryStorage is IAVSDirectory {
    using LibBit for uint256;

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @notice The EIP-712 typehash for the `Registration` struct used by the contract
    bytes32 public constant OPERATOR_AVS_REGISTRATION_TYPEHASH =
        keccak256("OperatorAVSRegistration(address operator,address avs,bytes32 salt,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `OperatorSetRegistration` struct used by the contract
    bytes32 public constant OPERATOR_SET_REGISTRATION_TYPEHASH =
        keccak256("OperatorSetRegistration(address avs,uint32[] operatorSetIds,bytes32 salt,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `StandbyParams` struct used by the contract
    bytes32 public constant OPERATOR_STANDBY_UPDATE =
        keccak256("OperatorStandbyUpdate(StandbyParam[] standbyParams,bytes32 salt,uint256 expiry)");

    uint256 internal constant EPOCH_LENGTH = 7 days;

    uint256 internal constant EPOCHS_PER_EPOCH_SET = 128;

    uint256 internal constant EPOCH_SET_LENGTH = EPOCH_LENGTH * EPOCHS_PER_EPOCH_SET;

    uint256 internal immutable START_TIME = block.timestamp;

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegation;

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /// @notice Mapping: AVS => operator => enum of operator status to the AVS
    mapping(address => mapping(address => OperatorAVSRegistrationStatus)) public avsOperatorStatus;

    /// @notice Mapping: operator => 32-byte salt => whether or not the salt has already been used by the operator.
    /// @dev Salt is used in the `registerOperatorToAVS` and `registerOperatorToOperatorSet` function.
    mapping(address => mapping(bytes32 => bool)) public operatorSaltIsSpent;

    /// @notice Mapping: AVS => whether or not the AVS uses operator set
    mapping(address => bool) public isOperatorSetAVS;

    // NOTE: Removed original isOperatorInOperator(avs, operator, operatorSetId) => bool mapping...

    /// @notice Mapping: (avs, operator, operatorSetId, epochSetId) => isOperatorInOperatorSet
    mapping(address => mapping(address => mapping(uint32 => mapping(uint256 => uint256)))) internal
        _isOperatorInOperatorSet;

    /// @notice Mapping: avs => operator => number of operator sets the operator is registered for the AVS
    mapping(address => mapping(address => uint256)) public operatorAVSOperatorSetCount;

    /// @notice Mapping: avs = operator => operatorSetId => Whether the given operator set in standby mode or not
    mapping(address => mapping(address => mapping(uint32 => bool))) public onStandby;

    constructor(IDelegationManager _delegation, IStrategyManager _strategyManager) {
        delegation = _delegation;
        strategyManager = _strategyManager;
    }

    function currentEpoch() public view returns (uint256) {
        // The parameter `START_TIME` must be less than or equal to `block.timestamp`.
        uint256 elapsed = block.timestamp - START_TIME;
        unchecked {
            // If `elapsed` is less than `EPOCH_LENGTH` the quotient is rounded down to zero.
            return elapsed / EPOCH_LENGTH;
        }
    }

    function currentEpochSet() public view returns (uint256 epochSet, uint8 epochSetIndex) {
        // The parameter `START_TIME` must be less than or equal to `block.timestamp`.
        uint256 elapsed = block.timestamp - START_TIME;
        unchecked {
            // If `elapsed` is less than `EPOCH_SET_LENGTH` the quotient is rounded down to zero.
            epochSet = elapsed / EPOCH_SET_LENGTH;

            // If `elapsed` is less than `EPOCH_LENGTH` the quotient is rounded down to zero.
            epochSetIndex = uint8((elapsed / EPOCH_LENGTH) % EPOCHS_PER_EPOCH_SET);
        }
    }

    function isOperatorInOperatorSet(
        address avs,
        address operator,
        uint32 operatorSetId
    ) public virtual returns (bool included) {
        unchecked {
            // This loop should probably from now backwards...

            // If `currentEpoch()` is less than `EPOCHS_PER_EPOCH_SET` the quotient is rounded down to zero.
            uint256 currentEpochSetId = currentEpoch() / EPOCHS_PER_EPOCH_SET;

            // Iterate through epoch sets and see if operator is registered to the given operator set.
            for (uint256 i; i < currentEpochSetId; ++i) {
                // If operator has registered...
                if (_isOperatorInOperatorSet[avs][operator][operatorSetId][i].getLeft() != 0) {
                    // Iterate through the remaining epoch sets up until now and
                    // assert a deregistration has not occured.
                    for (uint256 j = i; j < currentEpochSetId; ++j) {
                        if (_isOperatorInOperatorSet[avs][operator][operatorSetId][j].getRight() != 0) {
                            return false;
                        }
                    }

                    return true;
                }
            }
        }
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[43] private __gap;
}
