package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder provides operations to manage the calendar property of the microsoft.graph.event entity.
type ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderGetQueryParameters the calendar that contains the event. Navigation property. Read-only.
type ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderGetQueryParameters
}
// NewItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderInternal instantiates a new ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder and sets the default values.
func NewItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder) {
    m := &ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/calendar{?%24select}", pathParameters),
    }
    return m
}
// NewItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder instantiates a new ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder and sets the default values.
func NewItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the calendar that contains the event. Navigation property. Read-only.
// returns a Calendarable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCalendarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable), nil
}
// ToGetRequestInformation the calendar that contains the event. Navigation property. Read-only.
// returns a *RequestInformation when successful
func (m *ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder) WithUrl(rawUrl string)(*ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
