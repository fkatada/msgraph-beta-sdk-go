package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody instantiates a new DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody and sets the default values.
func NewDeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody()(*DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) {
    m := &DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["importedAppleDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedAppleDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable)
                }
            }
            m.SetImportedAppleDeviceIdentities(res)
        }
        return nil
    }
    res["overwriteImportedDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteImportedDeviceIdentities(val)
        }
        return nil
    }
    return res
}
// GetImportedAppleDeviceIdentities gets the importedAppleDeviceIdentities property value. The importedAppleDeviceIdentities property
// returns a []ImportedAppleDeviceIdentityable when successful
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) GetImportedAppleDeviceIdentities()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable) {
    val, err := m.GetBackingStore().Get("importedAppleDeviceIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable)
    }
    return nil
}
// GetOverwriteImportedDeviceIdentities gets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
// returns a *bool when successful
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) GetOverwriteImportedDeviceIdentities()(*bool) {
    val, err := m.GetBackingStore().Get("overwriteImportedDeviceIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetImportedAppleDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedAppleDeviceIdentities()))
        for i, v := range m.GetImportedAppleDeviceIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("importedAppleDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteImportedDeviceIdentities", m.GetOverwriteImportedDeviceIdentities())
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
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetImportedAppleDeviceIdentities sets the importedAppleDeviceIdentities property value. The importedAppleDeviceIdentities property
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) SetImportedAppleDeviceIdentities(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable)() {
    err := m.GetBackingStore().Set("importedAppleDeviceIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetOverwriteImportedDeviceIdentities sets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBody) SetOverwriteImportedDeviceIdentities(value *bool)() {
    err := m.GetBackingStore().Set("overwriteImportedDeviceIdentities", value)
    if err != nil {
        panic(err)
    }
}
type DeponboardingsettingsItemImportedappledeviceidentitiesImportappledeviceidentitylistImportAppleDeviceIdentityListPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetImportedAppleDeviceIdentities()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable)
    GetOverwriteImportedDeviceIdentities()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetImportedAppleDeviceIdentities(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable)()
    SetOverwriteImportedDeviceIdentities(value *bool)()
}
