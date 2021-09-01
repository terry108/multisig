package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/terry108/multisig/eth/contracts"
)

// DeployDynamicMultiSigContract 创建部署一个多签合约，返回合约地址
// address: 签名的参与者
// return: address_hex,txid
func DeployDynamicMultiSigContract(chainID *big.Int, backend bind.ContractBackend, privkHex string, hexAddress []string) (common.Address, error) {
	var (
		err          error
		privk        *ecdsa.PrivateKey
		owners       []common.Address
		contractAddr common.Address
	)

	// 预处理: owners需要保持incr序，合约内的简单查重依赖于排序性
	for i := 0; i < len(hexAddress); i++ {
		hexAddress[i] = strings.ToLower(hexAddress[i])
	}
	sort.Slice(hexAddress, func(i, j int) bool {
		return hexAddress[i] < hexAddress[j]
	})

	// init vars
	privk, err = crypto.HexToECDSA(privkHex)
	if err != nil {
		return contractAddr, fmt.Errorf("转换私钥时发生错误,%v", err)
	}
	for _, ha := range hexAddress {
		owners = append(owners, common.HexToAddress(ha))
	}

	// 部署多签合约
	auth, err := bind.NewKeyedTransactorWithChainID(privk, chainID)
	if err != nil {
		return contractAddr, fmt.Errorf("TransactOpts error,%v", err)
	}
	//TODO set gas,gasPrice
	contractAddr, tx, _, err := contracts.DeployDynamicMultiSig(auth, backend, owners)
	if err != nil {
		return contractAddr, fmt.Errorf("_部署多签合约失败, %v", err)
	}
	_ = tx
	return contractAddr, nil
}

// DynamicMultiSigExecuteSign return sig byte
func DynamicMultiSigExecuteSign(txKey, signerPrivkHex string, erc20ContractAddr, destinationAddr string, value *big.Int) ([]byte, error) {
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
		common.HexToAddress(destinationAddr),
		value,
		false,
		common.HexToAddress(erc20ContractAddr),
		big.NewInt(2),
	)
	if err != nil {
		fmt.Println(err)
	}
	hashBytes := crypto.Keccak256(
		bytes,
	)

	privk, err := crypto.HexToECDSA(signerPrivkHex)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(hashBytes, privk)
	if err != nil {
		return nil, errors.Wrap(err, "crypto sign failed")
	}
	return sig, nil
}

// TxParams 交易参数
type TxDynamicParams struct {
	Backend                              bind.ContractBackend
	TxKey                                string
	Signs                                []byte //签名
	PrivkHex                             string
	MultisigContractAddress, FromAddress string //多签合约地址，发起地址
	ERC20ContractAddress                 string
	Destination, Executor                string //toAddress
	Value, GasLimit                      *big.Int
	ChainID                              *big.Int
}

// ExecuteTX .
func ExecuteDynamicTX(txp *TxDynamicParams) (string, error) {
	// init vars
	privk, err := crypto.HexToECDSA(txp.PrivkHex)
	if err != nil {
		return "", err
	}
	multisigContract, err := contracts.NewDynamicMultiSig(common.HexToAddress(txp.MultisigContractAddress), txp.Backend)
	if err != nil {
		return "", fmt.Errorf("构建多签合约调用时异常,检查合约地址和rpc server,%v", err)
	}

	{ // 调用合约方法
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		nonce, err := txp.Backend.PendingNonceAt(ctx, common.HexToAddress(txp.FromAddress))
		if err != nil {
			return "", fmt.Errorf("获取多签地址nonce时发生错误, %v", err)
		}
		dst := make([]byte, hex.EncodedLen(len(txp.Signs)))
		hex.Encode(dst, txp.Signs)
		signer := types.LatestSignerForChainID(txp.ChainID)
		// fmt.Printf("txp.TxKey: %s,txp.Destination: %s, txp.Value: %v, txp.MultisigContractAddress: %s\n",
		// 	txp.TxKey, txp.Destination, txp.Value, txp.MultisigContractAddress)
		// fmt.Printf("txp.Sig: %s\n", dst)
		tx, err := multisigContract.CreateOrSignWithdraw(&bind.TransactOpts{
			From:     common.HexToAddress(txp.FromAddress),
			Nonce:    big.NewInt(int64(nonce)),
			GasLimit: uint64(txp.GasLimit.Int64()),
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privk)
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, signature)
			},
			Context: context.Background(),
		}, //txKey string, to common.Address, amount *big.Int, isERC20 bool, ERC20 common.Address, signatures []byte)
			txp.TxKey,
			common.HexToAddress(txp.Destination),
			txp.Value,
			false,
			common.HexToAddress(txp.ERC20ContractAddress),
			txp.Signs)
		if err != nil {
			return "", fmt.Errorf("调用合约交易方法时发生错误, %v", err)
		}
		return tx.Hash().Hex(), nil
	}

}
