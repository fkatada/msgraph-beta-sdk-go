package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse struct {
    AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse
}
// NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse instantiates a new AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse()(*AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse) {
    m := &AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse{
        AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse: *NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable interface {
    AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
