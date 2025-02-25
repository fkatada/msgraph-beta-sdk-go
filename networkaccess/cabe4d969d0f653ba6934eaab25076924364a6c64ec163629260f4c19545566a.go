package networkaccess

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the getAlertFrequencies method.
type AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters invoke function getAlertFrequencies
type AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
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
// AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewAlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewAlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/alerts/microsoft.graph.networkaccess.getAlertFrequencies(startDateTime={startDateTime},endDateTime={endDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewAlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewAlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getAlertFrequencies
// Deprecated: This method is obsolete. Use GetAsGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetResponse instead.
// returns a AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetAlertFrequenciesWithStartDateTimeWithEndDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetAlertFrequenciesWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetAlertFrequenciesWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetAlertFrequenciesWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetResponse invoke function getAlertFrequencies
// Deprecated:  as of 2022-06/PrivatePreview:NetworkAccess on 2024-09-09 and will be removed 2024-11-01
// returns a AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetAlertFrequenciesWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation invoke function getAlertFrequencies
// Deprecated:  as of 2022-06/PrivatePreview:NetworkAccess on 2024-09-09 and will be removed 2024-11-01
// returns a *RequestInformation when successful
func (m *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2022-06/PrivatePreview:NetworkAccess on 2024-09-09 and will be removed 2024-11-01
// returns a *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*AlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewAlertsMicrosoftGraphNetworkaccessGetAlertFrequenciesWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
