package models
type CertificateAuthorityType int

const (
    ROOT_CERTIFICATEAUTHORITYTYPE CertificateAuthorityType = iota
    INTERMEDIATE_CERTIFICATEAUTHORITYTYPE
    UNKNOWNFUTUREVALUE_CERTIFICATEAUTHORITYTYPE
)

func (i CertificateAuthorityType) String() string {
    return []string{"root", "intermediate", "unknownFutureValue"}[i]
}
func ParseCertificateAuthorityType(v string) (any, error) {
    result := ROOT_CERTIFICATEAUTHORITYTYPE
    switch v {
        case "root":
            result = ROOT_CERTIFICATEAUTHORITYTYPE
        case "intermediate":
            result = INTERMEDIATE_CERTIFICATEAUTHORITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CERTIFICATEAUTHORITYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCertificateAuthorityType(values []CertificateAuthorityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CertificateAuthorityType) isMultiValue() bool {
    return false
}
