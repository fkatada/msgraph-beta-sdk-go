package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder provides operations to call the delta method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderGetQueryParameters get newly created, updated, or deleted tasks in either a Planner plan or assigned to the signed-in user without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderGetQueryParameters struct {
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
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderGetQueryParameters
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderInternal instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get newly created, updated, or deleted tasks in either a Planner plan or assigned to the signed-in user without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannertask-delta?view=graph-rest-beta
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderGetRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaResponseable), nil
}
// GetAsDeltaGetResponse get newly created, updated, or deleted tasks in either a Planner plan or assigned to the signed-in user without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// returns a TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannertask-delta?view=graph-rest-beta
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderGetRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaGetResponseable), nil
}
// ToGetRequestInformation get newly created, updated, or deleted tasks in either a Planner plan or assigned to the signed-in user without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
