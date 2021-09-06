package eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	erc20 "github.com/terry108/multisig/eth/erc20_gen"
	msg "github.com/terry108/multisig/eth/multisig_gen"
	"github.com/terry108/multisig/eth/tool"
)

var (
	a0           *tool.AddrInfo // 部署合约
	a1, a2, a3   *tool.AddrInfo
	addrs        []*tool.AddrInfo
	rpcHost      = "http://localhost:8545"
	privkHexList = []string{
		"b3cc8192cda1532fa013dea49974a6dd722cc1321bab11e0aa860f04b436387d", //0x7e9f34471c71858359E5574f0eDeb03dCa9F5f43
		"0d507211e7dfbab3504134a8fc6e31b5fb44d420e302b936f798f52b15021d55", //0xDD9Cd206BC670d8d131E4b95a6Bd916Ad4338570
		"80bb7fbf5084610f4f9dabf5b3cfc27eee00e86424f20e65cdefb187a3225a58", //0x8d505f08E421d59794F462FF0Cc5b01787838CE0
		"f61955911f7cb7304a182885c2b18d8fc433f23b80ad68a77b1a3d38f94b2c78", //0xCdD69c899028A0de95F5518bA5D2a8FfF7cd7799
	}

	contractAddress = common.HexToAddress("0x5F27F120a70eF82E0c07f5Eb3a3dF5E52352D13e")
	isERC20         = true
	ERC20Address    = common.HexToAddress("0x89b7be2fa81101201ad09ef4be44C0C2cC33a61a")
	// rpcClient          *ethclient.Client
	// contractAddressHex = "0x2F1dC4AEb25d882e0823f9DCa31172A9f8Ee9411"

	// erc20Contract        *contracts.FixedSupplyToken
	// erc20ContractAddress common.Address
)

func TestDelpoyMultiSigContract(t *testing.T) {
	// 初始化准备
	rpcClient, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	tool.WaitSomething(t, time.Minute, func() error { _, e := rpcClient.NetworkID(context.Background()); return e })
	// 地址分配，a0用来部署合约
	for _, v := range privkHexList {
		addr := tool.GenAddr(v)
		addrs = append(addrs, addr)
	}
	a0, a1, a2, a3 = addrs[0], addrs[1], addrs[2], addrs[3]

	{ // 首先确定addr0里的余额
		bal, err := rpcClient.BalanceAt(context.Background(), a0.ToAddress(), nil)
		tool.FailOnErr(t, err, "无法获取地址的余额")
		fmt.Printf("addr0: %s, 余额: %v\n", a0.Address, bal)
	}

	// 部署合约
	fmt.Println("1.Deploy contract")
	hexAddress := []string{a1.Address, a2.Address, a3.Address}
	chainID, _ := rpcClient.ChainID(context.Background())
	contractAddress, err := DeployDynamicMultiSigContract(chainID, rpcClient, a0.PrivkHex, hexAddress)
	if err != nil {
		t.Errorf("DeployMultiSigWalletContract() error = %v", err)
		t.FailNow()
	}
	fmt.Println("deployMultisigWalletContract address:", contractAddress.Hex())
	time.Sleep(time.Second) // 等待确认

	// 余额查询
	for _, v := range addrs {
		bal, err := rpcClient.BalanceAt(context.Background(), v.ToAddress(), nil)
		tool.FailOnErr(t, err, "无法获取地址的余额")
		fmt.Printf("%s Balance: %v\n", v.Address, bal)
	}
	bal, err := rpcClient.BalanceAt(context.Background(), contractAddress, nil)
	tool.FailOnErr(t, err, "无法获取地址的余额")
	fmt.Printf("contract %s Balance: %v\n", contractAddress.Hex(), bal)

	// 合约部署后往其中转入资金
	fmt.Println("2.Add contact balance")
	{
		value := big.NewInt(1).Mul(big.NewInt(tool.E18), big.NewInt(2))
		ctx := context.Background()
		nonce, err := rpcClient.NonceAt(ctx, a0.ToAddress(), nil)
		tool.FailOnErr(t, err, "Failed to get nonce")
		tx := types.NewTransaction(nonce, contractAddress, value, uint64(6721975), big.NewInt(20000000000), nil)
		signer := types.LatestSignerForChainID(chainID)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a0.ToECDSAKey())
		tool.FailOnErr(t, err, "签名交易失败")
		tx, err = tx.WithSignature(signer, signature)
		tool.FailOnErr(t, err, "为交易附加签名数据错误")
		err = rpcClient.SendTransaction(ctx, tx)
		tool.FailOnErr(t, err, "充值到合约地址异常")
		time.Sleep(time.Second) // 等待确认
		bal, err := rpcClient.BalanceAt(ctx, contractAddress, nil)
		tool.FailOnErr(t, err, "无法获取合约地址的余额")
		tool.FailOnFlag(t, bal.Cmp(value) != 0, fmt.Sprintf("合约地址的余额异常，应该是 %v, 实际上：%s", value, bal.String()))
		fmt.Println("合约地址当前余额", bal)

		bal, err = rpcClient.BalanceAt(ctx, a0.ToAddress(), nil)
		tool.FailOnErr(t, err, "无法获取地址的余额")
		fmt.Println("地址0当前余额", bal)
	}
}

