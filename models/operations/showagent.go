// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
)

type ShowAgentRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ShowAgentRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type ShowAgentResponse struct {
	HTTPMeta components.HTTPMetadata
	// successful
	Agent *components.Agent
}

func (o *ShowAgentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ShowAgentResponse) GetAgent() *components.Agent {
	if o == nil {
		return nil
	}
	return o.Agent
}
