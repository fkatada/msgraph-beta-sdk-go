package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder provides operations to call the decline method.
type ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilderInternal instantiates a new ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder) {
    m := &ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/decline", pathParameters),
    }
    return m
}
// NewItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder instantiates a new ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilderInternal(urlParams, requestAdapter)
}
// Post decline invitation to the specified event in a user calendar. If the event allows proposals for new times, on declining the event, an invitee can choose to suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-decline?view=graph-rest-beta
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder) Post(ctx context.Context, body ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclinePostRequestBodyable, requestConfiguration *ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation decline invitation to the specified event in a user calendar. If the event allows proposals for new times, on declining the event, an invitee can choose to suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclinePostRequestBodyable, requestConfiguration *ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
