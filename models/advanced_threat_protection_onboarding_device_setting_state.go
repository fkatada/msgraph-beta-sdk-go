package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdvancedThreatProtectionOnboardingDeviceSettingState aTP onboarding State for a given device.
type AdvancedThreatProtectionOnboardingDeviceSettingState struct {
    Entity
}
// NewAdvancedThreatProtectionOnboardingDeviceSettingState instantiates a new advancedThreatProtectionOnboardingDeviceSettingState and sets the default values.
func NewAdvancedThreatProtectionOnboardingDeviceSettingState()(*AdvancedThreatProtectionOnboardingDeviceSettingState) {
    m := &AdvancedThreatProtectionOnboardingDeviceSettingState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdvancedThreatProtectionOnboardingDeviceSettingState(), nil
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("complianceGracePeriodExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The Device Id that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceModel gets the deviceModel property value. The device model that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetDeviceModel()(*string) {
    val, err := m.GetBackingStore().Get("deviceModel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The Device Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceGracePeriodExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*DeviceType))
        }
        return nil
    }
    res["setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val)
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["userEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
        }
        return nil
    }
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
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetPlatformType gets the platformType property value. Device type.
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetPlatformType()(*DeviceType) {
    val, err := m.GetBackingStore().Get("platformType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceType)
    }
    return nil
}
// GetSetting gets the setting property value. The setting class name and property name.
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetSetting()(*string) {
    val, err := m.GetBackingStore().Get("setting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingName gets the settingName property value. The Setting Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetSettingName()(*string) {
    val, err := m.GetBackingStore().Get("settingName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetState gets the state property value. The state property
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetState()(*ComplianceStatus) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ComplianceStatus)
    }
    return nil
}
// GetUserEmail gets the userEmail property value. The User email address that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetUserEmail()(*string) {
    val, err := m.GetBackingStore().Get("userEmail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The user Id that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserName gets the userName property value. The User Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetUserName()(*string) {
    val, err := m.GetBackingStore().Get("userName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The User PrincipalName that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("complianceGracePeriodExpirationDateTime", m.GetComplianceGracePeriodExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userEmail", m.GetUserEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("complianceGracePeriodExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The Device Id that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceModel sets the deviceModel property value. The device model that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetDeviceModel(value *string)() {
    err := m.GetBackingStore().Set("deviceModel", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The Device Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatformType sets the platformType property value. Device type.
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetPlatformType(value *DeviceType)() {
    err := m.GetBackingStore().Set("platformType", value)
    if err != nil {
        panic(err)
    }
}
// SetSetting sets the setting property value. The setting class name and property name.
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetSetting(value *string)() {
    err := m.GetBackingStore().Set("setting", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingName sets the settingName property value. The Setting Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetSettingName(value *string)() {
    err := m.GetBackingStore().Set("settingName", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetState(value *ComplianceStatus)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetUserEmail sets the userEmail property value. The User email address that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetUserEmail(value *string)() {
    err := m.GetBackingStore().Set("userEmail", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The user Id that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserName sets the userName property value. The User Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetUserName(value *string)() {
    err := m.GetBackingStore().Set("userName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The User PrincipalName that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// AdvancedThreatProtectionOnboardingDeviceSettingStateable 
type AdvancedThreatProtectionOnboardingDeviceSettingStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceId()(*string)
    GetDeviceModel()(*string)
    GetDeviceName()(*string)
    GetPlatformType()(*DeviceType)
    GetSetting()(*string)
    GetSettingName()(*string)
    GetState()(*ComplianceStatus)
    GetUserEmail()(*string)
    GetUserId()(*string)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceId(value *string)()
    SetDeviceModel(value *string)()
    SetDeviceName(value *string)()
    SetPlatformType(value *DeviceType)()
    SetSetting(value *string)()
    SetSettingName(value *string)()
    SetState(value *ComplianceStatus)()
    SetUserEmail(value *string)()
    SetUserId(value *string)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}
