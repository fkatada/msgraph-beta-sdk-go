package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern and exceptions that have been modified. It doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCancelRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder instantiates a new ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern and exceptions that have been modified. It doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) PermanentDelete()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern and exceptions that have been modified. It doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
