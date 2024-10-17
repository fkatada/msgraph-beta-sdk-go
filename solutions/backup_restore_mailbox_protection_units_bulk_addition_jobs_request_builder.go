package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder provides operations to manage the mailboxProtectionUnitsBulkAdditionJobs property of the microsoft.graph.backupRestoreRoot entity.
type BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderGetQueryParameters get mailboxProtectionUnitsBulkAdditionJobs from solutions
type BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderGetQueryParameters struct {
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
// BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderGetQueryParameters
}
// BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMailboxProtectionUnitsBulkAdditionJobId provides operations to manage the mailboxProtectionUnitsBulkAdditionJobs property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsMailboxProtectionUnitsBulkAdditionJobItemRequestBuilder when successful
func (m *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) ByMailboxProtectionUnitsBulkAdditionJobId(mailboxProtectionUnitsBulkAdditionJobId string)(*BackupRestoreMailboxProtectionUnitsBulkAdditionJobsMailboxProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mailboxProtectionUnitsBulkAdditionJobId != "" {
        urlTplParams["mailboxProtectionUnitsBulkAdditionJob%2Did"] = mailboxProtectionUnitsBulkAdditionJobId
    }
    return NewBackupRestoreMailboxProtectionUnitsBulkAdditionJobsMailboxProtectionUnitsBulkAdditionJobItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderInternal instantiates a new BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder and sets the default values.
func NewBackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) {
    m := &BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/mailboxProtectionUnitsBulkAdditionJobs{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder instantiates a new BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder and sets the default values.
func NewBackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsCountRequestBuilder when successful
func (m *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) Count()(*BackupRestoreMailboxProtectionUnitsBulkAdditionJobsCountRequestBuilder) {
    return NewBackupRestoreMailboxProtectionUnitsBulkAdditionJobsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get mailboxProtectionUnitsBulkAdditionJobs from solutions
// returns a MailboxProtectionUnitsBulkAdditionJobCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionUnitsBulkAdditionJobCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxProtectionUnitsBulkAdditionJobCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionUnitsBulkAdditionJobCollectionResponseable), nil
}
// Post create new navigation property to mailboxProtectionUnitsBulkAdditionJobs for solutions
// returns a MailboxProtectionUnitsBulkAdditionJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionUnitsBulkAdditionJobable, requestConfiguration *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionUnitsBulkAdditionJobable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxProtectionUnitsBulkAdditionJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionUnitsBulkAdditionJobable), nil
}
// ToGetRequestInformation get mailboxProtectionUnitsBulkAdditionJobs from solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to mailboxProtectionUnitsBulkAdditionJobs for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionUnitsBulkAdditionJobable, requestConfiguration *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder when successful
func (m *BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) {
    return NewBackupRestoreMailboxProtectionUnitsBulkAdditionJobsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
