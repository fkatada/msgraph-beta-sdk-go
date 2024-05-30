package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder provides operations to call the restoreVersion method.
type ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderInternal instantiates a new ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) {
    m := &ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/versions/{listItemVersion%2Did}/restoreVersion", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder instantiates a new ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action restoreVersion
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action restoreVersion
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder when successful
func (m *ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
