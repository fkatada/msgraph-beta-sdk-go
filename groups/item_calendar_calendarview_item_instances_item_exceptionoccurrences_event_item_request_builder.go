package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder instantiates a new ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from groups
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from groups
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
