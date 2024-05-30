package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
type EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderGetQueryParameters operations that an authorized principal is allowed to perform.
type EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderGetQueryParameters
}
// EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationContext provides operations to manage the authenticationContext property of the microsoft.graph.unifiedRbacResourceAction entity.
// returns a *EnterpriseappsItemResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) AuthenticationContext()(*EnterpriseappsItemResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder) {
    return NewEnterpriseappsItemResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderInternal instantiates a new EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder and sets the default values.
func NewEnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) {
    m := &EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}/resourceActions/{unifiedRbacResourceAction%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder instantiates a new EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder and sets the default values.
func NewEnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resourceActions for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get operations that an authorized principal is allowed to perform.
// returns a UnifiedRbacResourceActionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceActionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable), nil
}
// Patch update the navigation property resourceActions in roleManagement
// returns a UnifiedRbacResourceActionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceActionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable), nil
}
// ResourceScope provides operations to manage the resourceScope property of the microsoft.graph.unifiedRbacResourceAction entity.
// returns a *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) ResourceScope()(*EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) {
    return NewEnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property resourceActions for roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation operations that an authorized principal is allowed to perform.
// returns a *RequestInformation when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resourceActions in roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) WithUrl(rawUrl string)(*EnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder) {
    return NewEnterpriseappsItemResourcenamespacesItemResourceactionsUnifiedRbacResourceActionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
