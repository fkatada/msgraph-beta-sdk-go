package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable instead.
type EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponse struct {
    EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse
}
// NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponse instantiates a new EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponse and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponse()(*EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponse) {
    m := &EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponse{
        EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse: *NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse(),
    }
    return m
}
// CreateEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponse(), nil
}
// Deprecated: This class is obsolete. Use EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable instead.
type EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable interface {
    EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
