# multisig

使用solidity和golang来实现以太坊的多签功能。

## 简单多签

可以设置任意N:M的签名交易，最开始设置完N:M的多签之后是不能更改的。

具体实现请参考`SimpleMultiSig.sol`和`simplemultisig.go`。

**适用场景：**

多个账户共同管理一笔资产。

## 复杂多签

1. 实现任意修改可以发起交易的多个签名者；
2. 可以随时控制交易暂停或者启动；
3. 实现权限控制；

具体实现请参考`DynamicMultiSig.sol`和`dynamicmultisig.go`。

**适用场景：**

跨链桥资产管理。

## TODO

1. Ganache自动化部署和测试；
2. ERC20的测试。
