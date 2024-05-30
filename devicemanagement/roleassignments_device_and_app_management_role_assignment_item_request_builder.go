package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
type RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderGetQueryParameters the Role Assignments.
type RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderGetQueryParameters
}
// RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal instantiates a new RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder and sets the default values.
func NewRoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) {
    m := &RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder instantiates a new RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder and sets the default values.
func NewRoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the Role Assignments.
// returns a DeviceAndAppManagementRoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable), nil
}
// Patch update the navigation property roleAssignments in deviceManagement
// returns a DeviceAndAppManagementRoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable, requestConfiguration *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable), nil
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.roleAssignment entity.
// returns a *RoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) RoleDefinition()(*RoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleScopeTags provides operations to manage the roleScopeTags property of the microsoft.graph.deviceAndAppManagementRoleAssignment entity.
// returns a *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder when successful
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) RoleScopeTags()(*RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) {
    return NewRoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignments for deviceManagement
// returns a *RequestInformation when successful
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the Role Assignments.
// returns a *RequestInformation when successful
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignments in deviceManagement
// returns a *RequestInformation when successful
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable, requestConfiguration *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder when successful
func (m *RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*RoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) {
    return NewRoleassignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
