package security
type AntispamTeamsDirection int

const (
    UNKNOWN_ANTISPAMTEAMSDIRECTION AntispamTeamsDirection = iota
    INBOUND_ANTISPAMTEAMSDIRECTION
    OUTBOUND_ANTISPAMTEAMSDIRECTION
    INTRAORG_ANTISPAMTEAMSDIRECTION
    UNKNOWNFUTUREVALUE_ANTISPAMTEAMSDIRECTION
)

func (i AntispamTeamsDirection) String() string {
    return []string{"unknown", "inbound", "outbound", "intraorg", "unknownFutureValue"}[i]
}
func ParseAntispamTeamsDirection(v string) (any, error) {
    result := UNKNOWN_ANTISPAMTEAMSDIRECTION
    switch v {
        case "unknown":
            result = UNKNOWN_ANTISPAMTEAMSDIRECTION
        case "inbound":
            result = INBOUND_ANTISPAMTEAMSDIRECTION
        case "outbound":
            result = OUTBOUND_ANTISPAMTEAMSDIRECTION
        case "intraorg":
            result = INTRAORG_ANTISPAMTEAMSDIRECTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANTISPAMTEAMSDIRECTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAntispamTeamsDirection(values []AntispamTeamsDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AntispamTeamsDirection) isMultiValue() bool {
    return false
}
