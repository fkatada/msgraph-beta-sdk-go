package models
// Possible ways of adding a mobile device to management.
type DeviceEnrollmentType int

const (
    // Default value, enrollment type was not collected.
    UNKNOWN_DEVICEENROLLMENTTYPE DeviceEnrollmentType = iota
    // User driven enrollment through BYOD channel.
    USERENROLLMENT_DEVICEENROLLMENTTYPE
    // User enrollment with a device enrollment manager account.
    DEVICEENROLLMENTMANAGER_DEVICEENROLLMENTTYPE
    // Apple bulk enrollment with user challenge. (DEP, Apple Configurator)
    APPLEBULKWITHUSER_DEVICEENROLLMENTTYPE
    // Apple bulk enrollment without user challenge. (DEP, Apple Configurator, Mobile Config)
    APPLEBULKWITHOUTUSER_DEVICEENROLLMENTTYPE
    // Windows 10 Entra ID (Azure AD) Join.
    WINDOWSAZUREADJOIN_DEVICEENROLLMENTTYPE
    // Windows 10 Bulk enrollment through ICD with certificate.
    WINDOWSBULKUSERLESS_DEVICEENROLLMENTTYPE
    // Windows 10 automatic enrollment. (Add work account)
    WINDOWSAUTOENROLLMENT_DEVICEENROLLMENTTYPE
    // Windows 10 bulk Entra ID (Azure AD) Join.
    WINDOWSBULKAZUREDOMAINJOIN_DEVICEENROLLMENTTYPE
    // Windows 10 Co-Management triggered by AutoPilot or Group Policy.
    WINDOWSCOMANAGEMENT_DEVICEENROLLMENTTYPE
    // Windows 10 Entra ID (Azure AD) Join using Device Auth.
    WINDOWSAZUREADJOINUSINGDEVICEAUTH_DEVICEENROLLMENTTYPE
    // Indicates the device is enrolled via Apple User Enrollment with Company Portal. It results in an enrollment with a new partition for managed apps and data and which supports a limited set of management capabilities
    APPLEUSERENROLLMENT_DEVICEENROLLMENTTYPE
    // Indicates the device is enrolled via Apple User Enrollment with Company Portal using a device enrollment manager user. It results in an enrollment with a new partition for managed apps and data and which supports a limited set of management capabilities
    APPLEUSERENROLLMENTWITHSERVICEACCOUNT_DEVICEENROLLMENTTYPE
    // Entra ID (Azure AD) Join enrollment when an Azure VM is provisioned
    AZUREADJOINUSINGAZUREVMEXTENSION_DEVICEENROLLMENTTYPE
    // Android Enterprise Dedicated Device
    ANDROIDENTERPRISEDEDICATEDDEVICE_DEVICEENROLLMENTTYPE
    // Android Enterprise Fully Managed
    ANDROIDENTERPRISEFULLYMANAGED_DEVICEENROLLMENTTYPE
    // Android Enterprise Corporate Work Profile
    ANDROIDENTERPRISECORPORATEWORKPROFILE_DEVICEENROLLMENTTYPE
    // Indicates the device enrollment is for android device owned by/associated with user using Android Open Source Project (AOSP) on a non-Google mobile services.
    ANDROIDAOSPUSEROWNEDDEVICEENROLLMENT_DEVICEENROLLMENTTYPE
    // Indicates the device enrollment is for user less android device using Android Open Source Project (AOSP) on a non-Google mobile services.
    ANDROIDAOSPUSERLESSDEVICEENROLLMENT_DEVICEENROLLMENTTYPE
    // Indicates the device is enrolled via Apple Account Driven User Enrollment, a form of enrollment where the user enrolls via iOS Settings without using the iOS Company Portal. It results in an enrollment with a new partition for managed apps and data and which supports a limited set of management capabilities.
    APPLEACCOUNTDRIVENUSERENROLLMENT_DEVICEENROLLMENTTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEENROLLMENTTYPE
)

