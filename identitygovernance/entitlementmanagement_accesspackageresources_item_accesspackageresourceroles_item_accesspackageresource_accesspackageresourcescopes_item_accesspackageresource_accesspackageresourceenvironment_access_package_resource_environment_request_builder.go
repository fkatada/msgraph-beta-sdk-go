package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
type EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters contains the environment information for the resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports $expand.
type EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) {
    m := &EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResources/{accessPackageResource%2Did}/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource/accessPackageResourceScopes/{accessPackageResourceScope%2Did}/accessPackageResource/accessPackageResourceEnvironment{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder instantiates a new EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the environment information for the resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports $expand.
// returns a AccessPackageResourceEnvironmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable), nil
}
// ToGetRequestInformation contains the environment information for the resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourcesItemAccesspackageresourcerolesItemAccesspackageresourceAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourceenvironmentAccessPackageResourceEnvironmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
