package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessPackageResourceRoleId provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) ByAccessPackageResourceRoleId(accessPackageResourceRoleId string)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceRoleId != "" {
        urlTplParams["accessPackageResourceRole%2Did"] = accessPackageResourceRoleId
    }
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) {
    m := &EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackageResourceScopes/{accessPackageResourceScope%2Did}/accessPackageResource/accessPackageResourceRoles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder instantiates a new EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) Count()(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageResourceRoleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleCollectionResponseable), nil
}
// Post create new navigation property to accessPackageResourceRoles for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageResourceRoleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable), nil
}
// ToGetRequestInformation read-only. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to accessPackageResourceRoles for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRolesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
