package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder provides operations to call the delta method.
type ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderGetQueryParameters get a set of event resources that are added, deleted, or updated in one or more calendars. You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a calendarView (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar that belongs to the user. When getting incremental changes on calendarView, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or calendarView in a local store entails a round of multiple delta function calls. The initial call is a full synchronization, and every subsequent delta call in the same round gets the incremental changes (additions, deletions, or updates). Using deltas allows you to incrementally maintain and synchronize a local store of events in the specified calendar. The following table lists the differences between the delta function on events and the delta function on a calendarView in a calendar.
type ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // The end date and time of the time range in the function, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // The start date and time of the time range in the function, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderGetQueryParameters
}
// NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderInternal instantiates a new ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder) {
    m := &ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/delta()?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder instantiates a new ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of event resources that are added, deleted, or updated in one or more calendars. You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a calendarView (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar that belongs to the user. When getting incremental changes on calendarView, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or calendarView in a local store entails a round of multiple delta function calls. The initial call is a full synchronization, and every subsequent delta call in the same round gets the incremental changes (additions, deletions, or updates). Using deltas allows you to incrementally maintain and synchronize a local store of events in the specified calendar. The following table lists the differences between the delta function on events and the delta function on a calendarView in a calendar.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-delta?view=graph-rest-beta
func (m *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderGetRequestConfiguration)(ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseable), nil
}
// GetAsDeltaGetResponse get a set of event resources that are added, deleted, or updated in one or more calendars. You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a calendarView (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar that belongs to the user. When getting incremental changes on calendarView, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or calendarView in a local store entails a round of multiple delta function calls. The initial call is a full synchronization, and every subsequent delta call in the same round gets the incremental changes (additions, deletions, or updates). Using deltas allows you to incrementally maintain and synchronize a local store of events in the specified calendar. The following table lists the differences between the delta function on events and the delta function on a calendarView in a calendar.
// returns a ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-delta?view=graph-rest-beta
func (m *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderGetRequestConfiguration)(ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponseable), nil
}
// ToGetRequestInformation get a set of event resources that are added, deleted, or updated in one or more calendars. You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a calendarView (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar that belongs to the user. When getting incremental changes on calendarView, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or calendarView in a local store entails a round of multiple delta function calls. The initial call is a full synchronization, and every subsequent delta call in the same round gets the incremental changes (additions, deletions, or updates). Using deltas allows you to incrementally maintain and synchronize a local store of events in the specified calendar. The following table lists the differences between the delta function on events and the delta function on a calendarView in a calendar.
// returns a *RequestInformation when successful
func (m *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder when successful
func (m *ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder) {
    return NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
