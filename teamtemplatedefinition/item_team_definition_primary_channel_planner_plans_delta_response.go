package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaGetResponseable instead.
type ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponse struct {
    ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaGetResponse
}
// NewItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponse instantiates a new ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponse()(*ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponse) {
    m := &ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponse{
        ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaGetResponse: *NewItemTeamDefinitionPrimaryChannelPlannerPlansDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaGetResponseable instead.
type ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaResponseable interface {
    ItemTeamDefinitionPrimaryChannelPlannerPlansDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
