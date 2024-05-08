package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEventPresenter struct {
    Entity
}
// NewVirtualEventPresenter instantiates a new VirtualEventPresenter and sets the default values.
func NewVirtualEventPresenter()(*VirtualEventPresenter) {
    m := &VirtualEventPresenter{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventPresenterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventPresenterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventPresenter(), nil
}
// GetEmail gets the email property value. Email address of the presenter.
// returns a *string when successful
func (m *VirtualEventPresenter) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
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
func (m *VirtualEventPresenter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(Identityable))
        }
        return nil
    }
    res["presenterDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEventPresenterDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresenterDetails(val.(VirtualEventPresenterDetailsable))
        }
        return nil
    }
    res["sessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventSessionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventSessionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventSessionable)
                }
            }
            m.SetSessions(res)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. Identity information of the presenter. The supported identities are: communicationsGuestIdentity and communicationsUserIdentity.
// returns a Identityable when successful
func (m *VirtualEventPresenter) GetIdentity()(Identityable) {
    val, err := m.GetBackingStore().Get("identity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetPresenterDetails gets the presenterDetails property value. Other detail information of the presenter.
// returns a VirtualEventPresenterDetailsable when successful
func (m *VirtualEventPresenter) GetPresenterDetails()(VirtualEventPresenterDetailsable) {
    val, err := m.GetBackingStore().Get("presenterDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VirtualEventPresenterDetailsable)
    }
    return nil
}
// GetSessions gets the sessions property value. The sessions property
// returns a []VirtualEventSessionable when successful
func (m *VirtualEventPresenter) GetSessions()([]VirtualEventSessionable) {
    val, err := m.GetBackingStore().Get("sessions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventSessionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEventPresenter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("presenterDetails", m.GetPresenterDetails())
        if err != nil {
            return err
        }
    }
    if m.GetSessions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmail sets the email property value. Email address of the presenter.
func (m *VirtualEventPresenter) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentity sets the identity property value. Identity information of the presenter. The supported identities are: communicationsGuestIdentity and communicationsUserIdentity.
func (m *VirtualEventPresenter) SetIdentity(value Identityable)() {
    err := m.GetBackingStore().Set("identity", value)
    if err != nil {
        panic(err)
    }
}
// SetPresenterDetails sets the presenterDetails property value. Other detail information of the presenter.
func (m *VirtualEventPresenter) SetPresenterDetails(value VirtualEventPresenterDetailsable)() {
    err := m.GetBackingStore().Set("presenterDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetSessions sets the sessions property value. The sessions property
func (m *VirtualEventPresenter) SetSessions(value []VirtualEventSessionable)() {
    err := m.GetBackingStore().Set("sessions", value)
    if err != nil {
        panic(err)
    }
}
type VirtualEventPresenterable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetIdentity()(Identityable)
    GetPresenterDetails()(VirtualEventPresenterDetailsable)
    GetSessions()([]VirtualEventSessionable)
    SetEmail(value *string)()
    SetIdentity(value Identityable)()
    SetPresenterDetails(value VirtualEventPresenterDetailsable)()
    SetSessions(value []VirtualEventSessionable)()
}
