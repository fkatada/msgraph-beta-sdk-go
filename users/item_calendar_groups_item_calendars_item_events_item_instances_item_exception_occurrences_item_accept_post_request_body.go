package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody 
type ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The SendResponse property
    sendResponse *bool
}
// NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody instantiates a new ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) {
    m := &ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["sendResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendResponse(val)
        }
        return nil
    }
    return res
}
// GetSendResponse gets the sendResponse property value. The SendResponse property
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) GetSendResponse()(*bool) {
    return m.sendResponse
}
// Serialize serializes information the current object
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendResponse", m.GetSendResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetSendResponse sets the sendResponse property value. The SendResponse property
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBody) SetSendResponse(value *bool)() {
    m.sendResponse = value
}
