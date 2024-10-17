package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable instead.
type ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponse struct {
    ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse
}
// NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponse instantiates a new ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponse and sets the default values.
func NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponse()(*ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponse) {
    m := &ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponse{
        ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse: *NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse(),
    }
    return m
}
// CreateManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponse(), nil
}
// Deprecated: This class is obsolete. Use ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable instead.
type ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseable interface {
    ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
