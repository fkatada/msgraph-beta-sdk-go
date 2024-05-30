package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters represents a collection of worksheets associated with the workbook. Read-only.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CellWithRowWithColumn provides operations to call the cell method.
// returns a *ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) CellWithRowWithColumn(column *int32, row *int32)(*ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, column, row)
}
// Charts provides operations to manage the charts property of the microsoft.graph.workbookWorksheet entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Charts()(*ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property worksheets for drives
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents a collection of worksheets associated with the workbook. Read-only.
// returns a WorkbookWorksheetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookWorksheetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookWorksheetable), nil
}
// Names provides operations to manage the names property of the microsoft.graph.workbookWorksheet entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemNamesRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Names()(*ItemItemsItemWorkbookWorksheetsItemNamesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemNamesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property worksheets in drives
// returns a WorkbookWorksheetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookWorksheetable, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookWorksheetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookWorksheetable), nil
}
// PivotTables provides operations to manage the pivotTables property of the microsoft.graph.workbookWorksheet entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemPivottablesPivotTablesRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) PivotTables()(*ItemItemsItemWorkbookWorksheetsItemPivottablesPivotTablesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemPivottablesPivotTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Protection provides operations to manage the protection property of the microsoft.graph.workbookWorksheet entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemProtectionRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Protection()(*ItemItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemProtectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *ItemItemsItemWorkbookWorksheetsItemRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) RangeEscaped()(*ItemItemsItemWorkbookWorksheetsItemRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeWithAddress provides operations to call the range method.
// returns a *ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) RangeWithAddress(address *string)(*ItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, address)
}
// Tables provides operations to manage the tables property of the microsoft.graph.workbookWorksheet entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Tables()(*ItemItemsItemWorkbookWorksheetsItemTablesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.workbookWorksheet entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemTasksRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Tasks()(*ItemItemsItemWorkbookWorksheetsItemTasksRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property worksheets for drives
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property worksheets in drives
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookWorksheetable, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UsedRange provides operations to call the usedRange method.
// returns a *ItemItemsItemWorkbookWorksheetsItemUsedrangeUsedRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) UsedRange()(*ItemItemsItemWorkbookWorksheetsItemUsedrangeUsedRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemUsedrangeUsedRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsedRangeWithValuesOnly provides operations to call the usedRange method.
// returns a *ItemItemsItemWorkbookWorksheetsItemUsedrangewithvaluesonlyUsedRangeWithValuesOnlyRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ItemItemsItemWorkbookWorksheetsItemUsedrangewithvaluesonlyUsedRangeWithValuesOnlyRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemUsedrangewithvaluesonlyUsedRangeWithValuesOnlyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, valuesOnly)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
