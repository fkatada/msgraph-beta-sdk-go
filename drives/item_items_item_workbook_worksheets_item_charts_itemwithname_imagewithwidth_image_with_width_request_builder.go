package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder provides operations to call the image method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/item(name='{name}')/image(width={width})", pathParameters),
    }
    if width != nil {
        m.BaseRequestBuilder.PathParameters["width"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*width), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function image
// Deprecated: This method is obsolete. Use GetAsImageWithWidthGetResponse instead.
// returns a ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponseable), nil
}
// GetAsImageWithWidthGetResponse invoke function image
// returns a ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder) GetAsImageWithWidthGetResponse(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponseable), nil
}
// ToGetRequestInformation invoke function image
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
