package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse struct {
    AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse
}
// NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse instantiates a new AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse()(*AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse) {
    m := &AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse{
        AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse: *NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable interface {
    AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
