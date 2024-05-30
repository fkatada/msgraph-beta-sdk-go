package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder provides operations to manage the claimsMappingPolicies property of the microsoft.graph.servicePrincipal entity.
type ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderGetQueryParameters list the claimsMappingPolicy objects that are assigned to a servicePrincipal.
type ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderGetQueryParameters struct {
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
// ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderGetQueryParameters
}
// ByClaimsMappingPolicyId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.claimsMappingPolicies.item collection
// returns a *ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder when successful
func (m *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) ByClaimsMappingPolicyId(claimsMappingPolicyId string)(*ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if claimsMappingPolicyId != "" {
        urlTplParams["claimsMappingPolicy%2Did"] = claimsMappingPolicyId
    }
    return NewItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderInternal instantiates a new ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder and sets the default values.
func NewItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) {
    m := &ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/claimsMappingPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder instantiates a new ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder and sets the default values.
func NewItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemClaimsmappingpoliciesCountRequestBuilder when successful
func (m *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) Count()(*ItemClaimsmappingpoliciesCountRequestBuilder) {
    return NewItemClaimsmappingpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the claimsMappingPolicy objects that are assigned to a servicePrincipal.
// returns a ClaimsMappingPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-list-claimsmappingpolicies?view=graph-rest-beta
func (m *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ClaimsMappingPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateClaimsMappingPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ClaimsMappingPolicyCollectionResponseable), nil
}
// Ref provides operations to manage the collection of servicePrincipal entities.
// returns a *ItemClaimsmappingpoliciesRefRequestBuilder when successful
func (m *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) Ref()(*ItemClaimsmappingpoliciesRefRequestBuilder) {
    return NewItemClaimsmappingpoliciesRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list the claimsMappingPolicy objects that are assigned to a servicePrincipal.
// returns a *RequestInformation when successful
func (m *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder when successful
func (m *ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder) {
    return NewItemClaimsmappingpoliciesClaimsMappingPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
