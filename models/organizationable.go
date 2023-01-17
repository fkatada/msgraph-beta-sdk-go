package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Organizationable 
type Organizationable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedPlans()([]AssignedPlanable)
    GetBranding()(OrganizationalBrandingable)
    GetBusinessPhones()([]string)
    GetCertificateBasedAuthConfiguration()([]CertificateBasedAuthConfigurationable)
    GetCertificateConnectorSetting()(CertificateConnectorSettingable)
    GetCity()(*string)
    GetCountry()(*string)
    GetCountryLetterCode()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDefaultUsageLocation()(*string)
    GetDirectorySizeQuota()(DirectorySizeQuotaable)
    GetDisplayName()(*string)
    GetExtensions()([]Extensionable)
    GetIsMultipleDataLocationsForServicesEnabled()(*bool)
    GetMarketingNotificationEmails()([]string)
    GetMobileDeviceManagementAuthority()(*MdmAuthority)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesSyncEnabled()(*bool)
    GetPartnerTenantType()(*PartnerTenantType)
    GetPostalCode()(*string)
    GetPreferredLanguage()(*string)
    GetPrivacyProfile()(PrivacyProfileable)
    GetProvisionedPlans()([]ProvisionedPlanable)
    GetSecurityComplianceNotificationMails()([]string)
    GetSecurityComplianceNotificationPhones()([]string)
    GetSettings()(OrganizationSettingsable)
    GetState()(*string)
    GetStreet()(*string)
    GetTechnicalNotificationMails()([]string)
    GetVerifiedDomains()([]VerifiedDomainable)
    SetAssignedPlans(value []AssignedPlanable)()
    SetBranding(value OrganizationalBrandingable)()
    SetBusinessPhones(value []string)()
    SetCertificateBasedAuthConfiguration(value []CertificateBasedAuthConfigurationable)()
    SetCertificateConnectorSetting(value CertificateConnectorSettingable)()
    SetCity(value *string)()
    SetCountry(value *string)()
    SetCountryLetterCode(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDefaultUsageLocation(value *string)()
    SetDirectorySizeQuota(value DirectorySizeQuotaable)()
    SetDisplayName(value *string)()
    SetExtensions(value []Extensionable)()
    SetIsMultipleDataLocationsForServicesEnabled(value *bool)()
    SetMarketingNotificationEmails(value []string)()
    SetMobileDeviceManagementAuthority(value *MdmAuthority)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetPartnerTenantType(value *PartnerTenantType)()
    SetPostalCode(value *string)()
    SetPreferredLanguage(value *string)()
    SetPrivacyProfile(value PrivacyProfileable)()
    SetProvisionedPlans(value []ProvisionedPlanable)()
    SetSecurityComplianceNotificationMails(value []string)()
    SetSecurityComplianceNotificationPhones(value []string)()
    SetSettings(value OrganizationSettingsable)()
    SetState(value *string)()
    SetStreet(value *string)()
    SetTechnicalNotificationMails(value []string)()
    SetVerifiedDomains(value []VerifiedDomainable)()
}
