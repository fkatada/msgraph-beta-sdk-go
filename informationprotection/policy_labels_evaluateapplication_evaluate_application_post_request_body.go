package informationprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody instantiates a new PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody and sets the default values.
func NewPolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody()(*PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) {
    m := &PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
// returns a ContentInfoable when successful
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) GetContentInfo()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentInfoable) {
    val, err := m.GetBackingStore().Get("contentInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentInfoable))
        }
        return nil
    }
    res["labelingOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLabelingOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelingOptions(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LabelingOptionsable))
        }
        return nil
    }
    return res
}
// GetLabelingOptions gets the labelingOptions property value. The labelingOptions property
// returns a LabelingOptionsable when successful
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) GetLabelingOptions()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LabelingOptionsable) {
    val, err := m.GetBackingStore().Get("labelingOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LabelingOptionsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("labelingOptions", m.GetLabelingOptions())
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
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) SetContentInfo(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentInfoable)() {
    err := m.GetBackingStore().Set("contentInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetLabelingOptions sets the labelingOptions property value. The labelingOptions property
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBody) SetLabelingOptions(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LabelingOptionsable)() {
    err := m.GetBackingStore().Set("labelingOptions", value)
    if err != nil {
        panic(err)
    }
}
type PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetContentInfo()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentInfoable)
    GetLabelingOptions()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LabelingOptionsable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetContentInfo(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentInfoable)()
    SetLabelingOptions(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LabelingOptionsable)()
}
