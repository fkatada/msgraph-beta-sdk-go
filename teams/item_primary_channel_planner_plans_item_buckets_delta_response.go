package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimaryChannelPlannerPlansItemBucketsDeltaGetResponseable instead.
type ItemPrimaryChannelPlannerPlansItemBucketsDeltaResponse struct {
    ItemPrimaryChannelPlannerPlansItemBucketsDeltaGetResponse
}
// NewItemPrimaryChannelPlannerPlansItemBucketsDeltaResponse instantiates a new ItemPrimaryChannelPlannerPlansItemBucketsDeltaResponse and sets the default values.
func NewItemPrimaryChannelPlannerPlansItemBucketsDeltaResponse()(*ItemPrimaryChannelPlannerPlansItemBucketsDeltaResponse) {
    m := &ItemPrimaryChannelPlannerPlansItemBucketsDeltaResponse{
        ItemPrimaryChannelPlannerPlansItemBucketsDeltaGetResponse: *NewItemPrimaryChannelPlannerPlansItemBucketsDeltaGetResponse(),
    }
    return m
}
// CreateItemPrimaryChannelPlannerPlansItemBucketsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimaryChannelPlannerPlansItemBucketsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelPlannerPlansItemBucketsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimaryChannelPlannerPlansItemBucketsDeltaGetResponseable instead.
type ItemPrimaryChannelPlannerPlansItemBucketsDeltaResponseable interface {
    ItemPrimaryChannelPlannerPlansItemBucketsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
