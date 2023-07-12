package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftStoreForBusinessAppAssignmentSettings contains properties used to assign an Microsoft Store for Business mobile app to a group.
type MicrosoftStoreForBusinessAppAssignmentSettings struct {
    MobileAppAssignmentSettings
}
// NewMicrosoftStoreForBusinessAppAssignmentSettings instantiates a new microsoftStoreForBusinessAppAssignmentSettings and sets the default values.
func NewMicrosoftStoreForBusinessAppAssignmentSettings()(*MicrosoftStoreForBusinessAppAssignmentSettings) {
    m := &MicrosoftStoreForBusinessAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.microsoftStoreForBusinessAppAssignmentSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftStoreForBusinessAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftStoreForBusinessAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftStoreForBusinessAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftStoreForBusinessAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["useDeviceContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseDeviceContext(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MicrosoftStoreForBusinessAppAssignmentSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUseDeviceContext gets the useDeviceContext property value. Whether or not to use device execution context for Microsoft Store for Business mobile app.
func (m *MicrosoftStoreForBusinessAppAssignmentSettings) GetUseDeviceContext()(*bool) {
    val, err := m.GetBackingStore().Get("useDeviceContext")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MicrosoftStoreForBusinessAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useDeviceContext", m.GetUseDeviceContext())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MicrosoftStoreForBusinessAppAssignmentSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUseDeviceContext sets the useDeviceContext property value. Whether or not to use device execution context for Microsoft Store for Business mobile app.
func (m *MicrosoftStoreForBusinessAppAssignmentSettings) SetUseDeviceContext(value *bool)() {
    err := m.GetBackingStore().Set("useDeviceContext", value)
    if err != nil {
        panic(err)
    }
}
// MicrosoftStoreForBusinessAppAssignmentSettingsable 
type MicrosoftStoreForBusinessAppAssignmentSettingsable interface {
    MobileAppAssignmentSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetUseDeviceContext()(*bool)
    SetOdataType(value *string)()
    SetUseDeviceContext(value *bool)()
}
