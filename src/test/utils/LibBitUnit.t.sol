// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../../contracts/libraries/LibBit.sol";

contract LibBitTest is Test {
    function test_Correctness_Flip() public {
        uint256 x = 0;
        uint8 n = 0;
        uint256 result = LibBit.flip(x, n);
        assertEq(result, 1);

        x = 1;
        result = LibBit.flip(x, n);
        assertEq(result, 0);

        x = 0;
        n = 255;
        result = LibBit.flip(x, n);
        assertEq(result, 1 << 255);
    }

    function test_Correctness_getLeft() public {
        uint256 x = 0x1234567890abcdef1234567890abcdefffffffffffffffffffffffffffffffff;
        uint256 result = LibBit.getLeft(x);
        assertEq(result, 0x1234567890abcdef1234567890abcdef);
    }

    function test_Correctness_getRight() public {
        uint256 x = 0xffffffffffffffffffffffffffffffff1234567890abcdef1234567890abcdef;
        uint256 result = LibBit.getRight(x);
        assertEq(result, 0x1234567890abcdef1234567890abcdef);
    }

    function test_Correctness_setLeft() public {
        uint256 x = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
        uint256 y = 0xffffffffffffffffffffffffffffffff;
        uint256 result = LibBit.setLeft(x, y);
        assertEq(result, 0xffffffffffffffffffffffffffffffff1234567890abcdef1234567890abcdef);
    }

    function test_Correctness_setRight() public {
        uint256 x = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
        uint256 y = 0xffffffffffffffffffffffffffffffff;
        uint256 result = LibBit.setRight(x, y);
        assertEq(result, 0x1234567890abcdef1234567890abcdefffffffffffffffffffffffffffffffff);
    }

    function test_Correctness_clz() public {
        uint256 x = 0;
        uint256 result = LibBit.clz(x);
        assertEq(result, 256);

        x = 1;
        result = LibBit.clz(x);
        assertEq(result, 255);

        x = 0x8000000000000000000000000000000000000000000000000000000000000000;
        result = LibBit.clz(x);
        assertEq(result, 0);

        x = 0x0800000000000000000000000000000000000000000000000000000000000000;
        result = LibBit.clz(x);
        assertEq(result, 4);
    }
}
