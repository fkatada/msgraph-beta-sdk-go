package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder provides operations to call the stopApplyDecisions method.
type ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilderInternal instantiates a new ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder) {
    m := &ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stopApplyDecisions", pathParameters),
    }
    return m
}
// NewItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder instantiates a new ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action stopApplyDecisions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder) WithUrl(rawUrl string)(*ItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemStopapplydecisionsStopApplyDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
