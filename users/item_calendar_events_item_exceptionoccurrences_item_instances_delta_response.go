package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaGetResponseable instead.
type ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponse struct {
    ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaGetResponse
}
// NewItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponse instantiates a new ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponse()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponse) {
    m := &ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponse{
        ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaGetResponse: *NewItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaGetResponseable instead.
type ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaResponseable interface {
    ItemCalendarEventsItemExceptionoccurrencesItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
