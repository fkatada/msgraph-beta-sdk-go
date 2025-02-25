package models
// Represents various schedule modes for Zebra FOTA deployment.
type ZebraFotaScheduleMode int

const (
    // Instructs the device to install the update as soon as it is received.
    INSTALLNOW_ZEBRAFOTASCHEDULEMODE ZebraFotaScheduleMode = iota
    // Schedule an update to be installed at a specified date and time.
    SCHEDULED_ZEBRAFOTASCHEDULEMODE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ZEBRAFOTASCHEDULEMODE
)

func (i ZebraFotaScheduleMode) String() string {
    return []string{"installNow", "scheduled", "unknownFutureValue"}[i]
}
func ParseZebraFotaScheduleMode(v string) (any, error) {
    result := INSTALLNOW_ZEBRAFOTASCHEDULEMODE
    switch v {
        case "installNow":
            result = INSTALLNOW_ZEBRAFOTASCHEDULEMODE
        case "scheduled":
            result = SCHEDULED_ZEBRAFOTASCHEDULEMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTASCHEDULEMODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeZebraFotaScheduleMode(values []ZebraFotaScheduleMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ZebraFotaScheduleMode) isMultiValue() bool {
    return false
}
