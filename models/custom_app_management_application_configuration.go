package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CustomAppManagementApplicationConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCustomAppManagementApplicationConfiguration instantiates a new CustomAppManagementApplicationConfiguration and sets the default values.
func NewCustomAppManagementApplicationConfiguration()(*CustomAppManagementApplicationConfiguration) {
    m := &CustomAppManagementApplicationConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomAppManagementApplicationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomAppManagementApplicationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomAppManagementApplicationConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomAppManagementApplicationConfiguration) GetAdditionalData()(map[string]any) {
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
// GetAudiences gets the audiences property value. Property to restrict creation or update of apps based on their target signInAudience types.
// returns a AudiencesConfigurationable when successful
func (m *CustomAppManagementApplicationConfiguration) GetAudiences()(AudiencesConfigurationable) {
    val, err := m.GetBackingStore().Get("audiences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AudiencesConfigurationable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CustomAppManagementApplicationConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomAppManagementApplicationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["audiences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAudiencesConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudiences(val.(AudiencesConfigurationable))
        }
        return nil
    }
    res["identifierUris"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentifierUriConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifierUris(val.(IdentifierUriConfigurationable))
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
    return res
}
// GetIdentifierUris gets the identifierUris property value. Configuration for identifierUris restrictions.
// returns a IdentifierUriConfigurationable when successful
func (m *CustomAppManagementApplicationConfiguration) GetIdentifierUris()(IdentifierUriConfigurationable) {
    val, err := m.GetBackingStore().Get("identifierUris")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentifierUriConfigurationable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CustomAppManagementApplicationConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomAppManagementApplicationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("audiences", m.GetAudiences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("identifierUris", m.GetIdentifierUris())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomAppManagementApplicationConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAudiences sets the audiences property value. Property to restrict creation or update of apps based on their target signInAudience types.
func (m *CustomAppManagementApplicationConfiguration) SetAudiences(value AudiencesConfigurationable)() {
    err := m.GetBackingStore().Set("audiences", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CustomAppManagementApplicationConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIdentifierUris sets the identifierUris property value. Configuration for identifierUris restrictions.
func (m *CustomAppManagementApplicationConfiguration) SetIdentifierUris(value IdentifierUriConfigurationable)() {
    err := m.GetBackingStore().Set("identifierUris", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomAppManagementApplicationConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type CustomAppManagementApplicationConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudiences()(AudiencesConfigurationable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIdentifierUris()(IdentifierUriConfigurationable)
    GetOdataType()(*string)
    SetAudiences(value AudiencesConfigurationable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIdentifierUris(value IdentifierUriConfigurationable)()
    SetOdataType(value *string)()
}
