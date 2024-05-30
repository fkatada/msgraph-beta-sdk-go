package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder provides operations to call the stopApplyDecisions method.
type AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/stopApplyDecisions", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action stopApplyDecisions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action stopApplyDecisions
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
