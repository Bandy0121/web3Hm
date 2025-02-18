/*所有人都可以存钱
    ETH
    只有合约 owner 才可以取钱
    只要取钱，合约就销毁掉 selfdestruct
    扩展：支持主币以外的资产
    ERC20
    ERC721
    根据以上要求写一个存钱罐合约的例子
*/
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract PiggyBank {
    address public owner;

    // 构造函数，初始化合约所有者
    constructor() {
        owner = msg.sender;
    }

    // 修饰器，确保只有合约所有者可以调用某些函数
    modifier onlyOwner() {
        require(msg.sender == owner, "Not the owner");
        _;
    }

    // 接收 ETH 存款的函数
    receive() external payable {}

    // 存入 ERC20 代币的函数
    function depositERC20(address _tokenAddress, uint256 _amount) external {
        IERC20 token = IERC20(_tokenAddress);
        require(
            token.transferFrom(msg.sender, address(this), _amount),
            "Transfer failed"
        );
    }

    // 存入 ERC721 代币的函数
    function depositERC721(address _tokenAddress, uint256 _tokenId) external {
        IERC721 token = IERC721(_tokenAddress);
        token.safeTransferFrom(msg.sender, address(this), _tokenId);
    }

    // 取出 ETH、ERC20 和 ERC721 资产并销毁合约的函数
    function withdrawAndDestroy(
        address[] memory _erc20Tokens,
        address[] memory _erc721Tokens,
        uint256[] memory _erc721TokenIds
    ) external onlyOwner {
        // 取出 ETH
        uint256 ethBalance = address(this).balance;
        if (ethBalance > 0) {
            payable(owner).transfer(ethBalance);
        }

        // 取出 ERC20 代币
        for (uint256 i = 0; i < _erc20Tokens.length; i++) {
            IERC20 token = IERC20(_erc20Tokens[i]);
            uint256 tokenBalance = token.balanceOf(address(this));
            if (tokenBalance > 0) {
                token.transfer(owner, tokenBalance);
            }
        }

        // 取出 ERC721 代币
        for (uint256 i = 0; i < _erc721Tokens.length; i++) {
            IERC721 token = IERC721(_erc721Tokens[i]);
            token.safeTransferFrom(address(this), owner, _erc721TokenIds[i]);
        }

        // 销毁合约
        // selfdestruct(payable(owner));  已经被弃用了
    }
}

