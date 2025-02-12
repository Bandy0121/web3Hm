// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Immutable {
    address public immutable MY_ADDRESS;
    uint256 public immutable MY_UINT;

    constructor(uint256 myUint){
        MY_ADDRESS = msg.sender;
        MY_UINT = myUint;
    }
}
