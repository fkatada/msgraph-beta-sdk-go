package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder provides operations to call the recordDecisions method.
type ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilderInternal instantiates a new ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder and sets the default values.
func NewItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder) {
    m := &ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}/businessFlowsWithRequestsAwaitingMyDecision/{businessFlow%2Did}/recordDecisions", pathParameters),
    }
    return m
}
// NewItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder instantiates a new ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder and sets the default values.
func NewItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action recordDecisions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder) Post(ctx context.Context, body ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsPostRequestBodyable, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action recordDecisions
// returns a *RequestInformation when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsPostRequestBodyable, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder) WithUrl(rawUrl string)(*ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder) {
    return NewItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
