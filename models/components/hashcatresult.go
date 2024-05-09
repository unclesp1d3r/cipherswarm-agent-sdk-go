// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/internal/utils"
	"time"
)

type HashcatResult struct {
	Timestamp time.Time `json:"timestamp"`
	Hash      string    `json:"hash"`
	PlainText string    `json:"plain_text"`
}

func (h HashcatResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HashcatResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HashcatResult) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *HashcatResult) GetHash() string {
	if o == nil {
		return ""
	}
	return o.Hash
}

func (o *HashcatResult) GetPlainText() string {
	if o == nil {
		return ""
	}
	return o.PlainText
}
