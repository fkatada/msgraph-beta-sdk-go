package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponse struct {
    ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaGetResponse
}
// NewItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponse instantiates a new ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponse and sets the default values.
func NewItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponse()(*ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponse) {
    m := &ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponse{
        ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaGetResponse: *NewItemTeamPrimaryChannelPlannerPlansItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaResponseable interface {
    ItemTeamPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
