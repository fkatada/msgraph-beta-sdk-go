package networkaccess

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the crossTenantAccessReport method.
type ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters invoke function crossTenantAccessReport
type ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/reports/microsoft.graph.networkaccess.crossTenantAccessReport(startDateTime={startDateTime},endDateTime={endDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function crossTenantAccessReport
// Deprecated: This method is obsolete. Use GetAsCrossTenantAccessReportWithStartDateTimeWithEndDateTimeGetResponse instead.
// returns a ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeCrossTenantAccessReportWithStartDateTimeWithEndDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeCrossTenantAccessReportWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeCrossTenantAccessReportWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeCrossTenantAccessReportWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsCrossTenantAccessReportWithStartDateTimeWithEndDateTimeGetResponse invoke function crossTenantAccessReport
// returns a ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeCrossTenantAccessReportWithStartDateTimeWithEndDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsCrossTenantAccessReportWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeCrossTenantAccessReportWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeCrossTenantAccessReportWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeCrossTenantAccessReportWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation invoke function crossTenantAccessReport
// returns a *RequestInformation when successful
func (m *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*ReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewReportsMicrosoftgraphnetworkaccesscrosstenantaccessreportwithstartdatetimewithenddatetimeMicrosoftGraphNetworkaccessCrossTenantAccessReportWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
