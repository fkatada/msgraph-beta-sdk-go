package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// SolutionsRoot 
type SolutionsRoot struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSolutionsRoot instantiates a new solutionsRoot and sets the default values.
func NewSolutionsRoot()(*SolutionsRoot) {
    m := &SolutionsRoot{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSolutionsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSolutionsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSolutionsRoot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SolutionsRoot) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *SolutionsRoot) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBusinessScenarios gets the businessScenarios property value. The businessScenarios property
func (m *SolutionsRoot) GetBusinessScenarios()([]BusinessScenarioable) {
    val, err := m.GetBackingStore().Get("businessScenarios")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]BusinessScenarioable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SolutionsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["businessScenarios"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBusinessScenarioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessScenarioable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BusinessScenarioable)
                }
            }
            m.SetBusinessScenarios(res)
        }
        return nil
    }
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
    res["virtualEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEventsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualEvents(val.(VirtualEventsRootable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SolutionsRoot) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVirtualEvents gets the virtualEvents property value. The virtualEvents property
func (m *SolutionsRoot) GetVirtualEvents()(VirtualEventsRootable) {
    val, err := m.GetBackingStore().Get("virtualEvents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VirtualEventsRootable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SolutionsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBusinessScenarios() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBusinessScenarios()))
        for i, v := range m.GetBusinessScenarios() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("businessScenarios", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("virtualEvents", m.GetVirtualEvents())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SolutionsRoot) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *SolutionsRoot) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBusinessScenarios sets the businessScenarios property value. The businessScenarios property
func (m *SolutionsRoot) SetBusinessScenarios(value []BusinessScenarioable)() {
    err := m.GetBackingStore().Set("businessScenarios", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SolutionsRoot) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualEvents sets the virtualEvents property value. The virtualEvents property
func (m *SolutionsRoot) SetVirtualEvents(value VirtualEventsRootable)() {
    err := m.GetBackingStore().Set("virtualEvents", value)
    if err != nil {
        panic(err)
    }
}
// SolutionsRootable 
type SolutionsRootable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBusinessScenarios()([]BusinessScenarioable)
    GetOdataType()(*string)
    GetVirtualEvents()(VirtualEventsRootable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBusinessScenarios(value []BusinessScenarioable)()
    SetOdataType(value *string)()
    SetVirtualEvents(value VirtualEventsRootable)()
}
