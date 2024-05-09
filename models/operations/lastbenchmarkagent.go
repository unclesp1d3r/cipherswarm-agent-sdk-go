// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/models/components"
)

type LastBenchmarkAgentRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *LastBenchmarkAgentRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type LastBenchmarkAgentResponse struct {
	HTTPMeta components.HTTPMetadata
	// successful
	AgentLastBenchmark *components.AgentLastBenchmark
}

func (o *LastBenchmarkAgentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *LastBenchmarkAgentResponse) GetAgentLastBenchmark() *components.AgentLastBenchmark {
	if o == nil {
		return nil
	}
	return o.AgentLastBenchmark
}
