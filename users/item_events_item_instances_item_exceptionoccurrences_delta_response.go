package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemEventsItemInstancesItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemEventsItemInstancesItemExceptionoccurrencesDeltaResponse struct {
    ItemEventsItemInstancesItemExceptionoccurrencesDeltaGetResponse
}
// NewItemEventsItemInstancesItemExceptionoccurrencesDeltaResponse instantiates a new ItemEventsItemInstancesItemExceptionoccurrencesDeltaResponse and sets the default values.
func NewItemEventsItemInstancesItemExceptionoccurrencesDeltaResponse()(*ItemEventsItemInstancesItemExceptionoccurrencesDeltaResponse) {
    m := &ItemEventsItemInstancesItemExceptionoccurrencesDeltaResponse{
        ItemEventsItemInstancesItemExceptionoccurrencesDeltaGetResponse: *NewItemEventsItemInstancesItemExceptionoccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemEventsItemInstancesItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemEventsItemInstancesItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemEventsItemInstancesItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemEventsItemInstancesItemExceptionoccurrencesDeltaResponseable interface {
    ItemEventsItemInstancesItemExceptionoccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
