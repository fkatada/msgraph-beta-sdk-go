package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCriteria gets the criteria property value. The criteria property
// returns a WorkbookFilterCriteriaable when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) GetCriteria()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterCriteriaable) {
    val, err := m.GetBackingStore().Get("criteria")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterCriteriaable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["criteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFilterCriteriaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterCriteriaable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("criteria", m.GetCriteria())
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
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCriteria sets the criteria property value. The criteria property
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBody) SetCriteria(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterCriteriaable)() {
    err := m.GetBackingStore().Set("criteria", value)
    if err != nil {
        panic(err)
    }
}
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCriteria()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterCriteriaable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCriteria(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterCriteriaable)()
}
