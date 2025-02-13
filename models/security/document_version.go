package security
type DocumentVersion int

const (
    LATEST_DOCUMENTVERSION DocumentVersion = iota
    RECENT10_DOCUMENTVERSION
    RECENT100_DOCUMENTVERSION
    ALL_DOCUMENTVERSION
    UNKNOWNFUTUREVALUE_DOCUMENTVERSION
)

func (i DocumentVersion) String() string {
    return []string{"latest", "recent10", "recent100", "all", "unknownFutureValue"}[i]
}
func ParseDocumentVersion(v string) (any, error) {
    result := LATEST_DOCUMENTVERSION
    switch v {
        case "latest":
            result = LATEST_DOCUMENTVERSION
        case "recent10":
            result = RECENT10_DOCUMENTVERSION
        case "recent100":
            result = RECENT100_DOCUMENTVERSION
        case "all":
            result = ALL_DOCUMENTVERSION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DOCUMENTVERSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDocumentVersion(values []DocumentVersion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DocumentVersion) isMultiValue() bool {
    return false
}
