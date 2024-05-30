package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
type ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderGetQueryParameters get alertLogs from tenantRelationships
type ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderGetQueryParameters
}
// ByManagedTenantAlertLogId provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
// returns a *ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder when successful
func (m *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder) ByManagedTenantAlertLogId(managedTenantAlertLogId string)(*ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedTenantAlertLogId != "" {
        urlTplParams["managedTenantAlertLog%2Did"] = managedTenantAlertLogId
    }
    return NewManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderInternal instantiates a new ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder and sets the default values.
func NewManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder) {
    m := &ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}/alertLogs{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder instantiates a new ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder and sets the default values.
func NewManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsManagedtenantalertsItemAlertlogsCountRequestBuilder when successful
func (m *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder) Count()(*ManagedtenantsManagedtenantalertsItemAlertlogsCountRequestBuilder) {
    return NewManagedtenantsManagedtenantalertsItemAlertlogsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get alertLogs from tenantRelationships
// returns a ManagedTenantAlertLogCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertLogCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertLogCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertLogCollectionResponseable), nil
}
// ToGetRequestInformation get alertLogs from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder when successful
func (m *ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder) {
    return NewManagedtenantsManagedtenantalertsItemAlertlogsAlertLogsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
