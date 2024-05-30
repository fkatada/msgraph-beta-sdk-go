package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementEntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.roleManagement entity.
type EntitlementmanagementEntitlementManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters the RbacApplication for Entitlement Management
type EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters
}
// EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementEntitlementManagementRequestBuilderInternal instantiates a new EntitlementmanagementEntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    m := &EntitlementmanagementEntitlementManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementEntitlementManagementRequestBuilder instantiates a new EntitlementmanagementEntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementmanagementEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property entitlementManagement for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the RbacApplication for Entitlement Management
// returns a RbacApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable), nil
}
// Patch update the navigation property entitlementManagement in roleManagement
// returns a RbacApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable), nil
}
// ResourceNamespaces provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementResourcenamespacesResourceNamespacesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ResourceNamespaces()(*EntitlementmanagementResourcenamespacesResourceNamespacesRequestBuilder) {
    return NewEntitlementmanagementResourcenamespacesResourceNamespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignmentApprovals provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentapprovalsRoleAssignmentApprovalsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignmentApprovals()(*EntitlementmanagementRoleassignmentapprovalsRoleAssignmentApprovalsRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentapprovalsRoleAssignmentApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentsRoleAssignmentsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignments()(*EntitlementmanagementRoleassignmentsRoleAssignmentsRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentsRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignmentScheduleInstances provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignmentScheduleInstances()(*EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignmentScheduleRequests provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentschedulerequestsRoleAssignmentScheduleRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignmentScheduleRequests()(*EntitlementmanagementRoleassignmentschedulerequestsRoleAssignmentScheduleRequestsRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulerequestsRoleAssignmentScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignmentSchedules provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentschedulesRoleAssignmentSchedulesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignmentSchedules()(*EntitlementmanagementRoleassignmentschedulesRoleAssignmentSchedulesRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesRoleAssignmentSchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleDefinitions()(*EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) {
    return NewEntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleEligibilityScheduleInstances provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleEligibilityScheduleInstances()(*EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleEligibilityScheduleRequests provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleEligibilityScheduleRequests()(*EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleEligibilitySchedules provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleEligibilitySchedules()(*EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId provides operations to call the roleScheduleInstances method.
// returns a *EntitlementmanagementRolescheduleinstancesdirectoryscopeiddirectoryscopeidappscopeidappscopeidprincipalidprincipalidroledefinitionidroledefinitionidRoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId()(*EntitlementmanagementRolescheduleinstancesdirectoryscopeiddirectoryscopeidappscopeidappscopeidprincipalidprincipalidroledefinitionidroledefinitionidRoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return NewEntitlementmanagementRolescheduleinstancesdirectoryscopeiddirectoryscopeidappscopeidappscopeidprincipalidprincipalidroledefinitionidroledefinitionidRoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId provides operations to call the roleSchedules method.
// returns a *EntitlementmanagementRoleschedulesdirectoryscopeiddirectoryscopeidappscopeidappscopeidprincipalidprincipalidroledefinitionidroledefinitionidRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId()(*EntitlementmanagementRoleschedulesdirectoryscopeiddirectoryscopeidappscopeidappscopeidprincipalidprincipalidroledefinitionidroledefinitionidRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return NewEntitlementmanagementRoleschedulesdirectoryscopeiddirectoryscopeidappscopeidappscopeidprincipalidprincipalidroledefinitionidroledefinitionidRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property entitlementManagement for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the RbacApplication for Entitlement Management
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property entitlementManagement in roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TransitiveRoleAssignments provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) TransitiveRoleAssignments()(*EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) {
    return NewEntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementEntitlementManagementRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    return NewEntitlementmanagementEntitlementManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
