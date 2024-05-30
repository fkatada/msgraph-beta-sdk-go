package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder provides operations to call the setSolidColor method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/format/fill/setSolidColor", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sets the fill formatting of a chart element to a uniform color.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartfill-setsolidcolor?view=graph-rest-beta
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation sets the fill formatting of a chart element to a uniform color.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemFormatFillSetsolidcolorSetSolidColorRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
