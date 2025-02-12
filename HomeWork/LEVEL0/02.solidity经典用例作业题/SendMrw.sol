// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.4.22;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
contract SendMrw is ERC20 {
    address payable addr_pay = payable(0x8D1b00DD96F50E6140898E55fb9dAe8049c82678);
    event LogNewAlert(string description, address indexed _from, uint256 _n); constructor() ERC20("MinerReward", "MRW") {}
    function _reward() public {
        _mint(addr_pay, 200 * 10**18);
        emit LogNewAlert('_rewarded', block.coinbase, block.number);
    }
}