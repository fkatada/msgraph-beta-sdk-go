package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskActiveDirectoryGroup the class used to identify an Azure Directory group for the kiosk configuration
type WindowsKioskActiveDirectoryGroup struct {
    WindowsKioskUser
    // The OdataType property
    OdataType *string
}
// NewWindowsKioskActiveDirectoryGroup instantiates a new windowsKioskActiveDirectoryGroup and sets the default values.
func NewWindowsKioskActiveDirectoryGroup()(*WindowsKioskActiveDirectoryGroup) {
    m := &WindowsKioskActiveDirectoryGroup{
        WindowsKioskUser: *NewWindowsKioskUser(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskActiveDirectoryGroup"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskActiveDirectoryGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskActiveDirectoryGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskActiveDirectoryGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskActiveDirectoryGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskUser.GetFieldDeserializers()
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    return res
}
// GetGroupName gets the groupName property value. The name of the AD group that will be locked to this kiosk configuration
func (m *WindowsKioskActiveDirectoryGroup) GetGroupName()(*string) {
    val, err := m.GetBackingStore().Get("groupName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsKioskActiveDirectoryGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskUser.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupName sets the groupName property value. The name of the AD group that will be locked to this kiosk configuration
func (m *WindowsKioskActiveDirectoryGroup) SetGroupName(value *string)() {
    err := m.GetBackingStore().Set("groupName", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskActiveDirectoryGroupable 
type WindowsKioskActiveDirectoryGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskUserable
    GetGroupName()(*string)
    SetGroupName(value *string)()
}
