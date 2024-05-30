package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceRole entity.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    m := &EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackageResourceScopes/{accessPackageResourceScope%2Did}/accessPackageResource/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder instantiates a new EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResource for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Refresh()(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get accessPackageResource from identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
