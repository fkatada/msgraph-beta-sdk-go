package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder provides operations to call the search method.
type FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderGetQueryParameters search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
type FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderGetQueryParameters struct {
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
// FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderGetQueryParameters
}
// NewFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, q *string)(*FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/search(q='{q}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if q != nil {
        m.BaseRequestBuilder.PathParameters["q"] = *q
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// Deprecated: This method is obsolete. Use GetAsSearchWithQGetResponse instead.
// returns a FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-search?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQResponseable), nil
}
// GetAsSearchWithQGetResponse search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// returns a FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-search?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) GetAsSearchWithQGetResponse(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQGetResponseable), nil
}
// ToGetRequestInformation search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
