package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// NetworkAccessRequestBuilder provides operations to manage the networkAccessRoot singleton.
type NetworkAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NetworkAccessRequestBuilderGetQueryParameters get networkAccess
type NetworkAccessRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// NetworkAccessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NetworkAccessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *NetworkAccessRequestBuilderGetQueryParameters
}
// NetworkAccessRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NetworkAccessRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Alerts provides operations to manage the alerts property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *AlertsRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) Alerts()(*AlertsRequestBuilder) {
    return NewAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Connectivity provides operations to manage the connectivity property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *ConnectivityRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) Connectivity()(*ConnectivityRequestBuilder) {
    return NewConnectivityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewNetworkAccessRequestBuilderInternal instantiates a new NetworkAccessRequestBuilder and sets the default values.
func NewNetworkAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NetworkAccessRequestBuilder) {
    m := &NetworkAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewNetworkAccessRequestBuilder instantiates a new NetworkAccessRequestBuilder and sets the default values.
func NewNetworkAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NetworkAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNetworkAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// FilteringPolicies provides operations to manage the filteringPolicies property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *FilteringpoliciesFilteringPoliciesRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) FilteringPolicies()(*FilteringpoliciesFilteringPoliciesRequestBuilder) {
    return NewFilteringpoliciesFilteringPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilteringProfiles provides operations to manage the filteringProfiles property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *FilteringprofilesFilteringProfilesRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) FilteringProfiles()(*FilteringprofilesFilteringProfilesRequestBuilder) {
    return NewFilteringprofilesFilteringProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ForwardingPolicies provides operations to manage the forwardingPolicies property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *ForwardingpoliciesForwardingPoliciesRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) ForwardingPolicies()(*ForwardingpoliciesForwardingPoliciesRequestBuilder) {
    return NewForwardingpoliciesForwardingPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ForwardingProfiles provides operations to manage the forwardingProfiles property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *ForwardingprofilesForwardingProfilesRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) ForwardingProfiles()(*ForwardingprofilesForwardingProfilesRequestBuilder) {
    return NewForwardingprofilesForwardingProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get networkAccess
// returns a NetworkAccessRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *NetworkAccessRequestBuilder) Get(ctx context.Context, requestConfiguration *NetworkAccessRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.NetworkAccessRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateNetworkAccessRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.NetworkAccessRootable), nil
}
// Logs provides operations to manage the logs property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *LogsRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) Logs()(*LogsRequestBuilder) {
    return NewLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphNetworkaccessOnboard provides operations to call the onboard method.
// returns a *MicrosoftgraphnetworkaccessonboardMicrosoftGraphNetworkaccessOnboardRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) MicrosoftGraphNetworkaccessOnboard()(*MicrosoftgraphnetworkaccessonboardMicrosoftGraphNetworkaccessOnboardRequestBuilder) {
    return NewMicrosoftgraphnetworkaccessonboardMicrosoftGraphNetworkaccessOnboardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update networkAccess
// returns a NetworkAccessRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *NetworkAccessRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.NetworkAccessRootable, requestConfiguration *NetworkAccessRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.NetworkAccessRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateNetworkAccessRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.NetworkAccessRootable), nil
}
// Reports provides operations to manage the reports property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *ReportsRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) Reports()(*ReportsRequestBuilder) {
    return NewReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *SettingsRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) Settings()(*SettingsRequestBuilder) {
    return NewSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantStatus provides operations to manage the tenantStatus property of the microsoft.graph.networkaccess.networkAccessRoot entity.
// returns a *TenantstatusTenantStatusRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) TenantStatus()(*TenantstatusTenantStatusRequestBuilder) {
    return NewTenantstatusTenantStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get networkAccess
// returns a *RequestInformation when successful
func (m *NetworkAccessRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NetworkAccessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update networkAccess
// returns a *RequestInformation when successful
func (m *NetworkAccessRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.NetworkAccessRootable, requestConfiguration *NetworkAccessRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *NetworkAccessRequestBuilder when successful
func (m *NetworkAccessRequestBuilder) WithUrl(rawUrl string)(*NetworkAccessRequestBuilder) {
    return NewNetworkAccessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
