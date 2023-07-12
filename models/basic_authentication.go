package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BasicAuthentication 
type BasicAuthentication struct {
    ApiAuthenticationConfigurationBase
}
// NewBasicAuthentication instantiates a new basicAuthentication and sets the default values.
func NewBasicAuthentication()(*BasicAuthentication) {
    m := &BasicAuthentication{
        ApiAuthenticationConfigurationBase: *NewApiAuthenticationConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.basicAuthentication"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBasicAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBasicAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBasicAuthentication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BasicAuthentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApiAuthenticationConfigurationBase.GetFieldDeserializers()
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
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BasicAuthentication) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPassword gets the password property value. The password. It is not returned in the responses.
func (m *BasicAuthentication) GetPassword()(*string) {
    val, err := m.GetBackingStore().Get("password")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUsername gets the username property value. The username.
func (m *BasicAuthentication) GetUsername()(*string) {
    val, err := m.GetBackingStore().Get("username")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BasicAuthentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ApiAuthenticationConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("username", m.GetUsername())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BasicAuthentication) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPassword sets the password property value. The password. It is not returned in the responses.
func (m *BasicAuthentication) SetPassword(value *string)() {
    err := m.GetBackingStore().Set("password", value)
    if err != nil {
        panic(err)
    }
}
// SetUsername sets the username property value. The username.
func (m *BasicAuthentication) SetUsername(value *string)() {
    err := m.GetBackingStore().Set("username", value)
    if err != nil {
        panic(err)
    }
}
// BasicAuthenticationable 
type BasicAuthenticationable interface {
    ApiAuthenticationConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPassword()(*string)
    GetUsername()(*string)
    SetOdataType(value *string)()
    SetPassword(value *string)()
    SetUsername(value *string)()
}
