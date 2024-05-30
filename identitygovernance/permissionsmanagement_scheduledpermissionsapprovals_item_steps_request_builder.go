package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder provides operations to manage the steps property of the microsoft.graph.approval entity.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderGetQueryParameters used to represent the decision associated with a single step in the approval process configured in approvalStage.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderGetQueryParameters struct {
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
// PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderGetQueryParameters
}
// PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByApprovalStepId provides operations to manage the steps property of the microsoft.graph.approval entity.
// returns a *PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) ByApprovalStepId(approvalStepId string)(*PermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if approvalStepId != "" {
        urlTplParams["approvalStep%2Did"] = approvalStepId
    }
    return NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsApprovalStepItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderInternal instantiates a new PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder and sets the default values.
func NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) {
    m := &PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/{approval%2Did}/steps{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder instantiates a new PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder and sets the default values.
func NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PermissionsmanagementScheduledpermissionsapprovalsItemStepsCountRequestBuilder when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) Count()(*PermissionsmanagementScheduledpermissionsapprovalsItemStepsCountRequestBuilder) {
    return NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get used to represent the decision associated with a single step in the approval process configured in approvalStage.
// returns a ApprovalStepCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepCollectionResponseable), nil
}
// Post create new navigation property to steps for identityGovernance
// returns a ApprovalStepable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation used to represent the decision associated with a single step in the approval process configured in approvalStage.
// returns a *RequestInformation when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to steps for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder when successful
func (m *PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) WithUrl(rawUrl string)(*PermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder) {
    return NewPermissionsmanagementScheduledpermissionsapprovalsItemStepsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
