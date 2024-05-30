package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

type CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody instantiates a new CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody()(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) {
    m := &CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["purgeAreas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParsePurgeAreas)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurgeAreas(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas))
        }
        return nil
    }
    res["purgeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParsePurgeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurgeType(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType))
        }
        return nil
    }
    return res
}
// GetPurgeAreas gets the purgeAreas property value. The purgeAreas property
// returns a *PurgeAreas when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) GetPurgeAreas()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas) {
    val, err := m.GetBackingStore().Get("purgeAreas")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas)
    }
    return nil
}
// GetPurgeType gets the purgeType property value. The purgeType property
// returns a *PurgeType when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) GetPurgeType()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType) {
    val, err := m.GetBackingStore().Get("purgeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPurgeAreas() != nil {
        cast := (*m.GetPurgeAreas()).String()
        err := writer.WriteStringValue("purgeAreas", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurgeType() != nil {
        cast := (*m.GetPurgeType()).String()
        err := writer.WriteStringValue("purgeType", &cast)
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
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetPurgeAreas sets the purgeAreas property value. The purgeAreas property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) SetPurgeAreas(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas)() {
    err := m.GetBackingStore().Set("purgeAreas", value)
    if err != nil {
        panic(err)
    }
}
// SetPurgeType sets the purgeType property value. The purgeType property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBody) SetPurgeType(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType)() {
    err := m.GetBackingStore().Set("purgeType", value)
    if err != nil {
        panic(err)
    }
}
type CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataPurgeDataPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetPurgeAreas()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas)
    GetPurgeType()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetPurgeAreas(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas)()
    SetPurgeType(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType)()
}
