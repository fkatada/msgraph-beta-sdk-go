package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder provides operations to call the delta method.
type ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarviewItemInstancesDeltaRequestBuilderGetQueryParameters get a set of event resources that have been added, deleted, or updated in one or more calendars. You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a calendarView (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar of the user's. In the case of getting incremental changes on calendarView, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or calendarView in a local store entails a round of multiple delta function calls. The initial call is a full synchronization, and every subsequent delta call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified calendar, without having to fetch all the events of that calendar from the server every time. The following table lists the differences between the delta function on events and the delta function on a calendarView in a calendar.
type ItemCalendarCalendarviewItemInstancesDeltaRequestBuilderGetQueryParameters struct {
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
// ItemCalendarCalendarviewItemInstancesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarviewItemInstancesDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilderGetQueryParameters
}
// NewItemCalendarCalendarviewItemInstancesDeltaRequestBuilderInternal instantiates a new ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewItemInstancesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder) {
    m := &ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/instances/delta()?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarviewItemInstancesDeltaRequestBuilder instantiates a new ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewItemInstancesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarviewItemInstancesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of event resources that have been added, deleted, or updated in one or more calendars. You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a calendarView (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar of the user's. In the case of getting incremental changes on calendarView, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or calendarView in a local store entails a round of multiple delta function calls. The initial call is a full synchronization, and every subsequent delta call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified calendar, without having to fetch all the events of that calendar from the server every time. The following table lists the differences between the delta function on events and the delta function on a calendarView in a calendar.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ItemCalendarCalendarviewItemInstancesDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-delta?view=graph-rest-beta
func (m *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilderGetRequestConfiguration)(ItemCalendarCalendarviewItemInstancesDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarCalendarviewItemInstancesDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarCalendarviewItemInstancesDeltaResponseable), nil
}
// GetAsDeltaGetResponse get a set of event resources that have been added, deleted, or updated in one or more calendars. You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a calendarView (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar of the user's. In the case of getting incremental changes on calendarView, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or calendarView in a local store entails a round of multiple delta function calls. The initial call is a full synchronization, and every subsequent delta call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified calendar, without having to fetch all the events of that calendar from the server every time. The following table lists the differences between the delta function on events and the delta function on a calendarView in a calendar.
// returns a ItemCalendarCalendarviewItemInstancesDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-delta?view=graph-rest-beta
func (m *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilderGetRequestConfiguration)(ItemCalendarCalendarviewItemInstancesDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCalendarCalendarviewItemInstancesDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCalendarCalendarviewItemInstancesDeltaGetResponseable), nil
}
// ToGetRequestInformation get a set of event resources that have been added, deleted, or updated in one or more calendars. You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a calendarView (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar of the user's. In the case of getting incremental changes on calendarView, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or calendarView in a local store entails a round of multiple delta function calls. The initial call is a full synchronization, and every subsequent delta call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified calendar, without having to fetch all the events of that calendar from the server every time. The following table lists the differences between the delta function on events and the delta function on a calendarView in a calendar.
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarviewItemInstancesDeltaRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
