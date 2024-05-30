package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

type CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody instantiates a new CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody()(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) {
    m := &CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAdditionalOptions gets the additionalOptions property value. The additionalOptions property
// returns a *AdditionalOptions when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetAdditionalOptions()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalOptions) {
    val, err := m.GetBackingStore().Get("additionalOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalOptions)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExportCriteria gets the exportCriteria property value. The exportCriteria property
// returns a *ExportCriteria when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetExportCriteria()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportCriteria) {
    val, err := m.GetBackingStore().Get("exportCriteria")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportCriteria)
    }
    return nil
}
// GetExportFormat gets the exportFormat property value. The exportFormat property
// returns a *ExportFormat when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetExportFormat()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFormat) {
    val, err := m.GetBackingStore().Get("exportFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFormat)
    }
    return nil
}
// GetExportLocation gets the exportLocation property value. The exportLocation property
// returns a *ExportLocation when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetExportLocation()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportLocation) {
    val, err := m.GetBackingStore().Get("exportLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportLocation)
    }
    return nil
}
// GetExportSingleItems gets the exportSingleItems property value. The exportSingleItems property
// returns a *bool when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetExportSingleItems()(*bool) {
    val, err := m.GetBackingStore().Get("exportSingleItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseAdditionalOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalOptions(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalOptions))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["exportCriteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseExportCriteria)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportCriteria(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportCriteria))
        }
        return nil
    }
    res["exportFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseExportFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportFormat(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFormat))
        }
        return nil
    }
    res["exportLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseExportLocation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportLocation(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportLocation))
        }
        return nil
    }
    res["exportSingleItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportSingleItems(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalOptions() != nil {
        cast := (*m.GetAdditionalOptions()).String()
        err := writer.WriteStringValue("additionalOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExportCriteria() != nil {
        cast := (*m.GetExportCriteria()).String()
        err := writer.WriteStringValue("exportCriteria", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportFormat() != nil {
        cast := (*m.GetExportFormat()).String()
        err := writer.WriteStringValue("exportFormat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportLocation() != nil {
        cast := (*m.GetExportLocation()).String()
        err := writer.WriteStringValue("exportLocation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exportSingleItems", m.GetExportSingleItems())
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
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalOptions sets the additionalOptions property value. The additionalOptions property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetAdditionalOptions(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalOptions)() {
    err := m.GetBackingStore().Set("additionalOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDescription sets the description property value. The description property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExportCriteria sets the exportCriteria property value. The exportCriteria property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetExportCriteria(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportCriteria)() {
    err := m.GetBackingStore().Set("exportCriteria", value)
    if err != nil {
        panic(err)
    }
}
// SetExportFormat sets the exportFormat property value. The exportFormat property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetExportFormat(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFormat)() {
    err := m.GetBackingStore().Set("exportFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetExportLocation sets the exportLocation property value. The exportLocation property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetExportLocation(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportLocation)() {
    err := m.GetBackingStore().Set("exportLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetExportSingleItems sets the exportSingleItems property value. The exportSingleItems property
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBody) SetExportSingleItems(value *bool)() {
    err := m.GetBackingStore().Set("exportSingleItems", value)
    if err != nil {
        panic(err)
    }
}
type CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityexportresultExportResultPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalOptions()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalOptions)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExportCriteria()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportCriteria)
    GetExportFormat()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFormat)
    GetExportLocation()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportLocation)
    GetExportSingleItems()(*bool)
    SetAdditionalOptions(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalOptions)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExportCriteria(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportCriteria)()
    SetExportFormat(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFormat)()
    SetExportLocation(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportLocation)()
    SetExportSingleItems(value *bool)()
}
