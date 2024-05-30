package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder provides operations to count the resources in the collection.
type UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderGetQueryParameters get the number of the resource
type UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderGetQueryParameters
}
// NewUserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderInternal instantiates a new UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder and sets the default values.
func NewUserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder) {
    m := &UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/inactiveUsersByApplication/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewUserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder instantiates a new UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder and sets the default values.
func NewUserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder) Get(ctx context.Context, requestConfiguration *UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *RequestInformation when successful
func (m *UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder when successful
func (m *UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder) WithUrl(rawUrl string)(*UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder) {
    return NewUserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
