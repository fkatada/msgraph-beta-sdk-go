package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody 
type EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody instantiates a new EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody and sets the default values.
func NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody()(*EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) {
    m := &EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIds(res)
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) GetIds()([]string) {
    val, err := m.GetBackingStore().Get("ids")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
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
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIds sets the ids property value. The ids property
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) SetIds(value []string)() {
    err := m.GetBackingStore().Set("ids", value)
    if err != nil {
        panic(err)
    }
}
// EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBodyable 
type EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIds()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIds(value []string)()
}
