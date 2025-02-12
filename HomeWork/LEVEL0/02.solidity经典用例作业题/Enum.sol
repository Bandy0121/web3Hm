// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Enum {
     enum Status{
         Pending,
         Shipped,
         Accepted,
         Rejected,
         Canceled
     }

    Status public status;

    function set(Status _status) public  {
         status = _status;
    }

    function cancel() public  {
        status = Status.Canceled;
    }
// delete resets the enum to its first value, 0
    function reset() public {
        delete status;
    }
}
