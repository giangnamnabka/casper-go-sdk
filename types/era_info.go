package types

import (
	"github.com/giangnamnabka/casper-go-sdk/types/clvalue"
	"github.com/giangnamnabka/casper-go-sdk/types/keypair"
)

// EraInfo stores an auction metadata. Intended to be recorded at each era.
type EraInfo struct {
	// List of rewards allocated to delegators and validators.
	SeigniorageAllocations []SeigniorageAllocation `json:"seigniorage_allocations"`
}

// SeigniorageAllocation sores information about a seigniorage allocation
type SeigniorageAllocation struct {
	Validator *ValidatorAllocation `json:"Validator,omitempty"`
	Delegator *DelegatorAllocation `json:"Delegator,omitempty"`
}

type ValidatorAllocation struct {
	// Public key of the validator
	ValidatorPublicKey keypair.PublicKey `json:"validator_public_key"`
	// Amount allocated as a reward.
	Amount clvalue.UInt512 `json:"amount"`
}

type DelegatorAllocation struct {
	// Public key of the delegator
	DelegatorPublicKey keypair.PublicKey `json:"delegator_public_key"`
	// Public key of the validator
	ValidatorPublicKey keypair.PublicKey `json:"validator_public_key"`
	// Amount allocated as a reward.
	Amount clvalue.UInt512 `json:"amount"`
}
