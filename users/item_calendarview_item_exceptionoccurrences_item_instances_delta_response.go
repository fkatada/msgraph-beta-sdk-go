package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaGetResponseable instead.
type ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponse struct {
    ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaGetResponse
}
// NewItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponse instantiates a new ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponse()(*ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponse) {
    m := &ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponse{
        ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaGetResponse: *NewItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaGetResponseable instead.
type ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaResponseable interface {
    ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
