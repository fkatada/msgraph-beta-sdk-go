package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceRole entity.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    m := &EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackageResources/{accessPackageResource%2Did}/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder instantiates a new EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResource for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) Refresh()(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
