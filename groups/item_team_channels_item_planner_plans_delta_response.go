package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsItemPlannerPlansDeltaGetResponseable instead.
type ItemTeamChannelsItemPlannerPlansDeltaResponse struct {
    ItemTeamChannelsItemPlannerPlansDeltaGetResponse
}
// NewItemTeamChannelsItemPlannerPlansDeltaResponse instantiates a new ItemTeamChannelsItemPlannerPlansDeltaResponse and sets the default values.
func NewItemTeamChannelsItemPlannerPlansDeltaResponse()(*ItemTeamChannelsItemPlannerPlansDeltaResponse) {
    m := &ItemTeamChannelsItemPlannerPlansDeltaResponse{
        ItemTeamChannelsItemPlannerPlansDeltaGetResponse: *NewItemTeamChannelsItemPlannerPlansDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemPlannerPlansDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsItemPlannerPlansDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemPlannerPlansDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsItemPlannerPlansDeltaGetResponseable instead.
type ItemTeamChannelsItemPlannerPlansDeltaResponseable interface {
    ItemTeamChannelsItemPlannerPlansDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
