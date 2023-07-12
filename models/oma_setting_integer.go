package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingInteger oMA Settings Integer definition.
type OmaSettingInteger struct {
    OmaSetting
    // The OdataType property
    OdataType *string
}
// NewOmaSettingInteger instantiates a new omaSettingInteger and sets the default values.
func NewOmaSettingInteger()(*OmaSettingInteger) {
    m := &OmaSettingInteger{
        OmaSetting: *NewOmaSetting(),
    }
    odataTypeValue := "#microsoft.graph.omaSettingInteger"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOmaSettingIntegerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingIntegerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingInteger(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingInteger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OmaSetting.GetFieldDeserializers()
    res["isReadOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReadOnly(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetIsReadOnly gets the isReadOnly property value. By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
func (m *OmaSettingInteger) GetIsReadOnly()(*bool) {
    val, err := m.GetBackingStore().Get("isReadOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetValue gets the value property value. Value.
func (m *OmaSettingInteger) GetValue()(*int32) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OmaSettingInteger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OmaSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isReadOnly", m.GetIsReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsReadOnly sets the isReadOnly property value. By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
func (m *OmaSettingInteger) SetIsReadOnly(value *bool)() {
    err := m.GetBackingStore().Set("isReadOnly", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. Value.
func (m *OmaSettingInteger) SetValue(value *int32)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// OmaSettingIntegerable 
type OmaSettingIntegerable interface {
    OmaSettingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsReadOnly()(*bool)
    GetValue()(*int32)
    SetIsReadOnly(value *bool)()
    SetValue(value *int32)()
}