/*
    WETH 是包装 ETH 主币，作为 ERC20 的合约。 标准的 ERC20 合约包括如下几个

    3 个查询
        balanceOf: 查询指定地址的 Token 数量
        allowance: 查询指定地址对另外一个地址的剩余授权额度
        totalSupply: 查询当前合约的 Token 总量
    2 个交易
        transfer: 从当前调用者地址发送指定数量的 Token 到指定地址。
                  这是一个写入方法，所以还会抛出一个 Transfer 事件。
        transferFrom: 当向另外一个合约地址存款时，对方合约必须调用 transferFrom 才可以把 Token 拿到它自己的合约中。
    2 个事件
        Transfer
        Approval
    1 个授权
        approve: 授权指定地址可以操作调用者的最大 Token 数量。
    根据以上要求写一个WETH 合约的例子
*/
contract WETH {
    string public constant name = "Wrapped Ether";
    string public constant symbol = "WETH";
    uint8 public constant decimals = 18;

    // 存储每个地址的 Token 余额
    mapping(address => uint256) private _balances;
    // 存储授权信息，[所有者地址][被授权地址] => 授权额度
    mapping(address => mapping(address => uint256)) private _allowances;
    // 总供应量
    uint256 private _totalSupply;

    // Transfer 事件，当发生 Token 转移时触发
    event Transfer(address indexed from, address indexed to, uint256 value);
    // Approval 事件，当发生授权操作时触发
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    // 查询指定地址的 Token 数量
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    // 查询指定地址对另外一个地址的剩余授权额度
    function allowance(address owner, address spender)
        public
        view
        returns (uint256)
    {
        return _allowances[owner][spender];
    }

    // 查询当前合约的 Token 总量
    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    // 从当前调用者地址发送指定数量的 Token 到指定地址
    function transfer(address to, uint256 value) public returns (bool) {
        address sender = msg.sender;
        _transfer(sender, to, value);
        return true;
    }

    // 授权指定地址可以操作调用者的最大 Token 数量
    function approve(address spender, uint256 value) public returns (bool) {
        address owner = msg.sender;
        _approve(owner, spender, value);
        return true;
    }

    // 当向另外一个合约地址存款时，对方合约必须调用 transferFrom 才可以把 Token 拿到它自己的合约中
    function transferFrom(
        address from,
        address to,
        uint256 value
    ) public returns (bool) {
        address spender = msg.sender;
        _spendAllowance(from, spender, value);
        _transfer(from, to, value);
        return true;
    }

    // 内部转账函数
    function _transfer(
        address from,
        address to,
        uint256 value
    ) internal {
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");

        _beforeTokenTransfer(from, to, value);

        uint256 fromBalance = _balances[from];
        require(fromBalance >= value, "ERC20: transfer amount exceeds balance");
        unchecked {
            _balances[from] = fromBalance - value;
        }
        _balances[to] += value;

        emit Transfer(from, to, value);
    }

    // 内部授权函数
    function _approve(
        address owner,
        address spender,
        uint256 value
    ) internal {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = value;
        emit Approval(owner, spender, value);
    }

    // 花费授权额度
    function _spendAllowance(
        address owner,
        address spender,
        uint256 value
    ) internal {
        uint256 currentAllowance = allowance(owner, spender);
        if (currentAllowance != type(uint256).max) {
            require(currentAllowance >= value, "ERC20: insufficient allowance");
            unchecked {
                _approve(owner, spender, currentAllowance - value);
            }
        }
    }

    // 包装 ETH，将 ETH 存入合约并生成相应的 WETH
    receive() external payable {
        _mint(msg.sender, msg.value);
    }

    // 解包 WETH，将 WETH 销毁并提取相应的 ETH
    function withdraw(uint256 amount) public {
        _burn(msg.sender, amount);
        payable(msg.sender).transfer(amount);
    }

    // 内部铸造函数
    function _mint(address account, uint256 value) internal {
        require(account != address(0), "ERC20: mint to the zero address");

        _beforeTokenTransfer(address(0), account, value);

        _totalSupply += value;
        _balances[account] += value;
        emit Transfer(address(0), account, value);
    }

    // 内部销毁函数
    function _burn(address account, uint256 value) internal {
        require(account != address(0), "ERC20: burn from the zero address");

        _beforeTokenTransfer(account, address(0), value);

        uint256 accountBalance = _balances[account];
        require(accountBalance >= value, "ERC20: burn amount exceeds balance");
        unchecked {
            _balances[account] = accountBalance - value;
        }
        _totalSupply -= value;

        emit Transfer(account, address(0), value);
    }

    // 转账前的钩子函数，可用于扩展逻辑
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 value
    ) internal virtual {}
}

/*
    TodoList: 是类似便签一样功能的东西，记录我们需要做的事情，以及完成状态。 
    1.需要完成的功能
    创建任务
    修改任务名称
    任务名写错的时候
    修改完成状态：
    手动指定完成或者未完成
    自动切换
    如果未完成状态下，改为完成
    如果完成状态，改为未完成
    根据以上要求写一个todoList合约的例子
*/
contract TodoList {
    // 定义任务结构体
    struct Task {
        string name;
        bool isCompleted;
    }

    // 任务数组，用于存储所有任务
    Task[] public tasks; // 29414

    // 创建任务的函数
    function createTask(string memory _name) public {
        tasks.push(Task({name: _name, isCompleted: false}));
    }

    // 修改任务名称的函数
    function modifyTaskName(uint256 _taskIndex, string memory _newName) public {
        require(_taskIndex < tasks.length, "Task index out of bounds");
        // 方法1: 直接修改，修改一个属性时候比较省 gas
        tasks[_taskIndex].name = _newName;
    }

    function modiName2(uint256 index_, string memory name_) external {
        // 方法2: 先获取储存到 storage，在修改，在修改多个属性的时候比较省 gas
        Task storage temp = tasks[index_];
        temp.name = name_;
    }

    // 手动修改任务完成状态的函数
    function manuallyChangeCompletionStatus(
        uint256 _taskIndex,
        bool _isCompleted
    ) public {
        require(_taskIndex < tasks.length, "Task index out of bounds");
        tasks[_taskIndex].isCompleted = _isCompleted;
    }

    // 修改完成状态2:自动切换 toggle
    function modiStatus2(uint256 index_) external {
        tasks[index_].isCompleted = !tasks[index_].isCompleted;
    }

    // 自动切换任务完成状态的函数
    function automaticallyToggleCompletionStatus(uint256 _taskIndex) public {
        require(_taskIndex < tasks.length, "Task index out of bounds");
        tasks[_taskIndex].isCompleted = !tasks[_taskIndex].isCompleted;
    }

    // 获取任务数量的函数
    function getTaskCount() public view returns (uint256) {
        return tasks.length;
    }

    // 获取任务1: memory : 2次拷贝
    // 29448 gas
    function get1(uint256 index_)
        external
        view
        returns (string memory name_, bool status_)
    {
        Task memory temp = tasks[index_];
        return (temp.name, temp.isCompleted);
    }

    // 获取任务2: storage : 1次拷贝
    // 预期：get2 的 gas 费用比较低（相对 get1）
    // 29388 gas
    function get2(uint256 index_)
        external
        view
        returns (string memory name_, bool status_)
    {
        Task storage temp = tasks[index_];
        return (temp.name, temp.isCompleted);
    }
}

