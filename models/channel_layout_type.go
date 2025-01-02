package models
type ChannelLayoutType int

const (
    POST_CHANNELLAYOUTTYPE ChannelLayoutType = iota
    CHAT_CHANNELLAYOUTTYPE
    UNKNOWNFUTUREVALUE_CHANNELLAYOUTTYPE
)

func (i ChannelLayoutType) String() string {
    return []string{"post", "chat", "unknownFutureValue"}[i]
}
func ParseChannelLayoutType(v string) (any, error) {
    result := POST_CHANNELLAYOUTTYPE
    switch v {
        case "post":
            result = POST_CHANNELLAYOUTTYPE
        case "chat":
            result = CHAT_CHANNELLAYOUTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CHANNELLAYOUTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeChannelLayoutType(values []ChannelLayoutType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ChannelLayoutType) isMultiValue() bool {
    return false
}
