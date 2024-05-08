package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder provides operations to manage the tenants property of the microsoft.graph.multiTenantOrganization entity.
type MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetQueryParameters get a tenant and its properties in the multi-tenant organization.
type MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetQueryParameters
}
// MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderInternal instantiates a new MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder and sets the default values.
func NewMultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) {
    m := &MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/multiTenantOrganization/tenants/{multiTenantOrganizationMember%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder instantiates a new MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder and sets the default values.
func NewMultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a tenant from a multitenant organization. A tenant can be removed in the following scenarios: To allow for asynchronous processing, you must wait for up to 2 hours before removal of a tenant is completed.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganization-delete-tenants?view=graph-rest-beta
func (m *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get a tenant and its properties in the multi-tenant organization.
// returns a MultiTenantOrganizationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganizationmember-get?view=graph-rest-beta
func (m *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationMemberable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMultiTenantOrganizationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationMemberable), nil
}
// Patch update the navigation property tenants in tenantRelationships
// returns a MultiTenantOrganizationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationMemberable, requestConfiguration *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationMemberable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMultiTenantOrganizationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationMemberable), nil
}
// ToDeleteRequestInformation remove a tenant from a multitenant organization. A tenant can be removed in the following scenarios: To allow for asynchronous processing, you must wait for up to 2 hours before removal of a tenant is completed.
// returns a *RequestInformation when successful
func (m *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get a tenant and its properties in the multi-tenant organization.
// returns a *RequestInformation when successful
func (m *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tenants in tenantRelationships
// returns a *RequestInformation when successful
func (m *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationMemberable, requestConfiguration *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder when successful
func (m *MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) WithUrl(rawUrl string)(*MultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) {
    return NewMultiTenantOrganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
