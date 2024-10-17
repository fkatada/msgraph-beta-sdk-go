package models
type ProtectionUnitsBulkJobStatus int

const (
    UNKNOWN_PROTECTIONUNITSBULKJOBSTATUS ProtectionUnitsBulkJobStatus = iota
    ACTIVE_PROTECTIONUNITSBULKJOBSTATUS
    COMPLETED_PROTECTIONUNITSBULKJOBSTATUS
    COMPLETEDWITHERRORS_PROTECTIONUNITSBULKJOBSTATUS
    UNKNOWNFUTUREVALUE_PROTECTIONUNITSBULKJOBSTATUS
)

func (i ProtectionUnitsBulkJobStatus) String() string {
    return []string{"unknown", "active", "completed", "completedWithErrors", "unknownFutureValue"}[i]
}
func ParseProtectionUnitsBulkJobStatus(v string) (any, error) {
    result := UNKNOWN_PROTECTIONUNITSBULKJOBSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_PROTECTIONUNITSBULKJOBSTATUS
        case "active":
            result = ACTIVE_PROTECTIONUNITSBULKJOBSTATUS
        case "completed":
            result = COMPLETED_PROTECTIONUNITSBULKJOBSTATUS
        case "completedWithErrors":
            result = COMPLETEDWITHERRORS_PROTECTIONUNITSBULKJOBSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PROTECTIONUNITSBULKJOBSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProtectionUnitsBulkJobStatus(values []ProtectionUnitsBulkJobStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProtectionUnitsBulkJobStatus) isMultiValue() bool {
    return false
}
