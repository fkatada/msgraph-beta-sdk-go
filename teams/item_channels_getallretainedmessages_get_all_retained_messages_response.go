package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsGetallretainedmessagesGetAllRetainedMessagesGetResponseable instead.
type ItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponse struct {
    ItemChannelsGetallretainedmessagesGetAllRetainedMessagesGetResponse
}
// NewItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponse instantiates a new ItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponse and sets the default values.
func NewItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponse()(*ItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponse) {
    m := &ItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponse{
        ItemChannelsGetallretainedmessagesGetAllRetainedMessagesGetResponse: *NewItemChannelsGetallretainedmessagesGetAllRetainedMessagesGetResponse(),
    }
    return m
}
// CreateItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsGetallretainedmessagesGetAllRetainedMessagesGetResponseable instead.
type ItemChannelsGetallretainedmessagesGetAllRetainedMessagesResponseable interface {
    ItemChannelsGetallretainedmessagesGetAllRetainedMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
