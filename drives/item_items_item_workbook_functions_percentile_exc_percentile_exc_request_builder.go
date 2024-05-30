package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder provides operations to call the percentile_Exc method.
type ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/percentile_Exc", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action percentile_Exc
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action percentile_Exc
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsPercentile_excPercentile_ExcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
