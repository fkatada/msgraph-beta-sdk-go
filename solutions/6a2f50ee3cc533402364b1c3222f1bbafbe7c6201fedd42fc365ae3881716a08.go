package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder provides operations to manage the driveProtectionUnitsBulkAdditionJobs property of the microsoft.graph.backupRestoreRoot entity.
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters get driveProtectionUnitsBulkAdditionJobs from solutions
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters
}
// BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderInternal instantiates a new BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder and sets the default values.
func NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    m := &BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/driveProtectionUnitsBulkAdditionJobs/{driveProtectionUnitsBulkAdditionJob%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder instantiates a new BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder and sets the default values.
func NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property driveProtectionUnitsBulkAdditionJobs for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get driveProtectionUnitsBulkAdditionJobs from solutions
// returns a DriveProtectionUnitsBulkAdditionJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveProtectionUnitsBulkAdditionJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable), nil
}
// Patch update the navigation property driveProtectionUnitsBulkAdditionJobs in solutions
// returns a DriveProtectionUnitsBulkAdditionJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveProtectionUnitsBulkAdditionJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable), nil
}
// ToDeleteRequestInformation delete navigation property driveProtectionUnitsBulkAdditionJobs for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get driveProtectionUnitsBulkAdditionJobs from solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property driveProtectionUnitsBulkAdditionJobs in solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    return NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
