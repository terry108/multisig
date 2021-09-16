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
	game721 "github.com/terry108/multisig/go_abi/game721_abi"
	erc721 "github.com/terry108/multisig/go_abi/merc721_abi"
	"github.com/terry108/multisig/go_helper/tool"
)

func TestERC721DepositAndWithdraw(t *testing.T) {
	// prepare 4 account
	ctx := context.Background()
	PrepareAccounts(t)
	rpcClient, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")

	// 1.deploy mutlti sign contract
	fmt.Println("1.Deploy multisig contract")
	privkHex := a0.PrivkHex
	hexAddress := []string{a1.Address, a2.Address, a3.Address}
	fmt.Println("owners:", hexAddress)
	chainID, _ := rpcClient.ChainID(context.Background())
	contractAddress, err := DeployDynamicMultiSigContract(chainID, rpcClient, privkHex, hexAddress)
	if err != nil {
		t.Errorf("DeployMultiSigWalletContract() error = %v", err)
		t.FailNow()
	}
	fmt.Println("MultisigContract:", contractAddress.Hex())

	// 2. deploy game ERC721 and mint for a0
	fmt.Println("2.Deploy Game721 contract and mint for a0")
	auth, err := bind.NewKeyedTransactorWithChainID(a0.ToECDSAKey(), chainID)
	tool.FailOnErr(t, err, "NewKeyedTransactorWithChainID failed")
	game721Address, _, game721Contract, err := game721.DeployGameItem(auth, rpcClient)
	tool.FailOnErr(t, err, "deploy game721 failed")
	fmt.Println("game721Address address:", game721Address.Hex())
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
	tx, err := game721Contract.AwardItem(opts, a0.ToAddress(), "testURI")
	tool.FailOnErr(t, err, "mint game ERC721 failed")
	fmt.Println("mint game ERC721 success: ", tx.Hash().Hex())
	tokenId := big.NewInt(1)
	copts := &bind.CallOpts{Pending: true, Context: ctx}
	owner, err := game721Contract.OwnerOf(copts, tokenId)
	tool.FailOnErr(t, err, "get OwnerOf game721 failed")
	fmt.Println("Owner of NFT:", owner.Hex())
	tool.FailOnFlag(t, owner.Hex() != a0.Address, "CrossOutNFT game721 failed")

	// 3. cross chain deposite: send gameNFT to multisig contract
	fmt.Println("3.Cross chain deposite: send gameNFT to multisig contract")
	multisigContract, err := dms.NewDynamicMultiSig(contractAddress, rpcClient)
	tool.FailOnErr(t, err, "NewDynamicMultiSig failed")
	// approve
	nonce, _ = rpcClient.NonceAt(ctx, a0.ToAddress(), nil)
	opts.Nonce = big.NewInt(int64(nonce))
	tx, err = game721Contract.Approve(opts, contractAddress, tokenId)
	tool.FailOnErr(t, err, "approve failed")
	fmt.Println("approve tx: ", tx.Hash())
	// cross out
	nonce, _ = rpcClient.NonceAt(ctx, a0.ToAddress(), nil)
	opts.Nonce = big.NewInt(int64(nonce))
	tx, err = multisigContract.CrossOutNFT(opts, "cross_chain_to", game721Address, tokenId)
	tool.FailOnErr(t, err, "CrossOutNFT game721 failed")
	fmt.Println("CrossOutNFT game721 success: ", tx.Hash().Hex())
	owner, err = game721Contract.OwnerOf(copts, tokenId)
	tool.FailOnErr(t, err, "get OwnerOf game721 failed")
	fmt.Println("Owner of NFT:", owner.Hex())
	tool.FailOnFlag(t, owner.Hex() != contractAddress.Hex(), "CrossOutNFT game721 failed")

	// 4.Cross chain withdraw: get gameNFT from multsig contract
	// a3 withdraw game ERC721: signed by a2 and a3, a1 send withdraw tx
	fmt.Println("4.Cross chain withdraw: get gameNFT from multsig contract")
	var signs []byte // multisign
	privkHex = a1.PrivkHex
	fromAddress := a1.ToAddress()
	destination := a3.ToAddress()
	randomTx := types.NewTx(&types.AccessListTx{Nonce: uint64(time.Now().Unix())})
	randomTxKey := randomTx.Hash().Hex()
	tokenURI := "test_uri"
	// multi sign
	for _, add := range []*tool.AddrInfo{a2, a3} {
		s, err := DynamicMultiSignNFTTx(randomTxKey[2:], add.PrivkHex, tokenURI, game721Address, destination, tokenId)
		tool.FailOnErr(t, err, "create sig failed")
		signs = append(signs, s...)
	}
	params := &TxDynamicParams{
		Backend:                 rpcClient,
		TxKey:                   randomTxKey[2:],
		Signs:                   signs,
		PrivkHex:                privkHex,
		MultisigContractAddress: contractAddress,
		ERC721Addr:              game721Address,
		FromAddress:             fromAddress,
		Destination:             destination,
		TokenId:                 tokenId,
		TokenURI:                tokenURI,
		ChainID:                 chainID,
	}
	txid, err := ExecuteDynamicNFT(params)
	tool.FailOnErr(t, err, "Execute Failed")
	fmt.Println("execute txid", txid)
	copts = &bind.CallOpts{Pending: true, Context: ctx}
	owner, err = game721Contract.OwnerOf(copts, tokenId)
	tool.FailOnErr(t, err, "Get game721 contract owner failed")
	tool.FailOnFlag(t, owner.String() != a3.Address, fmt.Sprintf("Check owner failed, should %s, got %s", a3.Address, owner.String()))
	fmt.Println("CrossOut ERC721 success, owner:", owner)
}

