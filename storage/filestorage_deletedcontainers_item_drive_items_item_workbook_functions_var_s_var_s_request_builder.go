package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder provides operations to call the var_S method.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/functions/var_S", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action var_S
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder) Post(ctx context.Context, body FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SPostRequestBodyable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action var_S
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SPostRequestBodyable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookFunctionsVar_sVar_SRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
