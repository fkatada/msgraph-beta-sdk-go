package models
type SignUpIdentifierType int

const (
    EMAILADDRESS_SIGNUPIDENTIFIERTYPE SignUpIdentifierType = iota
    UNKNOWNFUTUREVALUE_SIGNUPIDENTIFIERTYPE
)

func (i SignUpIdentifierType) String() string {
    return []string{"emailAddress", "unknownFutureValue"}[i]
}
func ParseSignUpIdentifierType(v string) (any, error) {
    result := EMAILADDRESS_SIGNUPIDENTIFIERTYPE
    switch v {
        case "emailAddress":
            result = EMAILADDRESS_SIGNUPIDENTIFIERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SIGNUPIDENTIFIERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSignUpIdentifierType(values []SignUpIdentifierType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SignUpIdentifierType) isMultiValue() bool {
    return false
}
