package types

import (
	"github.com/giangnamnabka/casper-go-sdk/types/key"
	"github.com/giangnamnabka/casper-go-sdk/types/keypair"
)

// UnbondingPurse stores information of an unbonding or delegation withdrawal
type UnbondingPurse struct {
	// Unbonding Amount
	Amount uint64 `json:"amount,string"`
	// Bonding purse
	BondingPurse key.URef `json:"bonding_purse"`
	// Era ID in which this unbonding request was created.
	EraOfCreation uint32 `json:"era_of_creation"`
	// Unbonder public key.
	UnbonderPublicKey keypair.PublicKey `json:"unbonder_public_key"`
	// The original validator's public key.
	ValidatorPublicKey keypair.PublicKey `json:"validator_public_key"`
	// The re-delegated validator's public key.
	NewValidator keypair.PublicKey `json:"new_validator"`
}
