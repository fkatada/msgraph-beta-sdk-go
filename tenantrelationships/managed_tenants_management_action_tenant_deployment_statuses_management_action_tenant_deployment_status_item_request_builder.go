package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderGetQueryParameters read the properties and relationships of a managementActionTenantDeploymentStatus object.
type ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderGetQueryParameters
}
// ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderInternal instantiates a new ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) {
    m := &ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementActionTenantDeploymentStatuses/{managementActionTenantDeploymentStatus%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder instantiates a new ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managementActionTenantDeploymentStatuses for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a managementActionTenantDeploymentStatus object.
// returns a ManagementActionTenantDeploymentStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-managementactiontenantdeploymentstatus-get?view=graph-rest-beta
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable), nil
}
// Patch update the navigation property managementActionTenantDeploymentStatuses in tenantRelationships
// returns a ManagementActionTenantDeploymentStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable, requestConfiguration *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable), nil
}
// ToDeleteRequestInformation delete navigation property managementActionTenantDeploymentStatuses for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a managementActionTenantDeploymentStatus object.
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managementActionTenantDeploymentStatuses in tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable, requestConfiguration *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder when successful
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) {
    return NewManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
