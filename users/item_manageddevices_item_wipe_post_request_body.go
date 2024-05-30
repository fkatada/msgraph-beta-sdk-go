package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemManageddevicesItemWipePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemManageddevicesItemWipePostRequestBody instantiates a new ItemManageddevicesItemWipePostRequestBody and sets the default values.
func NewItemManageddevicesItemWipePostRequestBody()(*ItemManageddevicesItemWipePostRequestBody) {
    m := &ItemManageddevicesItemWipePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManageddevicesItemWipePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemManageddevicesItemWipePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManageddevicesItemWipePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemManageddevicesItemWipePostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemManageddevicesItemWipePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemManageddevicesItemWipePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keepEnrollmentData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepEnrollmentData(val)
        }
        return nil
    }
    res["keepUserData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    res["macOsUnlockCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOsUnlockCode(val)
        }
        return nil
    }
    res["obliterationBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseObliterationBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObliterationBehavior(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ObliterationBehavior))
        }
        return nil
    }
    res["persistEsimDataPlan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistEsimDataPlan(val)
        }
        return nil
    }
    res["useProtectedWipe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseProtectedWipe(val)
        }
        return nil
    }
    return res
}
// GetKeepEnrollmentData gets the keepEnrollmentData property value. The keepEnrollmentData property
// returns a *bool when successful
func (m *ItemManageddevicesItemWipePostRequestBody) GetKeepEnrollmentData()(*bool) {
    val, err := m.GetBackingStore().Get("keepEnrollmentData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKeepUserData gets the keepUserData property value. The keepUserData property
// returns a *bool when successful
func (m *ItemManageddevicesItemWipePostRequestBody) GetKeepUserData()(*bool) {
    val, err := m.GetBackingStore().Get("keepUserData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMacOsUnlockCode gets the macOsUnlockCode property value. The macOsUnlockCode property
// returns a *string when successful
func (m *ItemManageddevicesItemWipePostRequestBody) GetMacOsUnlockCode()(*string) {
    val, err := m.GetBackingStore().Get("macOsUnlockCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetObliterationBehavior gets the obliterationBehavior property value. The obliterationBehavior property
// returns a *ObliterationBehavior when successful
func (m *ItemManageddevicesItemWipePostRequestBody) GetObliterationBehavior()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ObliterationBehavior) {
    val, err := m.GetBackingStore().Get("obliterationBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ObliterationBehavior)
    }
    return nil
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. The persistEsimDataPlan property
// returns a *bool when successful
func (m *ItemManageddevicesItemWipePostRequestBody) GetPersistEsimDataPlan()(*bool) {
    val, err := m.GetBackingStore().Get("persistEsimDataPlan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUseProtectedWipe gets the useProtectedWipe property value. The useProtectedWipe property
// returns a *bool when successful
func (m *ItemManageddevicesItemWipePostRequestBody) GetUseProtectedWipe()(*bool) {
    val, err := m.GetBackingStore().Get("useProtectedWipe")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemManageddevicesItemWipePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("keepEnrollmentData", m.GetKeepEnrollmentData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macOsUnlockCode", m.GetMacOsUnlockCode())
        if err != nil {
            return err
        }
    }
    if m.GetObliterationBehavior() != nil {
        cast := (*m.GetObliterationBehavior()).String()
        err := writer.WriteStringValue("obliterationBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("persistEsimDataPlan", m.GetPersistEsimDataPlan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useProtectedWipe", m.GetUseProtectedWipe())
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
func (m *ItemManageddevicesItemWipePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemManageddevicesItemWipePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *ItemManageddevicesItemWipePostRequestBody) SetKeepEnrollmentData(value *bool)() {
    err := m.GetBackingStore().Set("keepEnrollmentData", value)
    if err != nil {
        panic(err)
    }
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *ItemManageddevicesItemWipePostRequestBody) SetKeepUserData(value *bool)() {
    err := m.GetBackingStore().Set("keepUserData", value)
    if err != nil {
        panic(err)
    }
}
// SetMacOsUnlockCode sets the macOsUnlockCode property value. The macOsUnlockCode property
func (m *ItemManageddevicesItemWipePostRequestBody) SetMacOsUnlockCode(value *string)() {
    err := m.GetBackingStore().Set("macOsUnlockCode", value)
    if err != nil {
        panic(err)
    }
}
// SetObliterationBehavior sets the obliterationBehavior property value. The obliterationBehavior property
func (m *ItemManageddevicesItemWipePostRequestBody) SetObliterationBehavior(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ObliterationBehavior)() {
    err := m.GetBackingStore().Set("obliterationBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *ItemManageddevicesItemWipePostRequestBody) SetPersistEsimDataPlan(value *bool)() {
    err := m.GetBackingStore().Set("persistEsimDataPlan", value)
    if err != nil {
        panic(err)
    }
}
// SetUseProtectedWipe sets the useProtectedWipe property value. The useProtectedWipe property
func (m *ItemManageddevicesItemWipePostRequestBody) SetUseProtectedWipe(value *bool)() {
    err := m.GetBackingStore().Set("useProtectedWipe", value)
    if err != nil {
        panic(err)
    }
}
type ItemManageddevicesItemWipePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetKeepEnrollmentData()(*bool)
    GetKeepUserData()(*bool)
    GetMacOsUnlockCode()(*string)
    GetObliterationBehavior()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ObliterationBehavior)
    GetPersistEsimDataPlan()(*bool)
    GetUseProtectedWipe()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetKeepEnrollmentData(value *bool)()
    SetKeepUserData(value *bool)()
    SetMacOsUnlockCode(value *string)()
    SetObliterationBehavior(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ObliterationBehavior)()
    SetPersistEsimDataPlan(value *bool)()
    SetUseProtectedWipe(value *bool)()
}
