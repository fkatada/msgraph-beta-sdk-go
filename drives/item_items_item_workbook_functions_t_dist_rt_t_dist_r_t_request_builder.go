package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder provides operations to call the t_Dist_RT method.
type ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/t_Dist_RT", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action t_Dist_RT
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action t_Dist_RT
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsT_dist_rtT_Dist_RTRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
