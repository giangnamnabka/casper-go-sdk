package types

import "github.com/giangnamnabka/casper-go-sdk/types/key"

// EraSummary is a summary of an era
type EraSummary struct {
	// The block hash
	BlockHash key.Hash `json:"block_hash"`
	// The Era Id
	EraID uint32 `json:"era_id"`
	// The StoredValue containing era information.
	StoredValue StoredValue `json:"stored_value"`
	// Hex-encoded hash of the state root.
	StateRootHash key.Hash `json:"state_root_hash"`
	// The merkle proof.
	MerkleProof string `json:"merkle_proof"`
}
