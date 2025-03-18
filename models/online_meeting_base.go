package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OnlineMeetingBase struct {
    Entity
}
// NewOnlineMeetingBase instantiates a new OnlineMeetingBase and sets the default values.
func NewOnlineMeetingBase()(*OnlineMeetingBase) {
    m := &OnlineMeetingBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnlineMeetingBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnlineMeetingBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.onlineMeeting":
                        return NewOnlineMeeting(), nil
                    case "#microsoft.graph.virtualEventSession":
                        return NewVirtualEventSession(), nil
                }
            }
        }
    }
    return NewOnlineMeetingBase(), nil
}
// GetAllowAttendeeToEnableCamera gets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowAttendeeToEnableCamera()(*bool) {
    val, err := m.GetBackingStore().Get("allowAttendeeToEnableCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowAttendeeToEnableMic gets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowAttendeeToEnableMic()(*bool) {
    val, err := m.GetBackingStore().Get("allowAttendeeToEnableMic")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowBreakoutRooms gets the allowBreakoutRooms property value. Indicates whether breakout rooms are enabled for the meeting.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowBreakoutRooms()(*bool) {
    val, err := m.GetBackingStore().Get("allowBreakoutRooms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowCopyingAndSharingMeetingContent gets the allowCopyingAndSharingMeetingContent property value. The allowCopyingAndSharingMeetingContent property
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowCopyingAndSharingMeetingContent()(*bool) {
    val, err := m.GetBackingStore().Get("allowCopyingAndSharingMeetingContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowedLobbyAdmitters gets the allowedLobbyAdmitters property value. Specifies the users who can admit from the lobby. Possible values are: organizerAndCoOrganizersAndPresenters, organizerAndCoOrganizers, unknownFutureValue.
// returns a *AllowedLobbyAdmitterRoles when successful
func (m *OnlineMeetingBase) GetAllowedLobbyAdmitters()(*AllowedLobbyAdmitterRoles) {
    val, err := m.GetBackingStore().Get("allowedLobbyAdmitters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AllowedLobbyAdmitterRoles)
    }
    return nil
}
// GetAllowedPresenters gets the allowedPresenters property value. Specifies who can be a presenter in a meeting. Possible values are: everyone, organization, roleIsPresenter, organizer, unknownFutureValue.
// returns a *OnlineMeetingPresenters when successful
func (m *OnlineMeetingBase) GetAllowedPresenters()(*OnlineMeetingPresenters) {
    val, err := m.GetBackingStore().Get("allowedPresenters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OnlineMeetingPresenters)
    }
    return nil
}
// GetAllowLiveShare gets the allowLiveShare property value. Indicates whether live share is enabled for the meeting. Possible values are: enabled, disabled, unknownFutureValue.
// returns a *MeetingLiveShareOptions when successful
func (m *OnlineMeetingBase) GetAllowLiveShare()(*MeetingLiveShareOptions) {
    val, err := m.GetBackingStore().Get("allowLiveShare")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MeetingLiveShareOptions)
    }
    return nil
}
// GetAllowMeetingChat gets the allowMeetingChat property value. Specifies the mode of meeting chat. Possible values are: enabled, disabled, limited, unknownFutureValue.
// returns a *MeetingChatMode when successful
func (m *OnlineMeetingBase) GetAllowMeetingChat()(*MeetingChatMode) {
    val, err := m.GetBackingStore().Get("allowMeetingChat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MeetingChatMode)
    }
    return nil
}
// GetAllowParticipantsToChangeName gets the allowParticipantsToChangeName property value. Specifies if participants are allowed to rename themselves in an instance of the meeting.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowParticipantsToChangeName()(*bool) {
    val, err := m.GetBackingStore().Get("allowParticipantsToChangeName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowPowerPointSharing gets the allowPowerPointSharing property value. Indicates whether PowerPoint live is enabled for the meeting.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowPowerPointSharing()(*bool) {
    val, err := m.GetBackingStore().Get("allowPowerPointSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowRecording gets the allowRecording property value. Indicates whether recording is enabled for the meeting.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowRecording()(*bool) {
    val, err := m.GetBackingStore().Get("allowRecording")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowTeamworkReactions gets the allowTeamworkReactions property value. Indicates if Teams reactions are enabled for the meeting.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowTeamworkReactions()(*bool) {
    val, err := m.GetBackingStore().Get("allowTeamworkReactions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowTranscription gets the allowTranscription property value. Indicates whether transcription is enabled for the meeting.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowTranscription()(*bool) {
    val, err := m.GetBackingStore().Get("allowTranscription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowWhiteboard gets the allowWhiteboard property value. Indicates whether whiteboard is enabled for the meeting.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetAllowWhiteboard()(*bool) {
    val, err := m.GetBackingStore().Get("allowWhiteboard")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAnonymizeIdentityForRoles gets the anonymizeIdentityForRoles property value. Specifies whose identity is anonymized in the meeting. Possible values are: attendee. The attendee value can't be removed through a PATCH operation once added.
// returns a []OnlineMeetingRole when successful
func (m *OnlineMeetingBase) GetAnonymizeIdentityForRoles()([]OnlineMeetingRole) {
    val, err := m.GetBackingStore().Get("anonymizeIdentityForRoles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OnlineMeetingRole)
    }
    return nil
}
// GetAttendanceReports gets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
// returns a []MeetingAttendanceReportable when successful
func (m *OnlineMeetingBase) GetAttendanceReports()([]MeetingAttendanceReportable) {
    val, err := m.GetBackingStore().Get("attendanceReports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MeetingAttendanceReportable)
    }
    return nil
}
// GetAudioConferencing gets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
// returns a AudioConferencingable when successful
func (m *OnlineMeetingBase) GetAudioConferencing()(AudioConferencingable) {
    val, err := m.GetBackingStore().Get("audioConferencing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AudioConferencingable)
    }
    return nil
}
// GetChatInfo gets the chatInfo property value. The chat information associated with this online meeting.
// returns a ChatInfoable when successful
func (m *OnlineMeetingBase) GetChatInfo()(ChatInfoable) {
    val, err := m.GetBackingStore().Get("chatInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ChatInfoable)
    }
    return nil
}
// GetChatRestrictions gets the chatRestrictions property value. Specifies the configuration settings for meeting chat restrictions.
// returns a ChatRestrictionsable when successful
func (m *OnlineMeetingBase) GetChatRestrictions()(ChatRestrictionsable) {
    val, err := m.GetBackingStore().Get("chatRestrictions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ChatRestrictionsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnlineMeetingBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowAttendeeToEnableCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableCamera(val)
        }
        return nil
    }
    res["allowAttendeeToEnableMic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableMic(val)
        }
        return nil
    }
    res["allowBreakoutRooms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowBreakoutRooms(val)
        }
        return nil
    }
    res["allowCopyingAndSharingMeetingContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCopyingAndSharingMeetingContent(val)
        }
        return nil
    }
    res["allowedLobbyAdmitters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowedLobbyAdmitterRoles)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedLobbyAdmitters(val.(*AllowedLobbyAdmitterRoles))
        }
        return nil
    }
    res["allowedPresenters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingPresenters)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedPresenters(val.(*OnlineMeetingPresenters))
        }
        return nil
    }
    res["allowLiveShare"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingLiveShareOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLiveShare(val.(*MeetingLiveShareOptions))
        }
        return nil
    }
    res["allowMeetingChat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingChatMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMeetingChat(val.(*MeetingChatMode))
        }
        return nil
    }
    res["allowParticipantsToChangeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowParticipantsToChangeName(val)
        }
        return nil
    }
    res["allowPowerPointSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPowerPointSharing(val)
        }
        return nil
    }
    res["allowRecording"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowRecording(val)
        }
        return nil
    }
    res["allowTeamworkReactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTeamworkReactions(val)
        }
        return nil
    }
    res["allowTranscription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTranscription(val)
        }
        return nil
    }
    res["allowWhiteboard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowWhiteboard(val)
        }
        return nil
    }
    res["anonymizeIdentityForRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOnlineMeetingRole)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnlineMeetingRole, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*OnlineMeetingRole))
                }
            }
            m.SetAnonymizeIdentityForRoles(res)
        }
        return nil
    }
    res["attendanceReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingAttendanceReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingAttendanceReportable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MeetingAttendanceReportable)
                }
            }
            m.SetAttendanceReports(res)
        }
        return nil
    }
    res["audioConferencing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAudioConferencingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioConferencing(val.(AudioConferencingable))
        }
        return nil
    }
    res["chatInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(ChatInfoable))
        }
        return nil
    }
    res["chatRestrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatRestrictionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatRestrictions(val.(ChatRestrictionsable))
        }
        return nil
    }
    res["isEndToEndEncryptionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEndToEndEncryptionEnabled(val)
        }
        return nil
    }
    res["isEntryExitAnnounced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEntryExitAnnounced(val)
        }
        return nil
    }
    res["joinInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinInformation(val.(ItemBodyable))
        }
        return nil
    }
    res["joinMeetingIdSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJoinMeetingIdSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinMeetingIdSettings(val.(JoinMeetingIdSettingsable))
        }
        return nil
    }
    res["joinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
        }
        return nil
    }
    res["lobbyBypassSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLobbyBypassSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLobbyBypassSettings(val.(LobbyBypassSettingsable))
        }
        return nil
    }
    res["recordAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordAutomatically(val)
        }
        return nil
    }
    res["shareMeetingChatHistoryDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingChatHistoryDefaultMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareMeetingChatHistoryDefault(val.(*MeetingChatHistoryDefaultMode))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["videoTeleconferenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoTeleconferenceId(val)
        }
        return nil
    }
    res["watermarkProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWatermarkProtectionValuesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatermarkProtection(val.(WatermarkProtectionValuesable))
        }
        return nil
    }
    return res
}
// GetIsEndToEndEncryptionEnabled gets the isEndToEndEncryptionEnabled property value. Indicates whether end-to-end encryption (E2EE) is enabled for the online meeting.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetIsEndToEndEncryptionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEndToEndEncryptionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsEntryExitAnnounced gets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetIsEntryExitAnnounced()(*bool) {
    val, err := m.GetBackingStore().Get("isEntryExitAnnounced")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetJoinInformation gets the joinInformation property value. The join information in the language and locale variant specified in 'Accept-Language' request HTTP header. Read-only.
