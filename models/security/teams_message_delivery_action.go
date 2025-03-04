package security
type TeamsMessageDeliveryAction int

const (
    UNKNOWN_TEAMSMESSAGEDELIVERYACTION TeamsMessageDeliveryAction = iota
    DELIVEREDASSPAM_TEAMSMESSAGEDELIVERYACTION
    DELIVERED_TEAMSMESSAGEDELIVERYACTION
    BLOCKED_TEAMSMESSAGEDELIVERYACTION
    REPLACED_TEAMSMESSAGEDELIVERYACTION
    UNKNOWNFUTUREVALUE_TEAMSMESSAGEDELIVERYACTION
)

func (i TeamsMessageDeliveryAction) String() string {
    return []string{"unknown", "deliveredAsSpam", "delivered", "blocked", "replaced", "unknownFutureValue"}[i]
}
func ParseTeamsMessageDeliveryAction(v string) (any, error) {
    result := UNKNOWN_TEAMSMESSAGEDELIVERYACTION
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMSMESSAGEDELIVERYACTION
        case "deliveredAsSpam":
            result = DELIVEREDASSPAM_TEAMSMESSAGEDELIVERYACTION
        case "delivered":
            result = DELIVERED_TEAMSMESSAGEDELIVERYACTION
        case "blocked":
            result = BLOCKED_TEAMSMESSAGEDELIVERYACTION
        case "replaced":
            result = REPLACED_TEAMSMESSAGEDELIVERYACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSMESSAGEDELIVERYACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTeamsMessageDeliveryAction(values []TeamsMessageDeliveryAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamsMessageDeliveryAction) isMultiValue() bool {
    return false
}
