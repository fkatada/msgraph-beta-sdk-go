package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder provides operations to call the getYammerDeviceUsageUserCounts method.
type GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal instantiates a new GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    m := &GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getYammerDeviceUsageUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder instantiates a new GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getYammerDeviceUsageUserCounts
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getYammerDeviceUsageUserCounts
// returns a *RequestInformation when successful
func (m *GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder when successful
func (m *GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewGetyammerdeviceusageusercountswithperiodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
