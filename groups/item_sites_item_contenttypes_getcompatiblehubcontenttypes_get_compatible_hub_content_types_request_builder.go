package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder provides operations to call the getCompatibleHubContentTypes method.
type ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters get compatible content types in the content type hub that can be added to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see addCopyFromContentTypeHub and the blog post Syntex Product Updates – August 2021.
type ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters
}
// NewItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderInternal instantiates a new ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) {
    m := &ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/getCompatibleHubContentTypes(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder instantiates a new ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get compatible content types in the content type hub that can be added to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see addCopyFromContentTypeHub and the blog post Syntex Product Updates – August 2021.
// Deprecated: This method is obsolete. Use GetAsGetCompatibleHubContentTypesGetResponse instead.
// returns a ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-getcompatiblehubcontenttypes?view=graph-rest-beta
func (m *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseable), nil
}
// GetAsGetCompatibleHubContentTypesGetResponse get compatible content types in the content type hub that can be added to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see addCopyFromContentTypeHub and the blog post Syntex Product Updates – August 2021.
// returns a ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-getcompatiblehubcontenttypes?view=graph-rest-beta
func (m *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) GetAsGetCompatibleHubContentTypesGetResponse(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable), nil
}
// ToGetRequestInformation get compatible content types in the content type hub that can be added to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see addCopyFromContentTypeHub and the blog post Syntex Product Updates – August 2021.
// returns a *RequestInformation when successful
func (m *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder when successful
func (m *ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) {
    return NewItemSitesItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
