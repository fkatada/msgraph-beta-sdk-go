package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarviewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type ItemCalendarCalendarviewEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarviewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type ItemCalendarCalendarviewEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarCalendarviewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarviewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarviewEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarCalendarviewItemAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Accept()(*ItemCalendarCalendarviewItemAcceptRequestBuilder) {
    return NewItemCalendarCalendarviewItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Attachments()(*ItemCalendarCalendarviewItemAttachmentsRequestBuilder) {
    return NewItemCalendarCalendarviewItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemCalendarRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Calendar()(*ItemCalendarCalendarviewItemCalendarRequestBuilder) {
    return NewItemCalendarCalendarviewItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarCalendarviewItemCancelRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Cancel()(*ItemCalendarCalendarviewItemCancelRequestBuilder) {
    return NewItemCalendarCalendarviewItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarviewEventItemRequestBuilderInternal instantiates a new ItemCalendarCalendarviewEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewEventItemRequestBuilder) {
    m := &ItemCalendarCalendarviewEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarviewEventItemRequestBuilder instantiates a new ItemCalendarCalendarviewEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarviewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarCalendarviewItemDeclineRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Decline()(*ItemCalendarCalendarviewItemDeclineRequestBuilder) {
    return NewItemCalendarCalendarviewItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarCalendarviewItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) DismissReminder()(*ItemCalendarCalendarviewItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarCalendarviewItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) ExceptionOccurrences()(*ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) {
    return NewItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemExtensionsRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Extensions()(*ItemCalendarCalendarviewItemExtensionsRequestBuilder) {
    return NewItemCalendarCalendarviewItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarCalendarviewItemForwardRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Forward()(*ItemCalendarCalendarviewItemForwardRequestBuilder) {
    return NewItemCalendarCalendarviewItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the calendar view for the calendar. Navigation property. Read-only.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarCalendarviewItemInstancesRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) Instances()(*ItemCalendarCalendarviewItemInstancesRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarCalendarviewEventItemRequestBuilder when successful
func (m *ItemCalendarCalendarviewEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarviewEventItemRequestBuilder) {
    return NewItemCalendarCalendarviewEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
