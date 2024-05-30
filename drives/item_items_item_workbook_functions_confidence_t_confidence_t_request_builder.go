package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder provides operations to call the confidence_T method.
type ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/confidence_T", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action confidence_T
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action confidence_T
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsConfidence_tConfidence_TRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
