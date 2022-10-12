package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Participant provides operations to manage the commsApplication singleton.
type Participant struct {
    Entity
    // The info property
    info ParticipantInfoable
    // The isIdentityAnonymized property
    isIdentityAnonymized *bool
    // true if the participant is in lobby.
    isInLobby *bool
    // true if the participant is muted (client or server muted).
    isMuted *bool
    // The list of media streams.
    mediaStreams []MediaStreamable
    // A blob of data provided by the participant in the roster.
    metadata *string
    // Information on whether the participant has recording capability.
    recordingInfo RecordingInfoable
}
// NewParticipant instantiates a new participant and sets the default values.
func NewParticipant()(*Participant) {
    m := &Participant{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.participant";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateParticipantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParticipantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipant(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Participant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["info"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateParticipantInfoFromDiscriminatorValue , m.SetInfo)
    res["isIdentityAnonymized"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsIdentityAnonymized)
    res["isInLobby"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsInLobby)
    res["isMuted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsMuted)
    res["mediaStreams"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMediaStreamFromDiscriminatorValue , m.SetMediaStreams)
    res["metadata"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMetadata)
    res["recordingInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRecordingInfoFromDiscriminatorValue , m.SetRecordingInfo)
    return res
}
// GetInfo gets the info property value. The info property
func (m *Participant) GetInfo()(ParticipantInfoable) {
    return m.info
}
// GetIsIdentityAnonymized gets the isIdentityAnonymized property value. The isIdentityAnonymized property
func (m *Participant) GetIsIdentityAnonymized()(*bool) {
    return m.isIdentityAnonymized
}
// GetIsInLobby gets the isInLobby property value. true if the participant is in lobby.
func (m *Participant) GetIsInLobby()(*bool) {
    return m.isInLobby
}
// GetIsMuted gets the isMuted property value. true if the participant is muted (client or server muted).
func (m *Participant) GetIsMuted()(*bool) {
    return m.isMuted
}
// GetMediaStreams gets the mediaStreams property value. The list of media streams.
func (m *Participant) GetMediaStreams()([]MediaStreamable) {
    return m.mediaStreams
}
// GetMetadata gets the metadata property value. A blob of data provided by the participant in the roster.
func (m *Participant) GetMetadata()(*string) {
    return m.metadata
}
// GetRecordingInfo gets the recordingInfo property value. Information on whether the participant has recording capability.
func (m *Participant) GetRecordingInfo()(RecordingInfoable) {
    return m.recordingInfo
}
// Serialize serializes information the current object
func (m *Participant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isIdentityAnonymized", m.GetIsIdentityAnonymized())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInLobby", m.GetIsInLobby())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMuted", m.GetIsMuted())
        if err != nil {
            return err
        }
    }
    if m.GetMediaStreams() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMediaStreams())
        err = writer.WriteCollectionOfObjectValues("mediaStreams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recordingInfo", m.GetRecordingInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInfo sets the info property value. The info property
func (m *Participant) SetInfo(value ParticipantInfoable)() {
    m.info = value
}
// SetIsIdentityAnonymized sets the isIdentityAnonymized property value. The isIdentityAnonymized property
func (m *Participant) SetIsIdentityAnonymized(value *bool)() {
    m.isIdentityAnonymized = value
}
// SetIsInLobby sets the isInLobby property value. true if the participant is in lobby.
func (m *Participant) SetIsInLobby(value *bool)() {
    m.isInLobby = value
}
// SetIsMuted sets the isMuted property value. true if the participant is muted (client or server muted).
func (m *Participant) SetIsMuted(value *bool)() {
    m.isMuted = value
}
// SetMediaStreams sets the mediaStreams property value. The list of media streams.
func (m *Participant) SetMediaStreams(value []MediaStreamable)() {
    m.mediaStreams = value
}
// SetMetadata sets the metadata property value. A blob of data provided by the participant in the roster.
func (m *Participant) SetMetadata(value *string)() {
    m.metadata = value
}
// SetRecordingInfo sets the recordingInfo property value. Information on whether the participant has recording capability.
func (m *Participant) SetRecordingInfo(value RecordingInfoable)() {
    m.recordingInfo = value
}