func TestMinerBurnERC721(t *testing.T) {
	// prepare 4 account
	ctx := context.Background()
	PrepareAccounts(t)
	rpcClient, err := ethclient.Dial(rpcHost)
	tool.FailOnErr(t, err, "dial failed")

	// 1.deploy mutlti sign contract
	fmt.Println("1.Deploy multisig contract")
	privkHex := a0.PrivkHex
	hexAddress := []string{a1.Address, a2.Address, a3.Address}
	fmt.Println("owners:", hexAddress)
	chainID, _ := rpcClient.ChainID(context.Background())
	contractAddress, err := DeployDynamicMultiSigContract(chainID, rpcClient, privkHex, hexAddress)
	if err != nil {
		t.Errorf("DeployMultiSigWalletContract() error = %v", err)
		t.FailNow()
	}
	fmt.Println("MultisigContract:", contractAddress.Hex())

	// 2.Deloy ERC721 Miner Token: MyToken
	fmt.Println("2.Deloy ERC721 Miner Token: MyToken")
	auth, err := bind.NewKeyedTransactorWithChainID(a0.ToECDSAKey(), chainID)
	tool.FailOnErr(t, err, "NewKeyedTransactorWithChainID failed")
	ERC721Address, _, ERC721Contract, err := erc721.DeployERC721Minter(auth, rpcClient, "NFTToken", "NT", contractAddress)
	tool.FailOnErr(t, err, "deploy ERC721 Miner failed")
	fmt.Println("DeployERC721Minter address:", ERC721Address.Hex())

	// 3.regist  ERC721 Miner Token to multisig contract
	fmt.Println("3.Regist ERC721 Miner Token to multisig contract")
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
	tx, err := multisigContract.RegisterMinterERC721(opts, ERC721Address)
	tool.FailOnErr(t, err, "register minter ERC721 failed")
	fmt.Println("register minter ERC721 success: ", tx.Hash().Hex())

	// 4.Mint 1 NFT for a1
	fmt.Println("4.Mint 1 NFT for a1")
	var signs []byte // multisign
	privkHex = a1.PrivkHex
	fromAddress := a1.ToAddress()
	destination := a1.ToAddress()
	randomTx := types.NewTx(&types.AccessListTx{Nonce: uint64(time.Now().Unix())})
	randomTxKey := randomTx.Hash().Hex()
	tokenURI := "test_token_URI"
	tokenId := common.Big1
	// multi sign
	for _, add := range []*tool.AddrInfo{a2, a3} {
		s, err := DynamicMultiSignNFTTx(randomTxKey[2:], add.PrivkHex, tokenURI, ERC721Address, destination, tokenId)
		tool.FailOnErr(t, err, "create sig failed")
		signs = append(signs, s...)
	}
	params := &TxDynamicParams{
		Backend:                 rpcClient,
		TxKey:                   randomTxKey[2:],
		Signs:                   signs,
		PrivkHex:                privkHex,
		MultisigContractAddress: contractAddress,
		ERC721Addr:              ERC721Address,
		FromAddress:             fromAddress,
		Destination:             destination,
		TokenId:                 tokenId,
		TokenURI:                tokenURI,
		ChainID:                 chainID,
	}
	txid, err := ExecuteDynamicNFT(params)
	tool.FailOnErr(t, err, "Execute Failed")
	fmt.Println("execute txid", txid)
	copts := &bind.CallOpts{Pending: true, Context: ctx}
	owner, err := ERC721Contract.OwnerOf(copts, tokenId)
	tool.FailOnErr(t, err, "Get erc721 tokenId owner failed")
	tool.FailOnFlag(t, owner.String() != a1.Address, fmt.Sprintf("Check tokenId owner failed, should %s, got %s", a1.Address, owner.String()))
	fmt.Println("CrossOut ERC721 success, owner:", owner)

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
	tx, err = ERC721Contract.Approve(opts, contractAddress, tokenId)
	tool.FailOnErr(t, err, "approve failed")
	fmt.Println("approve tx: ", tx.Hash())
	opts.Nonce = big.NewInt(int64(nonce + 1))
	tx, err = multisigContract.CrossOutNFT(opts, "test_address", ERC721Address, tokenId)
	tool.FailOnErr(t, err, "CrossOut NFT failed")
	fmt.Println("CrossOut NFT tx: ", tx.Hash().Hex())
	a1Bal, err := ERC721Contract.BalanceOf(copts, a1.ToAddress())
	tool.FailOnErr(t, err, "get NFT balance failed")
	fmt.Println("NFT balance of a1:", a1Bal.String())
	multsigBal, err := ERC721Contract.BalanceOf(copts, contractAddress)
	tool.FailOnErr(t, err, "get NFT balance failed")
	fmt.Println("NFT balance of a1:", multsigBal.String())
}
