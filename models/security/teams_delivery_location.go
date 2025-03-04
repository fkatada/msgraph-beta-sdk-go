package security
type TeamsDeliveryLocation int

const (
    UNKNOWN_TEAMSDELIVERYLOCATION TeamsDeliveryLocation = iota
    TEAMS_TEAMSDELIVERYLOCATION
    QUARANTINE_TEAMSDELIVERYLOCATION
    FAILED_TEAMSDELIVERYLOCATION
    UNKNOWNFUTUREVALUE_TEAMSDELIVERYLOCATION
)

func (i TeamsDeliveryLocation) String() string {
    return []string{"unknown", "teams", "quarantine", "failed", "unknownFutureValue"}[i]
}
func ParseTeamsDeliveryLocation(v string) (any, error) {
    result := UNKNOWN_TEAMSDELIVERYLOCATION
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMSDELIVERYLOCATION
        case "teams":
            result = TEAMS_TEAMSDELIVERYLOCATION
        case "quarantine":
            result = QUARANTINE_TEAMSDELIVERYLOCATION
        case "failed":
            result = FAILED_TEAMSDELIVERYLOCATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSDELIVERYLOCATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTeamsDeliveryLocation(values []TeamsDeliveryLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamsDeliveryLocation) isMultiValue() bool {
    return false
}