// returns a ItemBodyable when successful
func (m *OnlineMeetingBase) GetJoinInformation()(ItemBodyable) {
    val, err := m.GetBackingStore().Get("joinInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemBodyable)
    }
    return nil
}
// GetJoinMeetingIdSettings gets the joinMeetingIdSettings property value. Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is created, the joinMeetingIdSettings can't be modified. To make any changes to this property, the meeting needs to be canceled and a new one needs to be created.
// returns a JoinMeetingIdSettingsable when successful
func (m *OnlineMeetingBase) GetJoinMeetingIdSettings()(JoinMeetingIdSettingsable) {
    val, err := m.GetBackingStore().Get("joinMeetingIdSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(JoinMeetingIdSettingsable)
    }
    return nil
}
// GetJoinWebUrl gets the joinWebUrl property value. The join URL of the online meeting. Read-only.
// returns a *string when successful
func (m *OnlineMeetingBase) GetJoinWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("joinWebUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLobbyBypassSettings gets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting lobby.
// returns a LobbyBypassSettingsable when successful
func (m *OnlineMeetingBase) GetLobbyBypassSettings()(LobbyBypassSettingsable) {
    val, err := m.GetBackingStore().Get("lobbyBypassSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LobbyBypassSettingsable)
    }
    return nil
}
// GetRecordAutomatically gets the recordAutomatically property value. Indicates whether to record the meeting automatically.
// returns a *bool when successful
func (m *OnlineMeetingBase) GetRecordAutomatically()(*bool) {
    val, err := m.GetBackingStore().Get("recordAutomatically")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShareMeetingChatHistoryDefault gets the shareMeetingChatHistoryDefault property value. Specifies whether meeting chat history is shared with participants.  Possible values are: all, none, unknownFutureValue.
// returns a *MeetingChatHistoryDefaultMode when successful
func (m *OnlineMeetingBase) GetShareMeetingChatHistoryDefault()(*MeetingChatHistoryDefaultMode) {
    val, err := m.GetBackingStore().Get("shareMeetingChatHistoryDefault")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MeetingChatHistoryDefaultMode)
    }
    return nil
}
// GetSubject gets the subject property value. The subject of the online meeting.
// returns a *string when successful
func (m *OnlineMeetingBase) GetSubject()(*string) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVideoTeleconferenceId gets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
// returns a *string when successful
func (m *OnlineMeetingBase) GetVideoTeleconferenceId()(*string) {
    val, err := m.GetBackingStore().Get("videoTeleconferenceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWatermarkProtection gets the watermarkProtection property value. Specifies whether the client application should apply a watermark to a content type.
// returns a WatermarkProtectionValuesable when successful
func (m *OnlineMeetingBase) GetWatermarkProtection()(WatermarkProtectionValuesable) {
    val, err := m.GetBackingStore().Get("watermarkProtection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WatermarkProtectionValuesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnlineMeetingBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAttendeeToEnableCamera", m.GetAllowAttendeeToEnableCamera())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowAttendeeToEnableMic", m.GetAllowAttendeeToEnableMic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowBreakoutRooms", m.GetAllowBreakoutRooms())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowCopyingAndSharingMeetingContent", m.GetAllowCopyingAndSharingMeetingContent())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedLobbyAdmitters() != nil {
        cast := (*m.GetAllowedLobbyAdmitters()).String()
        err = writer.WriteStringValue("allowedLobbyAdmitters", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedPresenters() != nil {
        cast := (*m.GetAllowedPresenters()).String()
        err = writer.WriteStringValue("allowedPresenters", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowLiveShare() != nil {
        cast := (*m.GetAllowLiveShare()).String()
        err = writer.WriteStringValue("allowLiveShare", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowMeetingChat() != nil {
        cast := (*m.GetAllowMeetingChat()).String()
        err = writer.WriteStringValue("allowMeetingChat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowParticipantsToChangeName", m.GetAllowParticipantsToChangeName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowPowerPointSharing", m.GetAllowPowerPointSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowRecording", m.GetAllowRecording())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowTeamworkReactions", m.GetAllowTeamworkReactions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowTranscription", m.GetAllowTranscription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowWhiteboard", m.GetAllowWhiteboard())
        if err != nil {
            return err
        }
    }
    if m.GetAnonymizeIdentityForRoles() != nil {
        err = writer.WriteCollectionOfStringValues("anonymizeIdentityForRoles", SerializeOnlineMeetingRole(m.GetAnonymizeIdentityForRoles()))
        if err != nil {
            return err
        }
    }
    if m.GetAttendanceReports() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendanceReports()))
        for i, v := range m.GetAttendanceReports() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("attendanceReports", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("audioConferencing", m.GetAudioConferencing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("chatInfo", m.GetChatInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("chatRestrictions", m.GetChatRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEndToEndEncryptionEnabled", m.GetIsEndToEndEncryptionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEntryExitAnnounced", m.GetIsEntryExitAnnounced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("joinInformation", m.GetJoinInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("joinMeetingIdSettings", m.GetJoinMeetingIdSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lobbyBypassSettings", m.GetLobbyBypassSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("recordAutomatically", m.GetRecordAutomatically())
        if err != nil {
            return err
        }
    }
    if m.GetShareMeetingChatHistoryDefault() != nil {
        cast := (*m.GetShareMeetingChatHistoryDefault()).String()
        err = writer.WriteStringValue("shareMeetingChatHistoryDefault", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("videoTeleconferenceId", m.GetVideoTeleconferenceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("watermarkProtection", m.GetWatermarkProtection())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAttendeeToEnableCamera sets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeetingBase) SetAllowAttendeeToEnableCamera(value *bool)() {
    err := m.GetBackingStore().Set("allowAttendeeToEnableCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowAttendeeToEnableMic sets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeetingBase) SetAllowAttendeeToEnableMic(value *bool)() {
    err := m.GetBackingStore().Set("allowAttendeeToEnableMic", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowBreakoutRooms sets the allowBreakoutRooms property value. Indicates whether breakout rooms are enabled for the meeting.
func (m *OnlineMeetingBase) SetAllowBreakoutRooms(value *bool)() {
    err := m.GetBackingStore().Set("allowBreakoutRooms", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowCopyingAndSharingMeetingContent sets the allowCopyingAndSharingMeetingContent property value. The allowCopyingAndSharingMeetingContent property
func (m *OnlineMeetingBase) SetAllowCopyingAndSharingMeetingContent(value *bool)() {
    err := m.GetBackingStore().Set("allowCopyingAndSharingMeetingContent", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedLobbyAdmitters sets the allowedLobbyAdmitters property value. Specifies the users who can admit from the lobby. Possible values are: organizerAndCoOrganizersAndPresenters, organizerAndCoOrganizers, unknownFutureValue.
func (m *OnlineMeetingBase) SetAllowedLobbyAdmitters(value *AllowedLobbyAdmitterRoles)() {
    err := m.GetBackingStore().Set("allowedLobbyAdmitters", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedPresenters sets the allowedPresenters property value. Specifies who can be a presenter in a meeting. Possible values are: everyone, organization, roleIsPresenter, organizer, unknownFutureValue.
func (m *OnlineMeetingBase) SetAllowedPresenters(value *OnlineMeetingPresenters)() {
    err := m.GetBackingStore().Set("allowedPresenters", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowLiveShare sets the allowLiveShare property value. Indicates whether live share is enabled for the meeting. Possible values are: enabled, disabled, unknownFutureValue.
func (m *OnlineMeetingBase) SetAllowLiveShare(value *MeetingLiveShareOptions)() {
    err := m.GetBackingStore().Set("allowLiveShare", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowMeetingChat sets the allowMeetingChat property value. Specifies the mode of meeting chat. Possible values are: enabled, disabled, limited, unknownFutureValue.
func (m *OnlineMeetingBase) SetAllowMeetingChat(value *MeetingChatMode)() {
    err := m.GetBackingStore().Set("allowMeetingChat", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowParticipantsToChangeName sets the allowParticipantsToChangeName property value. Specifies if participants are allowed to rename themselves in an instance of the meeting.
func (m *OnlineMeetingBase) SetAllowParticipantsToChangeName(value *bool)() {
    err := m.GetBackingStore().Set("allowParticipantsToChangeName", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowPowerPointSharing sets the allowPowerPointSharing property value. Indicates whether PowerPoint live is enabled for the meeting.
func (m *OnlineMeetingBase) SetAllowPowerPointSharing(value *bool)() {
    err := m.GetBackingStore().Set("allowPowerPointSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowRecording sets the allowRecording property value. Indicates whether recording is enabled for the meeting.
func (m *OnlineMeetingBase) SetAllowRecording(value *bool)() {
    err := m.GetBackingStore().Set("allowRecording", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowTeamworkReactions sets the allowTeamworkReactions property value. Indicates if Teams reactions are enabled for the meeting.
func (m *OnlineMeetingBase) SetAllowTeamworkReactions(value *bool)() {
    err := m.GetBackingStore().Set("allowTeamworkReactions", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowTranscription sets the allowTranscription property value. Indicates whether transcription is enabled for the meeting.
func (m *OnlineMeetingBase) SetAllowTranscription(value *bool)() {
    err := m.GetBackingStore().Set("allowTranscription", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowWhiteboard sets the allowWhiteboard property value. Indicates whether whiteboard is enabled for the meeting.
func (m *OnlineMeetingBase) SetAllowWhiteboard(value *bool)() {
    err := m.GetBackingStore().Set("allowWhiteboard", value)
    if err != nil {
        panic(err)
    }
}
// SetAnonymizeIdentityForRoles sets the anonymizeIdentityForRoles property value. Specifies whose identity is anonymized in the meeting. Possible values are: attendee. The attendee value can't be removed through a PATCH operation once added.
func (m *OnlineMeetingBase) SetAnonymizeIdentityForRoles(value []OnlineMeetingRole)() {
    err := m.GetBackingStore().Set("anonymizeIdentityForRoles", value)
    if err != nil {
        panic(err)
    }
}
// SetAttendanceReports sets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeetingBase) SetAttendanceReports(value []MeetingAttendanceReportable)() {
    err := m.GetBackingStore().Set("attendanceReports", value)
    if err != nil {
        panic(err)
    }
}
// SetAudioConferencing sets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeetingBase) SetAudioConferencing(value AudioConferencingable)() {
    err := m.GetBackingStore().Set("audioConferencing", value)
    if err != nil {
        panic(err)
    }
}
// SetChatInfo sets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeetingBase) SetChatInfo(value ChatInfoable)() {
    err := m.GetBackingStore().Set("chatInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetChatRestrictions sets the chatRestrictions property value. Specifies the configuration settings for meeting chat restrictions.
func (m *OnlineMeetingBase) SetChatRestrictions(value ChatRestrictionsable)() {
    err := m.GetBackingStore().Set("chatRestrictions", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEndToEndEncryptionEnabled sets the isEndToEndEncryptionEnabled property value. Indicates whether end-to-end encryption (E2EE) is enabled for the online meeting.
func (m *OnlineMeetingBase) SetIsEndToEndEncryptionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEndToEndEncryptionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEntryExitAnnounced sets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeetingBase) SetIsEntryExitAnnounced(value *bool)() {
    err := m.GetBackingStore().Set("isEntryExitAnnounced", value)
    if err != nil {
        panic(err)
    }
}
// SetJoinInformation sets the joinInformation property value. The join information in the language and locale variant specified in 'Accept-Language' request HTTP header. Read-only.
func (m *OnlineMeetingBase) SetJoinInformation(value ItemBodyable)() {
    err := m.GetBackingStore().Set("joinInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetJoinMeetingIdSettings sets the joinMeetingIdSettings property value. Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is created, the joinMeetingIdSettings can't be modified. To make any changes to this property, the meeting needs to be canceled and a new one needs to be created.
func (m *OnlineMeetingBase) SetJoinMeetingIdSettings(value JoinMeetingIdSettingsable)() {
    err := m.GetBackingStore().Set("joinMeetingIdSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetJoinWebUrl sets the joinWebUrl property value. The join URL of the online meeting. Read-only.
func (m *OnlineMeetingBase) SetJoinWebUrl(value *string)() {
    err := m.GetBackingStore().Set("joinWebUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetLobbyBypassSettings sets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting lobby.
func (m *OnlineMeetingBase) SetLobbyBypassSettings(value LobbyBypassSettingsable)() {
    err := m.GetBackingStore().Set("lobbyBypassSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetRecordAutomatically sets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeetingBase) SetRecordAutomatically(value *bool)() {
    err := m.GetBackingStore().Set("recordAutomatically", value)
    if err != nil {
        panic(err)
    }
}
// SetShareMeetingChatHistoryDefault sets the shareMeetingChatHistoryDefault property value. Specifies whether meeting chat history is shared with participants.  Possible values are: all, none, unknownFutureValue.
func (m *OnlineMeetingBase) SetShareMeetingChatHistoryDefault(value *MeetingChatHistoryDefaultMode)() {
    err := m.GetBackingStore().Set("shareMeetingChatHistoryDefault", value)
    if err != nil {
        panic(err)
    }
}
// SetSubject sets the subject property value. The subject of the online meeting.
func (m *OnlineMeetingBase) SetSubject(value *string)() {
    err := m.GetBackingStore().Set("subject", value)
    if err != nil {
        panic(err)
    }
}
// SetVideoTeleconferenceId sets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeetingBase) SetVideoTeleconferenceId(value *string)() {
    err := m.GetBackingStore().Set("videoTeleconferenceId", value)
    if err != nil {
        panic(err)
    }
}
// SetWatermarkProtection sets the watermarkProtection property value. Specifies whether the client application should apply a watermark to a content type.
func (m *OnlineMeetingBase) SetWatermarkProtection(value WatermarkProtectionValuesable)() {
    err := m.GetBackingStore().Set("watermarkProtection", value)
    if err != nil {
        panic(err)
    }
}
type OnlineMeetingBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAttendeeToEnableCamera()(*bool)
    GetAllowAttendeeToEnableMic()(*bool)
    GetAllowBreakoutRooms()(*bool)
    GetAllowCopyingAndSharingMeetingContent()(*bool)
    GetAllowedLobbyAdmitters()(*AllowedLobbyAdmitterRoles)
    GetAllowedPresenters()(*OnlineMeetingPresenters)
    GetAllowLiveShare()(*MeetingLiveShareOptions)
    GetAllowMeetingChat()(*MeetingChatMode)
    GetAllowParticipantsToChangeName()(*bool)
    GetAllowPowerPointSharing()(*bool)
    GetAllowRecording()(*bool)
    GetAllowTeamworkReactions()(*bool)
    GetAllowTranscription()(*bool)
    GetAllowWhiteboard()(*bool)
    GetAnonymizeIdentityForRoles()([]OnlineMeetingRole)
    GetAttendanceReports()([]MeetingAttendanceReportable)
    GetAudioConferencing()(AudioConferencingable)
    GetChatInfo()(ChatInfoable)
    GetChatRestrictions()(ChatRestrictionsable)
    GetIsEndToEndEncryptionEnabled()(*bool)
    GetIsEntryExitAnnounced()(*bool)
    GetJoinInformation()(ItemBodyable)
    GetJoinMeetingIdSettings()(JoinMeetingIdSettingsable)
    GetJoinWebUrl()(*string)
    GetLobbyBypassSettings()(LobbyBypassSettingsable)
    GetRecordAutomatically()(*bool)
    GetShareMeetingChatHistoryDefault()(*MeetingChatHistoryDefaultMode)
    GetSubject()(*string)
    GetVideoTeleconferenceId()(*string)
    GetWatermarkProtection()(WatermarkProtectionValuesable)
    SetAllowAttendeeToEnableCamera(value *bool)()
    SetAllowAttendeeToEnableMic(value *bool)()
    SetAllowBreakoutRooms(value *bool)()
    SetAllowCopyingAndSharingMeetingContent(value *bool)()
    SetAllowedLobbyAdmitters(value *AllowedLobbyAdmitterRoles)()
    SetAllowedPresenters(value *OnlineMeetingPresenters)()
    SetAllowLiveShare(value *MeetingLiveShareOptions)()
    SetAllowMeetingChat(value *MeetingChatMode)()
    SetAllowParticipantsToChangeName(value *bool)()
    SetAllowPowerPointSharing(value *bool)()
    SetAllowRecording(value *bool)()
    SetAllowTeamworkReactions(value *bool)()
    SetAllowTranscription(value *bool)()
    SetAllowWhiteboard(value *bool)()
    SetAnonymizeIdentityForRoles(value []OnlineMeetingRole)()
    SetAttendanceReports(value []MeetingAttendanceReportable)()
    SetAudioConferencing(value AudioConferencingable)()
    SetChatInfo(value ChatInfoable)()
    SetChatRestrictions(value ChatRestrictionsable)()
    SetIsEndToEndEncryptionEnabled(value *bool)()
    SetIsEntryExitAnnounced(value *bool)()
    SetJoinInformation(value ItemBodyable)()
    SetJoinMeetingIdSettings(value JoinMeetingIdSettingsable)()
    SetJoinWebUrl(value *string)()
    SetLobbyBypassSettings(value LobbyBypassSettingsable)()
    SetRecordAutomatically(value *bool)()
    SetShareMeetingChatHistoryDefault(value *MeetingChatHistoryDefaultMode)()
    SetSubject(value *string)()
    SetVideoTeleconferenceId(value *string)()
    SetWatermarkProtection(value WatermarkProtectionValuesable)()
}
