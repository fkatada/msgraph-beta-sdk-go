package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreProtectionUnitsRequestBuilder provides operations to manage the protectionUnits property of the microsoft.graph.backupRestoreRoot entity.
type BackupRestoreProtectionUnitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreProtectionUnitsRequestBuilderGetQueryParameters read the properties and relationships of a protectionUnitBase object.
type BackupRestoreProtectionUnitsRequestBuilderGetQueryParameters struct {
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
// BackupRestoreProtectionUnitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreProtectionUnitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreProtectionUnitsRequestBuilderGetQueryParameters
}
// ByProtectionUnitBaseId provides operations to manage the protectionUnits property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackupRestoreProtectionUnitsProtectionUnitBaseItemRequestBuilder when successful
func (m *BackupRestoreProtectionUnitsRequestBuilder) ByProtectionUnitBaseId(protectionUnitBaseId string)(*BackupRestoreProtectionUnitsProtectionUnitBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if protectionUnitBaseId != "" {
        urlTplParams["protectionUnitBase%2Did"] = protectionUnitBaseId
    }
    return NewBackupRestoreProtectionUnitsProtectionUnitBaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackupRestoreProtectionUnitsRequestBuilderInternal instantiates a new BackupRestoreProtectionUnitsRequestBuilder and sets the default values.
func NewBackupRestoreProtectionUnitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreProtectionUnitsRequestBuilder) {
    m := &BackupRestoreProtectionUnitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/protectionUnits{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackupRestoreProtectionUnitsRequestBuilder instantiates a new BackupRestoreProtectionUnitsRequestBuilder and sets the default values.
func NewBackupRestoreProtectionUnitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreProtectionUnitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreProtectionUnitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackupRestoreProtectionUnitsCountRequestBuilder when successful
func (m *BackupRestoreProtectionUnitsRequestBuilder) Count()(*BackupRestoreProtectionUnitsCountRequestBuilder) {
    return NewBackupRestoreProtectionUnitsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a protectionUnitBase object.
// returns a ProtectionUnitBaseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreProtectionUnitsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreProtectionUnitsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionUnitBaseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProtectionUnitBaseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionUnitBaseCollectionResponseable), nil
}
// GraphDriveProtectionUnit casts the previous resource to driveProtectionUnit.
// returns a *BackupRestoreProtectionUnitsGraphDriveProtectionUnitRequestBuilder when successful
func (m *BackupRestoreProtectionUnitsRequestBuilder) GraphDriveProtectionUnit()(*BackupRestoreProtectionUnitsGraphDriveProtectionUnitRequestBuilder) {
    return NewBackupRestoreProtectionUnitsGraphDriveProtectionUnitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMailboxProtectionUnit casts the previous resource to mailboxProtectionUnit.
// returns a *BackupRestoreProtectionUnitsGraphMailboxProtectionUnitRequestBuilder when successful
func (m *BackupRestoreProtectionUnitsRequestBuilder) GraphMailboxProtectionUnit()(*BackupRestoreProtectionUnitsGraphMailboxProtectionUnitRequestBuilder) {
    return NewBackupRestoreProtectionUnitsGraphMailboxProtectionUnitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphSiteProtectionUnit casts the previous resource to siteProtectionUnit.
// returns a *BackupRestoreProtectionUnitsGraphSiteProtectionUnitRequestBuilder when successful
func (m *BackupRestoreProtectionUnitsRequestBuilder) GraphSiteProtectionUnit()(*BackupRestoreProtectionUnitsGraphSiteProtectionUnitRequestBuilder) {
    return NewBackupRestoreProtectionUnitsGraphSiteProtectionUnitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read the properties and relationships of a protectionUnitBase object.
// returns a *RequestInformation when successful
func (m *BackupRestoreProtectionUnitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreProtectionUnitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BackupRestoreProtectionUnitsRequestBuilder when successful
func (m *BackupRestoreProtectionUnitsRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreProtectionUnitsRequestBuilder) {
    return NewBackupRestoreProtectionUnitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
