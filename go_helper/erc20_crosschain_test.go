package eth

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	dms "github.com/terry108/multisig/go_abi/dms_abi"
	game20 "github.com/terry108/multisig/go_abi/game20_abi"
	"github.com/terry108/multisig/go_helper/tool"
)

func TestERC20DepositAndWithdraw(t *testing.T) {
	// prepare 4 account
	ctx := context.Background()
	PrepareAccounts(t)
	rpcClient, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")

	// 1.Deploy multisig contract
	fmt.Println("1.Deploy multisig contract")
	managerAddrs := []string{a1.Address, a2.Address, a3.Address}
	chainID, _ := rpcClient.ChainID(context.Background())
	contractAddress, err := DeployDynamicMultiSigContract(chainID, rpcClient, a0.PrivkHex, managerAddrs)
	if err != nil {
		t.Errorf("DeployMultiSigWalletContract() error = %v", err)
		t.FailNow()
	}
	fmt.Println("DeployMultisigWalletContract address:", contractAddress.Hex())

	// 2.Deploy game ERC20 contract
	fmt.Println("2.Deploy game ERC20 contract")
	auth, _ := bind.NewKeyedTransactorWithChainID(a0.ToECDSAKey(), chainID)
	game20Addr, tx, game20Contract, err := game20.DeployERC20FixedSupply(auth, rpcClient)
	tool.FailOnErr(t, err, "deploy game ERC20 failed")
	fmt.Println("tx:", tx.Hash().String(), "game20Addr:", game20Addr)

	// 3.Cross chain deposit: send game ERC20 to multsig contract
	fmt.Println("3.Cross chain deposit: send 1 game ERC20 to multsig contract")
	multisigContract, err := dms.NewDynamicMultiSig(contractAddress, rpcClient)
	tool.FailOnErr(t, err, "NewDynamicMultiSig failed")
	nonce, err := rpcClient.NonceAt(ctx, a0.ToAddress(), nil)
	tool.FailOnErr(t, err, "Failed to get nonce")
	signer := types.LatestSignerForChainID(chainID)
	amt := big.NewInt(100)
	opts := &bind.TransactOpts{
		From:  a0.ToAddress(),
		Nonce: big.NewInt(int64(nonce)),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a0.ToECDSAKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	tx, err = game20Contract.Approve(opts, contractAddress, amt)
	tool.FailOnErr(t, err, "approve failed")
	fmt.Println("approve tx: ", tx.Hash())
	nonce += 1
	opts = &bind.TransactOpts{
		From:  a0.ToAddress(),
		Nonce: big.NewInt(int64(nonce)),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a0.ToECDSAKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	tx, err = multisigContract.CrossOut(opts, "cross_chain_address", amt, game20Addr)
	tool.FailOnErr(t, err, "CrossOut eth failed")
	copts := &bind.CallOpts{Pending: true, Context: ctx}
	bal, err := game20Contract.BalanceOf(copts, contractAddress)
	tool.FailOnErr(t, err, "Get multisig contract balance failed")
	fmt.Println("CrossOut ERC20 tx: ", tx.Hash().Hex())
	fmt.Println("CrossOut ERC20 success, multisig contract balance:", bal)

	// 4.Cross chain withdraw: get game ERC20 from multsig contract
	// a3 withdraw game ERC20: signed by a2 and a3, a1 send withdraw tx
	var signs []byte // multisign
	isERC20 := true
	privkHex := a1.PrivkHex
	fromAddress := a1.ToAddress()
	destination := a3.ToAddress()
	randomTx := types.NewTx(&types.AccessListTx{Nonce: uint64(time.Now().Unix())})
	randomTxKey := randomTx.Hash().Hex()
	// multi sign
	for _, add := range []*tool.AddrInfo{a2, a3} {
		s, err := DynamicMultiSignTx(randomTxKey[2:], add.PrivkHex, isERC20, game20Addr, destination, amt)
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
		ERC20Addr:               game20Addr,
		FromAddress:             fromAddress,
		Destination:             destination,
		Value:                   amt,
		ChainID:                 chainID,
	}
	txid, err := ExecuteDynamicTX(params)
	tool.FailOnErr(t, err, "Execute Failed")
	fmt.Println("execute txid", txid)
	copts = &bind.CallOpts{Pending: true, Context: ctx}
	bal, err = game20Contract.BalanceOf(copts, contractAddress)
	tool.FailOnErr(t, err, "Get multisig contract balance failed")
	fmt.Println("CrossOut ERC20 success, multisig contract balance:", bal)
	bal, _ = game20Contract.BalanceOf(copts, a3.ToAddress())
	fmt.Println("CrossOut ERC20 success, a3 balance:", bal)
}
