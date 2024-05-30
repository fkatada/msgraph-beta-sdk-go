package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder provides operations to manage the steps property of the microsoft.graph.approval entity.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderGetQueryParameters used to represent the decision associated with a single step in the approval process configured in approvalStage.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderGetQueryParameters
}
// PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderInternal instantiates a new PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder and sets the default values.
func NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) {
    m := &PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/{approval%2Did}/steps/{approvalStep%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder instantiates a new PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder and sets the default values.
func NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property steps for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get used to represent the decision associated with a single step in the approval process configured in approvalStage.
// returns a ApprovalStepable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable), nil
}
// Patch update the navigation property steps in identityGovernance
// returns a ApprovalStepable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable), nil
}
// ToDeleteRequestInformation delete navigation property steps for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation used to represent the decision associated with a single step in the approval process configured in approvalStage.
// returns a *RequestInformation when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property steps in identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) WithUrl(rawUrl string)(*PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) {
    return NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
