// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Loop {
    function loop() public {

        for(uint256 i =0;i<10;i++){
              if(i == 3){
                   continue;
              }
              if(i == 5){
                  break;
              }
        }
        uint256 j;
        while(j<10){
            j++;
        }
    }

}
