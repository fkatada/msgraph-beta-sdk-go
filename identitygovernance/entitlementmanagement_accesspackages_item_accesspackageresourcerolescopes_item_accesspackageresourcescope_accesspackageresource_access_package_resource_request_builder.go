package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceScope entity.
type EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceRoles provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceScope/accessPackageResource{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResource for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceRefreshRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) Refresh()(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceRefreshRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
