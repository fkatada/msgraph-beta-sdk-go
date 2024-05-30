package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceRole entity.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageResourceScope/accessPackageResource/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResource for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Refresh()(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