/*
    众筹合约是一个募集资金的合约，在区块链上，我们是募集以太币，类似互联网业务的水滴筹。区块链早起的 ICO 就是类似业务。

    1.需求分析
    众筹合约分为两种角色：一个是受益人，一个是资助者。

    // 两种角色:
    //      受益人   beneficiary => address         => address 类型
    //      资助者   funders     => address:amount  => mapping 类型 或者 struct 类型
    状态变量按照众筹的业务：
    // 状态变量
    //      筹资目标数量    fundingGoal
    //      当前募集数量    fundingAmount
    //      资助者列表      funders
    //      资助者人数      fundersKey
    需要部署时候传入的数据:
    //      受益人
    //      筹资目标数量
    根据以上要求写一个众筹合约的例子
*/
contract Crowdfunding {
    // 受益人地址
    address public beneficiary;
    // 筹资目标数量（以 Wei 为单位）
    uint256 public fundingGoal;
    // 当前募集数量（以 Wei 为单位）
    uint256 public fundingAmount;
    // 资助者列表，记录每个资助者的资助金额
    mapping(address => uint256) public funders;
    // 资助者人数
    uint256 public fundersCount;

    // 构造函数，部署合约时传入受益人和筹资目标数量
    constructor(address _beneficiary, uint256 _fundingGoal) {
        beneficiary = _beneficiary;
        fundingGoal = _fundingGoal;
        fundingAmount = 0;
        fundersCount = 0;
    }

    // 资助函数，资助者调用此函数进行资金资助
    function contribute() external payable {
        require(msg.value > 0, "Contribution amount must be greater than 0");
        if (funders[msg.sender] == 0) {
            fundersCount++;
        }
        funders[msg.sender] += msg.value;
        fundingAmount += msg.value;

        // 检查是否达到筹资目标
        if (fundingAmount >= fundingGoal) {
            // 达到筹资目标，将资金转给受益人
            payable(beneficiary).transfer(fundingAmount);
        }
    }

    // 获取资助者的资助金额
    function getContribution(address _funder) external view returns (uint256) {
        return funders[_funder];
    }
}

/*
    这一个实战主要是加深大家对 3 个取钱方法的使用。

    任何人都可以发送金额到合约
    只有 owner 可以取款
    3 种取钱方式

    在以太坊智能合约开发中，这三种取钱方法 withdraw1、withdraw2 和 withdraw3 分别使用了 transfer、send 和 call 
    这三种不同的方式来转移以太币，它们在功能、安全性、异常处理和 gas 消耗等方面存在明显的差异
*/

