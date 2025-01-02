package models
type CloudPCTroubleshootReportType int

const (
    TROUBLESHOOTDETAILSREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE CloudPCTroubleshootReportType = iota
    TROUBLESHOOTTRENDCOUNTREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE
    TROUBLESHOOTREGIONALREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCTROUBLESHOOTREPORTTYPE
    TROUBLESHOOTISSUECOUNTREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE
)

func (i CloudPCTroubleshootReportType) String() string {
    return []string{"troubleshootDetailsReport", "troubleshootTrendCountReport", "troubleshootRegionalReport", "unknownFutureValue", "troubleshootIssueCountReport"}[i]
}
func ParseCloudPCTroubleshootReportType(v string) (any, error) {
    result := TROUBLESHOOTDETAILSREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE
    switch v {
        case "troubleshootDetailsReport":
            result = TROUBLESHOOTDETAILSREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE
        case "troubleshootTrendCountReport":
            result = TROUBLESHOOTTRENDCOUNTREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE
        case "troubleshootRegionalReport":
            result = TROUBLESHOOTREGIONALREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCTROUBLESHOOTREPORTTYPE
        case "troubleshootIssueCountReport":
            result = TROUBLESHOOTISSUECOUNTREPORT_CLOUDPCTROUBLESHOOTREPORTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPCTroubleshootReportType(values []CloudPCTroubleshootReportType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPCTroubleshootReportType) isMultiValue() bool {
    return false
}
