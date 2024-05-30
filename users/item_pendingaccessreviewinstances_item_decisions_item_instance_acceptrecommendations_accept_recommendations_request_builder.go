package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder provides operations to call the acceptRecommendations method.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilderInternal instantiates a new ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder) {
    m := &ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/acceptRecommendations", pathParameters),
    }
    return m
}
// NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder instantiates a new ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post allows the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that haven't been reviewed for an accessReviewInstance object for which the calling user is a reviewer. Recommendations are generated if recommendationsEnabled is true on the accessReviewScheduleDefinition object. If there isn't a recommendation on an accessReviewInstanceDecisionItem object no decision will be recorded.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-acceptrecommendations?view=graph-rest-beta
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation allows the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that haven't been reviewed for an accessReviewInstance object for which the calling user is a reviewer. Recommendations are generated if recommendationsEnabled is true on the accessReviewScheduleDefinition object. If there isn't a recommendation on an accessReviewInstanceDecisionItem object no decision will be recorded.
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder) WithUrl(rawUrl string)(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