func (i DeviceEnrollmentType) String() string {
    return []string{"unknown", "userEnrollment", "deviceEnrollmentManager", "appleBulkWithUser", "appleBulkWithoutUser", "windowsAzureADJoin", "windowsBulkUserless", "windowsAutoEnrollment", "windowsBulkAzureDomainJoin", "windowsCoManagement", "windowsAzureADJoinUsingDeviceAuth", "appleUserEnrollment", "appleUserEnrollmentWithServiceAccount", "azureAdJoinUsingAzureVmExtension", "androidEnterpriseDedicatedDevice", "androidEnterpriseFullyManaged", "androidEnterpriseCorporateWorkProfile", "androidAOSPUserOwnedDeviceEnrollment", "androidAOSPUserlessDeviceEnrollment", "appleAccountDrivenUserEnrollment", "unknownFutureValue"}[i]
}
func ParseDeviceEnrollmentType(v string) (any, error) {
    result := UNKNOWN_DEVICEENROLLMENTTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEENROLLMENTTYPE
        case "userEnrollment":
            result = USERENROLLMENT_DEVICEENROLLMENTTYPE
        case "deviceEnrollmentManager":
            result = DEVICEENROLLMENTMANAGER_DEVICEENROLLMENTTYPE
        case "appleBulkWithUser":
            result = APPLEBULKWITHUSER_DEVICEENROLLMENTTYPE
        case "appleBulkWithoutUser":
            result = APPLEBULKWITHOUTUSER_DEVICEENROLLMENTTYPE
        case "windowsAzureADJoin":
            result = WINDOWSAZUREADJOIN_DEVICEENROLLMENTTYPE
        case "windowsBulkUserless":
            result = WINDOWSBULKUSERLESS_DEVICEENROLLMENTTYPE
        case "windowsAutoEnrollment":
            result = WINDOWSAUTOENROLLMENT_DEVICEENROLLMENTTYPE
        case "windowsBulkAzureDomainJoin":
            result = WINDOWSBULKAZUREDOMAINJOIN_DEVICEENROLLMENTTYPE
        case "windowsCoManagement":
            result = WINDOWSCOMANAGEMENT_DEVICEENROLLMENTTYPE
        case "windowsAzureADJoinUsingDeviceAuth":
            result = WINDOWSAZUREADJOINUSINGDEVICEAUTH_DEVICEENROLLMENTTYPE
        case "appleUserEnrollment":
            result = APPLEUSERENROLLMENT_DEVICEENROLLMENTTYPE
        case "appleUserEnrollmentWithServiceAccount":
            result = APPLEUSERENROLLMENTWITHSERVICEACCOUNT_DEVICEENROLLMENTTYPE
        case "azureAdJoinUsingAzureVmExtension":
            result = AZUREADJOINUSINGAZUREVMEXTENSION_DEVICEENROLLMENTTYPE
        case "androidEnterpriseDedicatedDevice":
            result = ANDROIDENTERPRISEDEDICATEDDEVICE_DEVICEENROLLMENTTYPE
        case "androidEnterpriseFullyManaged":
            result = ANDROIDENTERPRISEFULLYMANAGED_DEVICEENROLLMENTTYPE
        case "androidEnterpriseCorporateWorkProfile":
            result = ANDROIDENTERPRISECORPORATEWORKPROFILE_DEVICEENROLLMENTTYPE
        case "androidAOSPUserOwnedDeviceEnrollment":
            result = ANDROIDAOSPUSEROWNEDDEVICEENROLLMENT_DEVICEENROLLMENTTYPE
        case "androidAOSPUserlessDeviceEnrollment":
            result = ANDROIDAOSPUSERLESSDEVICEENROLLMENT_DEVICEENROLLMENTTYPE
        case "appleAccountDrivenUserEnrollment":
            result = APPLEACCOUNTDRIVENUSERENROLLMENT_DEVICEENROLLMENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEENROLLMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceEnrollmentType(values []DeviceEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceEnrollmentType) isMultiValue() bool {
    return false
}
