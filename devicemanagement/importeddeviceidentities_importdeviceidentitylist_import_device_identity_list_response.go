package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponseable instead.
type ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponse struct {
    ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponse
}
// NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponse instantiates a new ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponse and sets the default values.
func NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponse()(*ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponse) {
    m := &ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponse{
        ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponse: *NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponse(),
    }
    return m
}
// CreateImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponse(), nil
}
// Deprecated: This class is obsolete. Use ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponseable instead.
type ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponseable interface {
    ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
