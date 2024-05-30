package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody instantiates a new ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody and sets the default values.
func NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody()(*ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) {
    m := &ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExpression gets the expression property value. The expression property
// returns a *string when successful
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) GetExpression()(*string) {
    val, err := m.GetBackingStore().Get("expression")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expression"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpression(val)
        }
        return nil
    }
    res["targetAttributeDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetAttributeDefinition(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable))
        }
        return nil
    }
    res["testInputObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExpressionInputObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestInputObject(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable))
        }
        return nil
    }
    return res
}
// GetTargetAttributeDefinition gets the targetAttributeDefinition property value. The targetAttributeDefinition property
// returns a AttributeDefinitionable when successful
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) GetTargetAttributeDefinition()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable) {
    val, err := m.GetBackingStore().Get("targetAttributeDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable)
    }
    return nil
}
// GetTestInputObject gets the testInputObject property value. The testInputObject property
// returns a ExpressionInputObjectable when successful
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) GetTestInputObject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable) {
    val, err := m.GetBackingStore().Get("testInputObject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("targetAttributeDefinition", m.GetTargetAttributeDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("testInputObject", m.GetTestInputObject())
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
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExpression sets the expression property value. The expression property
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) SetExpression(value *string)() {
    err := m.GetBackingStore().Set("expression", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetAttributeDefinition sets the targetAttributeDefinition property value. The targetAttributeDefinition property
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) SetTargetAttributeDefinition(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable)() {
    err := m.GetBackingStore().Set("targetAttributeDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetTestInputObject sets the testInputObject property value. The testInputObject property
func (m *ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBody) SetTestInputObject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable)() {
    err := m.GetBackingStore().Set("testInputObject", value)
    if err != nil {
        panic(err)
    }
}
type ItemSynchronizationTemplatesItemSchemaParseexpressionParseExpressionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExpression()(*string)
    GetTargetAttributeDefinition()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable)
    GetTestInputObject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExpression(value *string)()
    SetTargetAttributeDefinition(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable)()
    SetTestInputObject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable)()
}
