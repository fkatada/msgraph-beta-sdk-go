package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
type ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderGetQueryParameters the to [term] of the relation. The term to which the relationship is defined.
type ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderGetQueryParameters
}
// NewItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderInternal instantiates a new ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder) {
    m := &ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/groups/{group%2Did1}/sets/{set%2Did}/terms/{term%2Did}/children/{term%2Did1}/relations/{relation%2Did}/toTerm{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder instantiates a new ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the to [term] of the relation. The term to which the relationship is defined.
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// ToGetRequestInformation the to [term] of the relation. The term to which the relationship is defined.
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder when successful
func (m *ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder) {
    return NewItemSitesItemTermstoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemTotermToTermRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
