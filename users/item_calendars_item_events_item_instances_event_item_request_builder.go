package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemInstancesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarsItemEventsItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemInstancesItemCalendarRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarsItemEventsItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarsItemEventsItemInstancesItemCancelRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarsItemEventsItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarsItemEventsItemInstancesEventItemRequestBuilderInternal instantiates a new ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarsItemEventsItemInstancesEventItemRequestBuilder instantiates a new ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemEventsItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarsItemEventsItemInstancesItemDeclineRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarsItemEventsItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarsItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarsItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) ExceptionOccurrences()(*ItemCalendarsItemEventsItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarsItemEventsItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarsItemEventsItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarsItemEventsItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarsItemEventsItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarsItemEventsItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarsItemEventsItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarsItemEventsItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
