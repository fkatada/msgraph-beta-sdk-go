package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeGetResponseable instead.
type ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponse struct {
    ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeGetResponse
}
// NewItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponse instantiates a new ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponse and sets the default values.
func NewItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponse()(*ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponse) {
    m := &ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponse{
        ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeGetResponse: *NewItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeGetResponse(),
    }
    return m
}
// CreateItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeGetResponseable instead.
type ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeResponseable interface {
    ItemOnlineMeetingsGetAllRecordingsmeetingOrganizerUserIdMeetingOrganizerUserIdWithStartDateTimeWithEndDateTimeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
