package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder provides operations to manage the plans property of the microsoft.graph.teamsChannelPlanner entity.
type ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters a collection of plannerPlan objects owned by the Teams channel. Currently, only shared channels are supported. Read-only. Nullable.
type ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetQueryParameters
}
// ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Archive provides operations to call the archive method.
// returns a *ItemChannelsItemPlannerPlansItemArchiveRequestBuilder when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Archive()(*ItemChannelsItemPlannerPlansItemArchiveRequestBuilder) {
    return NewItemChannelsItemPlannerPlansItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Buckets provides operations to manage the buckets property of the microsoft.graph.plannerPlan entity.
// returns a *ItemChannelsItemPlannerPlansItemBucketsRequestBuilder when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Buckets()(*ItemChannelsItemPlannerPlansItemBucketsRequestBuilder) {
    return NewItemChannelsItemPlannerPlansItemBucketsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderInternal instantiates a new ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder and sets the default values.
func NewItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) {
    m := &ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/channels/{channel%2Did}/planner/plans/{plannerPlan%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder instantiates a new ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder and sets the default values.
func NewItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property plans for teams
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemChannelsItemPlannerPlansItemDetailsRequestBuilder when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Details()(*ItemChannelsItemPlannerPlansItemDetailsRequestBuilder) {
    return NewItemChannelsItemPlannerPlansItemDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of plannerPlan objects owned by the Teams channel. Currently, only shared channels are supported. Read-only. Nullable.
// returns a PlannerPlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
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
// returns a *ItemChannelsItemPlannerPlansItemMoveToContainerRequestBuilder when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) MoveToContainer()(*ItemChannelsItemPlannerPlansItemMoveToContainerRequestBuilder) {
    return NewItemChannelsItemPlannerPlansItemMoveToContainerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property plans in teams
// returns a PlannerPlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, requestConfiguration *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
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
// returns a *ItemChannelsItemPlannerPlansItemTasksRequestBuilder when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Tasks()(*ItemChannelsItemPlannerPlansItemTasksRequestBuilder) {
    return NewItemChannelsItemPlannerPlansItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property plans for teams
// returns a *RequestInformation when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property plans in teams
// returns a *RequestInformation when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, requestConfiguration *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChannelsItemPlannerPlansItemUnarchiveRequestBuilder when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) Unarchive()(*ItemChannelsItemPlannerPlansItemUnarchiveRequestBuilder) {
    return NewItemChannelsItemPlannerPlansItemUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder when successful
func (m *ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) WithUrl(rawUrl string)(*ItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder) {
    return NewItemChannelsItemPlannerPlansPlannerPlanItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
