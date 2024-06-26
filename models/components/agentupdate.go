// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AgentUpdate struct {
	// The id of the agent
	ID int64 `json:"id"`
	// The hostname of the agent
	Name string `json:"name"`
	// The signature of the client
	ClientSignature string `json:"client_signature"`
	// The operating system of the agent
	OperatingSystem string   `json:"operating_system"`
	Devices         []string `json:"devices"`
}

func (o *AgentUpdate) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *AgentUpdate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AgentUpdate) GetClientSignature() string {
	if o == nil {
		return ""
	}
	return o.ClientSignature
}

func (o *AgentUpdate) GetOperatingSystem() string {
	if o == nil {
		return ""
	}
	return o.OperatingSystem
}

func (o *AgentUpdate) GetDevices() []string {
	if o == nil {
		return []string{}
	}
	return o.Devices
}
