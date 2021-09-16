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
	erc20 "github.com/terry108/multisig/go_abi/merc20_abi"
	"github.com/terry108/multisig/go_helper/tool"
)

var (
	a0           *tool.AddrInfo // 部署合约
	a1, a2, a3   *tool.AddrInfo
	addrs        []*tool.AddrInfo
	rpcHost      = "http://localhost:7545"
	privkHexList = []string{
		"b3cc8192cda1532fa013dea49974a6dd722cc1321bab11e0aa860f04b436387d", //0x7e9f34471c71858359E5574f0eDeb03dCa9F5f43
		"0d507211e7dfbab3504134a8fc6e31b5fb44d420e302b936f798f52b15021d55", //0xDD9Cd206BC670d8d131E4b95a6Bd916Ad4338570
		"80bb7fbf5084610f4f9dabf5b3cfc27eee00e86424f20e65cdefb187a3225a58", //0x8d505f08E421d59794F462FF0Cc5b01787838CE0
		"f61955911f7cb7304a182885c2b18d8fc433f23b80ad68a77b1a3d38f94b2c78", //0xCdD69c899028A0de95F5518bA5D2a8FfF7cd7799
	}

	contractAddress = common.HexToAddress("0x5F27F120a70eF82E0c07f5Eb3a3dF5E52352D13e")
	isERC20         = true
	ERC20Address    = common.HexToAddress("0x89b7be2fa81101201ad09ef4be44C0C2cC33a61a")
)

func PrepareAccounts(t *testing.T) {
	for i := 0; i < 4; i++ {
		addr := tool.GenAddr("")
		fmt.Println(i, addr.Address)
		addrs = append(addrs, addr)
	}
	a0, a1, a2, a3 = addrs[0], addrs[1], addrs[2], addrs[3]
	rpcClient, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")
	tool.WaitSomething(t, time.Minute, func() error { _, e := rpcClient.NetworkID(context.Background()); return e })
	tool.PrepareFunds4address(t, rpcHost, a0.Address, 2)
	bal, err := rpcClient.BalanceAt(context.Background(), a0.ToAddress(), nil)
	tool.FailOnErr(t, err, "get balance failed")
	fmt.Println("addr0 balance:", bal)
	tool.PrepareFunds4address(t, rpcHost, a1.Address, 1)
	bal, err = rpcClient.BalanceAt(context.Background(), a1.ToAddress(), nil)
	tool.FailOnErr(t, err, "get balance failed")
	fmt.Println("addr1 balance:", bal)
}

