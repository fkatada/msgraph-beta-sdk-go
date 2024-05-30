package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody instantiates a new ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody and sets the default values.
func NewItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody()(*ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) {
    m := &ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["updateWindowsDeviceAccountActionParameter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUpdateWindowsDeviceAccountActionParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindowsDeviceAccountActionParameter(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable))
        }
        return nil
    }
    return res
}
// GetUpdateWindowsDeviceAccountActionParameter gets the updateWindowsDeviceAccountActionParameter property value. The updateWindowsDeviceAccountActionParameter property
// returns a UpdateWindowsDeviceAccountActionParameterable when successful
func (m *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) GetUpdateWindowsDeviceAccountActionParameter()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable) {
    val, err := m.GetBackingStore().Get("updateWindowsDeviceAccountActionParameter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("updateWindowsDeviceAccountActionParameter", m.GetUpdateWindowsDeviceAccountActionParameter())
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
func (m *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetUpdateWindowsDeviceAccountActionParameter sets the updateWindowsDeviceAccountActionParameter property value. The updateWindowsDeviceAccountActionParameter property
func (m *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBody) SetUpdateWindowsDeviceAccountActionParameter(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable)() {
    err := m.GetBackingStore().Set("updateWindowsDeviceAccountActionParameter", value)
    if err != nil {
        panic(err)
    }
}
type ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetUpdateWindowsDeviceAccountActionParameter()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetUpdateWindowsDeviceAccountActionParameter(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateWindowsDeviceAccountActionParameterable)()
}
