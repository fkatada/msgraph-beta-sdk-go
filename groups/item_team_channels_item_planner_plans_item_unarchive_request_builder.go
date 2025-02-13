package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder provides operations to call the unarchive method.
type ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilderInternal instantiates a new ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder and sets the default values.
func NewItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder) {
    m := &ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}/planner/plans/{plannerPlan%2Did}/unarchive", pathParameters),
    }
    return m
}
// NewItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder instantiates a new ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder and sets the default values.
func NewItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unarchive a plannerPlan object. Unarchiving a plan, also unarchives the plannerTasks and plannerBuckets in the plan.  Only a plan that is archived can be unarchived.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannerplan-unarchive?view=graph-rest-beta
func (m *ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder) Post(ctx context.Context, body ItemTeamChannelsItemPlannerPlansItemUnarchivePostRequestBodyable, requestConfiguration *ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation unarchive a plannerPlan object. Unarchiving a plan, also unarchives the plannerTasks and plannerBuckets in the plan.  Only a plan that is archived can be unarchived.
// returns a *RequestInformation when successful
func (m *ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamChannelsItemPlannerPlansItemUnarchivePostRequestBodyable, requestConfiguration *ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder when successful
func (m *ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder) WithUrl(rawUrl string)(*ItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder) {
    return NewItemTeamChannelsItemPlannerPlansItemUnarchiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
