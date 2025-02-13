package security
import (
    "math"
    "strings"
)
type AdditionalDataOptions int

const (
    ALLVERSIONS_ADDITIONALDATAOPTIONS = 1
    LINKEDFILES_ADDITIONALDATAOPTIONS = 2
    UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS = 4
    ADVANCEDINDEXING_ADDITIONALDATAOPTIONS = 8
    LISTATTACHMENTS_ADDITIONALDATAOPTIONS = 16
    HTMLTRANSCRIPTS_ADDITIONALDATAOPTIONS = 32
    MESSAGECONVERSATIONEXPANSION_ADDITIONALDATAOPTIONS = 64
    LOCATIONSWITHOUTHITS_ADDITIONALDATAOPTIONS = 128
    ALLITEMSINFOLDER_ADDITIONALDATAOPTIONS = 256
)

func (i AdditionalDataOptions) String() string {
    var values []string
    options := []string{"allVersions", "linkedFiles", "unknownFutureValue", "advancedIndexing", "listAttachments", "htmlTranscripts", "messageConversationExpansion", "locationsWithoutHits", "allItemsInFolder"}
    for p := 0; p < 9; p++ {
        mantis := AdditionalDataOptions(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAdditionalDataOptions(v string) (any, error) {
    var result AdditionalDataOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "allVersions":
                result |= ALLVERSIONS_ADDITIONALDATAOPTIONS
            case "linkedFiles":
                result |= LINKEDFILES_ADDITIONALDATAOPTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS
            case "advancedIndexing":
                result |= ADVANCEDINDEXING_ADDITIONALDATAOPTIONS
            case "listAttachments":
                result |= LISTATTACHMENTS_ADDITIONALDATAOPTIONS
            case "htmlTranscripts":
                result |= HTMLTRANSCRIPTS_ADDITIONALDATAOPTIONS
            case "messageConversationExpansion":
                result |= MESSAGECONVERSATIONEXPANSION_ADDITIONALDATAOPTIONS
            case "locationsWithoutHits":
                result |= LOCATIONSWITHOUTHITS_ADDITIONALDATAOPTIONS
            case "allItemsInFolder":
                result |= ALLITEMSINFOLDER_ADDITIONALDATAOPTIONS
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeAdditionalDataOptions(values []AdditionalDataOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AdditionalDataOptions) isMultiValue() bool {
    return true
}