contract EtherWallet {
    address payable public immutable owner;
    event Log(string funName, address from, uint256 value, bytes data);

    constructor() {
        owner = payable(msg.sender);
    }

    receive() external payable {
        emit Log("receive", msg.sender, msg.value, "");
    }

    function withdraw1() external {
        require(msg.sender == owner, "Not owner");
        // owner.transfer 相比 msg.sender 更消耗Gas
        // owner.transfer(address(this).balance);
        payable(msg.sender).transfer(100);
    }

    /*
        功能和使用场景
            该方法使用 transfer 函数将 100 Wei（以太坊的最小单位）的以太币从合约转移到调用者（也就是合约所有者，因为有 require(msg.sender == owner) 检查）的账户。transfer 是一种简单直接的以太币转移方式，通常用于安全要求较高、金额较小且确定性较强的转账操作。
        安全性
            内置检查：transfer 函数内置了一些安全检查，它会检查接收方地址是否有效，并且会检查合约的余额是否足够进行转账。如果余额不足或者接收方地址无效，transfer 会抛出异常并回滚交易，确保不会发生意外的资金转移。
            固定 gas 限制：transfer 函数会固定发送 2300 gas 给接收方，这个 gas 量通常只够接收方执行一些简单的日志记录操作，防止接收方的回退函数（fallback 或 receive）消耗过多的 gas 导致合约出现问题，从而避免了一些潜在的重入攻击风险。
        异常处理
            失败，transfer 会自动抛出异常，交易将会被回滚，调用者可以通过交易的失败状态得知转账未成功。
        gas 消耗
            由于 transfer 函数内置了安全检查和固定的 gas 限制，它的 gas 消耗相对较高。同时，注释中提到 owner.transfer 相比 msg.sender 更消耗 gas，这是因为 owner 是一个存储变量，访问存储变量会比直接使用 msg.sender （一个局部变量）消耗更多的 gas。
    */

    function withdraw2() external {
        require(msg.sender == owner, "Not owner");
        bool success = payable(msg.sender).send(200);
        require(success, "Send Failed");
    }

    /*
        功能和使用场景
            此方法使用 send 函数尝试将 200 Wei 的以太币从合约转移到调用者的账户。send 通常用于对转账结果的处理比较灵活的场景，因为它会返回一个布尔值表示转账是否成功。
        安全性
            有限检查：send 函数也会进行一些基本的检查，如检查接收方地址和合约余额，但它不会像 transfer 那样在转账失败时自动抛出异常，而是返回 false。
            固定 gas 限制：和 transfer 一样，send 也会固定发送 2300 gas 给接收方，以防止接收方的回退函数消耗过多 gas。
        异常处理
            send 函数返回一个布尔值 success，表示转账是否成功。调用者需要手动检查这个返回值，如果 success 为 false，则需要进行相应的错误处理，例如在代码中使用 require(success, "Send Failed"); 来确保转账成功。
        gas 消耗
            send 的 gas 消耗和 transfer 类似，因为它们都有固定的 2300 gas 限制，不过由于 send 不会自动抛出异常，在某些情况下可能会稍微节省一点 gas，但总体差异不大。       
    */

    function withdraw3() external {
        require(msg.sender == owner, "Not owner");
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success, "Call Failed");
    }

    /*
    功能和使用场景
        该方法使用 call 函数将合约的全部余额转移到调用者的账户。call 是一种更底层、更灵活的调用方式，它可以用于调用合约的函数，也可以用于转移以太币，适用于需要更多灵活性和控制的场景。
    安全性
        无固定 gas 限制：call 函数不会像 transfer 和 send 那样固定发送 2300 gas，而是可以通过 {value: ...} 语法指定要发送的以太币数量，并且可以在调用时传递额外的 gas。这使得接收方可以执行更复杂的操作，但也增加了重入攻击的风险，需要开发者自己进行安全检查和防范。
        自定义调用：call 可以调用接收方的任何函数，不仅仅是回退函数，因此需要确保接收方的代码是可信的，避免因调用恶意合约函数而导致资金损失。
    异常处理
        call 函数返回一个布尔值 success 和一个 bytes 类型的数据（在代码中使用 (bool success, ) 忽略了这个数据），表示调用是否成功。和 send 一样，调用者需要手动检查 success 的值，如果为 false，则需要进行相应的错误处理。
    gas 消耗
        call 的 gas 消耗相对更灵活，因为可以根据需要指定发送的 gas 量。在某些情况下，如果接收方需要执行复杂的操作，使用 call 可以避免因固定的 2300 gas 限制而导致操作失败，但如果不小心传递了过多的 gas，也会增加不必要的成本。
    */

    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
    /*
        总结
        transfer：安全性高，内置异常处理，固定 gas 限制，适用于安全要求高、金额小且操作简单的转账场景。
        send：相对灵活，返回布尔值表示转账结果，同样有固定 gas 限制，需要手动处理转账失败的情况。
        call：最灵活，但安全性较低，无固定 gas 限制，可用于更复杂的调用和转账场景，需要开发者自行处理安全问题和异常情况。
    */
}

