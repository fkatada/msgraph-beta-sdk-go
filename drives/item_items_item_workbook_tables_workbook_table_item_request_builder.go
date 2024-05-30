package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder provides operations to manage the tables property of the microsoft.graph.workbook entity.
type ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetQueryParameters represents a collection of tables associated with the workbook. Read-only.
type ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearFilters provides operations to call the clearFilters method.
// returns a *ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ClearFilters()(*ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookTablesItemColumnsRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Columns()(*ItemItemsItemWorkbookTablesItemColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder instantiates a new ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ConvertToRange provides operations to call the convertToRange method.
// returns a *ItemItemsItemWorkbookTablesItemConverttorangeConvertToRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ConvertToRange()(*ItemItemsItemWorkbookTablesItemConverttorangeConvertToRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemConverttorangeConvertToRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *ItemItemsItemWorkbookTablesItemDatabodyrangeDataBodyRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) DataBodyRange()(*ItemItemsItemWorkbookTablesItemDatabodyrangeDataBodyRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemDatabodyrangeDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property tables for drives
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get represents a collection of tables associated with the workbook. Read-only.
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// HeaderRowRange provides operations to call the headerRowRange method.
// returns a *ItemItemsItemWorkbookTablesItemHeaderrowrangeHeaderRowRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) HeaderRowRange()(*ItemItemsItemWorkbookTablesItemHeaderrowrangeHeaderRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemHeaderrowrangeHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property tables in drives
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, requestConfiguration *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// RangeEscaped provides operations to call the range method.
// returns a *ItemItemsItemWorkbookTablesItemRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) RangeEscaped()(*ItemItemsItemWorkbookTablesItemRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReapplyFilters provides operations to call the reapplyFilters method.
// returns a *ItemItemsItemWorkbookTablesItemReapplyfiltersReapplyFiltersRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ReapplyFilters()(*ItemItemsItemWorkbookTablesItemReapplyfiltersReapplyFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemReapplyfiltersReapplyFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookTablesItemRowsRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Rows()(*ItemItemsItemWorkbookTablesItemRowsRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookTablesItemSortRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Sort()(*ItemItemsItemWorkbookTablesItemSortRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemSortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property tables for drives
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents a collection of tables associated with the workbook. Read-only.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property tables in drives
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, requestConfiguration *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *ItemItemsItemWorkbookTablesItemTotalrowrangeTotalRowRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) TotalRowRange()(*ItemItemsItemWorkbookTablesItemTotalrowrangeTotalRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemTotalrowrangeTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) {
    return NewItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookTablesItemWorksheetRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Worksheet()(*ItemItemsItemWorkbookTablesItemWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
