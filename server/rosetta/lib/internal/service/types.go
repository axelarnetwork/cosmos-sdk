package service

import (
	"encoding/json"

	crgerrs "github.com/cosmos/cosmos-sdk/server/rosetta/lib/errors"
)

type SubAccountMetaData struct {
	ValidatorAddress string `json:"validator_address,omitempty"`
}

func (c *SubAccountMetaData) FromMetadata(meta map[string]interface{}) error {
	b, err := json.Marshal(meta)
	if err != nil {
		return crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	err = json.Unmarshal(b, c)
	if err != nil {
		return crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	return nil
}
