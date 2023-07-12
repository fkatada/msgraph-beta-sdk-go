package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileContainedApp an abstract class that represents a contained app in a mobileApp acting as a package.
type MobileContainedApp struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewMobileContainedApp instantiates a new mobileContainedApp and sets the default values.
func NewMobileContainedApp()(*MobileContainedApp) {
    m := &MobileContainedApp{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileContainedAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileContainedAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.microsoftStoreForBusinessContainedApp":
                        return NewMicrosoftStoreForBusinessContainedApp(), nil
                    case "#microsoft.graph.windowsUniversalAppXContainedApp":
                        return NewWindowsUniversalAppXContainedApp(), nil
                }
            }
        }
    }
    return NewMobileContainedApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileContainedApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MobileContainedApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// MobileContainedAppable 
type MobileContainedAppable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
