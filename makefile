genSimpleMultisigGo:
	abigen -sol ./contracts/SimpleMultiSig.sol -pkg sms_abi --out ./go_abi/sms_abi/SimpleMultiSig.g.go
genDynamicMultisigGo:
	abigen -sol ./flatten/DynamicMultiSig_Fla.sol -pkg dms_abi --out ./go_abi/dms_abi/DynamicMultiSig.g.go
genERC20Go:
	abigen -sol ./flatten/ERC20Minter_Fla.sol -pkg merc20_abi --out ./go_abi/merc20_abi/ERC20Minter.g.go
genERC721Go:
	abigen -sol ./flatten/ERC721Minter_Fla.sol -pkg merc721_abi --out ./go_abi/merc721_abi/ERC721Minter.g.go