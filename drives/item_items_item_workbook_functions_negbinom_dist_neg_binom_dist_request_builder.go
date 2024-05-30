package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder provides operations to call the negBinom_Dist method.
type ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/negBinom_Dist", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action negBinom_Dist
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action negBinom_Dist
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsNegbinom_distNegBinom_DistRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
