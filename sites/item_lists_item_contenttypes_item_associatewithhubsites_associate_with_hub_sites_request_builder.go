package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder provides operations to call the associateWithHubSites method.
type ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal instantiates a new ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    m := &ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/associateWithHubSites", pathParameters),
    }
    return m
}
// NewItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder instantiates a new ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action associateWithHubSites
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-associatewithhubsites?view=graph-rest-beta
func (m *ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) Post(ctx context.Context, body ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action associateWithHubSites
// returns a *RequestInformation when successful
func (m *ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder when successful
func (m *ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    return NewItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
