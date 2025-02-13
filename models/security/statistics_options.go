package security
import (
    "math"
    "strings"
)
type StatisticsOptions int

const (
    INCLUDEREFINERS_STATISTICSOPTIONS = 1
    INCLUDEQUERYSTATS_STATISTICSOPTIONS = 2
    INCLUDEUNINDEXEDSTATS_STATISTICSOPTIONS = 4
    ADVANCEDINDEXING_STATISTICSOPTIONS = 8
    LOCATIONSWITHOUTHITS_STATISTICSOPTIONS = 16
    UNKNOWNFUTUREVALUE_STATISTICSOPTIONS = 32
)

func (i StatisticsOptions) String() string {
    var values []string
    options := []string{"includeRefiners", "includeQueryStats", "includeUnindexedStats", "advancedIndexing", "locationsWithoutHits", "unknownFutureValue"}
    for p := 0; p < 6; p++ {
        mantis := StatisticsOptions(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseStatisticsOptions(v string) (any, error) {
    var result StatisticsOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "includeRefiners":
                result |= INCLUDEREFINERS_STATISTICSOPTIONS
            case "includeQueryStats":
                result |= INCLUDEQUERYSTATS_STATISTICSOPTIONS
            case "includeUnindexedStats":
                result |= INCLUDEUNINDEXEDSTATS_STATISTICSOPTIONS
            case "advancedIndexing":
                result |= ADVANCEDINDEXING_STATISTICSOPTIONS
            case "locationsWithoutHits":
                result |= LOCATIONSWITHOUTHITS_STATISTICSOPTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_STATISTICSOPTIONS
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeStatisticsOptions(values []StatisticsOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StatisticsOptions) isMultiValue() bool {
    return true
}
