package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComplexSettingDefinition entity representing the defintion for a complex setting
type DeviceManagementComplexSettingDefinition struct {
    DeviceManagementSettingDefinition
}
// NewDeviceManagementComplexSettingDefinition instantiates a new deviceManagementComplexSettingDefinition and sets the default values.
func NewDeviceManagementComplexSettingDefinition()(*DeviceManagementComplexSettingDefinition) {
    m := &DeviceManagementComplexSettingDefinition{
        DeviceManagementSettingDefinition: *NewDeviceManagementSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementComplexSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComplexSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComplexSettingDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplexSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingDefinition.GetFieldDeserializers()
    res["propertyDefinitionIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPropertyDefinitionIds(res)
        }
        return nil
    }
    return res
}
// GetPropertyDefinitionIds gets the propertyDefinitionIds property value. The definitions of each property of the complex setting
func (m *DeviceManagementComplexSettingDefinition) GetPropertyDefinitionIds()([]string) {
    val, err := m.GetBackingStore().Get("propertyDefinitionIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementComplexSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPropertyDefinitionIds() != nil {
        err = writer.WriteCollectionOfStringValues("propertyDefinitionIds", m.GetPropertyDefinitionIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPropertyDefinitionIds sets the propertyDefinitionIds property value. The definitions of each property of the complex setting
func (m *DeviceManagementComplexSettingDefinition) SetPropertyDefinitionIds(value []string)() {
    err := m.GetBackingStore().Set("propertyDefinitionIds", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementComplexSettingDefinitionable 
type DeviceManagementComplexSettingDefinitionable interface {
    DeviceManagementSettingDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPropertyDefinitionIds()([]string)
    SetPropertyDefinitionIds(value []string)()
}
