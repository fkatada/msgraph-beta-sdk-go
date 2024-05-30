package app

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody instantiates a new OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody and sets the default values.
func NewOnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody()(*OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) {
    m := &OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAttendees gets the attendees property value. The attendees property
// returns a []AttendeeNotificationInfoable when successful
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetAttendees()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeNotificationInfoable) {
    val, err := m.GetBackingStore().Get("attendees")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeNotificationInfoable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendeeNotificationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeNotificationInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeNotificationInfoable)
                }
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["remindBeforeTimeInMinutesType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseRemindBeforeTimeInMinutesType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemindBeforeTimeInMinutesType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RemindBeforeTimeInMinutesType))
        }
        return nil
    }
    return res
}
// GetRemindBeforeTimeInMinutesType gets the remindBeforeTimeInMinutesType property value. The remindBeforeTimeInMinutesType property
// returns a *RemindBeforeTimeInMinutesType when successful
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetRemindBeforeTimeInMinutesType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RemindBeforeTimeInMinutesType) {
    val, err := m.GetBackingStore().Get("remindBeforeTimeInMinutesType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RemindBeforeTimeInMinutesType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemindBeforeTimeInMinutesType() != nil {
        cast := (*m.GetRemindBeforeTimeInMinutesType()).String()
        err := writer.WriteStringValue("remindBeforeTimeInMinutesType", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAttendees sets the attendees property value. The attendees property
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) SetAttendees(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeNotificationInfoable)() {
    err := m.GetBackingStore().Set("attendees", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetRemindBeforeTimeInMinutesType sets the remindBeforeTimeInMinutesType property value. The remindBeforeTimeInMinutesType property
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) SetRemindBeforeTimeInMinutesType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RemindBeforeTimeInMinutesType)() {
    err := m.GetBackingStore().Set("remindBeforeTimeInMinutesType", value)
    if err != nil {
        panic(err)
    }
}
type OnlinemeetingswithjoinweburlSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendees()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeNotificationInfoable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetRemindBeforeTimeInMinutesType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RemindBeforeTimeInMinutesType)
    SetAttendees(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeNotificationInfoable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetRemindBeforeTimeInMinutesType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RemindBeforeTimeInMinutesType)()
}
