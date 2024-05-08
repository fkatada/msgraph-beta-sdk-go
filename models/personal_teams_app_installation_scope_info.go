package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PersonalTeamsAppInstallationScopeInfo struct {
    TeamsAppInstallationScopeInfo
}
// NewPersonalTeamsAppInstallationScopeInfo instantiates a new PersonalTeamsAppInstallationScopeInfo and sets the default values.
func NewPersonalTeamsAppInstallationScopeInfo()(*PersonalTeamsAppInstallationScopeInfo) {
    m := &PersonalTeamsAppInstallationScopeInfo{
        TeamsAppInstallationScopeInfo: *NewTeamsAppInstallationScopeInfo(),
    }
    odataTypeValue := "#microsoft.graph.personalTeamsAppInstallationScopeInfo"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePersonalTeamsAppInstallationScopeInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePersonalTeamsAppInstallationScopeInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonalTeamsAppInstallationScopeInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PersonalTeamsAppInstallationScopeInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamsAppInstallationScopeInfo.GetFieldDeserializers()
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetUserId gets the userId property value. The userId property
// returns a *string when successful
func (m *PersonalTeamsAppInstallationScopeInfo) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PersonalTeamsAppInstallationScopeInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamsAppInstallationScopeInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserId sets the userId property value. The userId property
func (m *PersonalTeamsAppInstallationScopeInfo) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
type PersonalTeamsAppInstallationScopeInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamsAppInstallationScopeInfoable
    GetUserId()(*string)
    SetUserId(value *string)()
}
