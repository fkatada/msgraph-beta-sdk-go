package security
type EventSource int

const (
    SYSTEM_EVENTSOURCE EventSource = iota
    ADMIN_EVENTSOURCE
    USER_EVENTSOURCE
    UNKNOWNFUTUREVALUE_EVENTSOURCE
)

func (i EventSource) String() string {
    return []string{"system", "admin", "user", "unknownFutureValue"}[i]
}
func ParseEventSource(v string) (any, error) {
    result := SYSTEM_EVENTSOURCE
    switch v {
        case "system":
            result = SYSTEM_EVENTSOURCE
        case "admin":
            result = ADMIN_EVENTSOURCE
        case "user":
            result = USER_EVENTSOURCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EVENTSOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEventSource(values []EventSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EventSource) isMultiValue() bool {
    return false
}
