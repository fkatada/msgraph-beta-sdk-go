package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoveWatermarkAction 
type RemoveWatermarkAction struct {
    InformationProtectionAction
}
// NewRemoveWatermarkAction instantiates a new removeWatermarkAction and sets the default values.
func NewRemoveWatermarkAction()(*RemoveWatermarkAction) {
    m := &RemoveWatermarkAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.removeWatermarkAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRemoveWatermarkActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoveWatermarkActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoveWatermarkAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoveWatermarkAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
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
    res["uiElementNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetUiElementNames(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RemoveWatermarkAction) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUiElementNames gets the uiElementNames property value. The name of the UI element of footer to be removed.
func (m *RemoveWatermarkAction) GetUiElementNames()([]string) {
    val, err := m.GetBackingStore().Get("uiElementNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RemoveWatermarkAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetUiElementNames() != nil {
        err = writer.WriteCollectionOfStringValues("uiElementNames", m.GetUiElementNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RemoveWatermarkAction) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUiElementNames sets the uiElementNames property value. The name of the UI element of footer to be removed.
func (m *RemoveWatermarkAction) SetUiElementNames(value []string)() {
    err := m.GetBackingStore().Set("uiElementNames", value)
    if err != nil {
        panic(err)
    }
}
// RemoveWatermarkActionable 
type RemoveWatermarkActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetUiElementNames()([]string)
    SetOdataType(value *string)()
    SetUiElementNames(value []string)()
}
