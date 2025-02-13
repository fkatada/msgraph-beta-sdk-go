package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder provides operations to manage the plans property of the microsoft.graph.teamsChannelPlanner entity.
type DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters a collection of plannerPlan objects owned by the Teams channel. Currently, only shared channels are supported. Read-only. Nullable.
type DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters
}
// DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Archive provides operations to call the archive method.
// returns a *DeletedTeamsItemChannelsItemPlannerPlansItemArchiveRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Archive()(*DeletedTeamsItemChannelsItemPlannerPlansItemArchiveRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Buckets provides operations to manage the buckets property of the microsoft.graph.plannerPlan entity.
// returns a *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Buckets()(*DeletedTeamsItemChannelsItemPlannerPlansItemBucketsRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemBucketsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderInternal instantiates a new DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/planner/plans/{plannerPlan%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder instantiates a new DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property plans for teamwork
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Details provides operations to manage the details property of the microsoft.graph.plannerPlan entity.
// returns a *DeletedTeamsItemChannelsItemPlannerPlansItemDetailsRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Details()(*DeletedTeamsItemChannelsItemPlannerPlansItemDetailsRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of plannerPlan objects owned by the Teams channel. Currently, only shared channels are supported. Read-only. Nullable.
// returns a PlannerPlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable), nil
}
// MoveToContainer provides operations to call the moveToContainer method.
// returns a *DeletedTeamsItemChannelsItemPlannerPlansItemMoveToContainerRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) MoveToContainer()(*DeletedTeamsItemChannelsItemPlannerPlansItemMoveToContainerRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemMoveToContainerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property plans in teamwork
// returns a PlannerPlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable), nil
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
// returns a *DeletedTeamsItemChannelsItemPlannerPlansItemTasksRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Tasks()(*DeletedTeamsItemChannelsItemPlannerPlansItemTasksRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property plans for teamwork
// returns a *RequestInformation when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of plannerPlan objects owned by the Teams channel. Currently, only shared channels are supported. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property plans in teamwork
// returns a *RequestInformation when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unarchive provides operations to call the unarchive method.
// returns a *DeletedTeamsItemChannelsItemPlannerPlansItemUnarchiveRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Unarchive()(*DeletedTeamsItemChannelsItemPlannerPlansItemUnarchiveRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) WithUrl(rawUrl string)(*DeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
