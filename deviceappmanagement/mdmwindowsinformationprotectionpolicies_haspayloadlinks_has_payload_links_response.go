package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponseable instead.
type MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponse struct {
    MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponse
}
// NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponse instantiates a new MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponse and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponse()(*MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponse) {
    m := &MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponse{
        MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponse: *NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponseable instead.
type MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponseable interface {
    MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
