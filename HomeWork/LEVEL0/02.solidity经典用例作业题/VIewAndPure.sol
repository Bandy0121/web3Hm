// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract VIewAndPure {
     uint256 public x = 1;

    function addToX(uint256 y) public view returns (uint256) {
       return x + y;
    }

    function add(uint256 i,uint256 j)public pure returns (uint256)  {
       return i + j;
    }
}
