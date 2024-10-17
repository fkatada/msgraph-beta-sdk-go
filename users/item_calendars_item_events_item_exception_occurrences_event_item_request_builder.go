package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder instantiates a new ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) PermanentDelete()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
