package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody instantiates a new GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody and sets the default values.
func NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody()(*GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) {
    m := &GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody(), nil
}
// GetAdded gets the added property value. The added property
// returns a []GroupPolicyDefinitionValueable when successful
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) GetAdded()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable) {
    val, err := m.GetBackingStore().Get("added")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeletedIds gets the deletedIds property value. The deletedIds property
// returns a []string when successful
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) GetDeletedIds()([]string) {
    val, err := m.GetBackingStore().Get("deletedIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["added"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)
                }
            }
            m.SetAdded(res)
        }
        return nil
    }
    res["deletedIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDeletedIds(res)
        }
        return nil
    }
    res["updated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)
                }
            }
            m.SetUpdated(res)
        }
        return nil
    }
    return res
}
// GetUpdated gets the updated property value. The updated property
// returns a []GroupPolicyDefinitionValueable when successful
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) GetUpdated()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable) {
    val, err := m.GetBackingStore().Get("updated")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdded() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdded()))
        for i, v := range m.GetAdded() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("added", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedIds() != nil {
        err := writer.WriteCollectionOfStringValues("deletedIds", m.GetDeletedIds())
        if err != nil {
            return err
        }
    }
    if m.GetUpdated() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUpdated()))
        for i, v := range m.GetUpdated() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("updated", cast)
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
// SetAdded sets the added property value. The added property
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) SetAdded(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)() {
    err := m.GetBackingStore().Set("added", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeletedIds sets the deletedIds property value. The deletedIds property
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) SetDeletedIds(value []string)() {
    err := m.GetBackingStore().Set("deletedIds", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdated sets the updated property value. The updated property
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBody) SetUpdated(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)() {
    err := m.GetBackingStore().Set("updated", value)
    if err != nil {
        panic(err)
    }
}
type GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdded()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeletedIds()([]string)
    GetUpdated()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)
    SetAdded(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeletedIds(value []string)()
    SetUpdated(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)()
}
