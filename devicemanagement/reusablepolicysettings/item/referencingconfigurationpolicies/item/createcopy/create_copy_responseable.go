package createcopy

import (
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CreateCopyResponseable 
type CreateCopyResponseable interface {
    AdditionalDataHolder
    Parsable
    GetDeviceManagementConfigurationPolicy()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicyable)
    SetDeviceManagementConfigurationPolicy(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicyable)()
}
