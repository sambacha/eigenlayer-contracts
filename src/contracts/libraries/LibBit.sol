// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

library LibBit {
    /// @dev Flips the nth bit (0-indexed) of `x`.
    function flip(uint256 x, uint8 n) internal pure returns (uint256 r) {
        /// @solidity memory-safe-assembly
        assembly {
            r := xor(x, shl(n, 1))
        }
    }

    /// @dev Returns the first 128 bits of `x`.
    function getLeft(uint256 x) internal pure returns (uint256 r) {
        /// @solidity memory-safe-assembly
        assembly {
            r := shr(128, x)
        }
    }

    /// @dev Returns the second 128 bits of `x`.
    function getRight(uint256 x) internal pure returns (uint256 r) {
        /// @solidity memory-safe-assembly
        assembly {
            r := and(shr(128, not(0)), x)
        }
    }

    /// @dev Updates the first 128 bits of `x` with `leftValue`.
    function setLeft(uint256 x, uint256 y) internal pure returns (uint256 r) {
        /// @solidity memory-safe-assembly
        assembly {
            r := or(and(0xffffffffffffffffffffffffffffffff, x), shl(128, y))
        }
    }

    /// @dev Updates the second 128 bits of `x` with `y`.
    function setRight(uint256 x, uint256 y) internal pure returns (uint256 r) {
        /// @solidity memory-safe-assembly
        assembly {
            r := or(and(x, shl(128, not(0))), y)
        }
    }

    /// @dev Returns the number of zeros preceding the most significant one bit.
    /// If `x` is zero, returns `256`.
    function clz(uint256 x) internal pure returns (uint256 r) {
        /// @solidity memory-safe-assembly
        assembly {
            r := shl(7, lt(0xffffffffffffffffffffffffffffffff, x))
            r := or(r, shl(6, lt(0xffffffffffffffff, shr(r, x))))
            r := or(r, shl(5, lt(0xffffffff, shr(r, x))))
            r := or(r, shl(4, lt(0xffff, shr(r, x))))
            r := or(r, shl(3, lt(0xff, shr(r, x))))
            // forgefmt: disable-next-item
            r := add(xor(r, byte(and(0x1f, shr(shr(r, x), 0x8421084210842108cc6318c6db6d54be)),
                0xf8f9f9faf9fdfafbf9fdfcfdfafbfcfef9fafdfafcfcfbfefafafcfbffffffff)), iszero(x))
        }
    }
}
