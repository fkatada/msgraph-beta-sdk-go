package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemAllMembersAddPostResponseable instead.
type DeletedTeamsItemChannelsItemAllMembersAddResponse struct {
    DeletedTeamsItemChannelsItemAllMembersAddPostResponse
}
// NewDeletedTeamsItemChannelsItemAllMembersAddResponse instantiates a new DeletedTeamsItemChannelsItemAllMembersAddResponse and sets the default values.
func NewDeletedTeamsItemChannelsItemAllMembersAddResponse()(*DeletedTeamsItemChannelsItemAllMembersAddResponse) {
    m := &DeletedTeamsItemChannelsItemAllMembersAddResponse{
        DeletedTeamsItemChannelsItemAllMembersAddPostResponse: *NewDeletedTeamsItemChannelsItemAllMembersAddPostResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsItemAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsItemAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsItemAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemAllMembersAddPostResponseable instead.
type DeletedTeamsItemChannelsItemAllMembersAddResponseable interface {
    DeletedTeamsItemChannelsItemAllMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
