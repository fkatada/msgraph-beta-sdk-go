package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaGetResponseable instead.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponse struct {
    TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaGetResponse
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponse instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponse and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponse()(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponse) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponse{
        TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaGetResponse: *NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaGetResponseable instead.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemPlannerPlansItemBucketsItemTasksDeltaGetResponseable
}
