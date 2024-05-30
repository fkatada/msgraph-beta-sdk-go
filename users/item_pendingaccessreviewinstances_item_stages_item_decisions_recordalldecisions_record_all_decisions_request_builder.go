package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder provides operations to call the recordAllDecisions method.
type ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilderInternal instantiates a new ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder) {
    m := &ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/recordAllDecisions", pathParameters),
    }
    return m
}
// NewItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder instantiates a new ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post as a reviewer of an access review, record a decision for an accessReviewInstanceDecisionItem that is assigned to you and that matches the principal or resource IDs specified. If no IDs are specified, the decisions will apply to every accessReviewInstanceDecisionItem for which you are the reviewer.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstancedecisionitem-recordalldecisions?view=graph-rest-beta
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder) Post(ctx context.Context, body ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsPostRequestBodyable, requestConfiguration *ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation as a reviewer of an access review, record a decision for an accessReviewInstanceDecisionItem that is assigned to you and that matches the principal or resource IDs specified. If no IDs are specified, the decisions will apply to every accessReviewInstanceDecisionItem for which you are the reviewer.
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsPostRequestBodyable, requestConfiguration *ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder) WithUrl(rawUrl string)(*ItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemStagesItemDecisionsRecordalldecisionsRecordAllDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
