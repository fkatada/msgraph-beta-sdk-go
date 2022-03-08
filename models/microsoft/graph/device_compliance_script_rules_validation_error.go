package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the validateComplianceScript method.
type DeviceComplianceScriptRulesValidationError int

const (
    NONE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR DeviceComplianceScriptRulesValidationError = iota
    JSONFILEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    JSONFILEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    JSONFILETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    RULESMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    DUPLICATERULES_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    TOOMANYRULESSPECIFIED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    OPERATORMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    OPERATORNOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    DATATYPEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    DATATYPENOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    OPERATORDATATYPECOMBINATIONNOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    MOREINFOURIMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    MOREINFOURIINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    MOREINFOURITOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    DESCRIPTIONMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    DESCRIPTIONINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    DESCRIPTIONTOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    TITLEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    TITLEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    TITLETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    OPERANDMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    OPERANDINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    OPERANDTOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    SETTINGNAMEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    SETTINGNAMEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    SETTINGNAMETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    ENGLISHLOCALEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    DUPLICATELOCALES_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    UNRECOGNIZEDLOCALE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    UNKNOWN_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    REMEDIATIONSTRINGSMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
)

func (i DeviceComplianceScriptRulesValidationError) String() string {
    return []string{"NONE", "JSONFILEINVALID", "JSONFILEMISSING", "JSONFILETOOLARGE", "RULESMISSING", "DUPLICATERULES", "TOOMANYRULESSPECIFIED", "OPERATORMISSING", "OPERATORNOTSUPPORTED", "DATATYPEMISSING", "DATATYPENOTSUPPORTED", "OPERATORDATATYPECOMBINATIONNOTSUPPORTED", "MOREINFOURIMISSING", "MOREINFOURIINVALID", "MOREINFOURITOOLARGE", "DESCRIPTIONMISSING", "DESCRIPTIONINVALID", "DESCRIPTIONTOOLARGE", "TITLEMISSING", "TITLEINVALID", "TITLETOOLARGE", "OPERANDMISSING", "OPERANDINVALID", "OPERANDTOOLARGE", "SETTINGNAMEMISSING", "SETTINGNAMEINVALID", "SETTINGNAMETOOLARGE", "ENGLISHLOCALEMISSING", "DUPLICATELOCALES", "UNRECOGNIZEDLOCALE", "UNKNOWN", "REMEDIATIONSTRINGSMISSING"}[i]
}
func ParseDeviceComplianceScriptRulesValidationError(v string) (interface{}, error) {
    result := NONE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "JSONFILEINVALID":
            result = JSONFILEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "JSONFILEMISSING":
            result = JSONFILEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "JSONFILETOOLARGE":
            result = JSONFILETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "RULESMISSING":
            result = RULESMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "DUPLICATERULES":
            result = DUPLICATERULES_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "TOOMANYRULESSPECIFIED":
            result = TOOMANYRULESSPECIFIED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "OPERATORMISSING":
            result = OPERATORMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "OPERATORNOTSUPPORTED":
            result = OPERATORNOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "DATATYPEMISSING":
            result = DATATYPEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "DATATYPENOTSUPPORTED":
            result = DATATYPENOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "OPERATORDATATYPECOMBINATIONNOTSUPPORTED":
            result = OPERATORDATATYPECOMBINATIONNOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "MOREINFOURIMISSING":
            result = MOREINFOURIMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "MOREINFOURIINVALID":
            result = MOREINFOURIINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "MOREINFOURITOOLARGE":
            result = MOREINFOURITOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "DESCRIPTIONMISSING":
            result = DESCRIPTIONMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "DESCRIPTIONINVALID":
            result = DESCRIPTIONINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "DESCRIPTIONTOOLARGE":
            result = DESCRIPTIONTOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "TITLEMISSING":
            result = TITLEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "TITLEINVALID":
            result = TITLEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "TITLETOOLARGE":
            result = TITLETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "OPERANDMISSING":
            result = OPERANDMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "OPERANDINVALID":
            result = OPERANDINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "OPERANDTOOLARGE":
            result = OPERANDTOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "SETTINGNAMEMISSING":
            result = SETTINGNAMEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "SETTINGNAMEINVALID":
            result = SETTINGNAMEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "SETTINGNAMETOOLARGE":
            result = SETTINGNAMETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "ENGLISHLOCALEMISSING":
            result = ENGLISHLOCALEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "DUPLICATELOCALES":
            result = DUPLICATELOCALES_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "UNRECOGNIZEDLOCALE":
            result = UNRECOGNIZEDLOCALE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "UNKNOWN":
            result = UNKNOWN_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "REMEDIATIONSTRINGSMISSING":
            result = REMEDIATIONSTRINGSMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        default:
            return 0, errors.New("Unknown DeviceComplianceScriptRulesValidationError value: " + v)
    }
    return &result, nil
}
func SerializeDeviceComplianceScriptRulesValidationError(values []DeviceComplianceScriptRulesValidationError) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
