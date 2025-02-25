package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MonthlyPrintUsageSummariesByPrinterCountRequestBuilder provides operations to count the resources in the collection.
type MonthlyPrintUsageSummariesByPrinterCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MonthlyPrintUsageSummariesByPrinterCountRequestBuilderGetQueryParameters get the number of the resource
type MonthlyPrintUsageSummariesByPrinterCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// MonthlyPrintUsageSummariesByPrinterCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonthlyPrintUsageSummariesByPrinterCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MonthlyPrintUsageSummariesByPrinterCountRequestBuilderGetQueryParameters
}
// NewMonthlyPrintUsageSummariesByPrinterCountRequestBuilderInternal instantiates a new MonthlyPrintUsageSummariesByPrinterCountRequestBuilder and sets the default values.
func NewMonthlyPrintUsageSummariesByPrinterCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonthlyPrintUsageSummariesByPrinterCountRequestBuilder) {
    m := &MonthlyPrintUsageSummariesByPrinterCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/monthlyPrintUsageSummariesByPrinter/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewMonthlyPrintUsageSummariesByPrinterCountRequestBuilder instantiates a new MonthlyPrintUsageSummariesByPrinterCountRequestBuilder and sets the default values.
func NewMonthlyPrintUsageSummariesByPrinterCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonthlyPrintUsageSummariesByPrinterCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonthlyPrintUsageSummariesByPrinterCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MonthlyPrintUsageSummariesByPrinterCountRequestBuilder) Get(ctx context.Context, requestConfiguration *MonthlyPrintUsageSummariesByPrinterCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
// returns a *RequestInformation when successful
func (m *MonthlyPrintUsageSummariesByPrinterCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MonthlyPrintUsageSummariesByPrinterCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
// returns a *MonthlyPrintUsageSummariesByPrinterCountRequestBuilder when successful
func (m *MonthlyPrintUsageSummariesByPrinterCountRequestBuilder) WithUrl(rawUrl string)(*MonthlyPrintUsageSummariesByPrinterCountRequestBuilder) {
    return NewMonthlyPrintUsageSummariesByPrinterCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
