package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EntitlementmanagementAccesspackagesSearchGetResponseable instead.
type EntitlementmanagementAccesspackagesSearchResponse struct {
    EntitlementmanagementAccesspackagesSearchGetResponse
}
// NewEntitlementmanagementAccesspackagesSearchResponse instantiates a new EntitlementmanagementAccesspackagesSearchResponse and sets the default values.
func NewEntitlementmanagementAccesspackagesSearchResponse()(*EntitlementmanagementAccesspackagesSearchResponse) {
    m := &EntitlementmanagementAccesspackagesSearchResponse{
        EntitlementmanagementAccesspackagesSearchGetResponse: *NewEntitlementmanagementAccesspackagesSearchGetResponse(),
    }
    return m
}
// CreateEntitlementmanagementAccesspackagesSearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementmanagementAccesspackagesSearchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementmanagementAccesspackagesSearchResponse(), nil
}
// Deprecated: This class is obsolete. Use EntitlementmanagementAccesspackagesSearchGetResponseable instead.
type EntitlementmanagementAccesspackagesSearchResponseable interface {
    EntitlementmanagementAccesspackagesSearchGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
