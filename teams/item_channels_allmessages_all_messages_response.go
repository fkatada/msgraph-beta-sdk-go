package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsAllmessagesAllMessagesGetResponseable instead.
type ItemChannelsAllmessagesAllMessagesResponse struct {
    ItemChannelsAllmessagesAllMessagesGetResponse
}
// NewItemChannelsAllmessagesAllMessagesResponse instantiates a new ItemChannelsAllmessagesAllMessagesResponse and sets the default values.
func NewItemChannelsAllmessagesAllMessagesResponse()(*ItemChannelsAllmessagesAllMessagesResponse) {
    m := &ItemChannelsAllmessagesAllMessagesResponse{
        ItemChannelsAllmessagesAllMessagesGetResponse: *NewItemChannelsAllmessagesAllMessagesGetResponse(),
    }
    return m
}
// CreateItemChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsAllmessagesAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsAllmessagesAllMessagesGetResponseable instead.
type ItemChannelsAllmessagesAllMessagesResponseable interface {
    ItemChannelsAllmessagesAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
