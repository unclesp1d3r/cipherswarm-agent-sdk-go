// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/internal/utils"
	"net/http"
	"time"
)

// Metadata - Additional metadata about the error
type Metadata struct {
	// The date of the error
	ErrorDate time.Time `json:"error_date"`
	// Other metadata
	Other map[string]any `json:"other,omitempty"`
}

func (m Metadata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *Metadata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Metadata) GetErrorDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ErrorDate
}

func (o *Metadata) GetOther() map[string]any {
	if o == nil {
		return nil
	}
	return o.Other
}

// Severity - The severity of the error:
//   - `info` - Informational message, no action required.
//   - `warning` - Non-critical error, no action required. Anticipated, but not necessarily problematic.
//   - `minor` - Minor error, no action required. Should be investigated, but the task can continue.
//   - `major` - Major error, action required. The task should be investigated and possibly restarted.
//   - `critical` - Critical error, action required. The task should be stopped and investigated.
//   - `fatal` - Fatal error, action required. The agent cannot continue with the task and should not be reattempted.
type Severity string

const (
	SeverityInfo     Severity = "info"
	SeverityWarning  Severity = "warning"
	SeverityMinor    Severity = "minor"
	SeverityMajor    Severity = "major"
	SeverityCritical Severity = "critical"
	SeverityFatal    Severity = "fatal"
)

func (e Severity) ToPointer() *Severity {
	return &e
}
func (e *Severity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "info":
		fallthrough
	case "warning":
		fallthrough
	case "minor":
		fallthrough
	case "major":
		fallthrough
	case "critical":
		fallthrough
	case "fatal":
		*e = Severity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Severity: %v", v)
	}
}

type SubmitErrorAgentRequestBody struct {
	// The error message
	Message string `json:"message"`
	// Additional metadata about the error
	Metadata *Metadata `json:"metadata,omitempty"`
	// The severity of the error:
	//                        * `info` - Informational message, no action required.
	//                        * `warning` - Non-critical error, no action required. Anticipated, but not necessarily problematic.
	//                        * `minor` - Minor error, no action required. Should be investigated, but the task can continue.
	//                        * `major` - Major error, action required. The task should be investigated and possibly restarted.
	//                        * `critical` - Critical error, action required. The task should be stopped and investigated.
	//                         * `fatal` - Fatal error, action required. The agent cannot continue with the task and should not be reattempted.
	Severity Severity `json:"severity"`
	// The agent that caused the error
	AgentID int64 `json:"agent_id"`
	// The task that caused the error, if any
	TaskID *int64 `json:"task_id,omitempty"`
}

func (o *SubmitErrorAgentRequestBody) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *SubmitErrorAgentRequestBody) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *SubmitErrorAgentRequestBody) GetSeverity() Severity {
	if o == nil {
		return Severity("")
	}
	return o.Severity
}

func (o *SubmitErrorAgentRequestBody) GetAgentID() int64 {
	if o == nil {
		return 0
	}
	return o.AgentID
}

func (o *SubmitErrorAgentRequestBody) GetTaskID() *int64 {
	if o == nil {
		return nil
	}
	return o.TaskID
}

type SubmitErrorAgentRequest struct {
	// id
	ID          int64                        `pathParam:"style=simple,explode=false,name=id"`
	RequestBody *SubmitErrorAgentRequestBody `request:"mediaType=application/json"`
}

func (o *SubmitErrorAgentRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *SubmitErrorAgentRequest) GetRequestBody() *SubmitErrorAgentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type SubmitErrorAgentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SubmitErrorAgentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SubmitErrorAgentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SubmitErrorAgentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
