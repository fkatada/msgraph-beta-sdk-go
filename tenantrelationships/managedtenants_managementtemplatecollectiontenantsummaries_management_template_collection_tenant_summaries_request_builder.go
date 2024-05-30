package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderGetQueryParameters get managementTemplateCollectionTenantSummaries from tenantRelationships
type ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderGetQueryParameters
}
// ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagementTemplateCollectionTenantSummaryId provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummaryItemRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) ByManagementTemplateCollectionTenantSummaryId(managementTemplateCollectionTenantSummaryId string)(*ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managementTemplateCollectionTenantSummaryId != "" {
        urlTplParams["managementTemplateCollectionTenantSummary%2Did"] = managementTemplateCollectionTenantSummaryId
    }
    return NewManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderInternal instantiates a new ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) {
    m := &ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateCollectionTenantSummaries{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder instantiates a new ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsManagementtemplatecollectiontenantsummariesCountRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) Count()(*ManagedtenantsManagementtemplatecollectiontenantsummariesCountRequestBuilder) {
    return NewManagedtenantsManagementtemplatecollectiontenantsummariesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get managementTemplateCollectionTenantSummaries from tenantRelationships
// returns a ManagementTemplateCollectionTenantSummaryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionTenantSummaryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateCollectionTenantSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionTenantSummaryCollectionResponseable), nil
}
// Post create new navigation property to managementTemplateCollectionTenantSummaries for tenantRelationships
// returns a ManagementTemplateCollectionTenantSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionTenantSummaryable, requestConfiguration *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionTenantSummaryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateCollectionTenantSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionTenantSummaryable), nil
}
// ToGetRequestInformation get managementTemplateCollectionTenantSummaries from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managementTemplateCollectionTenantSummaries for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionTenantSummaryable, requestConfiguration *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) {
    return NewManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
