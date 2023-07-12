package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDerivedCredentialSettings entity that describes tenant level settings for derived credentials
type DeviceManagementDerivedCredentialSettings struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewDeviceManagementDerivedCredentialSettings instantiates a new deviceManagementDerivedCredentialSettings and sets the default values.
func NewDeviceManagementDerivedCredentialSettings()(*DeviceManagementDerivedCredentialSettings) {
    m := &DeviceManagementDerivedCredentialSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDerivedCredentialSettings(), nil
}
// GetDisplayName gets the displayName property value. The display name for the profile.
func (m *DeviceManagementDerivedCredentialSettings) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDerivedCredentialSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["helpUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpUrl(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementDerivedCredentialIssuer)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val.(*DeviceManagementDerivedCredentialIssuer))
        }
        return nil
    }
    res["notificationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementDerivedCredentialNotificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationType(val.(*DeviceManagementDerivedCredentialNotificationType))
        }
        return nil
    }
    res["renewalThresholdPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenewalThresholdPercentage(val)
        }
        return nil
    }
    return res
}
// GetHelpUrl gets the helpUrl property value. The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
func (m *DeviceManagementDerivedCredentialSettings) GetHelpUrl()(*string) {
    val, err := m.GetBackingStore().Get("helpUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIssuer gets the issuer property value. Supported values for the derived credential issuer.
func (m *DeviceManagementDerivedCredentialSettings) GetIssuer()(*DeviceManagementDerivedCredentialIssuer) {
    val, err := m.GetBackingStore().Get("issuer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementDerivedCredentialIssuer)
    }
    return nil
}
// GetNotificationType gets the notificationType property value. Supported values for the notification type to use.
func (m *DeviceManagementDerivedCredentialSettings) GetNotificationType()(*DeviceManagementDerivedCredentialNotificationType) {
    val, err := m.GetBackingStore().Get("notificationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementDerivedCredentialNotificationType)
    }
    return nil
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. The nominal percentage of time before certificate renewal is initiated by the client.
func (m *DeviceManagementDerivedCredentialSettings) GetRenewalThresholdPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("renewalThresholdPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementDerivedCredentialSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("helpUrl", m.GetHelpUrl())
        if err != nil {
            return err
        }
    }
    if m.GetIssuer() != nil {
        cast := (*m.GetIssuer()).String()
        err = writer.WriteStringValue("issuer", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationType() != nil {
        cast := (*m.GetNotificationType()).String()
        err = writer.WriteStringValue("notificationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("renewalThresholdPercentage", m.GetRenewalThresholdPercentage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name for the profile.
func (m *DeviceManagementDerivedCredentialSettings) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetHelpUrl sets the helpUrl property value. The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
func (m *DeviceManagementDerivedCredentialSettings) SetHelpUrl(value *string)() {
    err := m.GetBackingStore().Set("helpUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuer sets the issuer property value. Supported values for the derived credential issuer.
func (m *DeviceManagementDerivedCredentialSettings) SetIssuer(value *DeviceManagementDerivedCredentialIssuer)() {
    err := m.GetBackingStore().Set("issuer", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationType sets the notificationType property value. Supported values for the notification type to use.
func (m *DeviceManagementDerivedCredentialSettings) SetNotificationType(value *DeviceManagementDerivedCredentialNotificationType)() {
    err := m.GetBackingStore().Set("notificationType", value)
    if err != nil {
        panic(err)
    }
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. The nominal percentage of time before certificate renewal is initiated by the client.
func (m *DeviceManagementDerivedCredentialSettings) SetRenewalThresholdPercentage(value *int32)() {
    err := m.GetBackingStore().Set("renewalThresholdPercentage", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementDerivedCredentialSettingsable 
type DeviceManagementDerivedCredentialSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetHelpUrl()(*string)
    GetIssuer()(*DeviceManagementDerivedCredentialIssuer)
    GetNotificationType()(*DeviceManagementDerivedCredentialNotificationType)
    GetRenewalThresholdPercentage()(*int32)
    SetDisplayName(value *string)()
    SetHelpUrl(value *string)()
    SetIssuer(value *DeviceManagementDerivedCredentialIssuer)()
    SetNotificationType(value *DeviceManagementDerivedCredentialNotificationType)()
    SetRenewalThresholdPercentage(value *int32)()
}
