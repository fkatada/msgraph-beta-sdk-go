package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder provides operations to call the range method.
type ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, address *string)(*ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/range(address='{address}')", pathParameters),
    }
    if address != nil {
        m.BaseRequestBuilder.PathParameters["address"] = *address
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function range
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookRangeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookRangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookRangeable), nil
}
// ToGetRequestInformation invoke function range
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
