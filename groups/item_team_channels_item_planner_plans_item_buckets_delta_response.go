package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsItemPlannerPlansItemBucketsDeltaGetResponseable instead.
type ItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponse struct {
    ItemTeamChannelsItemPlannerPlansItemBucketsDeltaGetResponse
}
// NewItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponse instantiates a new ItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponse and sets the default values.
func NewItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponse()(*ItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponse) {
    m := &ItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponse{
        ItemTeamChannelsItemPlannerPlansItemBucketsDeltaGetResponse: *NewItemTeamChannelsItemPlannerPlansItemBucketsDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsItemPlannerPlansItemBucketsDeltaGetResponseable instead.
type ItemTeamChannelsItemPlannerPlansItemBucketsDeltaResponseable interface {
    ItemTeamChannelsItemPlannerPlansItemBucketsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
