package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse struct {
    EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse
}
// NewEntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse instantiates a new EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse()(*EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse) {
    m := &EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse{
        EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse: *NewEntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable interface {
    EntitlementmanagementAccesspackageassignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
