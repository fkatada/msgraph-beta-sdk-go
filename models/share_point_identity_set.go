package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharePointIdentitySet 
type SharePointIdentitySet struct {
    IdentitySet
}
// NewSharePointIdentitySet instantiates a new sharePointIdentitySet and sets the default values.
func NewSharePointIdentitySet()(*SharePointIdentitySet) {
    m := &SharePointIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    odataTypeValue := "#microsoft.graph.sharePointIdentitySet"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSharePointIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharePointIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointIdentitySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharePointIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Identityable))
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
    res["siteGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteGroup(val.(SharePointIdentityable))
        }
        return nil
    }
    res["siteUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteUser(val.(SharePointIdentityable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group associated with this action. Optional.
func (m *SharePointIdentitySet) GetGroup()(Identityable) {
    val, err := m.GetBackingStore().Get("group")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SharePointIdentitySet) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSiteGroup gets the siteGroup property value. The SharePoint group associated with this action. Optional.
func (m *SharePointIdentitySet) GetSiteGroup()(SharePointIdentityable) {
    val, err := m.GetBackingStore().Get("siteGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharePointIdentityable)
    }
    return nil
}
// GetSiteUser gets the siteUser property value. The SharePoint user associated with this action. Optional.
func (m *SharePointIdentitySet) GetSiteUser()(SharePointIdentityable) {
    val, err := m.GetBackingStore().Get("siteUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharePointIdentityable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SharePointIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
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
    {
        err = writer.WriteObjectValue("siteGroup", m.GetSiteGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteUser", m.GetSiteUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. The group associated with this action. Optional.
func (m *SharePointIdentitySet) SetGroup(value Identityable)() {
    err := m.GetBackingStore().Set("group", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SharePointIdentitySet) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteGroup sets the siteGroup property value. The SharePoint group associated with this action. Optional.
func (m *SharePointIdentitySet) SetSiteGroup(value SharePointIdentityable)() {
    err := m.GetBackingStore().Set("siteGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteUser sets the siteUser property value. The SharePoint user associated with this action. Optional.
func (m *SharePointIdentitySet) SetSiteUser(value SharePointIdentityable)() {
    err := m.GetBackingStore().Set("siteUser", value)
    if err != nil {
        panic(err)
    }
}
// SharePointIdentitySetable 
type SharePointIdentitySetable interface {
    IdentitySetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroup()(Identityable)
    GetOdataType()(*string)
    GetSiteGroup()(SharePointIdentityable)
    GetSiteUser()(SharePointIdentityable)
    SetGroup(value Identityable)()
    SetOdataType(value *string)()
    SetSiteGroup(value SharePointIdentityable)()
    SetSiteUser(value SharePointIdentityable)()
}
