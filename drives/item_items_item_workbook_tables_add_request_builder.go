package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesAddRequestBuilder provides operations to call the add method.
type ItemItemsItemWorkbookTablesAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesAddRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesAddRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesAddRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/add", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookTablesAddRequestBuilder instantiates a new ItemItemsItemWorkbookTablesAddRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a new table. The range source address determines the worksheet under which the table will be added. If the table can't be added (for example, because the address is invalid, or the table would overlap with another table), an error is thrown.
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablecollection-add?view=graph-rest-beta
func (m *ItemItemsItemWorkbookTablesAddRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookTablesAddPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesAddRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable), nil
}
// ToPostRequestInformation create a new table. The range source address determines the worksheet under which the table will be added. If the table can't be added (for example, because the address is invalid, or the table would overlap with another table), an error is thrown.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookTablesAddPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookTablesAddRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesAddRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesAddRequestBuilder) {
    return NewItemItemsItemWorkbookTablesAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
