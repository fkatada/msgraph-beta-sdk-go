package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderGetQueryParameters
}
// ByEventId1 provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemExceptionoccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) ByEventId1(eventId1 string)(*ItemCalendarCalendarviewItemExceptionoccurrencesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if eventId1 != "" {
        urlTplParams["event%2Did1"] = eventId1
    }
    return NewItemCalendarCalendarviewItemExceptionoccurrencesEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderInternal instantiates a new ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) {
    m := &ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder instantiates a new ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemCalendarCalendarviewItemExceptionoccurrencesCountRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) Count()(*ItemCalendarCalendarviewItemExceptionoccurrencesCountRequestBuilder) {
    return NewItemCalendarCalendarviewItemExceptionoccurrencesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *ItemCalendarCalendarviewItemExceptionoccurrencesDeltaRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) Delta()(*ItemCalendarCalendarviewItemExceptionoccurrencesDeltaRequestBuilder) {
    return NewItemCalendarCalendarviewItemExceptionoccurrencesDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a EventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable), nil
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) {
    return NewItemCalendarCalendarviewItemExceptionoccurrencesExceptionOccurrencesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
