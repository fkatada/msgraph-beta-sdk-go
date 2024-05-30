package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
type EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessPackageResourceScopeId provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopeItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) ByAccessPackageResourceScopeId(accessPackageResourceScopeId string)(*EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceScopeId != "" {
        urlTplParams["accessPackageResourceScope%2Did"] = accessPackageResourceScopeId
    }
    return NewEntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) {
    m := &EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceScope/accessPackageResource/accessPackageResourceScopes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder instantiates a new EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) Count()(*EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. Supports $expand.
// returns a AccessPackageResourceScopeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceScopeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeCollectionResponseable), nil
}
// Post create new navigation property to accessPackageResourceScopes for identityGovernance
// returns a AccessPackageResourceScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, requestConfiguration *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable), nil
}
// ToGetRequestInformation read-only. Nullable. Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to accessPackageResourceScopes for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, requestConfiguration *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcescopesAccessPackageResourceScopesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