func TestETHDepositAndWithdraw(t *testing.T) {
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

	// 2.Cross chain deposit: send eth to multsig contract
	fmt.Println("2.Cross chain deposit: send 1 eth to multsig contract")
	multisigContract, err := dms.NewDynamicMultiSig(contractAddress, rpcClient)
	tool.FailOnErr(t, err, "NewDynamicMultiSig failed")
	nonce, err := rpcClient.NonceAt(ctx, a0.ToAddress(), nil)
	tool.FailOnErr(t, err, "Failed to get nonce")
	signer := types.LatestSignerForChainID(chainID)
	amt := big.NewInt(1).Mul(big.NewInt(tool.E18), common.Big1)
	opts := &bind.TransactOpts{
		From:  a0.ToAddress(),
		Nonce: big.NewInt(int64(nonce)),
		Value: amt,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a0.ToECDSAKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	erc20Addr := common.Address{}
	tx, err := multisigContract.CrossOut(opts, "test_address", amt, erc20Addr)
	tool.FailOnErr(t, err, "CrossOut eth failed")
	bal, err := rpcClient.BalanceAt(ctx, contractAddress, nil)
	tool.FailOnErr(t, err, "Get multisig contract balance failed")
	fmt.Println("CrossOut eth tx: ", tx.Hash().Hex())
	tool.FailOnFlag(t, bal.Cmp(amt) != 0, fmt.Sprintf("Balance not equal，should be %v, get %s", amt, bal.String()))
	fmt.Println("CrossOut eth success, multisig contract balance:", bal)

	// 3.Cross chain withdraw: get eth from multsig contract
	// a3 withdraw eth: signed by a2 and a3, a1 send withdraw tx
	fmt.Println("3.Cross chain withdraw: get 1 eth from multsig contract")
	var signs []byte // multisign
	isERC20 := false
	privkHex := a1.PrivkHex
	fromAddress := a2.ToAddress()
	destination := a3.ToAddress()
	randomTx := types.NewTx(&types.AccessListTx{Nonce: uint64(time.Now().Unix())})
	randomTxKey := randomTx.Hash().Hex()
	// multi sign
	for _, add := range []*tool.AddrInfo{a2, a3} {
		s, err := DynamicMultiSignTx(randomTxKey[2:], add.PrivkHex, isERC20, ERC20Address, destination, amt)
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
		ERC20Addr:               ERC20Address,
		FromAddress:             fromAddress,
		Destination:             destination,
		Value:                   amt,
		ChainID:                 chainID,
	}
	txid, err := ExecuteDynamicTX(params)
	tool.FailOnErr(t, err, "Execute Failed")
	fmt.Println("execute txid", txid)
	bal, err = rpcClient.BalanceAt(ctx, contractAddress, nil)
	tool.FailOnErr(t, err, "Get multisig contract balance failed")
	tool.FailOnFlag(t, bal.Cmp(big.NewInt(0)) != 0, fmt.Sprintf("Balance not equal，should be 0, get %s", bal.String()))
	fmt.Println("CrossOut withdraw eth success, multisig contract balance:", bal)
	bal, _ = rpcClient.BalanceAt(ctx, a3.ToAddress(), nil)
	fmt.Println("CrossOut withdraw eth success, a3 balance:", bal)
}

func TestMinerBurnERC20(t *testing.T) {
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

	// 2.Deloy ERC20 Token: MyToken
	fmt.Println("2.Deploy ERC20 contract: MyToken")
	auth, err := bind.NewKeyedTransactorWithChainID(a0.ToECDSAKey(), chainID)
	tool.FailOnErr(t, err, "NewKeyedTransactorWithChainID failed")
	ERC20Address, _, _, err = erc20.DeployERC20Minter(auth, rpcClient, "ETHToken", "ETHT", contractAddress)
	tool.FailOnErr(t, err, "deploy erc20 failed")
	fmt.Println("DeployERC20Minter address:", ERC20Address.Hex())

	// 3.regist MyToken to multisig contract
	fmt.Println("3.Regist MyToken to multisig contract")
	multisigContract, err := dms.NewDynamicMultiSig(contractAddress, rpcClient)
	tool.FailOnErr(t, err, "NewDynamicMultiSig failed")
	nonce, err := rpcClient.NonceAt(ctx, a0.ToAddress(), nil)
	tool.FailOnErr(t, err, "Failed to get nonce")
	signer := types.LatestSignerForChainID(chainID)
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
	tx, err := multisigContract.RegisterMinterERC20(opts, ERC20Address)
	tool.FailOnErr(t, err, "register minter ERC20 failed")
	fmt.Println("register minter ERC20 success: ", tx.Hash().Hex())

	// 4.Mint 1 Mytoke for a1
	fmt.Println("4.Mint 1 Mytoke for a1")
	var signs []byte // multisign
	isERC20 := true
	privkHex := a1.PrivkHex
	fromAddress := a2.ToAddress()
	destination := a1.ToAddress()
	randomTx := types.NewTx(&types.AccessListTx{Nonce: uint64(time.Now().Unix())})
	randomTxKey := randomTx.Hash().Hex()
	amt := big.NewInt(1).Mul(big.NewInt(tool.E18), common.Big1)
	// multi sign
	for _, add := range []*tool.AddrInfo{a2, a3} {
		s, err := DynamicMultiSignTx(randomTxKey[2:], add.PrivkHex, isERC20, ERC20Address, destination, amt)
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
		ERC20Addr:               ERC20Address,
		FromAddress:             fromAddress,
		Destination:             destination,
		Value:                   amt,
		ChainID:                 chainID,
	}
	txid, err := ExecuteDynamicTX(params)
	tool.FailOnErr(t, err, "Execute Failed")
	fmt.Println("execute txid", txid)
	erc20Minter, err := erc20.NewERC20Minter(ERC20Address, rpcClient)
	tool.FailOnErr(t, err, "NewERC20Minter failed")
	copts := &bind.CallOpts{Pending: true, Context: ctx}
	bal, err := erc20Minter.BalanceOf(copts, a1.ToAddress())
	tool.FailOnErr(t, err, "get erc20 balance failed")
	fmt.Println("erc20 balance: ", bal)

	// 5.cross chain transfer: burn
	fmt.Println("5.cross chain transfer: burn")
	nonce, err = rpcClient.NonceAt(ctx, a1.ToAddress(), nil)
	tool.FailOnErr(t, err, "Failed to get nonce")
	opts = &bind.TransactOpts{
		From:  a1.ToAddress(),
		Nonce: big.NewInt(int64(nonce)),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a1.ToECDSAKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	tx, err = erc20Minter.Approve(opts, contractAddress, amt)
	tool.FailOnErr(t, err, "approve failed")
	fmt.Println("approve tx: ", tx.Hash())
	nonce += 1
	opts = &bind.TransactOpts{
		From:  a1.ToAddress(),
		Nonce: big.NewInt(int64(nonce)),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a1.ToECDSAKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	tx, err = multisigContract.CrossOut(opts, "test_address", amt, ERC20Address)
	tool.FailOnErr(t, err, "CrossOut eth failed")
	bal, err = rpcClient.BalanceAt(ctx, contractAddress, nil)
	tool.FailOnErr(t, err, "Get multisig contract balance failed")
	fmt.Println("CrossOut eth tx: ", tx.Hash().Hex())
	fmt.Println("CrossOut eth success, multisig contract balance:", bal)
	copts = &bind.CallOpts{Pending: true, Context: ctx}
	totalSupply, err := erc20Minter.TotalSupply(copts)
	tool.FailOnErr(t, err, "Get total supply failed")
	fmt.Println("total supply:", totalSupply)
}
