package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder instantiates a new ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) PermanentDelete()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
