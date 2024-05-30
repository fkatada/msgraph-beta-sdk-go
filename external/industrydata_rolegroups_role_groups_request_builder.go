package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataRolegroupsRoleGroupsRequestBuilder provides operations to manage the roleGroups property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustrydataRolegroupsRoleGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataRolegroupsRoleGroupsRequestBuilderGetQueryParameters get a list of the roleGroup objects and their properties.
type IndustrydataRolegroupsRoleGroupsRequestBuilderGetQueryParameters struct {
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
// IndustrydataRolegroupsRoleGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataRolegroupsRoleGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataRolegroupsRoleGroupsRequestBuilderGetQueryParameters
}
// IndustrydataRolegroupsRoleGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataRolegroupsRoleGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRoleGroupId provides operations to manage the roleGroups property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataRolegroupsRoleGroupItemRequestBuilder when successful
func (m *IndustrydataRolegroupsRoleGroupsRequestBuilder) ByRoleGroupId(roleGroupId string)(*IndustrydataRolegroupsRoleGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if roleGroupId != "" {
        urlTplParams["roleGroup%2Did"] = roleGroupId
    }
    return NewIndustrydataRolegroupsRoleGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustrydataRolegroupsRoleGroupsRequestBuilderInternal instantiates a new IndustrydataRolegroupsRoleGroupsRequestBuilder and sets the default values.
func NewIndustrydataRolegroupsRoleGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRolegroupsRoleGroupsRequestBuilder) {
    m := &IndustrydataRolegroupsRoleGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/roleGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIndustrydataRolegroupsRoleGroupsRequestBuilder instantiates a new IndustrydataRolegroupsRoleGroupsRequestBuilder and sets the default values.
func NewIndustrydataRolegroupsRoleGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRolegroupsRoleGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataRolegroupsRoleGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IndustrydataRolegroupsCountRequestBuilder when successful
func (m *IndustrydataRolegroupsRoleGroupsRequestBuilder) Count()(*IndustrydataRolegroupsCountRequestBuilder) {
    return NewIndustrydataRolegroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the roleGroup objects and their properties.
// returns a RoleGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-rolegroup-list?view=graph-rest-beta
func (m *IndustrydataRolegroupsRoleGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataRolegroupsRoleGroupsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.RoleGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateRoleGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.RoleGroupCollectionResponseable), nil
}
// Post create new navigation property to roleGroups for external
// returns a RoleGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataRolegroupsRoleGroupsRequestBuilder) Post(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.RoleGroupable, requestConfiguration *IndustrydataRolegroupsRoleGroupsRequestBuilderPostRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.RoleGroupable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateRoleGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.RoleGroupable), nil
}
// ToGetRequestInformation get a list of the roleGroup objects and their properties.
// returns a *RequestInformation when successful
func (m *IndustrydataRolegroupsRoleGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataRolegroupsRoleGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleGroups for external
// returns a *RequestInformation when successful
func (m *IndustrydataRolegroupsRoleGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.RoleGroupable, requestConfiguration *IndustrydataRolegroupsRoleGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataRolegroupsRoleGroupsRequestBuilder when successful
func (m *IndustrydataRolegroupsRoleGroupsRequestBuilder) WithUrl(rawUrl string)(*IndustrydataRolegroupsRoleGroupsRequestBuilder) {
    return NewIndustrydataRolegroupsRoleGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
