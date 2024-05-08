package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppManagementPoliciesItemAppliesToRequestBuilder provides operations to manage the appliesTo property of the microsoft.graph.appManagementPolicy entity.
type AppManagementPoliciesItemAppliesToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppManagementPoliciesItemAppliesToRequestBuilderGetQueryParameters list application and service principal objects assigned an appManagementPolicy policy object.
type AppManagementPoliciesItemAppliesToRequestBuilderGetQueryParameters struct {
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
// AppManagementPoliciesItemAppliesToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppManagementPoliciesItemAppliesToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppManagementPoliciesItemAppliesToRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the appliesTo property of the microsoft.graph.appManagementPolicy entity.
// returns a *AppManagementPoliciesItemAppliesToDirectoryObjectItemRequestBuilder when successful
func (m *AppManagementPoliciesItemAppliesToRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*AppManagementPoliciesItemAppliesToDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewAppManagementPoliciesItemAppliesToDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAppManagementPoliciesItemAppliesToRequestBuilderInternal instantiates a new AppManagementPoliciesItemAppliesToRequestBuilder and sets the default values.
func NewAppManagementPoliciesItemAppliesToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManagementPoliciesItemAppliesToRequestBuilder) {
    m := &AppManagementPoliciesItemAppliesToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/appManagementPolicies/{appManagementPolicy%2Did}/appliesTo{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAppManagementPoliciesItemAppliesToRequestBuilder instantiates a new AppManagementPoliciesItemAppliesToRequestBuilder and sets the default values.
func NewAppManagementPoliciesItemAppliesToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManagementPoliciesItemAppliesToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppManagementPoliciesItemAppliesToRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AppManagementPoliciesItemAppliesToCountRequestBuilder when successful
func (m *AppManagementPoliciesItemAppliesToRequestBuilder) Count()(*AppManagementPoliciesItemAppliesToCountRequestBuilder) {
    return NewAppManagementPoliciesItemAppliesToCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list application and service principal objects assigned an appManagementPolicy policy object.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/appmanagementpolicy-list-appliesto?view=graph-rest-beta
func (m *AppManagementPoliciesItemAppliesToRequestBuilder) Get(ctx context.Context, requestConfiguration *AppManagementPoliciesItemAppliesToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
// ToGetRequestInformation list application and service principal objects assigned an appManagementPolicy policy object.
// returns a *RequestInformation when successful
func (m *AppManagementPoliciesItemAppliesToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppManagementPoliciesItemAppliesToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppManagementPoliciesItemAppliesToRequestBuilder when successful
func (m *AppManagementPoliciesItemAppliesToRequestBuilder) WithUrl(rawUrl string)(*AppManagementPoliciesItemAppliesToRequestBuilder) {
    return NewAppManagementPoliciesItemAppliesToRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