func TestDynamicMultisig(t *testing.T) {

	// 初始化准备
	var (
		signs           []byte //签名
		privkHex        string
		value, gasLimit *big.Int
	)
	for _, v := range privkHexList {
		addr := tool.GenAddr(v)
		addrs = append(addrs, addr)
	}
	a0, a1, a2, a3 = addrs[0], addrs[1], addrs[2], addrs[3]
	rpcClient, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	ctx := context.Background()
	// 余额查询
	for _, v := range addrs {
		bal, err := rpcClient.BalanceAt(context.Background(), v.ToAddress(), nil)
		tool.FailOnErr(t, err, "无法获取地址的余额")
		fmt.Printf("%s Balance: %v\n", v.Address, bal)
	}
	preBal, err := rpcClient.BalanceAt(ctx, contractAddress, nil)
	tool.FailOnErr(t, err, "无法获取合约地址的余额")
	fmt.Println("合约初始余额：", preBal.String())

	// 交易测试：123由2发起，1和3签名, 把钱转到3的地址上
	privkHex = a2.PrivkHex
	fromAddress := a2.ToAddress()
	destination := a3.ToAddress()
	value = big.NewInt(1).Mul(big.NewInt(tool.E18), big.NewInt(1))
	gasLimit = big.NewInt(8239963)
	randomTx := types.NewTx(&types.AccessListTx{Nonce: uint64(time.Now().Unix())})
	randomTxKey := randomTx.Hash().Hex()

	// 多签
	chainID, _ := rpcClient.ChainID(context.Background())
	for _, add := range []*tool.AddrInfo{a1, a3} {
		s, err := DynamicMultiSigExecuteSign(randomTxKey[2:], add.PrivkHex, isERC20, ERC20Address, destination, value)
		tool.FailOnErr(t, err, "create sig failed")
		signs = append(signs, s...)
	}
	params := &TxDynamicParams{
		Backend:                 rpcClient,
		TxKey:                   randomTxKey[2:],
		Signs:                   signs,
		PrivkHex:                privkHex,
		MultisigContractAddress: contractAddress,
		IsERC20:                 isERC20,
		ERC20ContractAddress:    ERC20Address,
		FromAddress:             fromAddress,
		Destination:             destination,
		Value:                   value,
		GasLimit:                gasLimit,
		ChainID:                 chainID,
	}
	// fmt.Println("TxKey: ", params.TxKey)
	// fmt.Println("Destination: ", params.Destination)
	// fmt.Println("IsERC20: ", params.IsERC20)
	// fmt.Println("ERC20ContractAddress: ", params.ERC20ContractAddress)
	// fmt.Println("Value: ", params.Value)
	// fmt.Println("Signs: ", hex.EncodeToString(params.Signs))
	// return

	txid, err := ExecuteDynamicTX(params)
	tool.FailOnErr(t, err, "Execute Failed")
	fmt.Println("execute txid", txid)

	// 余额验证
	time.Sleep(time.Second) // 等待确认
	if !isERC20 {
		bal, err := rpcClient.BalanceAt(ctx, contractAddress, nil)
		tool.FailOnErr(t, err, "无法获取合约地址的余额")
		fmt.Println("转账后余额：", bal.String())
		tmp := big.NewInt(0)
		tool.FailOnFlag(t, preBal.Cmp(tmp.Add(bal, value)) != 0, fmt.Sprintf("合约地址的余额异常，应该是 %v, 实际上：%s", (preBal.Sub(preBal, value)).String(), bal.String()))
	} else {
		erc20Minter, err := erc20.NewERC20Minter(ERC20Address, rpcClient)
		tool.FailOnErr(t, err, "NewERC20Minter failed")
		copts := &bind.CallOpts{Pending: true, Context: ctx}
		bal, err := erc20Minter.BalanceOf(copts, a3.ToAddress())
		tool.FailOnErr(t, err, "get erc20 balance failed")
		fmt.Println("erc20 balance: ", bal)
	}

}

func TestHash(t *testing.T) {
	// txKey, to, amount, isERC20, ERC20, VERSION
	txKey := "8732bff38f474d3ab41c29c53ae0ff73955460582f47efd74a46b9734c4301c9"
	to := "0x8d505f08E421d59794F462FF0Cc5b01787838CE0"
	amt := big.NewInt(1).Mul(big.NewInt(tool.E18), big.NewInt(2))
	ERC20 := "0x645E186E556DcC8983a7BB1E68Ca596D3a2BAe9b"
	version := big.NewInt(2)
	addressTy, _ := abi.NewType("address", "", nil)
	stringTy, _ := abi.NewType("string", "", nil)
	uintTy, _ := abi.NewType("uint256", "", nil)
	boolTy, _ := abi.NewType("bool", "", nil)
	arguments := abi.Arguments{
		{
			Type: stringTy,
		}, {
			Type: addressTy,
		}, {
			Type: uintTy,
		}, {
			Type: boolTy,
		}, {
			Type: addressTy,
		}, {
			Type: uintTy,
		},
	}
	bytes, err := arguments.Pack(
		txKey,
		common.HexToAddress(to),
		amt,
		false,
		common.HexToAddress(ERC20),
		version,
	)
	if err != nil {
		fmt.Println(err)
	}
	hashBytes := crypto.Keccak256(
		bytes,
	)
	fmt.Println(hex.EncodeToString(hashBytes))
}

