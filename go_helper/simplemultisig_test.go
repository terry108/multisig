package eth

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	sms "github.com/terry108/multisig/go_abi/sms_abi"
	"github.com/terry108/multisig/go_helper/tool"
)

var expireTime = big.NewInt(time.Now().Add(3 * 24 * time.Hour).Unix())

const (
	bucketNum = 10
	bucketIdx = 0
)

func TestDeploySimpleMultiSigContract(t *testing.T) {
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
	var (
		contractAddressHex string
		contractAddress    common.Address
	)
	{ // 部署合约测试
		type args struct {
			privkHex   string
			hexAddress []string
			mRequired  uint8
		}
		arg := args{
			privkHex:   a0.PrivkHex,
			hexAddress: []string{a0.Address, a1.Address, a2.Address},
			mRequired:  2,
		}

		fmt.Println("owners:", arg.hexAddress)
		chainID, _ := client.ChainID(context.Background())
		got, err := DeploySimpleMultiSigContract(bucketNum, chainID, client, arg.privkHex, arg.hexAddress, arg.mRequired)
		if err != nil {
			t.Errorf("DeployMultiSigWalletContract() error = %v", err)
			t.FailNow()
		}
		fmt.Println("deployMultisigWalletContract got:", got)

		contractAddressHex = got
		contractAddress = common.HexToAddress(contractAddressHex)
		fmt.Println("contractAddressHex", contractAddressHex)
		fmt.Println("contractAddress", contractAddress)
		// owners: [0x7e9f34471c71858359E5574f0eDeb03dCa9F5f43 0x8d505f08E421d59794F462FF0Cc5b01787838CE0 0xCdD69c899028A0de95F5518bA5D2a8FfF7cd7799]
		// deployMultisigWalletContract got: 0x27888Fb851bb9c5B9E802Ca2e1f2d13fc2e55909
		// contractAddressHex 0x27888Fb851bb9c5B9E802Ca2e1f2d13fc2e55909
	}
}

func TestGetAllBalance(t *testing.T) {

	client, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	for _, v := range privkHexList {
		addr := tool.GenAddr(v)
		bal, err := client.BalanceAt(context.Background(), addr.ToAddress(), nil)
		tool.FailOnErr(t, err, "无法获取合约地址的余额")
		fmt.Printf("%s, %v\n", addr.Address, bal)
	}

	bal, _ := client.BalanceAt(context.Background(), contractAddress, nil)
	fmt.Println("合约地址当前余额", bal)
}

func TestGetContractInfo(t *testing.T) {
	client, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	contractAddress := "0x27888Fb851bb9c5B9E802Ca2e1f2d13fc2e55909"
	addrs, mRequired, err := GetContractInfo(client, contractAddress)
	tool.FailOnErr(t, err, "get contract info faild")
	fmt.Printf("addrs: %v, mRequired: %d\n", addrs, mRequired)
}

func TestAddETH(t *testing.T) {
	a0 = tool.GenAddr(privkHexList[0])
	client, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")

	{ //合约部署后往其中转入资金
		value := big.NewInt(1).Mul(big.NewInt(tool.E18), big.NewInt(2))
		ctx := context.Background()
		nonce, err := client.NonceAt(ctx, contractAddress, nil)
		tool.FailOnErr(t, err, "Failed to get nonce")
		tx := types.NewTransaction(nonce, contractAddress, value, uint64(6721975), big.NewInt(20000000000), nil)
		signer := types.MakeSigner(params.TestChainConfig, nil)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a0.ToECDSAKey())
		tool.FailOnErr(t, err, "签名交易失败")
		tx, err = tx.WithSignature(signer, signature)
		tool.FailOnErr(t, err, "为交易附加签名数据错误")
		err = client.SendTransaction(ctx, tx)
		tool.FailOnErr(t, err, "充值到合约地址异常")

		bal, err := client.BalanceAt(ctx, contractAddress, nil)
		tool.FailOnErr(t, err, "无法获取合约地址的余额")
		tool.FailOnFlag(t, bal.Cmp(value) != 0, fmt.Sprintf("合约地址的余额异常，应该是 %v, 实际上：%s", value, bal.String()))
		fmt.Println("合约地址当前余额", bal)

		bal, err = client.BalanceAt(ctx, a0.ToAddress(), nil)
		tool.FailOnErr(t, err, "无法获取地址的余额")
		fmt.Println("地址0当前余额", bal)
	}
}

func TestMutisigWithdraw(t *testing.T) {
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
	outAddr := a3.Address
	transferValue := big.NewInt(tool.E18)
	{ // 交易测试
		multisigContract, err := sms.NewSimpleMultiSig(contractAddress, client)
		tool.FailOnErr(t, err, "构建多签合约调用时异常,检查合约地址和rpc server")

		nonce, err := multisigContract.NonceBucket(&bind.CallOpts{Pending: true}, big.NewInt(bucketIdx))
		// nonce, err := multisigContract.Nonce(&bind.CallOpts{Pending: true})
		tool.FailOnErr(t, err, "无法获取合约内部nonce")
		var (
			sigV                                 []uint8    //签名
			sigR, sigS                           [][32]byte //签名
			privkHex                             string
			multisigContractAddress, fromAddress string //多签合约地址，发起地址
			destination, executor                string //toAddress
			value, gasLimit                      *big.Int
			data                                 []byte
		)
		// 012由0发起，0和2签名, 把钱赚到3的地址上，executor 为0
		privkHex = a0.PrivkHex
		multisigContractAddress = contractAddress.Hex()
		fromAddress = a0.Address
		executor = a0.Address
		destination = outAddr
		value = transferValue
		gasLimit = big.NewInt(239963)
		data = []byte("")

		chainID, _ := client.ChainID(context.Background())
		for _, add := range []*tool.AddrInfo{a0, a2} {
			v, r, s, err := SimpleMultiSigExecuteSign(expireTime, chainID.Int64(), add.PrivkHex, multisigContractAddress, destination, executor, nonce, value, gasLimit, data)
			tool.FailOnErr(t, err, "create sig failed")
			sigV = append(sigV, v)
			sigR = append(sigR, r)
			sigS = append(sigS, s)
		}

		txid, err := ExecuteTX(&TxParams{
			Backend:                 client,
			BucketIdx:               0,
			ExpireTime:              expireTime,
			SigV:                    sigV,
			SigR:                    sigR,
			SigS:                    sigS,
			PrivkHex:                privkHex,
			MultisigContractAddress: multisigContractAddress,
			FromAddress:             fromAddress,
			Destination:             destination,
			Executor:                executor,
			Value:                   value,
			GasLimit:                gasLimit,
			Data:                    data,
			ChainID:                 chainID,
		})
		tool.FailOnErr(t, err, "Execute Failed")
		fmt.Println("execute txid", txid)
	}
}

// 0x26725dd91ed0c76a4fb151c5dd3503f343f4889162ee23a8bbfef92662044656
