package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable instead.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponse struct {
    ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponse
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponse instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponse and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponse()(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponse) {
    m := &ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponse{
        ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponse: *NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponse(),
    }
    return m
}
// CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable instead.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseable interface {
    ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
