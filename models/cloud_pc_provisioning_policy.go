package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcProvisioningPolicy struct {
    Entity
}
// NewCloudPcProvisioningPolicy instantiates a new CloudPcProvisioningPolicy and sets the default values.
func NewCloudPcProvisioningPolicy()(*CloudPcProvisioningPolicy) {
    m := &CloudPcProvisioningPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcProvisioningPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcProvisioningPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcProvisioningPolicy(), nil
}
// GetAlternateResourceUrl gets the alternateResourceUrl property value. The URL of the alternate resource that links to this provisioning policy. Read-only.
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetAlternateResourceUrl()(*string) {
    val, err := m.GetBackingStore().Get("alternateResourceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignments gets the assignments property value. A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments relationship, see Get cloudPcProvisioningPolicy.
// returns a []CloudPcProvisioningPolicyAssignmentable when successful
func (m *CloudPcProvisioningPolicy) GetAssignments()([]CloudPcProvisioningPolicyAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcProvisioningPolicyAssignmentable)
    }
    return nil
}
// GetAutopatch gets the autopatch property value. Indicates the Windows Autopatch settings for Cloud PCs using this provisioning policy. The settings take effect when the tenant enrolls in Autopatch and the managedType of the microsoftManagedDesktop property is set as starterManaged. Supports $select.
// returns a CloudPcProvisioningPolicyAutopatchable when successful
func (m *CloudPcProvisioningPolicy) GetAutopatch()(CloudPcProvisioningPolicyAutopatchable) {
    val, err := m.GetBackingStore().Get("autopatch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcProvisioningPolicyAutopatchable)
    }
    return nil
}
// GetAutopilotConfiguration gets the autopilotConfiguration property value. The specific settings for Windows Autopilot that enable Windows 365 customers to experience it on Cloud PC. Supports $select.
// returns a CloudPcAutopilotConfigurationable when successful
func (m *CloudPcProvisioningPolicy) GetAutopilotConfiguration()(CloudPcAutopilotConfigurationable) {
    val, err := m.GetBackingStore().Get("autopilotConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcAutopilotConfigurationable)
    }
    return nil
}
// GetCloudPcGroupDisplayName gets the cloudPcGroupDisplayName property value. The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetCloudPcGroupDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("cloudPcGroupDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCloudPcNamingTemplate gets the cloudPcNamingTemplate property value. The template used to name Cloud PCs provisioned using this policy. The template can contain custom text and replacement tokens, including %USERNAME:x% and %RAND:x%, which represent the user's name and a randomly generated number, respectively. For example, CPC-%USERNAME:4%-%RAND:5% means that the name of the Cloud PC starts with CPC-, followed by a four-character username, a - character, and then five random characters. The total length of the text generated by the template can't exceed 15 characters. Supports $filter, $select, and $orderby.
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetCloudPcNamingTemplate()(*string) {
    val, err := m.GetBackingStore().Get("cloudPcNamingTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The provisioning policy description. Supports $filter, $select, and $orderBy.
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the provisioning policy.
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomainJoinConfigurations gets the domainJoinConfigurations property value. Specifies a list ordered by priority on how Cloud PCs join Microsoft Entra ID (Azure AD). Supports $select.
// returns a []CloudPcDomainJoinConfigurationable when successful
func (m *CloudPcProvisioningPolicy) GetDomainJoinConfigurations()([]CloudPcDomainJoinConfigurationable) {
    val, err := m.GetBackingStore().Get("domainJoinConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcDomainJoinConfigurationable)
    }
    return nil
}
// GetEnableSingleSignOn gets the enableSingleSignOn property value. True if single sign-on can access the provisioned Cloud PC. False indicates that the provisioned Cloud PC doesn't support this feature. The default value is false. Windows 365 users can use single sign-on to authenticate to Microsoft Entra ID with passwordless options (for example, FIDO keys) to access their Cloud PC. Optional.
// returns a *bool when successful
func (m *CloudPcProvisioningPolicy) GetEnableSingleSignOn()(*bool) {
    val, err := m.GetBackingStore().Get("enableSingleSignOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcProvisioningPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alternateResourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateResourceUrl(val)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcProvisioningPolicyAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcProvisioningPolicyAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["autopatch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcProvisioningPolicyAutopatchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutopatch(val.(CloudPcProvisioningPolicyAutopatchable))
        }
        return nil
    }
    res["autopilotConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcAutopilotConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutopilotConfiguration(val.(CloudPcAutopilotConfigurationable))
        }
        return nil
    }
    res["cloudPcGroupDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcGroupDisplayName(val)
        }
        return nil
    }
    res["cloudPcNamingTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcNamingTemplate(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["domainJoinConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDomainJoinConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcDomainJoinConfigurationable)
                }
            }
            m.SetDomainJoinConfigurations(res)
        }
        return nil
    }
    res["enableSingleSignOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleSignOn(val)
        }
        return nil
    }
    res["gracePeriodInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodInHours(val)
        }
        return nil
    }
    res["imageDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDisplayName(val)
        }
        return nil
    }
    res["imageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageId(val)
        }
        return nil
    }
    res["imageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningPolicyImageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageType(val.(*CloudPcProvisioningPolicyImageType))
        }
        return nil
    }
    res["localAdminEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalAdminEnabled(val)
        }
        return nil
    }
    res["managedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcManagementService)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBy(val.(*CloudPcManagementService))
        }
        return nil
    }
    res["microsoftManagedDesktop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMicrosoftManagedDesktopFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftManagedDesktop(val.(MicrosoftManagedDesktopable))
        }
        return nil
    }
    res["provisioningType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningType(val.(*CloudPcProvisioningType))
        }
        return nil
    }
    res["scopeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetScopeIds(res)
        }
        return nil
    }
    res["windowsSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcWindowsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSetting(val.(CloudPcWindowsSettingable))
        }
        return nil
    }
    res["windowsSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcWindowsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSettings(val.(CloudPcWindowsSettingsable))
        }
        return nil
    }
    return res
}
// GetGracePeriodInHours gets the gracePeriodInHours property value. The number of hours to wait before reprovisioning/deprovisioning happens. Read-only.
// returns a *int32 when successful
func (m *CloudPcProvisioningPolicy) GetGracePeriodInHours()(*int32) {
    val, err := m.GetBackingStore().Get("gracePeriodInHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetImageDisplayName gets the imageDisplayName property value. The display name of the operating system image that is used for provisioning. For example, Windows 11 Preview + Microsoft 365 Apps 23H2 23H2. Supports $filter, $select, and $orderBy.
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetImageDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("imageDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetImageId gets the imageId property value. The unique identifier that represents an operating system image that is used for provisioning new Cloud PCs. The format for a gallery type image is: {publisherNameofferNameskuName}. Supported values for each of the parameters are:publisher: Microsoftwindowsdesktop offer: windows-ent-cpc sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os, 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365, and 19h2-ent-cpc-os Supports $filter, $select, and $orderBy.
// returns a *string when successful
func (m *CloudPcProvisioningPolicy) GetImageId()(*string) {
    val, err := m.GetBackingStore().Get("imageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetImageType gets the imageType property value. The imageType property
// returns a *CloudPcProvisioningPolicyImageType when successful
func (m *CloudPcProvisioningPolicy) GetImageType()(*CloudPcProvisioningPolicyImageType) {
    val, err := m.GetBackingStore().Get("imageType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcProvisioningPolicyImageType)
    }
    return nil
}
// GetLocalAdminEnabled gets the localAdminEnabled property value. When true, the local admin is enabled for Cloud PCs; false indicates that the local admin isn't enabled for Cloud PCs. The default value is false. Supports $filter, $select, and $orderBy.
// returns a *bool when successful
func (m *CloudPcProvisioningPolicy) GetLocalAdminEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("localAdminEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagedBy gets the managedBy property value. The managedBy property
// returns a *CloudPcManagementService when successful
func (m *CloudPcProvisioningPolicy) GetManagedBy()(*CloudPcManagementService) {
    val, err := m.GetBackingStore().Get("managedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcManagementService)
    }
    return nil
}
// GetMicrosoftManagedDesktop gets the microsoftManagedDesktop property value. The specific settings to microsoftManagedDesktop that enables Microsoft Managed Desktop customers to get device managed experience for Cloud PC. To enable microsoftManagedDesktop to provide more value, an admin needs to specify certain settings in it. Supports $filter, $select, and $orderBy.
// returns a MicrosoftManagedDesktopable when successful
func (m *CloudPcProvisioningPolicy) GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable) {
    val, err := m.GetBackingStore().Get("microsoftManagedDesktop")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MicrosoftManagedDesktopable)
    }
    return nil
}
// GetProvisioningType gets the provisioningType property value. Specifies the type of licenses to be used when provisioning Cloud PCs using this policy. The possible values are dedicated, shared, unknownFutureValue, sharedByUser, sharedByEntraGroup. Use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: sharedByUser, sharedByEntraGroup. The shared member is deprecated and will stop returning on April 30, 2027; going forward, use the sharedByUser member. For example, a dedicated service plan can be assigned to only one user and provision only one Cloud PC. The shared and sharedByUser plans require customers to purchase a shared service plan. Each shared license purchased can enable up to three Cloud PCs, with only one user signed in at a time. The sharedByEntraGroup plan also requires the purchase of a shared service plan. Each shared license under this plan can enable one Cloud PC, which is shared for the group according to the assignments of this policy. By default, the license type is dedicated if the provisioningType isn't specified when you create the cloudPcProvisioningPolicy. You can't change this property after the cloudPcProvisioningPolicy is created.
// returns a *CloudPcProvisioningType when successful
func (m *CloudPcProvisioningPolicy) GetProvisioningType()(*CloudPcProvisioningType) {
    val, err := m.GetBackingStore().Get("provisioningType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcProvisioningType)
    }
    return nil
}
// GetScopeIds gets the scopeIds property value. The scopeIds property
// returns a []string when successful
func (m *CloudPcProvisioningPolicy) GetScopeIds()([]string) {
    val, err := m.GetBackingStore().Get("scopeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetWindowsSetting gets the windowsSetting property value. Indicates a specific Windows setting to configure during the creation of Cloud PCs for this provisioning policy. Supports $select.
// returns a CloudPcWindowsSettingable when successful
func (m *CloudPcProvisioningPolicy) GetWindowsSetting()(CloudPcWindowsSettingable) {
    val, err := m.GetBackingStore().Get("windowsSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcWindowsSettingable)
    }
    return nil
}
// GetWindowsSettings gets the windowsSettings property value. Specific Windows settings to configure during the creation of Cloud PCs for this provisioning policy. Supports $select. The windowsSettings property is deprecated and will stop returning data on January 31, 2024. Going forward, use the windowsSetting property.
// returns a CloudPcWindowsSettingsable when successful
func (m *CloudPcProvisioningPolicy) GetWindowsSettings()(CloudPcWindowsSettingsable) {
    val, err := m.GetBackingStore().Get("windowsSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcWindowsSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcProvisioningPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alternateResourceUrl", m.GetAlternateResourceUrl())
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("autopatch", m.GetAutopatch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("autopilotConfiguration", m.GetAutopilotConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cloudPcGroupDisplayName", m.GetCloudPcGroupDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cloudPcNamingTemplate", m.GetCloudPcNamingTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDomainJoinConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainJoinConfigurations()))
        for i, v := range m.GetDomainJoinConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("domainJoinConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSingleSignOn", m.GetEnableSingleSignOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("gracePeriodInHours", m.GetGracePeriodInHours())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageDisplayName", m.GetImageDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageId", m.GetImageId())
        if err != nil {
            return err
        }
    }
    if m.GetImageType() != nil {
        cast := (*m.GetImageType()).String()
        err = writer.WriteStringValue("imageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localAdminEnabled", m.GetLocalAdminEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetManagedBy() != nil {
        cast := (*m.GetManagedBy()).String()
        err = writer.WriteStringValue("managedBy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("microsoftManagedDesktop", m.GetMicrosoftManagedDesktop())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningType() != nil {
        cast := (*m.GetProvisioningType()).String()
        err = writer.WriteStringValue("provisioningType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopeIds() != nil {
        err = writer.WriteCollectionOfStringValues("scopeIds", m.GetScopeIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsSetting", m.GetWindowsSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsSettings", m.GetWindowsSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlternateResourceUrl sets the alternateResourceUrl property value. The URL of the alternate resource that links to this provisioning policy. Read-only.
func (m *CloudPcProvisioningPolicy) SetAlternateResourceUrl(value *string)() {
    err := m.GetBackingStore().Set("alternateResourceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignments sets the assignments property value. A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments relationship, see Get cloudPcProvisioningPolicy.
func (m *CloudPcProvisioningPolicy) SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetAutopatch sets the autopatch property value. Indicates the Windows Autopatch settings for Cloud PCs using this provisioning policy. The settings take effect when the tenant enrolls in Autopatch and the managedType of the microsoftManagedDesktop property is set as starterManaged. Supports $select.
func (m *CloudPcProvisioningPolicy) SetAutopatch(value CloudPcProvisioningPolicyAutopatchable)() {
    err := m.GetBackingStore().Set("autopatch", value)
    if err != nil {
        panic(err)
    }
}
// SetAutopilotConfiguration sets the autopilotConfiguration property value. The specific settings for Windows Autopilot that enable Windows 365 customers to experience it on Cloud PC. Supports $select.
func (m *CloudPcProvisioningPolicy) SetAutopilotConfiguration(value CloudPcAutopilotConfigurationable)() {
    err := m.GetBackingStore().Set("autopilotConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcGroupDisplayName sets the cloudPcGroupDisplayName property value. The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.
func (m *CloudPcProvisioningPolicy) SetCloudPcGroupDisplayName(value *string)() {
    err := m.GetBackingStore().Set("cloudPcGroupDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcNamingTemplate sets the cloudPcNamingTemplate property value. The template used to name Cloud PCs provisioned using this policy. The template can contain custom text and replacement tokens, including %USERNAME:x% and %RAND:x%, which represent the user's name and a randomly generated number, respectively. For example, CPC-%USERNAME:4%-%RAND:5% means that the name of the Cloud PC starts with CPC-, followed by a four-character username, a - character, and then five random characters. The total length of the text generated by the template can't exceed 15 characters. Supports $filter, $select, and $orderby.
func (m *CloudPcProvisioningPolicy) SetCloudPcNamingTemplate(value *string)() {
    err := m.GetBackingStore().Set("cloudPcNamingTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The provisioning policy description. Supports $filter, $select, and $orderBy.
func (m *CloudPcProvisioningPolicy) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the provisioning policy.
func (m *CloudPcProvisioningPolicy) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainJoinConfigurations sets the domainJoinConfigurations property value. Specifies a list ordered by priority on how Cloud PCs join Microsoft Entra ID (Azure AD). Supports $select.
func (m *CloudPcProvisioningPolicy) SetDomainJoinConfigurations(value []CloudPcDomainJoinConfigurationable)() {
    err := m.GetBackingStore().Set("domainJoinConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSingleSignOn sets the enableSingleSignOn property value. True if single sign-on can access the provisioned Cloud PC. False indicates that the provisioned Cloud PC doesn't support this feature. The default value is false. Windows 365 users can use single sign-on to authenticate to Microsoft Entra ID with passwordless options (for example, FIDO keys) to access their Cloud PC. Optional.
func (m *CloudPcProvisioningPolicy) SetEnableSingleSignOn(value *bool)() {
    err := m.GetBackingStore().Set("enableSingleSignOn", value)
    if err != nil {
        panic(err)
    }
}
// SetGracePeriodInHours sets the gracePeriodInHours property value. The number of hours to wait before reprovisioning/deprovisioning happens. Read-only.
func (m *CloudPcProvisioningPolicy) SetGracePeriodInHours(value *int32)() {
    err := m.GetBackingStore().Set("gracePeriodInHours", value)
    if err != nil {
        panic(err)
    }
}
// SetImageDisplayName sets the imageDisplayName property value. The display name of the operating system image that is used for provisioning. For example, Windows 11 Preview + Microsoft 365 Apps 23H2 23H2. Supports $filter, $select, and $orderBy.
func (m *CloudPcProvisioningPolicy) SetImageDisplayName(value *string)() {
    err := m.GetBackingStore().Set("imageDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetImageId sets the imageId property value. The unique identifier that represents an operating system image that is used for provisioning new Cloud PCs. The format for a gallery type image is: {publisherNameofferNameskuName}. Supported values for each of the parameters are:publisher: Microsoftwindowsdesktop offer: windows-ent-cpc sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os, 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365, and 19h2-ent-cpc-os Supports $filter, $select, and $orderBy.
func (m *CloudPcProvisioningPolicy) SetImageId(value *string)() {
    err := m.GetBackingStore().Set("imageId", value)
    if err != nil {
        panic(err)
    }
}
// SetImageType sets the imageType property value. The imageType property
func (m *CloudPcProvisioningPolicy) SetImageType(value *CloudPcProvisioningPolicyImageType)() {
    err := m.GetBackingStore().Set("imageType", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalAdminEnabled sets the localAdminEnabled property value. When true, the local admin is enabled for Cloud PCs; false indicates that the local admin isn't enabled for Cloud PCs. The default value is false. Supports $filter, $select, and $orderBy.
func (m *CloudPcProvisioningPolicy) SetLocalAdminEnabled(value *bool)() {
    err := m.GetBackingStore().Set("localAdminEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedBy sets the managedBy property value. The managedBy property
func (m *CloudPcProvisioningPolicy) SetManagedBy(value *CloudPcManagementService)() {
    err := m.GetBackingStore().Set("managedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftManagedDesktop sets the microsoftManagedDesktop property value. The specific settings to microsoftManagedDesktop that enables Microsoft Managed Desktop customers to get device managed experience for Cloud PC. To enable microsoftManagedDesktop to provide more value, an admin needs to specify certain settings in it. Supports $filter, $select, and $orderBy.
func (m *CloudPcProvisioningPolicy) SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)() {
    err := m.GetBackingStore().Set("microsoftManagedDesktop", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningType sets the provisioningType property value. Specifies the type of licenses to be used when provisioning Cloud PCs using this policy. The possible values are dedicated, shared, unknownFutureValue, sharedByUser, sharedByEntraGroup. Use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: sharedByUser, sharedByEntraGroup. The shared member is deprecated and will stop returning on April 30, 2027; going forward, use the sharedByUser member. For example, a dedicated service plan can be assigned to only one user and provision only one Cloud PC. The shared and sharedByUser plans require customers to purchase a shared service plan. Each shared license purchased can enable up to three Cloud PCs, with only one user signed in at a time. The sharedByEntraGroup plan also requires the purchase of a shared service plan. Each shared license under this plan can enable one Cloud PC, which is shared for the group according to the assignments of this policy. By default, the license type is dedicated if the provisioningType isn't specified when you create the cloudPcProvisioningPolicy. You can't change this property after the cloudPcProvisioningPolicy is created.
func (m *CloudPcProvisioningPolicy) SetProvisioningType(value *CloudPcProvisioningType)() {
    err := m.GetBackingStore().Set("provisioningType", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeIds sets the scopeIds property value. The scopeIds property
func (m *CloudPcProvisioningPolicy) SetScopeIds(value []string)() {
    err := m.GetBackingStore().Set("scopeIds", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsSetting sets the windowsSetting property value. Indicates a specific Windows setting to configure during the creation of Cloud PCs for this provisioning policy. Supports $select.
func (m *CloudPcProvisioningPolicy) SetWindowsSetting(value CloudPcWindowsSettingable)() {
    err := m.GetBackingStore().Set("windowsSetting", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsSettings sets the windowsSettings property value. Specific Windows settings to configure during the creation of Cloud PCs for this provisioning policy. Supports $select. The windowsSettings property is deprecated and will stop returning data on January 31, 2024. Going forward, use the windowsSetting property.
func (m *CloudPcProvisioningPolicy) SetWindowsSettings(value CloudPcWindowsSettingsable)() {
    err := m.GetBackingStore().Set("windowsSettings", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcProvisioningPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternateResourceUrl()(*string)
    GetAssignments()([]CloudPcProvisioningPolicyAssignmentable)
    GetAutopatch()(CloudPcProvisioningPolicyAutopatchable)
    GetAutopilotConfiguration()(CloudPcAutopilotConfigurationable)
    GetCloudPcGroupDisplayName()(*string)
    GetCloudPcNamingTemplate()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDomainJoinConfigurations()([]CloudPcDomainJoinConfigurationable)
    GetEnableSingleSignOn()(*bool)
    GetGracePeriodInHours()(*int32)
    GetImageDisplayName()(*string)
    GetImageId()(*string)
    GetImageType()(*CloudPcProvisioningPolicyImageType)
    GetLocalAdminEnabled()(*bool)
    GetManagedBy()(*CloudPcManagementService)
    GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable)
    GetProvisioningType()(*CloudPcProvisioningType)
    GetScopeIds()([]string)
    GetWindowsSetting()(CloudPcWindowsSettingable)
    GetWindowsSettings()(CloudPcWindowsSettingsable)
    SetAlternateResourceUrl(value *string)()
    SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)()
    SetAutopatch(value CloudPcProvisioningPolicyAutopatchable)()
    SetAutopilotConfiguration(value CloudPcAutopilotConfigurationable)()
    SetCloudPcGroupDisplayName(value *string)()
    SetCloudPcNamingTemplate(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDomainJoinConfigurations(value []CloudPcDomainJoinConfigurationable)()
    SetEnableSingleSignOn(value *bool)()
    SetGracePeriodInHours(value *int32)()
    SetImageDisplayName(value *string)()
    SetImageId(value *string)()
    SetImageType(value *CloudPcProvisioningPolicyImageType)()
    SetLocalAdminEnabled(value *bool)()
    SetManagedBy(value *CloudPcManagementService)()
    SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)()
    SetProvisioningType(value *CloudPcProvisioningType)()
    SetScopeIds(value []string)()
    SetWindowsSetting(value CloudPcWindowsSettingable)()
    SetWindowsSettings(value CloudPcWindowsSettingsable)()
}
