package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder provides operations to manage the alertDefinition property of the microsoft.graph.unifiedRoleManagementAlertConfiguration entity.
type RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderGetQueryParameters the definition of the alert that contains its description, impact, and measures to mitigate or prevent it. Supports $expand.
type RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderGetQueryParameters
}
// NewRolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderInternal instantiates a new RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder) {
    m := &RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alertConfigurations/{unifiedRoleManagementAlertConfiguration%2Did}/alertDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder instantiates a new RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the definition of the alert that contains its description, impact, and measures to mitigate or prevent it. Supports $expand.
// returns a UnifiedRoleManagementAlertDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertDefinitionable), nil
}
// ToGetRequestInformation the definition of the alert that contains its description, impact, and measures to mitigate or prevent it. Supports $expand.
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder when successful
func (m *RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder) {
    return NewRolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
