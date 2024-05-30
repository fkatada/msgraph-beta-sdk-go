package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody instantiates a new OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody and sets the default values.
func NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody()(*OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) {
    m := &OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEntityId gets the entityId property value. The entityId property
// returns a *string when successful
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) GetEntityId()(*string) {
    val, err := m.GetBackingStore().Get("entityId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEntityType gets the entityType property value. The entityType property
// returns a *string when successful
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) GetEntityType()(*string) {
    val, err := m.GetBackingStore().Get("entityType")
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
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["entityId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityId(val)
        }
        return nil
    }
    res["entityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("entityId", m.GetEntityId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityType", m.GetEntityType())
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
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEntityId sets the entityId property value. The entityId property
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) SetEntityId(value *string)() {
    err := m.GetBackingStore().Set("entityId", value)
    if err != nil {
        panic(err)
    }
}
// SetEntityType sets the entityType property value. The entityType property
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBody) SetEntityType(value *string)() {
    err := m.GetBackingStore().Set("entityType", value)
    if err != nil {
        panic(err)
    }
}
type OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEntityId()(*string)
    GetEntityType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEntityId(value *string)()
    SetEntityType(value *string)()
}
