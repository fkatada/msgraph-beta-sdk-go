package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceScope entity.
type EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceRoles provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageResourceScope/accessPackageResource{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResource for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get accessPackageResource from identityGovernance
// returns a AccessPackageResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
// Patch update the navigation property accessPackageResource in identityGovernance
// returns a AccessPackageResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
// Refresh provides operations to call the refresh method.
// returns a *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceRefreshRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) Refresh()(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceRefreshRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get accessPackageResource from identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property accessPackageResource in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