/*
    多签钱包的功能: 合约有多个 owner，一笔交易发出后，需要多个 owner 确认，确认数达到最低要求数之后，才可以真正的执行。

    部署时候传入地址参数和需要的签名数
        多个 owner 地址
        发起交易的最低签名数
    有接受 ETH 主币的方法，
    除了存款外，其他所有方法都需要 owner 地址才可以触发
    发送前需要检测是否获得了足够的签名数
    使用发出的交易数量值作为签名的凭据 ID（类似上么）
    每次修改状态变量都需要抛出事件
    允许批准的交易，在没有真正执行前取消。
    足够数量的 approve 后，才允许真正执行。
*/
contract MultiSigWallet {
    // 状态变量
    address[] public owners;
    mapping(address => bool) public isOwner;
    uint256 public required;
    struct Transaction {
        address to;
        uint256 value;
        bytes data;
        bool exected;
    }
    Transaction[] public transactions;
    mapping(uint256 => mapping(address => bool)) public approved;
    // 事件
    event Deposit(address indexed sender, uint256 amount);
    event Submit(uint256 indexed txId);
    event Approve(address indexed owner, uint256 indexed txId);
    event Revoke(address indexed owner, uint256 indexed txId);
    event Execute(uint256 indexed txId);

    // receive
    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }

    // 函数修改器
    modifier onlyOwner() {
        require(isOwner[msg.sender], "not owner");
        _;
    }
    modifier txExists(uint256 _txId) {
        require(_txId < transactions.length, "tx doesn't exist");
        _;
    }
    modifier notApproved(uint256 _txId) {
        require(!approved[_txId][msg.sender], "tx already approved");
        _;
    }
    modifier notExecuted(uint256 _txId) {
        require(!transactions[_txId].exected, "tx is exected");
        _;
    }

    // 构造函数
    constructor(address[] memory _owners, uint256 _required) {
        require(_owners.length > 0, "owner required");
        require(
            _required > 0 && _required <= _owners.length,
            "invalid required number of owners"
        );
        for (uint256 index = 0; index < _owners.length; index++) {
            address owner = _owners[index];
            require(owner != address(0), "invalid owner");
            require(!isOwner[owner], "owner is not unique"); // 如果重复会抛出错误
            isOwner[owner] = true;
            owners.push(owner);
        }
        required = _required;
    }

    // 函数
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
    //新增交易函数
    function submit(
        address _to,
        uint256 _value,
        bytes calldata _data
    ) external onlyOwner returns (uint256) {
        transactions.push(
            Transaction({to: _to, value: _value, data: _data, exected: false})
        );
        emit Submit(transactions.length - 1);
        return transactions.length - 1;
    }
    //owner签名函数
    function approv(uint256 _txId)
        external
        onlyOwner
        txExists(_txId)
        notApproved(_txId)
        notExecuted(_txId)
    {
        approved[_txId][msg.sender] = true;
        emit Approve(msg.sender, _txId);
    }
    //交易执行函数
    function execute(uint256 _txId)
        external
        onlyOwner
        txExists(_txId)
        notExecuted(_txId)
    {
        require(getApprovalCount(_txId) >= required, "approvals < required");
        Transaction storage transaction = transactions[_txId];
        transaction.exected = true;
        (bool sucess, ) = transaction.to.call{value: transaction.value}(
            transaction.data
        );
        require(sucess, "tx failed");
        emit Execute(_txId);
    }

    function getApprovalCount(uint256 _txId)
        public
        view
        returns (uint256 count)
    {
        for (uint256 index = 0; index < owners.length; index++) {
            if (approved[_txId][owners[index]]) {
                count += 1;
            }
        }
    }

    function revoke(uint256 _txId)
        external
        onlyOwner
        txExists(_txId)
        notExecuted(_txId)
    {
        require(approved[_txId][msg.sender], "tx not approved");
        approved[_txId][msg.sender] = false;
        emit Revoke(msg.sender, _txId);
    }
}
