package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponseable instead.
type ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponse struct {
    ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponse
}
// NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponse instantiates a new ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponse and sets the default values.
func NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponse()(*ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponse) {
    m := &ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponse{
        ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponse: *NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponse(),
    }
    return m
}
// CreateResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponse(), nil
}
// Deprecated: This class is obsolete. Use ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponseable instead.
type ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponseable
}
