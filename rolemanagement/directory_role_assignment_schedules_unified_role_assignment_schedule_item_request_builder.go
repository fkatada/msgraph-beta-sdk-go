package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetQueryParameters retrieve the schedule for an active role assignment operation.
type DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetQueryParameters
}
// DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentSchedule entity.
// returns a *DirectoryRoleAssignmentSchedulesItemActivatedUsingRequestBuilder when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) ActivatedUsing()(*DirectoryRoleAssignmentSchedulesItemActivatedUsingRequestBuilder) {
    return NewDirectoryRoleAssignmentSchedulesItemActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *DirectoryRoleAssignmentSchedulesItemAppScopeRequestBuilder when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) AppScope()(*DirectoryRoleAssignmentSchedulesItemAppScopeRequestBuilder) {
    return NewDirectoryRoleAssignmentSchedulesItemAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal instantiates a new DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    m := &DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentSchedules/{unifiedRoleAssignmentSchedule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder instantiates a new DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentSchedules for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *DirectoryRoleAssignmentSchedulesItemDirectoryScopeRequestBuilder when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) DirectoryScope()(*DirectoryRoleAssignmentSchedulesItemDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleAssignmentSchedulesItemDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the schedule for an active role assignment operation.
// returns a UnifiedRoleAssignmentScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentschedule-get?view=graph-rest-beta
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable), nil
}
// Patch update the navigation property roleAssignmentSchedules in roleManagement
// returns a UnifiedRoleAssignmentScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable, requestConfiguration *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *DirectoryRoleAssignmentSchedulesItemPrincipalRequestBuilder when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) Principal()(*DirectoryRoleAssignmentSchedulesItemPrincipalRequestBuilder) {
    return NewDirectoryRoleAssignmentSchedulesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *DirectoryRoleAssignmentSchedulesItemRoleDefinitionRequestBuilder when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) RoleDefinition()(*DirectoryRoleAssignmentSchedulesItemRoleDefinitionRequestBuilder) {
    return NewDirectoryRoleAssignmentSchedulesItemRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentSchedules for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the schedule for an active role assignment operation.
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignmentSchedules in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleable, requestConfiguration *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder when successful
func (m *DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    return NewDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
