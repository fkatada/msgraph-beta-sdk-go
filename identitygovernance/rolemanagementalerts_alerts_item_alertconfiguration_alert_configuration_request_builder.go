package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder provides operations to manage the alertConfiguration property of the microsoft.graph.unifiedRoleManagementAlert entity.
type RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderGetQueryParameters the configuration of the alert in PIM for Microsoft Entra roles. Alert configurations are pre-defined and cannot be created or deleted, but some configurations can be modified. Supports $filter for the isEnabled property and $expand.
type RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderGetQueryParameters
}
// NewRolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderInternal instantiates a new RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder) {
    m := &RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertConfiguration{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder instantiates a new RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the configuration of the alert in PIM for Microsoft Entra roles. Alert configurations are pre-defined and cannot be created or deleted, but some configurations can be modified. Supports $filter for the isEnabled property and $expand.
// returns a UnifiedRoleManagementAlertConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable), nil
}
// ToGetRequestInformation the configuration of the alert in PIM for Microsoft Entra roles. Alert configurations are pre-defined and cannot be created or deleted, but some configurations can be modified. Supports $filter for the isEnabled property and $expand.
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder when successful
func (m *RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder) {
    return NewRolemanagementalertsAlertsItemAlertconfigurationAlertConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
