package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"

	"github.com/giangnamnabka/casper-go-sdk/types/key"
)

// Transform is an enumeration of transformation types used in the execution of a `deploy`.
type TransformKey struct {
	Key       key.Key   `json:"key"`
	Transform Transform `json:"transform"`
}

type Transform json.RawMessage

type WriteTransfer struct {
	ID         *uint64          `json:"id"`
	To         *key.AccountHash `json:"to"`
	DeployHash key.Hash         `json:"deploy_hash"`
	From       key.AccountHash  `json:"from"`
	Amount     uint64           `json:"amount,string"`
	Source     key.URef         `json:"source"`
	Target     key.URef         `json:"target"`
	Gas        string           `json:"gas"`
}

// MarshalJSON returns m as the JSON encoding of m.
func (t Transform) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return t, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (t *Transform) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*t = append((*t)[0:0], data...)
	return nil
}

func (t *Transform) IsWriteTransfer() bool {
	return strings.Contains(string(*t), "WriteTransfer")
}

func (t *Transform) ParseAsWriteTransfer() (*WriteTransfer, error) {
	type RawWriteTransferTransform struct {
		WriteTransfer `json:"WriteTransfer"`
	}

	jsonRes := RawWriteTransferTransform{}
	if err := json.Unmarshal(*t, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes.WriteTransfer, nil
}

func (t *Transform) IsWriteUnbonding() bool {
	return strings.Contains(string(*t), "WriteUnbonding")
}

func (t *Transform) ParseAsWriteUnbonding() ([]UnbondingPurse, error) {
	type RawWriteUnbonding struct {
		WriteUnbonding []UnbondingPurse `json:"WriteUnbonding"`
	}

	jsonRes := RawWriteUnbonding{}
	if err := json.Unmarshal(*t, &jsonRes); err != nil {
		return nil, err
	}

	return jsonRes.WriteUnbonding, nil
}

func (t *Transform) IsWriteContract() bool {
	return bytes.Equal(*t, []byte("\"WriteContract\""))
}

func (t *Transform) IsWriteWithdraw() bool {
	return strings.Contains(string(*t), "WriteWithdraw")
}

func (t *Transform) IsWriteCLValue() bool {
	return bytes.Contains(*t, []byte("\"WriteCLValue\""))
}

func (t *Transform) IsWriteBid() bool {
	return strings.Contains(string(*t), "WriteBid")
}

func (t *Transform) ParseAsWriteWithdraws() ([]UnbondingPurse, error) {
	type RawWriteWithdrawals struct {
		Withdraws []UnbondingPurse `json:"WriteWithdraw"`
	}

	jsonRes := RawWriteWithdrawals{}
	if err := json.Unmarshal(*t, &jsonRes); err != nil {
		return nil, err
	}

	return jsonRes.Withdraws, nil
}

func (t *Transform) ParseAsWriteCLValue() (*Argument, error) {
	type RawWriteCLValue struct {
		WriteCLValue Argument `json:"WriteCLValue"`
	}

	jsonRes := RawWriteCLValue{}
	if err := json.Unmarshal(*t, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes.WriteCLValue, nil
}
