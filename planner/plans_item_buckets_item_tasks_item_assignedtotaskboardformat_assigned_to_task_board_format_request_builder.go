package planner

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder provides operations to manage the assignedToTaskBoardFormat property of the microsoft.graph.plannerTask entity.
type PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetQueryParameters read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
type PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetQueryParameters
}
// PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderInternal instantiates a new PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder and sets the default values.
func NewPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) {
    m := &PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks/{plannerTask%2Did}/assignedToTaskBoardFormat{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder instantiates a new PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder and sets the default values.
func NewPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignedToTaskBoardFormat for planner
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) Delete(ctx context.Context, requestConfiguration *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
// returns a PlannerAssignedToTaskBoardTaskFormatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) Get(ctx context.Context, requestConfiguration *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerAssignedToTaskBoardTaskFormatable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerAssignedToTaskBoardTaskFormatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerAssignedToTaskBoardTaskFormatable), nil
}
// Patch update the navigation property assignedToTaskBoardFormat in planner
// returns a PlannerAssignedToTaskBoardTaskFormatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerAssignedToTaskBoardTaskFormatable, requestConfiguration *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerAssignedToTaskBoardTaskFormatable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerAssignedToTaskBoardTaskFormatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerAssignedToTaskBoardTaskFormatable), nil
}
// ToDeleteRequestInformation delete navigation property assignedToTaskBoardFormat for planner
// returns a *RequestInformation when successful
func (m *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
// returns a *RequestInformation when successful
func (m *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignedToTaskBoardFormat in planner
// returns a *RequestInformation when successful
func (m *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerAssignedToTaskBoardTaskFormatable, requestConfiguration *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder when successful
func (m *PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) WithUrl(rawUrl string)(*PlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) {
    return NewPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
