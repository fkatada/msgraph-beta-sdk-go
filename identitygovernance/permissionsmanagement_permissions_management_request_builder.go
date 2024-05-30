package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsmanagementPermissionsManagementRequestBuilder provides operations to manage the permissionsManagement property of the microsoft.graph.identityGovernance entity.
type PermissionsmanagementPermissionsManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsmanagementPermissionsManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementPermissionsManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionsmanagementPermissionsManagementRequestBuilderGetQueryParameters get permissionsManagement from identityGovernance
type PermissionsmanagementPermissionsManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsmanagementPermissionsManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementPermissionsManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsmanagementPermissionsManagementRequestBuilderGetQueryParameters
}
// PermissionsmanagementPermissionsManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementPermissionsManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionsmanagementPermissionsManagementRequestBuilderInternal instantiates a new PermissionsmanagementPermissionsManagementRequestBuilder and sets the default values.
func NewPermissionsmanagementPermissionsManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsmanagementPermissionsManagementRequestBuilder) {
    m := &PermissionsmanagementPermissionsManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionsmanagementPermissionsManagementRequestBuilder instantiates a new PermissionsmanagementPermissionsManagementRequestBuilder and sets the default values.
func NewPermissionsmanagementPermissionsManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsmanagementPermissionsManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsmanagementPermissionsManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property permissionsManagement for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionsmanagementPermissionsManagementRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get permissionsManagement from identityGovernance
// returns a PermissionsManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsmanagementPermissionsManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsManagementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsManagementable), nil
}
// Patch update the navigation property permissionsManagement in identityGovernance
// returns a PermissionsManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsManagementable, requestConfiguration *PermissionsmanagementPermissionsManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsManagementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsManagementable), nil
}
// PermissionsRequestChanges provides operations to manage the permissionsRequestChanges property of the microsoft.graph.permissionsManagement entity.
// returns a *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder when successful
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) PermissionsRequestChanges()(*PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) {
    return NewPermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ScheduledPermissionsApprovals provides operations to manage the scheduledPermissionsApprovals property of the microsoft.graph.permissionsManagement entity.
// returns a *PermissionsmanagementScheduledpermissionsapprovalsScheduledPermissionsApprovalsRequestBuilder when successful
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) ScheduledPermissionsApprovals()(*PermissionsmanagementScheduledpermissionsapprovalsScheduledPermissionsApprovalsRequestBuilder) {
    return NewPermissionsmanagementScheduledpermissionsapprovalsScheduledPermissionsApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ScheduledPermissionsRequests provides operations to manage the scheduledPermissionsRequests property of the microsoft.graph.permissionsManagement entity.
// returns a *PermissionsmanagementScheduledpermissionsrequestsScheduledPermissionsRequestsRequestBuilder when successful
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) ScheduledPermissionsRequests()(*PermissionsmanagementScheduledpermissionsrequestsScheduledPermissionsRequestsRequestBuilder) {
    return NewPermissionsmanagementScheduledpermissionsrequestsScheduledPermissionsRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property permissionsManagement for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionsmanagementPermissionsManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get permissionsManagement from identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsmanagementPermissionsManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property permissionsManagement in identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsManagementable, requestConfiguration *PermissionsmanagementPermissionsManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissionsmanagementPermissionsManagementRequestBuilder when successful
func (m *PermissionsmanagementPermissionsManagementRequestBuilder) WithUrl(rawUrl string)(*PermissionsmanagementPermissionsManagementRequestBuilder) {
    return NewPermissionsmanagementPermissionsManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
