package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder provides operations to manage the media for the site entity.
type ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderInternal instantiates a new ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder and sets the default values.
func NewItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder) {
    m := &ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/activities/{itemActivityOLD%2Did}/driveItem/contentStream", pathParameters),
    }
    return m
}
// NewItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder instantiates a new ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder and sets the default values.
func NewItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the content stream, if the item represents a file.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put the content stream, if the item represents a file.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation the content stream, if the item represents a file.
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the content stream, if the item represents a file.
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder when successful
func (m *ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder) {
    return NewItemListsItemItemsItemActivitiesItemDriveitemContentstreamContentStreamRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
