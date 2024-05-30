package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder provides operations to manage the appliesTo property of the microsoft.graph.stsPolicy entity.
type TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TokenissuancepoliciesItemAppliestoAppliesToRequestBuilderGetQueryParameters get a list of directoryObject objects that a tokenIssuancePolicy object has been applied to. The tokenIssuancePolicy can only be applied to application and servicePrincipal resources.
type TokenissuancepoliciesItemAppliestoAppliesToRequestBuilderGetQueryParameters struct {
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
// TokenissuancepoliciesItemAppliestoAppliesToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TokenissuancepoliciesItemAppliestoAppliesToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the appliesTo property of the microsoft.graph.stsPolicy entity.
// returns a *TokenissuancepoliciesItemAppliestoDirectoryObjectItemRequestBuilder when successful
func (m *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*TokenissuancepoliciesItemAppliestoDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewTokenissuancepoliciesItemAppliestoDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTokenissuancepoliciesItemAppliestoAppliesToRequestBuilderInternal instantiates a new TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder and sets the default values.
func NewTokenissuancepoliciesItemAppliestoAppliesToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder) {
    m := &TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/tokenIssuancePolicies/{tokenIssuancePolicy%2Did}/appliesTo{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTokenissuancepoliciesItemAppliestoAppliesToRequestBuilder instantiates a new TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder and sets the default values.
func NewTokenissuancepoliciesItemAppliestoAppliesToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTokenissuancepoliciesItemAppliestoAppliesToRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TokenissuancepoliciesItemAppliestoCountRequestBuilder when successful
func (m *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder) Count()(*TokenissuancepoliciesItemAppliestoCountRequestBuilder) {
    return NewTokenissuancepoliciesItemAppliestoCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of directoryObject objects that a tokenIssuancePolicy object has been applied to. The tokenIssuancePolicy can only be applied to application and servicePrincipal resources.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tokenissuancepolicy-list-appliesto?view=graph-rest-beta
func (m *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder) Get(ctx context.Context, requestConfiguration *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// ToGetRequestInformation get a list of directoryObject objects that a tokenIssuancePolicy object has been applied to. The tokenIssuancePolicy can only be applied to application and servicePrincipal resources.
// returns a *RequestInformation when successful
func (m *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder when successful
func (m *TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder) WithUrl(rawUrl string)(*TokenissuancepoliciesItemAppliestoAppliesToRequestBuilder) {
    return NewTokenissuancepoliciesItemAppliestoAppliesToRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
