package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderGetQueryParameters get a list of the managementTemplate objects and their properties.
type ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderGetQueryParameters
}
// ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagementTemplateId provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatesManagementTemplateItemRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) ByManagementTemplateId(managementTemplateId string)(*ManagedtenantsManagementtemplatesManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managementTemplateId != "" {
        urlTplParams["managementTemplate%2Did"] = managementTemplateId
    }
    return NewManagedtenantsManagementtemplatesManagementTemplateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderInternal instantiates a new ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) {
    m := &ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder instantiates a new ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsManagementtemplatesCountRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) Count()(*ManagedtenantsManagementtemplatesCountRequestBuilder) {
    return NewManagedtenantsManagementtemplatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the managementTemplate objects and their properties.
// returns a ManagementTemplateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-managedtenant-list-managementtemplates?view=graph-rest-beta
func (m *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionResponseable), nil
}
// Post create new navigation property to managementTemplates for tenantRelationships
// returns a ManagementTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateable, requestConfiguration *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateable), nil
}
// ToGetRequestInformation get a list of the managementTemplate objects and their properties.
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managementTemplates for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateable, requestConfiguration *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) {
    return NewManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
