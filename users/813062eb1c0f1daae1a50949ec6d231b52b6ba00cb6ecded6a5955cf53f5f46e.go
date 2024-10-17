package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) PermanentDelete()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
