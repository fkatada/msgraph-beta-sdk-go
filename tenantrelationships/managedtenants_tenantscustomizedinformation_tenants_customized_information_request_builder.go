package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderGetQueryParameters get a list of the tenantCustomizedInformation objects and their properties.
type ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderGetQueryParameters
}
// ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTenantCustomizedInformationId provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsTenantscustomizedinformationTenantCustomizedInformationItemRequestBuilder when successful
func (m *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) ByTenantCustomizedInformationId(tenantCustomizedInformationId string)(*ManagedtenantsTenantscustomizedinformationTenantCustomizedInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tenantCustomizedInformationId != "" {
        urlTplParams["tenantCustomizedInformation%2Did"] = tenantCustomizedInformationId
    }
    return NewManagedtenantsTenantscustomizedinformationTenantCustomizedInformationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderInternal instantiates a new ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder and sets the default values.
func NewManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) {
    m := &ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenantsCustomizedInformation{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder instantiates a new ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder and sets the default values.
func NewManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsTenantscustomizedinformationCountRequestBuilder when successful
func (m *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) Count()(*ManagedtenantsTenantscustomizedinformationCountRequestBuilder) {
    return NewManagedtenantsTenantscustomizedinformationCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the tenantCustomizedInformation objects and their properties.
// returns a TenantCustomizedInformationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-managedtenant-list-tenantscustomizedinformation?view=graph-rest-beta
func (m *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantCustomizedInformationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationCollectionResponseable), nil
}
// Post create new navigation property to tenantsCustomizedInformation for tenantRelationships
// returns a TenantCustomizedInformationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, requestConfiguration *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantCustomizedInformationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable), nil
}
// ToGetRequestInformation get a list of the tenantCustomizedInformation objects and their properties.
// returns a *RequestInformation when successful
func (m *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to tenantsCustomizedInformation for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, requestConfiguration *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder when successful
func (m *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) {
    return NewManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
