package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable instead.
type DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponse struct {
    DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponse
}
// NewDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponse instantiates a new DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponse and sets the default values.
func NewDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponse()(*DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponse) {
    m := &DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponse{
        DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponse: *NewDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable instead.
type DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponseable interface {
    DeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
