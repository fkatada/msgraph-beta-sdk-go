package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody()(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) {
    m := &ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
// returns a ContentInfoable when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) GetContentInfo()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable) {
    val, err := m.GetBackingStore().Get("contentInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)
    }
    return nil
}
// GetDowngradeJustification gets the downgradeJustification property value. The downgradeJustification property
// returns a DowngradeJustificationable when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) GetDowngradeJustification()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DowngradeJustificationable) {
    val, err := m.GetBackingStore().Get("downgradeJustification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DowngradeJustificationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable))
        }
        return nil
    }
    res["downgradeJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateDowngradeJustificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeJustification(val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DowngradeJustificationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("downgradeJustification", m.GetDowngradeJustification())
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
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) SetContentInfo(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)() {
    err := m.GetBackingStore().Set("contentInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetDowngradeJustification sets the downgradeJustification property value. The downgradeJustification property
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBody) SetDowngradeJustification(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DowngradeJustificationable)() {
    err := m.GetBackingStore().Set("downgradeJustification", value)
    if err != nil {
        panic(err)
    }
}
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetContentInfo()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)
    GetDowngradeJustification()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DowngradeJustificationable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetContentInfo(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)()
    SetDowngradeJustification(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DowngradeJustificationable)()
}
