package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder provides operations to manage the changes property of the microsoft.graph.workbookDocumentTask entity.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderGetQueryParameters a collection of task change histories.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderGetQueryParameters
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWorkbookDocumentTaskChangeId provides operations to manage the changes property of the microsoft.graph.workbookDocumentTask entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesWorkbookDocumentTaskChangeItemRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) ByWorkbookDocumentTaskChangeId(workbookDocumentTaskChangeId string)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesWorkbookDocumentTaskChangeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workbookDocumentTaskChangeId != "" {
        urlTplParams["workbookDocumentTaskChange%2Did"] = workbookDocumentTaskChangeId
    }
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesWorkbookDocumentTaskChangeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tasks/{workbookDocumentTask%2Did}/changes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to call the count method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) Count()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of task change histories.
// returns a WorkbookDocumentTaskChangeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookDocumentTaskChangeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeCollectionResponseable), nil
}
// ItemAtWithIndex provides operations to call the itemAt method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) ItemAtWithIndex(index *int32)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesItematwithindexItemAtWithIndexRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesItematwithindexItemAtWithIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, index)
}
// Post create new navigation property to changes for storage
// returns a WorkbookDocumentTaskChangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookDocumentTaskChangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable), nil
}
// ToGetRequestInformation a collection of task change histories.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to changes for storage
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
