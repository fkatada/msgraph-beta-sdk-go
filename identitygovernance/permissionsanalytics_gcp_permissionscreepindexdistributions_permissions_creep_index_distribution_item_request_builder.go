package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder provides operations to manage the permissionsCreepIndexDistributions property of the microsoft.graph.permissionsAnalytics entity.
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters
}
// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthorizationSystem provides operations to manage the authorizationSystem property of the microsoft.graph.permissionsCreepIndexDistribution entity.
// returns a *PermissionsanalyticsGcpPermissionscreepindexdistributionsItemAuthorizationsystemAuthorizationSystemRequestBuilder when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) AuthorizationSystem()(*PermissionsanalyticsGcpPermissionscreepindexdistributionsItemAuthorizationsystemAuthorizationSystemRequestBuilder) {
    return NewPermissionsanalyticsGcpPermissionscreepindexdistributionsItemAuthorizationsystemAuthorizationSystemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal instantiates a new PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder and sets the default values.
func NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    m := &PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions/{permissionsCreepIndexDistribution%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder instantiates a new PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder and sets the default values.
func NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property permissionsCreepIndexDistributions for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
// returns a PermissionsCreepIndexDistributionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable), nil
}
// Patch update the navigation property permissionsCreepIndexDistributions in identityGovernance
// returns a PermissionsCreepIndexDistributionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable), nil
}
// ToDeleteRequestInformation delete navigation property permissionsCreepIndexDistributions for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property permissionsCreepIndexDistributions in identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) WithUrl(rawUrl string)(*PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    return NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
