// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
)

type AcceptTaskRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *AcceptTaskRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type AcceptTaskResponse struct {
	HTTPMeta components.HTTPMetadata
}

func (o *AcceptTaskResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
