package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse struct {
    ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse
}
// NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse instantiates a new ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse and sets the default values.
func NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse) {
    m := &ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse{
        ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse: *NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable interface {
    ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
