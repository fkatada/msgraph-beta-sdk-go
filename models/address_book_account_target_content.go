package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddressBookAccountTargetContent 
type AddressBookAccountTargetContent struct {
    AccountTargetContent
}
// NewAddressBookAccountTargetContent instantiates a new addressBookAccountTargetContent and sets the default values.
func NewAddressBookAccountTargetContent()(*AddressBookAccountTargetContent) {
    m := &AddressBookAccountTargetContent{
        AccountTargetContent: *NewAccountTargetContent(),
    }
    odataTypeValue := "#microsoft.graph.addressBookAccountTargetContent"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAddressBookAccountTargetContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddressBookAccountTargetContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddressBookAccountTargetContent(), nil
}
// GetAccountTargetEmails gets the accountTargetEmails property value. List of user emails targeted for an attack simulation training campaign.
func (m *AddressBookAccountTargetContent) GetAccountTargetEmails()([]string) {
    val, err := m.GetBackingStore().Get("accountTargetEmails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddressBookAccountTargetContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccountTargetContent.GetFieldDeserializers()
    res["accountTargetEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAccountTargetEmails(res)
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
func (m *AddressBookAccountTargetContent) GetOdataType()(*string) {
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
func (m *AddressBookAccountTargetContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccountTargetContent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccountTargetEmails() != nil {
        err = writer.WriteCollectionOfStringValues("accountTargetEmails", m.GetAccountTargetEmails())
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
// SetAccountTargetEmails sets the accountTargetEmails property value. List of user emails targeted for an attack simulation training campaign.
func (m *AddressBookAccountTargetContent) SetAccountTargetEmails(value []string)() {
    err := m.GetBackingStore().Set("accountTargetEmails", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AddressBookAccountTargetContent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AddressBookAccountTargetContentable 
type AddressBookAccountTargetContentable interface {
    AccountTargetContentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountTargetEmails()([]string)
    GetOdataType()(*string)
    SetAccountTargetEmails(value []string)()
    SetOdataType(value *string)()
}
