package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponseable instead.
type GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponse struct {
    GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponse
}
// NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponse instantiates a new GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponse and sets the default values.
func NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponse()(*GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponse) {
    m := &GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponse{
        GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponse: *NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponse(),
    }
    return m
}
// CreateGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponse(), nil
}
// Deprecated: This class is obsolete. Use GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponseable instead.
type GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceResponseable interface {
    GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
