package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsItemPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemTeamChannelsItemPlannerPlansItemTasksDeltaResponse struct {
    ItemTeamChannelsItemPlannerPlansItemTasksDeltaGetResponse
}
// NewItemTeamChannelsItemPlannerPlansItemTasksDeltaResponse instantiates a new ItemTeamChannelsItemPlannerPlansItemTasksDeltaResponse and sets the default values.
func NewItemTeamChannelsItemPlannerPlansItemTasksDeltaResponse()(*ItemTeamChannelsItemPlannerPlansItemTasksDeltaResponse) {
    m := &ItemTeamChannelsItemPlannerPlansItemTasksDeltaResponse{
        ItemTeamChannelsItemPlannerPlansItemTasksDeltaGetResponse: *NewItemTeamChannelsItemPlannerPlansItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsItemPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemPlannerPlansItemTasksDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsItemPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemTeamChannelsItemPlannerPlansItemTasksDeltaResponseable interface {
    ItemTeamChannelsItemPlannerPlansItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
