package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder provides operations to call the forward method.
type ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderInternal instantiates a new ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) {
    m := &ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/forward", pathParameters),
    }
    return m
}
// NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder instantiates a new ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this action allows the organizer or attendee of a meeting event to forward themeeting request to a new recipient. If the meeting event is forwarded from an attendee's Microsoft 365 mailbox to another recipient, this actionalso sends a message to notify the organizer of the forwarding, and adds the recipient to the organizer'scopy of the meeting event. This convenience is not available when forwarding from an Outlook.com account.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-forward?view=graph-rest-beta
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) Post(ctx context.Context, body ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardPostRequestBodyable, requestConfiguration *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation this action allows the organizer or attendee of a meeting event to forward themeeting request to a new recipient. If the meeting event is forwarded from an attendee's Microsoft 365 mailbox to another recipient, this actionalso sends a message to notify the organizer of the forwarding, and adds the recipient to the organizer'scopy of the meeting event. This convenience is not available when forwarding from an Outlook.com account.
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardPostRequestBodyable, requestConfiguration *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
