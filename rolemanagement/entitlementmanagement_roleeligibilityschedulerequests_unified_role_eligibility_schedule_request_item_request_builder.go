package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters get roleEligibilityScheduleRequests from roleManagement
type EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) AppScope()(*EntitlementmanagementRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsItemCancelRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Cancel()(*EntitlementmanagementRoleeligibilityschedulerequestsItemCancelRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal instantiates a new EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    m := &EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder instantiates a new EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleEligibilityScheduleRequests for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) DirectoryScope()(*EntitlementmanagementRoleeligibilityschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get roleEligibilityScheduleRequests from roleManagement
// returns a UnifiedRoleEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable), nil
}
// Patch update the navigation property roleEligibilityScheduleRequests in roleManagement
// returns a UnifiedRoleEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsItemPrincipalRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Principal()(*EntitlementmanagementRoleeligibilityschedulerequestsItemPrincipalRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) RoleDefinition()(*EntitlementmanagementRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) TargetSchedule()(*EntitlementmanagementRoleeligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleEligibilityScheduleRequests for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get roleEligibilityScheduleRequests from roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleEligibilityScheduleRequests in roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
