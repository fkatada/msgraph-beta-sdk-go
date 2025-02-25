package models
type ManagedDeviceRemoteAction int

const (
    // Name of the retire action.
    RETIRE_MANAGEDDEVICEREMOTEACTION ManagedDeviceRemoteAction = iota
    // Name of the delete action.
    DELETE_MANAGEDDEVICEREMOTEACTION
    // Name of the full Scan action.
    FULLSCAN_MANAGEDDEVICEREMOTEACTION
    // Name of the Quick Scan action.
    QUICKSCAN_MANAGEDDEVICEREMOTEACTION
    // Signature Update action
    SIGNATUREUPDATE_MANAGEDDEVICEREMOTEACTION
    // Name of the wipe action.
    WIPE_MANAGEDDEVICEREMOTEACTION
    // Name of the Custom Text Notification action.
    CUSTOMTEXTNOTIFICATION_MANAGEDDEVICEREMOTEACTION
    // Name of the reboot now action.
    REBOOTNOW_MANAGEDDEVICEREMOTEACTION
    // Set Device Name action.
    SETDEVICENAME_MANAGEDDEVICEREMOTEACTION
    // Sync Device action.
    SYNCDEVICE_MANAGEDDEVICEREMOTEACTION
    // Name of the deprovision action.
    DEPROVISION_MANAGEDDEVICEREMOTEACTION
    // Name of the disable action.
    DISABLE_MANAGEDDEVICEREMOTEACTION
    // Name of the reenable action.
    REENABLE_MANAGEDDEVICEREMOTEACTION
    // Name of the moveDevicesToOU action.
    MOVEDEVICETOORGANIZATIONALUNIT_MANAGEDDEVICEREMOTEACTION
    // Name of action to Activate eSIM on the device.
    ACTIVATEDEVICEESIM_MANAGEDDEVICEREMOTEACTION
    // Name of the collectDiagnostics action.
    COLLECTDIAGNOSTICS_MANAGEDDEVICEREMOTEACTION
    // Name of action to initiate MDM key recovery
    INITIATEMOBILEDEVICEMANAGEMENTKEYRECOVERY_MANAGEDDEVICEREMOTEACTION
    // Name of action to initiate On Demand Proactive Remediation
    INITIATEONDEMANDPROACTIVEREMEDIATION_MANAGEDDEVICEREMOTEACTION
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MANAGEDDEVICEREMOTEACTION
    // Indicates remote device action to intiate Mobile Device Management (MDM) attestation if device is capable for it
    INITIATEDEVICEATTESTATION_MANAGEDDEVICEREMOTEACTION
)

func (i ManagedDeviceRemoteAction) String() string {
    return []string{"retire", "delete", "fullScan", "quickScan", "signatureUpdate", "wipe", "customTextNotification", "rebootNow", "setDeviceName", "syncDevice", "deprovision", "disable", "reenable", "moveDeviceToOrganizationalUnit", "activateDeviceEsim", "collectDiagnostics", "initiateMobileDeviceManagementKeyRecovery", "initiateOnDemandProactiveRemediation", "unknownFutureValue", "initiateDeviceAttestation"}[i]
}
func ParseManagedDeviceRemoteAction(v string) (any, error) {
    result := RETIRE_MANAGEDDEVICEREMOTEACTION
    switch v {
        case "retire":
            result = RETIRE_MANAGEDDEVICEREMOTEACTION
        case "delete":
            result = DELETE_MANAGEDDEVICEREMOTEACTION
        case "fullScan":
            result = FULLSCAN_MANAGEDDEVICEREMOTEACTION
        case "quickScan":
            result = QUICKSCAN_MANAGEDDEVICEREMOTEACTION
        case "signatureUpdate":
            result = SIGNATUREUPDATE_MANAGEDDEVICEREMOTEACTION
        case "wipe":
            result = WIPE_MANAGEDDEVICEREMOTEACTION
        case "customTextNotification":
            result = CUSTOMTEXTNOTIFICATION_MANAGEDDEVICEREMOTEACTION
        case "rebootNow":
            result = REBOOTNOW_MANAGEDDEVICEREMOTEACTION
        case "setDeviceName":
            result = SETDEVICENAME_MANAGEDDEVICEREMOTEACTION
        case "syncDevice":
            result = SYNCDEVICE_MANAGEDDEVICEREMOTEACTION
        case "deprovision":
            result = DEPROVISION_MANAGEDDEVICEREMOTEACTION
        case "disable":
            result = DISABLE_MANAGEDDEVICEREMOTEACTION
        case "reenable":
            result = REENABLE_MANAGEDDEVICEREMOTEACTION
        case "moveDeviceToOrganizationalUnit":
            result = MOVEDEVICETOORGANIZATIONALUNIT_MANAGEDDEVICEREMOTEACTION
        case "activateDeviceEsim":
            result = ACTIVATEDEVICEESIM_MANAGEDDEVICEREMOTEACTION
        case "collectDiagnostics":
            result = COLLECTDIAGNOSTICS_MANAGEDDEVICEREMOTEACTION
        case "initiateMobileDeviceManagementKeyRecovery":
            result = INITIATEMOBILEDEVICEMANAGEMENTKEYRECOVERY_MANAGEDDEVICEREMOTEACTION
        case "initiateOnDemandProactiveRemediation":
            result = INITIATEONDEMANDPROACTIVEREMEDIATION_MANAGEDDEVICEREMOTEACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEDDEVICEREMOTEACTION
        case "initiateDeviceAttestation":
            result = INITIATEDEVICEATTESTATION_MANAGEDDEVICEREMOTEACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeManagedDeviceRemoteAction(values []ManagedDeviceRemoteAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManagedDeviceRemoteAction) isMultiValue() bool {
    return false
}
