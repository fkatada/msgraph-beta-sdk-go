package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder provides operations to manage the inactiveUsersByApplication property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyInactiveUsersByApplicationRequestBuilderGetQueryParameters get inactiveUsersByApplication from reports
type UserInsightsMonthlyInactiveUsersByApplicationRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UserInsightsMonthlyInactiveUsersByApplicationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyInactiveUsersByApplicationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilderGetQueryParameters
}
// ByMonthlyInactiveUsersByApplicationMetricId provides operations to manage the inactiveUsersByApplication property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs on 2024-02-02 and will be removed 2024-02-16
// returns a *UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder when successful
func (m *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder) ByMonthlyInactiveUsersByApplicationMetricId(monthlyInactiveUsersByApplicationMetricId string)(*UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if monthlyInactiveUsersByApplicationMetricId != "" {
        urlTplParams["monthlyInactiveUsersByApplicationMetric%2Did"] = monthlyInactiveUsersByApplicationMetricId
    }
    return NewUserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserInsightsMonthlyInactiveUsersByApplicationRequestBuilderInternal instantiates a new UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder and sets the default values.
func NewUserInsightsMonthlyInactiveUsersByApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder) {
    m := &UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/inactiveUsersByApplication{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyInactiveUsersByApplicationRequestBuilder instantiates a new UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder and sets the default values.
func NewUserInsightsMonthlyInactiveUsersByApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyInactiveUsersByApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserInsightsMonthlyInactiveUsersByApplicationCountRequestBuilder when successful
func (m *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder) Count()(*UserInsightsMonthlyInactiveUsersByApplicationCountRequestBuilder) {
    return NewUserInsightsMonthlyInactiveUsersByApplicationCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get inactiveUsersByApplication from reports
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs on 2024-02-02 and will be removed 2024-02-16
// returns a MonthlyInactiveUsersByApplicationMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyInactiveUsersByApplicationMetricCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMonthlyInactiveUsersByApplicationMetricCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyInactiveUsersByApplicationMetricCollectionResponseable), nil
}
// ToGetRequestInformation get inactiveUsersByApplication from reports
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs on 2024-02-02 and will be removed 2024-02-16
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs on 2024-02-02 and will be removed 2024-02-16
// returns a *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder when successful
func (m *UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyInactiveUsersByApplicationRequestBuilder) {
    return NewUserInsightsMonthlyInactiveUsersByApplicationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
