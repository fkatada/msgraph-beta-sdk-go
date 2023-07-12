package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnTokenIssuanceStartCustomExtensionHandler 
type OnTokenIssuanceStartCustomExtensionHandler struct {
    OnTokenIssuanceStartHandler
}
// NewOnTokenIssuanceStartCustomExtensionHandler instantiates a new onTokenIssuanceStartCustomExtensionHandler and sets the default values.
func NewOnTokenIssuanceStartCustomExtensionHandler()(*OnTokenIssuanceStartCustomExtensionHandler) {
    m := &OnTokenIssuanceStartCustomExtensionHandler{
        OnTokenIssuanceStartHandler: *NewOnTokenIssuanceStartHandler(),
    }
    odataTypeValue := "#microsoft.graph.onTokenIssuanceStartCustomExtensionHandler"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnTokenIssuanceStartCustomExtensionHandlerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnTokenIssuanceStartCustomExtensionHandlerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnTokenIssuanceStartCustomExtensionHandler(), nil
}
// GetCustomExtension gets the customExtension property value. The customExtension property
func (m *OnTokenIssuanceStartCustomExtensionHandler) GetCustomExtension()(OnTokenIssuanceStartCustomExtensionable) {
    val, err := m.GetBackingStore().Get("customExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnTokenIssuanceStartCustomExtensionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnTokenIssuanceStartCustomExtensionHandler) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnTokenIssuanceStartHandler.GetFieldDeserializers()
    res["customExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnTokenIssuanceStartCustomExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(OnTokenIssuanceStartCustomExtensionable))
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnTokenIssuanceStartCustomExtensionHandler) GetOdataType()(*string) {
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
func (m *OnTokenIssuanceStartCustomExtensionHandler) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnTokenIssuanceStartHandler.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomExtension sets the customExtension property value. The customExtension property
func (m *OnTokenIssuanceStartCustomExtensionHandler) SetCustomExtension(value OnTokenIssuanceStartCustomExtensionable)() {
    err := m.GetBackingStore().Set("customExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnTokenIssuanceStartCustomExtensionHandler) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// OnTokenIssuanceStartCustomExtensionHandlerable 
type OnTokenIssuanceStartCustomExtensionHandlerable interface {
    OnTokenIssuanceStartHandlerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomExtension()(OnTokenIssuanceStartCustomExtensionable)
    GetOdataType()(*string)
    SetCustomExtension(value OnTokenIssuanceStartCustomExtensionable)()
    SetOdataType(value *string)()
}
