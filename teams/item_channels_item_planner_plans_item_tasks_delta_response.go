package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsItemPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemChannelsItemPlannerPlansItemTasksDeltaResponse struct {
    ItemChannelsItemPlannerPlansItemTasksDeltaGetResponse
}
// NewItemChannelsItemPlannerPlansItemTasksDeltaResponse instantiates a new ItemChannelsItemPlannerPlansItemTasksDeltaResponse and sets the default values.
func NewItemChannelsItemPlannerPlansItemTasksDeltaResponse()(*ItemChannelsItemPlannerPlansItemTasksDeltaResponse) {
    m := &ItemChannelsItemPlannerPlansItemTasksDeltaResponse{
        ItemChannelsItemPlannerPlansItemTasksDeltaGetResponse: *NewItemChannelsItemPlannerPlansItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateItemChannelsItemPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsItemPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemPlannerPlansItemTasksDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsItemPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemChannelsItemPlannerPlansItemTasksDeltaResponseable interface {
    ItemChannelsItemPlannerPlansItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
