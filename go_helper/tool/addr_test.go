package tool

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestGenAddr(t *testing.T) {
	for i := 0; i < 4; i++ {
		addr := GenAddr("")
		fmt.Printf("%s,%s\n", addr.Address, addr.PrivkHex)
	}
}

func TestPrivHex(t *testing.T) {
	// pri 84b2022b46db46cbd507e983bb5ee2326e1f5dd85c425f680059d56b58dba605
	// pub 04334142b05da978df147c05b499a8c4f30eaf39c4a09a60309a7d188308cf7dee03d3d0fd51e350cbab3b704a6580b2219be69607ec2294e734f58683b7d0a054
	// addr 0x7e7366444b174656C148F80A8DBa9fAcc79be385
	privkHex := "0d507211e7dfbab3504134a8fc6e31b5fb44d420e302b936f798f52b15021d55"
	key, _ := crypto.HexToECDSA(privkHex)
	pubKHex := hexutil.Encode(crypto.FromECDSAPub(&key.PublicKey))
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println("pub", pubKHex)
	fmt.Println("add", address)
}
