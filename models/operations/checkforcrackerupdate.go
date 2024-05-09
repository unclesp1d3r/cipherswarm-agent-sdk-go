// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

type CheckForCrackerUpdateRequest struct {
	// operating_system
	OperatingSystem *string `queryParam:"style=form,explode=true,name=operating_system"`
	// version
	Version *string `queryParam:"style=form,explode=true,name=version"`
}

func (o *CheckForCrackerUpdateRequest) GetOperatingSystem() *string {
	if o == nil {
		return nil
	}
	return o.OperatingSystem
}

func (o *CheckForCrackerUpdateRequest) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
