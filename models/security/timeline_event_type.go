package security
type TimelineEventType int

const (
    ORIGINALDELIVERY_TIMELINEEVENTTYPE TimelineEventType = iota
    SYSTEMTIMETRAVEL_TIMELINEEVENTTYPE
    DYNAMICDELIVERY_TIMELINEEVENTTYPE
    USERURLCLICK_TIMELINEEVENTTYPE
    REPROCESSED_TIMELINEEVENTTYPE
    ZAP_TIMELINEEVENTTYPE
    QUARANTINERELEASE_TIMELINEEVENTTYPE
    AIR_TIMELINEEVENTTYPE
    UNKNOWN_TIMELINEEVENTTYPE
    UNKNOWNFUTUREVALUE_TIMELINEEVENTTYPE
)

func (i TimelineEventType) String() string {
    return []string{"originalDelivery", "systemTimeTravel", "dynamicDelivery", "userUrlClick", "reprocessed", "zap", "quarantineRelease", "air", "unknown", "unknownFutureValue"}[i]
}
func ParseTimelineEventType(v string) (any, error) {
    result := ORIGINALDELIVERY_TIMELINEEVENTTYPE
    switch v {
        case "originalDelivery":
            result = ORIGINALDELIVERY_TIMELINEEVENTTYPE
        case "systemTimeTravel":
            result = SYSTEMTIMETRAVEL_TIMELINEEVENTTYPE
        case "dynamicDelivery":
            result = DYNAMICDELIVERY_TIMELINEEVENTTYPE
        case "userUrlClick":
            result = USERURLCLICK_TIMELINEEVENTTYPE
        case "reprocessed":
            result = REPROCESSED_TIMELINEEVENTTYPE
        case "zap":
            result = ZAP_TIMELINEEVENTTYPE
        case "quarantineRelease":
            result = QUARANTINERELEASE_TIMELINEEVENTTYPE
        case "air":
            result = AIR_TIMELINEEVENTTYPE
        case "unknown":
            result = UNKNOWN_TIMELINEEVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TIMELINEEVENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTimelineEventType(values []TimelineEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TimelineEventType) isMultiValue() bool {
    return false
}
