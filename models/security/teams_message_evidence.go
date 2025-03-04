package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamsMessageEvidence struct {
    AlertEvidence
}
// NewTeamsMessageEvidence instantiates a new TeamsMessageEvidence and sets the default values.
func NewTeamsMessageEvidence()(*TeamsMessageEvidence) {
    m := &TeamsMessageEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.teamsMessageEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeamsMessageEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamsMessageEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsMessageEvidence(), nil
}
// GetCampaignId gets the campaignId property value. The identifier of the campaign that this Teams message is part of.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetCampaignId()(*string) {
    val, err := m.GetBackingStore().Get("campaignId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetChannelId gets the channelId property value. The channel ID associated with this Teams message.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetChannelId()(*string) {
    val, err := m.GetBackingStore().Get("channelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeliveryAction gets the deliveryAction property value. The delivery action of this Teams message. Possible values are: unknown, deliveredAsSpam, delivered, blocked, replaced, unknownFutureValue.
// returns a *TeamsMessageDeliveryAction when successful
func (m *TeamsMessageEvidence) GetDeliveryAction()(*TeamsMessageDeliveryAction) {
    val, err := m.GetBackingStore().Get("deliveryAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamsMessageDeliveryAction)
    }
    return nil
}
// GetDeliveryLocation gets the deliveryLocation property value. The delivery location of this Teams message. Possible values are: unknown, teams, quarantine, failed, unknownFutureValue.
// returns a *TeamsDeliveryLocation when successful
func (m *TeamsMessageEvidence) GetDeliveryLocation()(*TeamsDeliveryLocation) {
    val, err := m.GetBackingStore().Get("deliveryLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamsDeliveryLocation)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamsMessageEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["campaignId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCampaignId(val)
        }
        return nil
    }
    res["channelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelId(val)
        }
        return nil
    }
    res["deliveryAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsMessageDeliveryAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryAction(val.(*TeamsMessageDeliveryAction))
        }
        return nil
    }
    res["deliveryLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsDeliveryLocation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryLocation(val.(*TeamsDeliveryLocation))
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileEvidenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FileEvidenceable)
                }
            }
            m.SetFiles(res)
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["isExternal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternal(val)
        }
        return nil
    }
    res["isOwned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOwned(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["messageDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAntispamTeamsDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageDirection(val.(*AntispamTeamsDirection))
        }
        return nil
    }
    res["messageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["owningTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwningTenantId(val)
        }
        return nil
    }
    res["parentMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentMessageId(val)
        }
        return nil
    }
    res["receivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRecipients(res)
        }
        return nil
    }
    res["senderFromAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderFromAddress(val)
        }
        return nil
    }
    res["senderIP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderIP(val)
        }
        return nil
    }
    res["sourceAppName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceAppName(val)
        }
        return nil
    }
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
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
    res["suspiciousRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSuspiciousRecipients(res)
        }
        return nil
    }
    res["threadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreadId(val)
        }
        return nil
    }
    res["threadType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreadType(val)
        }
        return nil
    }
    res["urls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUrlEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UrlEvidenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UrlEvidenceable)
                }
            }
            m.SetUrls(res)
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. The list of file entities that are attached to this Teams message.
// returns a []FileEvidenceable when successful
func (m *TeamsMessageEvidence) GetFiles()([]FileEvidenceable) {
    val, err := m.GetBackingStore().Get("files")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FileEvidenceable)
    }
    return nil
}
// GetGroupId gets the groupId property value. The identifier of the team or group that this message is part of.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("groupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsExternal gets the isExternal property value. Indicates whether the message is owned by the organization that reported the security detection alert.
// returns a *bool when successful
func (m *TeamsMessageEvidence) GetIsExternal()(*bool) {
    val, err := m.GetBackingStore().Get("isExternal")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsOwned gets the isOwned property value. Indicates whether the message is owned by your organization.
// returns a *bool when successful
func (m *TeamsMessageEvidence) GetIsOwned()(*bool) {
    val, err := m.GetBackingStore().Get("isOwned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time when the message was last edited. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *TeamsMessageEvidence) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMessageDirection gets the messageDirection property value. The direction of the Teams message. The possible values are: unknown, inbound, outbound, intraorg, unknownFutureValue.
// returns a *AntispamTeamsDirection when successful
func (m *TeamsMessageEvidence) GetMessageDirection()(*AntispamTeamsDirection) {
    val, err := m.GetBackingStore().Get("messageDirection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AntispamTeamsDirection)
    }
    return nil
}
// GetMessageId gets the messageId property value. The message identifier, unique within the thread.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetMessageId()(*string) {
    val, err := m.GetBackingStore().Get("messageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwningTenantId gets the owningTenantId property value. Tenant ID (GUID) of the owner of the message.
// returns a *UUID when successful
func (m *TeamsMessageEvidence) GetOwningTenantId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("owningTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetParentMessageId gets the parentMessageId property value. Identifier of the message to which the current message is a reply; otherwise, it's the same as the messageId.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetParentMessageId()(*string) {
    val, err := m.GetBackingStore().Get("parentMessageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReceivedDateTime gets the receivedDateTime property value. The received date of this message. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *TeamsMessageEvidence) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("receivedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRecipients gets the recipients property value. The recipients of this Teams message.
// returns a []string when successful
func (m *TeamsMessageEvidence) GetRecipients()([]string) {
    val, err := m.GetBackingStore().Get("recipients")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSenderFromAddress gets the senderFromAddress property value. The SMTP format address of the sender.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetSenderFromAddress()(*string) {
    val, err := m.GetBackingStore().Get("senderFromAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSenderIP gets the senderIP property value. The IP address of the sender.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetSenderIP()(*string) {
    val, err := m.GetBackingStore().Get("senderIP")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSourceAppName gets the sourceAppName property value. Source of the message; for example, desktop and mobile.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetSourceAppName()(*string) {
    val, err := m.GetBackingStore().Get("sourceAppName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSourceId gets the sourceId property value. The source ID of this Teams message.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetSourceId()(*string) {
    val, err := m.GetBackingStore().Get("sourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubject gets the subject property value. The subject of this Teams message.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetSubject()(*string) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSuspiciousRecipients gets the suspiciousRecipients property value. The list of recipients who were detected as suspicious.
// returns a []string when successful
func (m *TeamsMessageEvidence) GetSuspiciousRecipients()([]string) {
    val, err := m.GetBackingStore().Get("suspiciousRecipients")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetThreadId gets the threadId property value. Identifier of the channel or chat that this message is part of.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetThreadId()(*string) {
    val, err := m.GetBackingStore().Get("threadId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThreadType gets the threadType property value. The Teams message type. Supported values are: Chat, Topic, Space, and Meeting.
// returns a *string when successful
func (m *TeamsMessageEvidence) GetThreadType()(*string) {
    val, err := m.GetBackingStore().Get("threadType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUrls gets the urls property value. The URLs contained in this Teams message.
// returns a []UrlEvidenceable when successful
func (m *TeamsMessageEvidence) GetUrls()([]UrlEvidenceable) {
    val, err := m.GetBackingStore().Get("urls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UrlEvidenceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamsMessageEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("campaignId", m.GetCampaignId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("channelId", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    if m.GetDeliveryAction() != nil {
        cast := (*m.GetDeliveryAction()).String()
        err = writer.WriteStringValue("deliveryAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeliveryLocation() != nil {
        cast := (*m.GetDeliveryLocation()).String()
        err = writer.WriteStringValue("deliveryLocation", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("files", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isExternal", m.GetIsExternal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOwned", m.GetIsOwned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMessageDirection() != nil {
        cast := (*m.GetMessageDirection()).String()
        err = writer.WriteStringValue("messageDirection", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("messageId", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("owningTenantId", m.GetOwningTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentMessageId", m.GetParentMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        err = writer.WriteCollectionOfStringValues("recipients", m.GetRecipients())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderFromAddress", m.GetSenderFromAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderIP", m.GetSenderIP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceAppName", m.GetSourceAppName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceId", m.GetSourceId())
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
    if m.GetSuspiciousRecipients() != nil {
        err = writer.WriteCollectionOfStringValues("suspiciousRecipients", m.GetSuspiciousRecipients())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("threadId", m.GetThreadId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("threadType", m.GetThreadType())
        if err != nil {
            return err
        }
    }
    if m.GetUrls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUrls()))
        for i, v := range m.GetUrls() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("urls", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCampaignId sets the campaignId property value. The identifier of the campaign that this Teams message is part of.
func (m *TeamsMessageEvidence) SetCampaignId(value *string)() {
    err := m.GetBackingStore().Set("campaignId", value)
    if err != nil {
        panic(err)
    }
}
// SetChannelId sets the channelId property value. The channel ID associated with this Teams message.
func (m *TeamsMessageEvidence) SetChannelId(value *string)() {
    err := m.GetBackingStore().Set("channelId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeliveryAction sets the deliveryAction property value. The delivery action of this Teams message. Possible values are: unknown, deliveredAsSpam, delivered, blocked, replaced, unknownFutureValue.
func (m *TeamsMessageEvidence) SetDeliveryAction(value *TeamsMessageDeliveryAction)() {
    err := m.GetBackingStore().Set("deliveryAction", value)
    if err != nil {
        panic(err)
    }
}
// SetDeliveryLocation sets the deliveryLocation property value. The delivery location of this Teams message. Possible values are: unknown, teams, quarantine, failed, unknownFutureValue.
func (m *TeamsMessageEvidence) SetDeliveryLocation(value *TeamsDeliveryLocation)() {
    err := m.GetBackingStore().Set("deliveryLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetFiles sets the files property value. The list of file entities that are attached to this Teams message.
func (m *TeamsMessageEvidence) SetFiles(value []FileEvidenceable)() {
    err := m.GetBackingStore().Set("files", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupId sets the groupId property value. The identifier of the team or group that this message is part of.
func (m *TeamsMessageEvidence) SetGroupId(value *string)() {
    err := m.GetBackingStore().Set("groupId", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExternal sets the isExternal property value. Indicates whether the message is owned by the organization that reported the security detection alert.
func (m *TeamsMessageEvidence) SetIsExternal(value *bool)() {
    err := m.GetBackingStore().Set("isExternal", value)
    if err != nil {
        panic(err)
    }
}
// SetIsOwned sets the isOwned property value. Indicates whether the message is owned by your organization.
func (m *TeamsMessageEvidence) SetIsOwned(value *bool)() {
    err := m.GetBackingStore().Set("isOwned", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time when the message was last edited. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *TeamsMessageEvidence) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMessageDirection sets the messageDirection property value. The direction of the Teams message. The possible values are: unknown, inbound, outbound, intraorg, unknownFutureValue.
func (m *TeamsMessageEvidence) SetMessageDirection(value *AntispamTeamsDirection)() {
    err := m.GetBackingStore().Set("messageDirection", value)
    if err != nil {
        panic(err)
    }
}
// SetMessageId sets the messageId property value. The message identifier, unique within the thread.
func (m *TeamsMessageEvidence) SetMessageId(value *string)() {
    err := m.GetBackingStore().Set("messageId", value)
    if err != nil {
        panic(err)
    }
}
// SetOwningTenantId sets the owningTenantId property value. Tenant ID (GUID) of the owner of the message.
func (m *TeamsMessageEvidence) SetOwningTenantId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("owningTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetParentMessageId sets the parentMessageId property value. Identifier of the message to which the current message is a reply; otherwise, it's the same as the messageId.
func (m *TeamsMessageEvidence) SetParentMessageId(value *string)() {
    err := m.GetBackingStore().Set("parentMessageId", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivedDateTime sets the receivedDateTime property value. The received date of this message. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *TeamsMessageEvidence) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("receivedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRecipients sets the recipients property value. The recipients of this Teams message.
func (m *TeamsMessageEvidence) SetRecipients(value []string)() {
    err := m.GetBackingStore().Set("recipients", value)
    if err != nil {
        panic(err)
    }
}
// SetSenderFromAddress sets the senderFromAddress property value. The SMTP format address of the sender.
func (m *TeamsMessageEvidence) SetSenderFromAddress(value *string)() {
    err := m.GetBackingStore().Set("senderFromAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetSenderIP sets the senderIP property value. The IP address of the sender.
func (m *TeamsMessageEvidence) SetSenderIP(value *string)() {
    err := m.GetBackingStore().Set("senderIP", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceAppName sets the sourceAppName property value. Source of the message; for example, desktop and mobile.
func (m *TeamsMessageEvidence) SetSourceAppName(value *string)() {
    err := m.GetBackingStore().Set("sourceAppName", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceId sets the sourceId property value. The source ID of this Teams message.
func (m *TeamsMessageEvidence) SetSourceId(value *string)() {
    err := m.GetBackingStore().Set("sourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetSubject sets the subject property value. The subject of this Teams message.
func (m *TeamsMessageEvidence) SetSubject(value *string)() {
    err := m.GetBackingStore().Set("subject", value)
    if err != nil {
        panic(err)
    }
}
// SetSuspiciousRecipients sets the suspiciousRecipients property value. The list of recipients who were detected as suspicious.
func (m *TeamsMessageEvidence) SetSuspiciousRecipients(value []string)() {
    err := m.GetBackingStore().Set("suspiciousRecipients", value)
    if err != nil {
        panic(err)
    }
}
// SetThreadId sets the threadId property value. Identifier of the channel or chat that this message is part of.
func (m *TeamsMessageEvidence) SetThreadId(value *string)() {
    err := m.GetBackingStore().Set("threadId", value)
    if err != nil {
        panic(err)
    }
}
// SetThreadType sets the threadType property value. The Teams message type. Supported values are: Chat, Topic, Space, and Meeting.
func (m *TeamsMessageEvidence) SetThreadType(value *string)() {
    err := m.GetBackingStore().Set("threadType", value)
    if err != nil {
        panic(err)
    }
}
// SetUrls sets the urls property value. The URLs contained in this Teams message.
func (m *TeamsMessageEvidence) SetUrls(value []UrlEvidenceable)() {
    err := m.GetBackingStore().Set("urls", value)
    if err != nil {
        panic(err)
    }
}
type TeamsMessageEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCampaignId()(*string)
    GetChannelId()(*string)
    GetDeliveryAction()(*TeamsMessageDeliveryAction)
    GetDeliveryLocation()(*TeamsDeliveryLocation)
    GetFiles()([]FileEvidenceable)
    GetGroupId()(*string)
    GetIsExternal()(*bool)
    GetIsOwned()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMessageDirection()(*AntispamTeamsDirection)
    GetMessageId()(*string)
    GetOwningTenantId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetParentMessageId()(*string)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipients()([]string)
    GetSenderFromAddress()(*string)
    GetSenderIP()(*string)
    GetSourceAppName()(*string)
    GetSourceId()(*string)
    GetSubject()(*string)
    GetSuspiciousRecipients()([]string)
    GetThreadId()(*string)
    GetThreadType()(*string)
    GetUrls()([]UrlEvidenceable)
    SetCampaignId(value *string)()
    SetChannelId(value *string)()
    SetDeliveryAction(value *TeamsMessageDeliveryAction)()
    SetDeliveryLocation(value *TeamsDeliveryLocation)()
    SetFiles(value []FileEvidenceable)()
    SetGroupId(value *string)()
    SetIsExternal(value *bool)()
    SetIsOwned(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMessageDirection(value *AntispamTeamsDirection)()
    SetMessageId(value *string)()
    SetOwningTenantId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetParentMessageId(value *string)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipients(value []string)()
    SetSenderFromAddress(value *string)()
    SetSenderIP(value *string)()
    SetSourceAppName(value *string)()
    SetSourceId(value *string)()
    SetSubject(value *string)()
    SetSuspiciousRecipients(value []string)()
    SetThreadId(value *string)()
    SetThreadType(value *string)()
    SetUrls(value []UrlEvidenceable)()
}
