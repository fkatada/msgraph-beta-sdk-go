package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
type EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
type EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters
}
// NewEnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderInternal instantiates a new EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder and sets the default values.
func NewEnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) {
    m := &EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}/roleDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder instantiates a new EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder and sets the default values.
func NewEnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
// ToGetRequestInformation detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
// returns a *RequestInformation when successful
func (m *EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) WithUrl(rawUrl string)(*EnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewEnterpriseappsItemRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
