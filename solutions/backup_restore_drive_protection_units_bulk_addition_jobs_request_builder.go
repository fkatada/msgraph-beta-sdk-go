package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder provides operations to manage the driveProtectionUnitsBulkAdditionJobs property of the microsoft.graph.backupRestoreRoot entity.
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderGetQueryParameters get driveProtectionUnitsBulkAdditionJobs from solutions
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderGetQueryParameters struct {
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
// BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderGetQueryParameters
}
// BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDriveProtectionUnitsBulkAdditionJobId provides operations to manage the driveProtectionUnitsBulkAdditionJobs property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) ByDriveProtectionUnitsBulkAdditionJobId(driveProtectionUnitsBulkAdditionJobId string)(*BackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if driveProtectionUnitsBulkAdditionJobId != "" {
        urlTplParams["driveProtectionUnitsBulkAdditionJob%2Did"] = driveProtectionUnitsBulkAdditionJobId
    }
    return NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsDriveProtectionUnitsBulkAdditionJobItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderInternal instantiates a new BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder and sets the default values.
func NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) {
    m := &BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/driveProtectionUnitsBulkAdditionJobs{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder instantiates a new BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder and sets the default values.
func NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackupRestoreDriveProtectionUnitsBulkAdditionJobsCountRequestBuilder when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) Count()(*BackupRestoreDriveProtectionUnitsBulkAdditionJobsCountRequestBuilder) {
    return NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get driveProtectionUnitsBulkAdditionJobs from solutions
// returns a DriveProtectionUnitsBulkAdditionJobCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveProtectionUnitsBulkAdditionJobCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobCollectionResponseable), nil
}
// Post create new navigation property to driveProtectionUnitsBulkAdditionJobs for solutions
// returns a DriveProtectionUnitsBulkAdditionJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get driveProtectionUnitsBulkAdditionJobs from solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to driveProtectionUnitsBulkAdditionJobs for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitsBulkAdditionJobable, requestConfiguration *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder when successful
func (m *BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder) {
    return NewBackupRestoreDriveProtectionUnitsBulkAdditionJobsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
