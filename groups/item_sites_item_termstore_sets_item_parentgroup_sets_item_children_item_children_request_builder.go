package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.term entity.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderGetQueryParameters children of current term.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderGetQueryParameters
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTermId1 provides operations to manage the children property of the microsoft.graph.termStore.term entity.
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenTermItemRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) ByTermId1(termId1 string)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if termId1 != "" {
        urlTplParams["term%2Did1"] = termId1
    }
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderInternal instantiates a new ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) {
    m := &ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}/children{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder instantiates a new ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenCountRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) Count()(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenCountRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get children of current term.
// returns a TermCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.TermCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateTermCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.TermCollectionResponseable), nil
}
// Post create new navigation property to children for groups
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) Post(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderPostRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable), nil
}
// ToGetRequestInformation children of current term.
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to children for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) ToPostRequestInformation(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
