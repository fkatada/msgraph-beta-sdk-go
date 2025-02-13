package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.plannerBucket entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters read-only. Nullable. The collection of tasks in the bucket.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedToTaskBoardFormat provides operations to manage the assignedToTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemAssignedToTaskBoardFormatRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) AssignedToTaskBoardFormat()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemAssignedToTaskBoardFormatRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemAssignedToTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BucketTaskBoardFormat provides operations to manage the bucketTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemBucketTaskBoardFormatRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) BucketTaskBoardFormat()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemBucketTaskBoardFormatRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemBucketTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks/{plannerTask%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tasks for teamwork
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Details provides operations to manage the details property of the microsoft.graph.plannerTask entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemDetailsRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Details()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemDetailsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. The collection of tasks in the bucket.
// returns a PlannerTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable), nil
}
// Patch update the navigation property tasks in teamwork
// returns a PlannerTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable), nil
}
// ProgressTaskBoardFormat provides operations to manage the progressTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemProgressTaskBoardFormatRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ProgressTaskBoardFormat()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemProgressTaskBoardFormatRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksItemProgressTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property tasks for teamwork
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. The collection of tasks in the bucket.
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tasks in teamwork
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
