package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder provides operations to manage the format property of the microsoft.graph.workbookChartAxis entity.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderGetQueryParameters represents the formatting of a chart object, which includes line and font formatting. Read-only.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderGetQueryParameters
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/axes/seriesAxis/format{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property format for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderDeleteRequestConfiguration)(error) {
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
// Font provides operations to manage the font property of the microsoft.graph.workbookChartAxisFormat entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatFontRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) Font()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatFontRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatFontRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the formatting of a chart object, which includes line and font formatting. Read-only.
// returns a WorkbookChartAxisFormatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartAxisFormatable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookChartAxisFormatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartAxisFormatable), nil
}
// Line provides operations to manage the line property of the microsoft.graph.workbookChartAxisFormat entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) Line()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property format in storage
// returns a WorkbookChartAxisFormatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartAxisFormatable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartAxisFormatable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookChartAxisFormatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartAxisFormatable), nil
}
// ToDeleteRequestInformation delete navigation property format for storage
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the formatting of a chart object, which includes line and font formatting. Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property format in storage
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartAxisFormatable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
