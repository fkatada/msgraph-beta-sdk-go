package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemPrimaryChannelPlannerPlansItemTasksDeltaResponse struct {
    ItemPrimaryChannelPlannerPlansItemTasksDeltaGetResponse
}
// NewItemPrimaryChannelPlannerPlansItemTasksDeltaResponse instantiates a new ItemPrimaryChannelPlannerPlansItemTasksDeltaResponse and sets the default values.
func NewItemPrimaryChannelPlannerPlansItemTasksDeltaResponse()(*ItemPrimaryChannelPlannerPlansItemTasksDeltaResponse) {
    m := &ItemPrimaryChannelPlannerPlansItemTasksDeltaResponse{
        ItemPrimaryChannelPlannerPlansItemTasksDeltaGetResponse: *NewItemPrimaryChannelPlannerPlansItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateItemPrimaryChannelPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimaryChannelPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelPlannerPlansItemTasksDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemPrimaryChannelPlannerPlansItemTasksDeltaResponseable interface {
    ItemPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
