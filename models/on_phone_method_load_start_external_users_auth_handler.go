package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OnPhoneMethodLoadStartExternalUsersAuthHandler struct {
    OnPhoneMethodLoadStartHandler
}
// NewOnPhoneMethodLoadStartExternalUsersAuthHandler instantiates a new OnPhoneMethodLoadStartExternalUsersAuthHandler and sets the default values.
func NewOnPhoneMethodLoadStartExternalUsersAuthHandler()(*OnPhoneMethodLoadStartExternalUsersAuthHandler) {
    m := &OnPhoneMethodLoadStartExternalUsersAuthHandler{
        OnPhoneMethodLoadStartHandler: *NewOnPhoneMethodLoadStartHandler(),
    }
    odataTypeValue := "#microsoft.graph.onPhoneMethodLoadStartExternalUsersAuthHandler"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnPhoneMethodLoadStartExternalUsersAuthHandlerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnPhoneMethodLoadStartExternalUsersAuthHandlerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPhoneMethodLoadStartExternalUsersAuthHandler(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnPhoneMethodLoadStartExternalUsersAuthHandler) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnPhoneMethodLoadStartHandler.GetFieldDeserializers()
    res["smsOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhoneOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsOptions(val.(PhoneOptionsable))
        }
        return nil
    }
    res["voiceOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhoneOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVoiceOptions(val.(PhoneOptionsable))
        }
        return nil
    }
    return res
}
// GetSmsOptions gets the smsOptions property value. The smsOptions property
// returns a PhoneOptionsable when successful
func (m *OnPhoneMethodLoadStartExternalUsersAuthHandler) GetSmsOptions()(PhoneOptionsable) {
    val, err := m.GetBackingStore().Get("smsOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PhoneOptionsable)
    }
    return nil
}
// GetVoiceOptions gets the voiceOptions property value. The voiceOptions property
// returns a PhoneOptionsable when successful
func (m *OnPhoneMethodLoadStartExternalUsersAuthHandler) GetVoiceOptions()(PhoneOptionsable) {
    val, err := m.GetBackingStore().Get("voiceOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PhoneOptionsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnPhoneMethodLoadStartExternalUsersAuthHandler) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnPhoneMethodLoadStartHandler.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("smsOptions", m.GetSmsOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("voiceOptions", m.GetVoiceOptions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSmsOptions sets the smsOptions property value. The smsOptions property
func (m *OnPhoneMethodLoadStartExternalUsersAuthHandler) SetSmsOptions(value PhoneOptionsable)() {
    err := m.GetBackingStore().Set("smsOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetVoiceOptions sets the voiceOptions property value. The voiceOptions property
func (m *OnPhoneMethodLoadStartExternalUsersAuthHandler) SetVoiceOptions(value PhoneOptionsable)() {
    err := m.GetBackingStore().Set("voiceOptions", value)
    if err != nil {
        panic(err)
    }
}
type OnPhoneMethodLoadStartExternalUsersAuthHandlerable interface {
    OnPhoneMethodLoadStartHandlerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSmsOptions()(PhoneOptionsable)
    GetVoiceOptions()(PhoneOptionsable)
    SetSmsOptions(value PhoneOptionsable)()
    SetVoiceOptions(value PhoneOptionsable)()
}
