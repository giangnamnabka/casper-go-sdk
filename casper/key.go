package casper

import (
	"github.com/giangnamnabka/casper-go-sdk/types/key"
	"github.com/giangnamnabka/casper-go-sdk/types/keypair"
)

type (
	Key                 = key.Key
	Hash                = key.Hash
	DeployHash          = key.Hash
	AccountHash         = key.AccountHash
	ContractHash        = key.ContractHash
	ContractPackageHash = key.ContractPackageHash
	TransferHash        = key.TransferHash
	Uref                = key.URef
	PublicKey           = keypair.PublicKey
	PublicKeyList       = keypair.PublicKeyList
	PrivateKey          = keypair.PrivateKey
)

var (
	NewAccountHash                    = key.NewAccountHash
	NewTransferHash                   = key.NewTransferHash
	NewHash                           = key.NewHash
	NewHashFromBytes                  = key.NewHashFromBytes
	NewContractHash                   = key.NewContract
	NewContractPackageHash            = key.NewContractPackage
	NewUref                           = key.NewURef
	NewKey                            = key.NewKey
	NewKeyFromBytes                   = key.NewKeyFromBytes
	NewPublicKey                      = keypair.NewPublicKey
	NewED25519PrivateKeyFromPEMFile   = keypair.NewPrivateKeyED25518
	NewSECP256k1PrivateKeyFromPEMFile = keypair.NewPrivateKeySECP256K1
)
