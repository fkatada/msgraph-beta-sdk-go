package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder provides operations to manage the permissionsCreepIndexDistributions property of the microsoft.graph.permissionsAnalytics entity.
type PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderGetQueryParameters get a list of the permissionsCreepIndexDistribution objects and their properties.
type PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderGetQueryParameters struct {
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
// PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderGetQueryParameters
}
// PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPermissionsCreepIndexDistributionId provides operations to manage the permissionsCreepIndexDistributions property of the microsoft.graph.permissionsAnalytics entity.
// returns a *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder when successful
func (m *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) ByPermissionsCreepIndexDistributionId(permissionsCreepIndexDistributionId string)(*PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if permissionsCreepIndexDistributionId != "" {
        urlTplParams["permissionsCreepIndexDistribution%2Did"] = permissionsCreepIndexDistributionId
    }
    return NewPermissionsAnalyticsAwsPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderInternal instantiates a new PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder and sets the default values.
func NewPermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) {
    m := &PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder instantiates a new PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder and sets the default values.
func NewPermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsCountRequestBuilder when successful
func (m *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) Count()(*PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsCountRequestBuilder) {
    return NewPermissionsAnalyticsAwsPermissionsCreepIndexDistributionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the permissionsCreepIndexDistribution objects and their properties.
// returns a PermissionsCreepIndexDistributionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissionsanalytics-list-permissionscreepindexdistributions?view=graph-rest-beta
func (m *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionCollectionResponseable), nil
}
// Post create new navigation property to permissionsCreepIndexDistributions for identityGovernance
// returns a PermissionsCreepIndexDistributionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable), nil
}
// ToGetRequestInformation get a list of the permissionsCreepIndexDistribution objects and their properties.
// returns a *RequestInformation when successful
func (m *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to permissionsCreepIndexDistributions for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder when successful
func (m *PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) WithUrl(rawUrl string)(*PermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder) {
    return NewPermissionsAnalyticsAwsPermissionsCreepIndexDistributionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
