package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemPlannerPlansDeltaGetResponseable instead.
type DeletedTeamsItemChannelsItemPlannerPlansDeltaResponse struct {
    DeletedTeamsItemChannelsItemPlannerPlansDeltaGetResponse
}
// NewDeletedTeamsItemChannelsItemPlannerPlansDeltaResponse instantiates a new DeletedTeamsItemChannelsItemPlannerPlansDeltaResponse and sets the default values.
func NewDeletedTeamsItemChannelsItemPlannerPlansDeltaResponse()(*DeletedTeamsItemChannelsItemPlannerPlansDeltaResponse) {
    m := &DeletedTeamsItemChannelsItemPlannerPlansDeltaResponse{
        DeletedTeamsItemChannelsItemPlannerPlansDeltaGetResponse: *NewDeletedTeamsItemChannelsItemPlannerPlansDeltaGetResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsItemPlannerPlansDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsItemPlannerPlansDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemPlannerPlansDeltaGetResponseable instead.
type DeletedTeamsItemChannelsItemPlannerPlansDeltaResponseable interface {
    DeletedTeamsItemChannelsItemPlannerPlansDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