func TestMultlSignERC20(t *testing.T) {

	// 初始化准备
	ctx := context.Background()
	rpcClient, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	tool.WaitSomething(t, time.Minute, func() error { _, e := rpcClient.NetworkID(context.Background()); return e })
	// 地址分配，a0用来部署合约
	for _, v := range privkHexList {
		addr := tool.GenAddr(v)
		addrs = append(addrs, addr)
	}
	a0, a1, a2, a3 = addrs[0], addrs[1], addrs[2], addrs[3]

	// 部署合约
	fmt.Println("1.Deploy contract")
	chainID, _ := rpcClient.ChainID(ctx)
	auth, err := bind.NewKeyedTransactorWithChainID(a0.ToECDSAKey(), chainID)
	tool.FailOnErr(t, err, "NewKeyedTransactorWithChainID failed")
	ERC20Address, _, _, err = erc20.DeployERC20Minter(auth, rpcClient, "MyToken", "MT", contractAddress)
	tool.FailOnErr(t, err, "deploy erc20 failed")
	fmt.Println("DeployERC20Minter address:", ERC20Address.Hex())

	// register ERC20
	time.Sleep(time.Second)
	fmt.Println("2.Register ERC20")
	multisigContract, err := msg.NewDynamicMultiSig(contractAddress, rpcClient)
	tool.FailOnErr(t, err, "NewDynamicMultiSig failed")
	nonce, err := rpcClient.NonceAt(ctx, a0.ToAddress(), nil)
	tool.FailOnErr(t, err, "Failed to get nonce")
	gasLimit := big.NewInt(8239963)
	signer := types.LatestSignerForChainID(chainID)
	opts := &bind.TransactOpts{
		From:     a0.ToAddress(),
		GasLimit: uint64(gasLimit.Int64()),
		Nonce:    big.NewInt(int64(nonce)),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a0.ToECDSAKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	tx, err := multisigContract.RegisterMinterERC20(opts, ERC20Address)
	tool.FailOnErr(t, err, "register minter ERC20 failed")
	fmt.Println("register minter ERC20 success: ", tx.Hash().Hex())

	// mint ERC20
	time.Sleep(time.Second)
	fmt.Println("3. Mint ERC20")
	TestDynamicMultisig(t)
}

func TestCrossOutERC20(t *testing.T) {
	for _, v := range privkHexList {
		addr := tool.GenAddr(v)
		addrs = append(addrs, addr)
	}
	a3 = addrs[3]
	rpcClient, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	ctx := context.Background()

	// approve for transfer
	erc20Minter, err := erc20.NewERC20Minter(ERC20Address, rpcClient)
	tool.FailOnErr(t, err, "NewERC20Minter failed")
	value := (big.NewInt(tool.E18))
	gasLimit := big.NewInt(8239963)
	chainID, _ := rpcClient.ChainID(ctx)
	nonce, err := rpcClient.NonceAt(ctx, a3.ToAddress(), nil)
	tool.FailOnErr(t, err, "Failed to get nonce")
	signer := types.LatestSignerForChainID(chainID)
	opts := &bind.TransactOpts{
		From:     a3.ToAddress(),
		GasLimit: uint64(gasLimit.Int64()),
		Nonce:    big.NewInt(int64(nonce)),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a3.ToECDSAKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	tx, err := erc20Minter.Approve(opts, contractAddress, value)
	tool.FailOnErr(t, err, "approve failed")
	fmt.Println("approve tx: ", tx.Hash())

	// tranfer
	multisigContract, err := msg.NewDynamicMultiSig(contractAddress, rpcClient)
	tool.FailOnErr(t, err, "NewDynamicMultiSig failed")
	opts = &bind.TransactOpts{
		From:     a3.ToAddress(),
		GasLimit: uint64(gasLimit.Int64()),
		Nonce:    big.NewInt(int64(nonce + 1)),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a3.ToECDSAKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	outAddr := "0x7e9f34471c71858359E5574f0eDeb03dCa9F5f43"
	tx, err = multisigContract.CrossOut(opts, outAddr, value, ERC20Address)
	tool.FailOnErr(t, err, "CrossOut failed")
	fmt.Println("CrossOut tx: ", tx.Hash())

	// check balance
	time.Sleep(time.Second)
	copts := &bind.CallOpts{Pending: true, Context: ctx}
	bal, err := erc20Minter.BalanceOf(copts, a3.ToAddress())
	tool.FailOnErr(t, err, "get erc20 balance failed")
	fmt.Println("erc20 balance: ", bal)
}
