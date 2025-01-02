package security
type VerdictCategory int

const (
    NONE_VERDICTCATEGORY VerdictCategory = iota
    MALWARE_VERDICTCATEGORY
    PHISH_VERDICTCATEGORY
    SITEUNAVAILABLE_VERDICTCATEGORY
    SPAM_VERDICTCATEGORY
    DECRYPTIONFAILED_VERDICTCATEGORY
    UNSUPPORTEDURISCHEME_VERDICTCATEGORY
    UNSUPPORTEDFILETYPE_VERDICTCATEGORY
    UNDEFINED_VERDICTCATEGORY
    UNKNOWNFUTUREVALUE_VERDICTCATEGORY
)

func (i VerdictCategory) String() string {
    return []string{"none", "malware", "phish", "siteUnavailable", "spam", "decryptionFailed", "unsupportedUriScheme", "unsupportedFileType", "undefined", "unknownFutureValue"}[i]
}
func ParseVerdictCategory(v string) (any, error) {
    result := NONE_VERDICTCATEGORY
    switch v {
        case "none":
            result = NONE_VERDICTCATEGORY
        case "malware":
            result = MALWARE_VERDICTCATEGORY
        case "phish":
            result = PHISH_VERDICTCATEGORY
        case "siteUnavailable":
            result = SITEUNAVAILABLE_VERDICTCATEGORY
        case "spam":
            result = SPAM_VERDICTCATEGORY
        case "decryptionFailed":
            result = DECRYPTIONFAILED_VERDICTCATEGORY
        case "unsupportedUriScheme":
            result = UNSUPPORTEDURISCHEME_VERDICTCATEGORY
        case "unsupportedFileType":
            result = UNSUPPORTEDFILETYPE_VERDICTCATEGORY
        case "undefined":
            result = UNDEFINED_VERDICTCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_VERDICTCATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVerdictCategory(values []VerdictCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VerdictCategory) isMultiValue() bool {
    return false
}
