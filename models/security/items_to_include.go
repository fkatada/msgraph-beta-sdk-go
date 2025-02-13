package security
import (
    "math"
    "strings"
)
type ItemsToInclude int

const (
    SEARCHHITS_ITEMSTOINCLUDE = 1
    PARTIALLYINDEXED_ITEMSTOINCLUDE = 2
    UNKNOWNFUTUREVALUE_ITEMSTOINCLUDE = 4
)

func (i ItemsToInclude) String() string {
    var values []string
    options := []string{"searchHits", "partiallyIndexed", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := ItemsToInclude(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseItemsToInclude(v string) (any, error) {
    var result ItemsToInclude
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "searchHits":
                result |= SEARCHHITS_ITEMSTOINCLUDE
            case "partiallyIndexed":
                result |= PARTIALLYINDEXED_ITEMSTOINCLUDE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ITEMSTOINCLUDE
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeItemsToInclude(values []ItemsToInclude) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ItemsToInclude) isMultiValue() bool {
    return true
}
