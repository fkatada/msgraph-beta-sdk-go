package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemAllMembersRemovePostResponseable instead.
type ItemTeamDefinitionChannelsItemAllMembersRemoveResponse struct {
    ItemTeamDefinitionChannelsItemAllMembersRemovePostResponse
}
// NewItemTeamDefinitionChannelsItemAllMembersRemoveResponse instantiates a new ItemTeamDefinitionChannelsItemAllMembersRemoveResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemAllMembersRemoveResponse()(*ItemTeamDefinitionChannelsItemAllMembersRemoveResponse) {
    m := &ItemTeamDefinitionChannelsItemAllMembersRemoveResponse{
        ItemTeamDefinitionChannelsItemAllMembersRemovePostResponse: *NewItemTeamDefinitionChannelsItemAllMembersRemovePostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsItemAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemAllMembersRemovePostResponseable instead.
type ItemTeamDefinitionChannelsItemAllMembersRemoveResponseable interface {
    ItemTeamDefinitionChannelsItemAllMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
