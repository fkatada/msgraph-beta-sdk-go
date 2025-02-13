package security
import (
    "math"
    "strings"
)
type AdditionalOptions int

const (
    NONE_ADDITIONALOPTIONS = 1
    TEAMSANDYAMMERCONVERSATIONS_ADDITIONALOPTIONS = 2
    CLOUDATTACHMENTS_ADDITIONALOPTIONS = 4
    ALLDOCUMENTVERSIONS_ADDITIONALOPTIONS = 8
    SUBFOLDERCONTENTS_ADDITIONALOPTIONS = 16
    LISTATTACHMENTS_ADDITIONALOPTIONS = 32
    UNKNOWNFUTUREVALUE_ADDITIONALOPTIONS = 64
    HTMLTRANSCRIPTS_ADDITIONALOPTIONS = 128
    ADVANCEDINDEXING_ADDITIONALOPTIONS = 256
    ALLITEMSINFOLDER_ADDITIONALOPTIONS = 512
    INCLUDEFOLDERANDPATH_ADDITIONALOPTIONS = 1024
    CONDENSEPATHS_ADDITIONALOPTIONS = 2048
    FRIENDLYNAME_ADDITIONALOPTIONS = 4096
    SPLITSOURCE_ADDITIONALOPTIONS = 8192
    OPTIMIZEDPARTITIONSIZE_ADDITIONALOPTIONS = 16384
    INCLUDEREPORT_ADDITIONALOPTIONS = 32768
)

func (i AdditionalOptions) String() string {
    var values []string
    options := []string{"none", "teamsAndYammerConversations", "cloudAttachments", "allDocumentVersions", "subfolderContents", "listAttachments", "unknownFutureValue", "htmlTranscripts", "advancedIndexing", "allItemsInFolder", "includeFolderAndPath", "condensePaths", "friendlyName", "splitSource", "optimizedPartitionSize", "includeReport"}
    for p := 0; p < 16; p++ {
        mantis := AdditionalOptions(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAdditionalOptions(v string) (any, error) {
    var result AdditionalOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_ADDITIONALOPTIONS
            case "teamsAndYammerConversations":
                result |= TEAMSANDYAMMERCONVERSATIONS_ADDITIONALOPTIONS
            case "cloudAttachments":
                result |= CLOUDATTACHMENTS_ADDITIONALOPTIONS
            case "allDocumentVersions":
                result |= ALLDOCUMENTVERSIONS_ADDITIONALOPTIONS
            case "subfolderContents":
                result |= SUBFOLDERCONTENTS_ADDITIONALOPTIONS
            case "listAttachments":
                result |= LISTATTACHMENTS_ADDITIONALOPTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ADDITIONALOPTIONS
            case "htmlTranscripts":
                result |= HTMLTRANSCRIPTS_ADDITIONALOPTIONS
            case "advancedIndexing":
                result |= ADVANCEDINDEXING_ADDITIONALOPTIONS
            case "allItemsInFolder":
                result |= ALLITEMSINFOLDER_ADDITIONALOPTIONS
            case "includeFolderAndPath":
                result |= INCLUDEFOLDERANDPATH_ADDITIONALOPTIONS
            case "condensePaths":
                result |= CONDENSEPATHS_ADDITIONALOPTIONS
            case "friendlyName":
                result |= FRIENDLYNAME_ADDITIONALOPTIONS
            case "splitSource":
                result |= SPLITSOURCE_ADDITIONALOPTIONS
            case "optimizedPartitionSize":
                result |= OPTIMIZEDPARTITIONSIZE_ADDITIONALOPTIONS
            case "includeReport":
                result |= INCLUDEREPORT_ADDITIONALOPTIONS
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeAdditionalOptions(values []AdditionalOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AdditionalOptions) isMultiValue() bool {
    return true
}
