package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
type AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderGetQueryParameters there's exactly one accessReviewScheduleDefinition associated with each instance. It's the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
type AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderGetQueryParameters
}
// NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/definition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get there's exactly one accessReviewScheduleDefinition associated with each instance. It's the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
// returns a AccessReviewScheduleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewScheduleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable), nil
}
// ToGetRequestInformation there's exactly one accessReviewScheduleDefinition associated with each instance. It's the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
