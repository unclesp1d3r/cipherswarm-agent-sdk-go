// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
)

type SubmitCrackRequest struct {
	// id
	ID            int64                     `pathParam:"style=simple,explode=false,name=id"`
	HashcatResult *components.HashcatResult `request:"mediaType=application/json"`
}

func (o *SubmitCrackRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *SubmitCrackRequest) GetHashcatResult() *components.HashcatResult {
	if o == nil {
		return nil
	}
	return o.HashcatResult
}

type SubmitCrackResponse struct {
	HTTPMeta components.HTTPMetadata
}

func (o *SubmitCrackResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
