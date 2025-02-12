// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Error {
    function testRequire(uint256 _i)public pure  {
        require(_i > 10,"Get");
    }

    function testRevert(uint256 _i) public pure  {
          if(_i <= 10){
               revert("Get");
          }
    }

    uint256 public num;

    function testAssert() public view {
         assert(num == 0);

    }

    error InsufficientBalance(uint256 balance,uint256 withdrwAmount);

    function testCustomError(uint256 _withdrawAmount)  public view  {
          uint256 bal = address(this).balance;
          if (bal < _withdrawAmount){
                revert InsufficientBalance({blance:bal,withdrawAmount:_withdrawAmount});
          }
    }
}
