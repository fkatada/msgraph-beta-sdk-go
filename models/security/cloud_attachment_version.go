package security
type CloudAttachmentVersion int

const (
    LATEST_CLOUDATTACHMENTVERSION CloudAttachmentVersion = iota
    RECENT10_CLOUDATTACHMENTVERSION
    RECENT100_CLOUDATTACHMENTVERSION
    ALL_CLOUDATTACHMENTVERSION
    UNKNOWNFUTUREVALUE_CLOUDATTACHMENTVERSION
)

func (i CloudAttachmentVersion) String() string {
    return []string{"latest", "recent10", "recent100", "all", "unknownFutureValue"}[i]
}
func ParseCloudAttachmentVersion(v string) (any, error) {
    result := LATEST_CLOUDATTACHMENTVERSION
    switch v {
        case "latest":
            result = LATEST_CLOUDATTACHMENTVERSION
        case "recent10":
            result = RECENT10_CLOUDATTACHMENTVERSION
        case "recent100":
            result = RECENT100_CLOUDATTACHMENTVERSION
        case "all":
            result = ALL_CLOUDATTACHMENTVERSION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDATTACHMENTVERSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudAttachmentVersion(values []CloudAttachmentVersion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudAttachmentVersion) isMultiValue() bool {
    return false
}
