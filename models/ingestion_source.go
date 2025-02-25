package models
// Category Ingestion source
type IngestionSource int

const (
    // Indicates unknown category
    UNKNOWN_INGESTIONSOURCE IngestionSource = iota
    // Indicates the category is ingested by IT admin with sufficient permissions through custom ingestion process
    CUSTOM_INGESTIONSOURCE
    // Indicates the category is ingested through system ingestion process
    BUILTIN_INGESTIONSOURCE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_INGESTIONSOURCE
)

func (i IngestionSource) String() string {
    return []string{"unknown", "custom", "builtIn", "unknownFutureValue"}[i]
}
func ParseIngestionSource(v string) (any, error) {
    result := UNKNOWN_INGESTIONSOURCE
    switch v {
        case "unknown":
            result = UNKNOWN_INGESTIONSOURCE
        case "custom":
            result = CUSTOM_INGESTIONSOURCE
        case "builtIn":
            result = BUILTIN_INGESTIONSOURCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INGESTIONSOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIngestionSource(values []IngestionSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IngestionSource) isMultiValue() bool {
    return false
}
