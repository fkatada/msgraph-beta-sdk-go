package models
import (
    "errors"
)
// Criteria which defines the type of device this access rule will apply to
type DeviceManagementExchangeAccessRuleType int

const (
    // Family of devices
    FAMILY_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE DeviceManagementExchangeAccessRuleType = iota
    // Specific model of device
    MODEL_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE
)

func (i DeviceManagementExchangeAccessRuleType) String() string {
    return []string{"family", "model"}[i]
}
func ParseDeviceManagementExchangeAccessRuleType(v string) (any, error) {
    result := FAMILY_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE
    switch v {
        case "family":
            result = FAMILY_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE
        case "model":
            result = MODEL_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeAccessRuleType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeAccessRuleType(values []DeviceManagementExchangeAccessRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
