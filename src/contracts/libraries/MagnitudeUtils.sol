// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/**
 * @title Library of utilities for making magnitudes compatible with OZ's checkpoints library.
 * Each Checkpoint struct contains a uint32 timestamp key and a uint224 value. We pack multiple
 * magnitude values into the uint224 value to ensure a single storage slot is used for each checkpoint.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
library MagnitudeUtils {
    /**
     * TotalAndNonslashableMagnitude uses Checkpoints struct as follows
     *
     *     struct Checkpoint {
     *         uint32 _blockNumber;
     *         uint224 _value;
     *     }
     *     =>
     *     struct Checkpoint {
     *         uint32 timestamp;
     *         uint64 totalMagnitude;
     *         uint64 nonslashableMagnitude;
     *     }
     */

    /**
     * @dev Packs two uint64 values into a uint224 value.
     * @param totalMagnitude The total magnitude value.
     * @param nonslashableMagnitude The nonslashable magnitude value.
     * @return packedValue uint224 value used in Checkpoints storage
     */
    function packTotalAndNonslashableMagnitude(
        uint64 totalMagnitude,
        uint64 nonslashableMagnitude
    ) internal pure returns (uint224 packedValue) {
        assembly {
            // Pack the two uint64 values into a uint224 value using bitwise operations
            packedValue := or(shl(64, totalMagnitude), nonslashableMagnitude)
        }
    }

    /**
     * @dev Unpacks a uint224 value into two uint64 values.
     * @param packedValue The packed uint224 value.
     * @return totalMagnitude The total magnitude value.
     * @return nonslashableMagnitude The nonslashable magnitude value.
     */
    function unpackTotalAndNonslashableMagnitude(uint224 packedValue)
        internal
        pure
        returns (uint64 totalMagnitude, uint64 nonslashableMagnitude)
    {
        assembly {
            // Extract the totalMagnitude by shifting right 64 bits and masking again
            totalMagnitude := and(shr(64, packedValue), 0xFFFFFFFFFFFFFFFF)
            // Mask the lower 64 bits to get nonslashableMagnitude
            nonslashableMagnitude := and(packedValue, 0xFFFFFFFFFFFFFFFF)
        }
    }

    /**
     * @dev Gets the totalMagnitude from a packed uint224 value.
     * @param packedValue The packed uint224 value.
     * @return totalMagnitude value
     */
    function totalMagnitude(uint224 packedValue) internal pure returns (uint64 totalMagnitude) {
        assembly {
            // Extract the totalMagnitude by shifting right 64 bits and masking to ensure only 64 bits are used
            totalMagnitude := and(shr(64, packedValue), 0xFFFFFFFFFFFFFFFF)
        }
    }

    /**
     * @dev Gets the nonslashableMagnitude from a packed uint224 value.
     * @param packedValue The packed uint224 value.
     * @return nonslashableMagnitude value
     */
    function nonslashableMagnitude(uint224 packedValue) internal pure returns (uint64 nonslashableMagnitude) {
        assembly {
            // Mask the lower 64 bits to get nonslashableMagnitude
            nonslashableMagnitude := and(packedValue, 0xFFFFFFFFFFFFFFFF)
        }
    }
}
