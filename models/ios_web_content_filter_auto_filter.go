package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosWebContentFilterAutoFilter represents an iOS Web Content Filter setting type, which enables iOS automatic filter feature and allows for additional URL access control. When constructed with no property values, the iOS device will enable the automatic filter regardless.
type IosWebContentFilterAutoFilter struct {
    IosWebContentFilterBase
}
// NewIosWebContentFilterAutoFilter instantiates a new iosWebContentFilterAutoFilter and sets the default values.
func NewIosWebContentFilterAutoFilter()(*IosWebContentFilterAutoFilter) {
    m := &IosWebContentFilterAutoFilter{
        IosWebContentFilterBase: *NewIosWebContentFilterBase(),
    }
    odataTypeValue := "#microsoft.graph.iosWebContentFilterAutoFilter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosWebContentFilterAutoFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosWebContentFilterAutoFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosWebContentFilterAutoFilter(), nil
}
// GetAllowedUrls gets the allowedUrls property value. Additional URLs allowed for access
func (m *IosWebContentFilterAutoFilter) GetAllowedUrls()([]string) {
    val, err := m.GetBackingStore().Get("allowedUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetBlockedUrls gets the blockedUrls property value. Additional URLs blocked for access
func (m *IosWebContentFilterAutoFilter) GetBlockedUrls()([]string) {
    val, err := m.GetBackingStore().Get("blockedUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosWebContentFilterAutoFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosWebContentFilterBase.GetFieldDeserializers()
    res["allowedUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAllowedUrls(res)
        }
        return nil
    }
    res["blockedUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBlockedUrls(res)
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
func (m *IosWebContentFilterAutoFilter) GetOdataType()(*string) {
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
func (m *IosWebContentFilterAutoFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosWebContentFilterBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedUrls() != nil {
        err = writer.WriteCollectionOfStringValues("allowedUrls", m.GetAllowedUrls())
        if err != nil {
            return err
        }
    }
    if m.GetBlockedUrls() != nil {
        err = writer.WriteCollectionOfStringValues("blockedUrls", m.GetBlockedUrls())
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
// SetAllowedUrls sets the allowedUrls property value. Additional URLs allowed for access
func (m *IosWebContentFilterAutoFilter) SetAllowedUrls(value []string)() {
    err := m.GetBackingStore().Set("allowedUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockedUrls sets the blockedUrls property value. Additional URLs blocked for access
func (m *IosWebContentFilterAutoFilter) SetBlockedUrls(value []string)() {
    err := m.GetBackingStore().Set("blockedUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosWebContentFilterAutoFilter) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// IosWebContentFilterAutoFilterable 
type IosWebContentFilterAutoFilterable interface {
    IosWebContentFilterBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedUrls()([]string)
    GetBlockedUrls()([]string)
    GetOdataType()(*string)
    SetAllowedUrls(value []string)()
    SetBlockedUrls(value []string)()
    SetOdataType(value *string)()
}
