// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AdvancedAgentConfiguration struct {
	// The interval in seconds to check for agent updates
	AgentUpdateInterval *int64 `json:"agent_update_interval"`
	// Use the hashcat binary already installed on the client system
	UseNativeHashcat *bool `json:"use_native_hashcat"`
	// The device to use for hashcat, separated by commas
	BackendDevice *string `json:"backend_device"`
	// The OpenCL device types to use for hashcat, separated by commas
	OpenclDevices *string `json:"opencl_devices,omitempty"`
	// Causes hashcat to perform benchmark-all, rather than just benchmark
	EnableAdditionalHashTypes bool `json:"enable_additional_hash_types"`
}

func (o *AdvancedAgentConfiguration) GetAgentUpdateInterval() *int64 {
	if o == nil {
		return nil
	}
	return o.AgentUpdateInterval
}

func (o *AdvancedAgentConfiguration) GetUseNativeHashcat() *bool {
	if o == nil {
		return nil
	}
	return o.UseNativeHashcat
}

func (o *AdvancedAgentConfiguration) GetBackendDevice() *string {
	if o == nil {
		return nil
	}
	return o.BackendDevice
}

func (o *AdvancedAgentConfiguration) GetOpenclDevices() *string {
	if o == nil {
		return nil
	}
	return o.OpenclDevices
}

func (o *AdvancedAgentConfiguration) GetEnableAdditionalHashTypes() bool {
	if o == nil {
		return false
	}
	return o.EnableAdditionalHashTypes
}
