package eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/terry108/multisig/eth/tool"
)

func TestDelpoyMultiSigContract(t *testing.T) {
	// 初始化准备
	{
		for _, v := range privkHexList {
			addr := tool.GenAddr(v)
			addrs = append(addrs, addr)
		}
		sort.Slice(addrs, func(i, j int) bool {
			return addrs[i].Address < addrs[j].Address
		})
		a0, a1, a2, a3 = addrs[0], addrs[1], addrs[2], addrs[3]
	}

	client, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	tool.WaitSomething(t, time.Minute, func() error { _, e := client.NetworkID(context.Background()); return e })
	{ // 首先确定addr0里的余额
		bal, err := client.BalanceAt(context.Background(), a0.ToAddress(), nil)
		tool.FailOnErr(t, err, "无法获取地址的余额")
		fmt.Printf("addr0: %s, 余额: %v\n", a0.Address, bal)
	}
	{ // 部署合约测试
		hexAddress := []string{a1.Address, a2.Address, a3.Address}
		chainID, _ := client.ChainID(context.Background())
		got, err := DeployDynamicMultiSigContract(chainID, client, a0.PrivkHex, hexAddress)
		if err != nil {
			t.Errorf("DeployMultiSigWalletContract() error = %v", err)
			t.FailNow()
		}
		fmt.Println("deployMultisigWalletContract got:", got)
	}
}

func TestDynamicMultisig(t *testing.T) {
	// 初始化准备
	{
		for _, v := range privkHexList {
			addr := tool.GenAddr(v)
			addrs = append(addrs, addr)
		}
		a0, a1, a2, a3 = addrs[0], addrs[1], addrs[2], addrs[3]
	}

	client, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	{ // 交易测试
		// multisigContract, err := contracts.NewDynamicMultiSig(contractAddress, client)
		// tool.FailOnErr(t, err, "构建多签合约调用时异常,检查合约地址和rpc server")
		var (
			sig                                  []byte //签名
			privkHex                             string
			multisigContractAddress, fromAddress string //多签合约地址，发起地址
			destination                          string //toAddress
			value, gasLimit                      *big.Int
			data                                 []byte
		)
		// 123由2发起，1和3签名, 把钱赚到3的地址上，executor为0
		privkHex = a2.PrivkHex
		multisigContractAddress = contractAddressHex
		fromAddress = a2.Address
		destination = a2.Address
		value = big.NewInt(1).Mul(big.NewInt(tool.E18), big.NewInt(2))
		gasLimit = big.NewInt(8239963)
		data = []byte("")
		// 8732bff38f474d3ab41c29c53ae0ff73955460582f47efd74a46b9734c4301c9
		// 5541a24b3f1d3ed8a01809c1ad5b657291e1f182c4d2e03ecf2cce583813513c60525ccbe1af9fbddc72b153850f79520c0f9a7b4ab1a81b02c9b8dbb377324100
		txKey := "8732bff38f474d3ab41c29c53ae0ff73955460582f47efd74a46b9734c4301c9"

		chainID, _ := client.ChainID(context.Background())
		for _, add := range []*tool.AddrInfo{a1, a3} {
			s, err := DynamicMultiSigExecuteSign(txKey, add.PrivkHex, multisigContractAddress, destination, value)
			fmt.Println("sig length:", len(s))
			fmt.Println("sig:", hex.EncodeToString(s))
			tool.FailOnErr(t, err, "create sig failed")
			sig = append(sig, s...)
		}
		par := &TxDynamicParams{
			Backend:                 client,
			TxKey:                   txKey,
			ExpireTime:              expireTime,
			Sig:                     sig,
			PrivkHex:                privkHex,
			MultisigContractAddress: multisigContractAddress,
			FromAddress:             fromAddress,
			Destination:             destination,
			Value:                   value,
			GasLimit:                gasLimit,
			Data:                    data,
			ChainID:                 chainID,
		}
		txid, err := ExecuteDynamicTX(par)

		tool.FailOnErr(t, err, "Execute Failed")
		fmt.Println("execute txid", txid)
	}
}

func TestHash(t *testing.T) {
	// leftPad64 := func(str string) string { return _AllZero[:64-len(str)] + str }
	// txKey := "8732bff38f474d3ab41c29c53ae0ff73955460582f47efd74a46b9734c4301c9"
	// to := "0x8d505f08E421d59794F462FF0Cc5b01787838CE0"
	// to = strings.ToLower(to)
	// amount := "000000000000000000000000645E186E556DcC8983a7BB1E68Ca596D3a2BAe9b"
	// isERC20 := "0000000000000000000000000000000000000000000000000000000000000000"
	// erc20 := "000000000000000000000000645E186E556DcC8983a7BB1E68Ca596D3a2BAe9b"
	// version := "0000000000000000000000000000000000000000000000000000000000000002"
	// hashBytes := crypto.Keccak256(
	// 	// []byte(txKey),
	// 	[]byte(leftPad64(to[2:])),
	// 	// []byte(amount),
	// 	// []byte(isERC20),
	// 	// []byte(erc20),
	// 	// []byte(version),
	// )
	// fmt.Println(hex.EncodeToString(hashBytes))
	// 0xc26fab2956146bc15979f8ef3e5b5bef54f12c6f780eef87f6fc31fe6cf6dac4

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
