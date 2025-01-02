package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AnalyzedEmailSenderDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAnalyzedEmailSenderDetail instantiates a new AnalyzedEmailSenderDetail and sets the default values.
func NewAnalyzedEmailSenderDetail()(*AnalyzedEmailSenderDetail) {
    m := &AnalyzedEmailSenderDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAnalyzedEmailSenderDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAnalyzedEmailSenderDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnalyzedEmailSenderDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AnalyzedEmailSenderDetail) GetAdditionalData()(map[string]any) {
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
func (m *AnalyzedEmailSenderDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomainCreationDateTime gets the domainCreationDateTime property value. The domainCreationDateTime property
// returns a *Time when successful
func (m *AnalyzedEmailSenderDetail) GetDomainCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("domainCreationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDomainName gets the domainName property value. The domainName property
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetDomainName()(*string) {
    val, err := m.GetBackingStore().Get("domainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomainOwner gets the domainOwner property value. The domainOwner property
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetDomainOwner()(*string) {
    val, err := m.GetBackingStore().Get("domainOwner")
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
func (m *AnalyzedEmailSenderDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["domainCreationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainCreationDateTime(val)
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["domainOwner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainOwner(val)
        }
        return nil
    }
    res["fromAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromAddress(val)
        }
        return nil
    }
    res["ipv4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpv4(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["mailFromAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailFromAddress(val)
        }
        return nil
    }
    res["mailFromDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailFromDomainName(val)
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
// GetFromAddress gets the fromAddress property value. The sender email address in the mail From header, also known as the envelope sender or the P1 sender.
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetFromAddress()(*string) {
    val, err := m.GetBackingStore().Get("fromAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIpv4 gets the ipv4 property value. The IPv4 address of the last detected mail server that relayed the message.
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetIpv4()(*string) {
    val, err := m.GetBackingStore().Get("ipv4")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocation gets the location property value. The location property
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetLocation()(*string) {
    val, err := m.GetBackingStore().Get("location")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMailFromAddress gets the mailFromAddress property value. The sender email address in the From header, which is visible to email recipients on their email clients. Also known as P2 sender.
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetMailFromAddress()(*string) {
    val, err := m.GetBackingStore().Get("mailFromAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMailFromDomainName gets the mailFromDomainName property value. The mailFromDomainName property
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetMailFromDomainName()(*string) {
    val, err := m.GetBackingStore().Get("mailFromDomainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AnalyzedEmailSenderDetail) GetOdataType()(*string) {
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
func (m *AnalyzedEmailSenderDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("domainCreationDateTime", m.GetDomainCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainOwner", m.GetDomainOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fromAddress", m.GetFromAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipv4", m.GetIpv4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mailFromAddress", m.GetMailFromAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mailFromDomainName", m.GetMailFromDomainName())
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
func (m *AnalyzedEmailSenderDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AnalyzedEmailSenderDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AnalyzedEmailSenderDetail) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainCreationDateTime sets the domainCreationDateTime property value. The domainCreationDateTime property
func (m *AnalyzedEmailSenderDetail) SetDomainCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("domainCreationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainName sets the domainName property value. The domainName property
func (m *AnalyzedEmailSenderDetail) SetDomainName(value *string)() {
    err := m.GetBackingStore().Set("domainName", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainOwner sets the domainOwner property value. The domainOwner property
func (m *AnalyzedEmailSenderDetail) SetDomainOwner(value *string)() {
    err := m.GetBackingStore().Set("domainOwner", value)
    if err != nil {
        panic(err)
    }
}
// SetFromAddress sets the fromAddress property value. The sender email address in the mail From header, also known as the envelope sender or the P1 sender.
func (m *AnalyzedEmailSenderDetail) SetFromAddress(value *string)() {
    err := m.GetBackingStore().Set("fromAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetIpv4 sets the ipv4 property value. The IPv4 address of the last detected mail server that relayed the message.
func (m *AnalyzedEmailSenderDetail) SetIpv4(value *string)() {
    err := m.GetBackingStore().Set("ipv4", value)
    if err != nil {
        panic(err)
    }
}
// SetLocation sets the location property value. The location property
func (m *AnalyzedEmailSenderDetail) SetLocation(value *string)() {
    err := m.GetBackingStore().Set("location", value)
    if err != nil {
        panic(err)
    }
}
// SetMailFromAddress sets the mailFromAddress property value. The sender email address in the From header, which is visible to email recipients on their email clients. Also known as P2 sender.
func (m *AnalyzedEmailSenderDetail) SetMailFromAddress(value *string)() {
    err := m.GetBackingStore().Set("mailFromAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetMailFromDomainName sets the mailFromDomainName property value. The mailFromDomainName property
func (m *AnalyzedEmailSenderDetail) SetMailFromDomainName(value *string)() {
    err := m.GetBackingStore().Set("mailFromDomainName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AnalyzedEmailSenderDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type AnalyzedEmailSenderDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDisplayName()(*string)
    GetDomainCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDomainName()(*string)
    GetDomainOwner()(*string)
    GetFromAddress()(*string)
    GetIpv4()(*string)
    GetLocation()(*string)
    GetMailFromAddress()(*string)
    GetMailFromDomainName()(*string)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDisplayName(value *string)()
    SetDomainCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDomainName(value *string)()
    SetDomainOwner(value *string)()
    SetFromAddress(value *string)()
    SetIpv4(value *string)()
    SetLocation(value *string)()
    SetMailFromAddress(value *string)()
    SetMailFromDomainName(value *string)()
    SetOdataType(value *string)()
}
