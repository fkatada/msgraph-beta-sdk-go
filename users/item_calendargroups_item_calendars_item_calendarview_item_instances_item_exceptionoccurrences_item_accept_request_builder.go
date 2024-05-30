package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder provides operations to call the accept method.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/accept", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post accept the specified event in a user calendar.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-accept?view=graph-rest-beta
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) Post(ctx context.Context, body ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation accept the specified event in a user calendar.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
