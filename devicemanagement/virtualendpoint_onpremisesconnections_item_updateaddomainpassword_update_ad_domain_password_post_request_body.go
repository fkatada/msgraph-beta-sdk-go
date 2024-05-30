package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody instantiates a new VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody and sets the default values.
func NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody()(*VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) {
    m := &VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAdDomainPassword gets the adDomainPassword property value. The adDomainPassword property
// returns a *string when successful
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) GetAdDomainPassword()(*string) {
    val, err := m.GetBackingStore().Get("adDomainPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adDomainPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainPassword(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("adDomainPassword", m.GetAdDomainPassword())
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
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdDomainPassword sets the adDomainPassword property value. The adDomainPassword property
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) SetAdDomainPassword(value *string)() {
    err := m.GetBackingStore().Set("adDomainPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
type VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdDomainPassword()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    SetAdDomainPassword(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
}
