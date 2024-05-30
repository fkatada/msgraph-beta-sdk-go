package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarsItemCalendarviewItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarsItemCalendarviewItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemCalendarRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarsItemCalendarviewItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemCancelRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarsItemCalendarviewItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderInternal instantiates a new ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder instantiates a new ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemDeclineRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarsItemCalendarviewItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarsItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) ExceptionOccurrences()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarsItemCalendarviewItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarsItemCalendarviewItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarsItemCalendarviewItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
