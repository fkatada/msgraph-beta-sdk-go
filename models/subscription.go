package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Subscription provides operations to manage the collection of activityStatistics entities.
type Subscription struct {
    Entity
    // Optional. Identifier of the application used to create the subscription. Read-only.
    applicationId *string
    // Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
    changeType *string
    // Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 255 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
    clientState *string
    // Optional. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the ID of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the ID of the service principal corresponding to the app. Read-only.
    creatorId *string
    // Optional. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
    encryptionCertificate *string
    // Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Required when includeResourceData is true.
    encryptionCertificateId *string
    // Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Optional. When set to true, change notifications include resource data (such as content of a chat message).
    includeResourceData *bool
    // Optional. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
    latestSupportedTlsVersion *string
    // Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol.
    lifecycleNotificationUrl *string
    // Optional. Desired content-type for Microsoft Graph change notifications for supported resource types. The default content-type is application/json.
    notificationContentType *string
    // Optional.  OData query options for specifying the value for the targeting resource. Clients receive notifications when the resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.
    notificationQueryOptions *string
    // Required. The URL of the endpoint that receives the change notifications. This URL must make use of the HTTPS protocol.
    notificationUrl *string
    // Optional. The app ID that the subscription service can use to generate the validation token. This allows the client to validate the authenticity of the notification received.
    notificationUrlAppId *string
    // Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/beta/). See the possible resource path values for each supported resource.
    resource *string
}
// NewSubscription instantiates a new subscription and sets the default values.
func NewSubscription()(*Subscription) {
    m := &Subscription{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.subscription";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSubscriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubscriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubscription(), nil
}
// GetApplicationId gets the applicationId property value. Optional. Identifier of the application used to create the subscription. Read-only.
func (m *Subscription) GetApplicationId()(*string) {
    return m.applicationId
}
// GetChangeType gets the changeType property value. Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
func (m *Subscription) GetChangeType()(*string) {
    return m.changeType
}
// GetClientState gets the clientState property value. Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 255 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
func (m *Subscription) GetClientState()(*string) {
    return m.clientState
}
// GetCreatorId gets the creatorId property value. Optional. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the ID of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the ID of the service principal corresponding to the app. Read-only.
func (m *Subscription) GetCreatorId()(*string) {
    return m.creatorId
}
// GetEncryptionCertificate gets the encryptionCertificate property value. Optional. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
func (m *Subscription) GetEncryptionCertificate()(*string) {
    return m.encryptionCertificate
}
// GetEncryptionCertificateId gets the encryptionCertificateId property value. Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Required when includeResourceData is true.
func (m *Subscription) GetEncryptionCertificateId()(*string) {
    return m.encryptionCertificateId
}
// GetExpirationDateTime gets the expirationDateTime property value. Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
func (m *Subscription) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Subscription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplicationId)
    res["changeType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetChangeType)
    res["clientState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetClientState)
    res["creatorId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatorId)
    res["encryptionCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEncryptionCertificate)
    res["encryptionCertificateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEncryptionCertificateId)
    res["expirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpirationDateTime)
    res["includeResourceData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIncludeResourceData)
    res["latestSupportedTlsVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLatestSupportedTlsVersion)
    res["lifecycleNotificationUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLifecycleNotificationUrl)
    res["notificationContentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotificationContentType)
    res["notificationQueryOptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotificationQueryOptions)
    res["notificationUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotificationUrl)
    res["notificationUrlAppId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotificationUrlAppId)
    res["resource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResource)
    return res
}
// GetIncludeResourceData gets the includeResourceData property value. Optional. When set to true, change notifications include resource data (such as content of a chat message).
func (m *Subscription) GetIncludeResourceData()(*bool) {
    return m.includeResourceData
}
// GetLatestSupportedTlsVersion gets the latestSupportedTlsVersion property value. Optional. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
func (m *Subscription) GetLatestSupportedTlsVersion()(*string) {
    return m.latestSupportedTlsVersion
}
// GetLifecycleNotificationUrl gets the lifecycleNotificationUrl property value. Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) GetLifecycleNotificationUrl()(*string) {
    return m.lifecycleNotificationUrl
}
// GetNotificationContentType gets the notificationContentType property value. Optional. Desired content-type for Microsoft Graph change notifications for supported resource types. The default content-type is application/json.
func (m *Subscription) GetNotificationContentType()(*string) {
    return m.notificationContentType
}
// GetNotificationQueryOptions gets the notificationQueryOptions property value. Optional.  OData query options for specifying the value for the targeting resource. Clients receive notifications when the resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.
func (m *Subscription) GetNotificationQueryOptions()(*string) {
    return m.notificationQueryOptions
}
// GetNotificationUrl gets the notificationUrl property value. Required. The URL of the endpoint that receives the change notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) GetNotificationUrl()(*string) {
    return m.notificationUrl
}
// GetNotificationUrlAppId gets the notificationUrlAppId property value. Optional. The app ID that the subscription service can use to generate the validation token. This allows the client to validate the authenticity of the notification received.
func (m *Subscription) GetNotificationUrlAppId()(*string) {
    return m.notificationUrlAppId
}
// GetResource gets the resource property value. Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/beta/). See the possible resource path values for each supported resource.
func (m *Subscription) GetResource()(*string) {
    return m.resource
}
// Serialize serializes information the current object
func (m *Subscription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("changeType", m.GetChangeType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientState", m.GetClientState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creatorId", m.GetCreatorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("encryptionCertificate", m.GetEncryptionCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("encryptionCertificateId", m.GetEncryptionCertificateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeResourceData", m.GetIncludeResourceData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("latestSupportedTlsVersion", m.GetLatestSupportedTlsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lifecycleNotificationUrl", m.GetLifecycleNotificationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationContentType", m.GetNotificationContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationQueryOptions", m.GetNotificationQueryOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationUrl", m.GetNotificationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationUrlAppId", m.GetNotificationUrlAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationId sets the applicationId property value. Optional. Identifier of the application used to create the subscription. Read-only.
func (m *Subscription) SetApplicationId(value *string)() {
    m.applicationId = value
}
// SetChangeType sets the changeType property value. Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
func (m *Subscription) SetChangeType(value *string)() {
    m.changeType = value
}
// SetClientState sets the clientState property value. Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 255 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
func (m *Subscription) SetClientState(value *string)() {
    m.clientState = value
}
// SetCreatorId sets the creatorId property value. Optional. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the ID of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the ID of the service principal corresponding to the app. Read-only.
func (m *Subscription) SetCreatorId(value *string)() {
    m.creatorId = value
}
// SetEncryptionCertificate sets the encryptionCertificate property value. Optional. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
func (m *Subscription) SetEncryptionCertificate(value *string)() {
    m.encryptionCertificate = value
}
// SetEncryptionCertificateId sets the encryptionCertificateId property value. Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Required when includeResourceData is true.
func (m *Subscription) SetEncryptionCertificateId(value *string)() {
    m.encryptionCertificateId = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
func (m *Subscription) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetIncludeResourceData sets the includeResourceData property value. Optional. When set to true, change notifications include resource data (such as content of a chat message).
func (m *Subscription) SetIncludeResourceData(value *bool)() {
    m.includeResourceData = value
}
// SetLatestSupportedTlsVersion sets the latestSupportedTlsVersion property value. Optional. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
func (m *Subscription) SetLatestSupportedTlsVersion(value *string)() {
    m.latestSupportedTlsVersion = value
}
// SetLifecycleNotificationUrl sets the lifecycleNotificationUrl property value. Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) SetLifecycleNotificationUrl(value *string)() {
    m.lifecycleNotificationUrl = value
}
// SetNotificationContentType sets the notificationContentType property value. Optional. Desired content-type for Microsoft Graph change notifications for supported resource types. The default content-type is application/json.
func (m *Subscription) SetNotificationContentType(value *string)() {
    m.notificationContentType = value
}
// SetNotificationQueryOptions sets the notificationQueryOptions property value. Optional.  OData query options for specifying the value for the targeting resource. Clients receive notifications when the resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.
func (m *Subscription) SetNotificationQueryOptions(value *string)() {
    m.notificationQueryOptions = value
}
// SetNotificationUrl sets the notificationUrl property value. Required. The URL of the endpoint that receives the change notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) SetNotificationUrl(value *string)() {
    m.notificationUrl = value
}
// SetNotificationUrlAppId sets the notificationUrlAppId property value. Optional. The app ID that the subscription service can use to generate the validation token. This allows the client to validate the authenticity of the notification received.
func (m *Subscription) SetNotificationUrlAppId(value *string)() {
    m.notificationUrlAppId = value
}
// SetResource sets the resource property value. Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/beta/). See the possible resource path values for each supported resource.
func (m *Subscription) SetResource(value *string)() {
    m.resource = value
}
