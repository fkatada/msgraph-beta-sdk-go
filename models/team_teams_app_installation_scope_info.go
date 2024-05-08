package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamTeamsAppInstallationScopeInfo struct {
    TeamsAppInstallationScopeInfo
}
// NewTeamTeamsAppInstallationScopeInfo instantiates a new TeamTeamsAppInstallationScopeInfo and sets the default values.
func NewTeamTeamsAppInstallationScopeInfo()(*TeamTeamsAppInstallationScopeInfo) {
    m := &TeamTeamsAppInstallationScopeInfo{
        TeamsAppInstallationScopeInfo: *NewTeamsAppInstallationScopeInfo(),
    }
    odataTypeValue := "#microsoft.graph.teamTeamsAppInstallationScopeInfo"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeamTeamsAppInstallationScopeInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamTeamsAppInstallationScopeInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTeamsAppInstallationScopeInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamTeamsAppInstallationScopeInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamsAppInstallationScopeInfo.GetFieldDeserializers()
    res["teamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    return res
}
// GetTeamId gets the teamId property value. The teamId property
// returns a *string when successful
func (m *TeamTeamsAppInstallationScopeInfo) GetTeamId()(*string) {
    val, err := m.GetBackingStore().Get("teamId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamTeamsAppInstallationScopeInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamsAppInstallationScopeInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTeamId sets the teamId property value. The teamId property
func (m *TeamTeamsAppInstallationScopeInfo) SetTeamId(value *string)() {
    err := m.GetBackingStore().Set("teamId", value)
    if err != nil {
        panic(err)
    }
}
type TeamTeamsAppInstallationScopeInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamsAppInstallationScopeInfoable
    GetTeamId()(*string)
    SetTeamId(value *string)()
}
