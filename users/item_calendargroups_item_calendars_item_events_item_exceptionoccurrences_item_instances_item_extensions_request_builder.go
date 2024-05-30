package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder provides operations to manage the extensions property of the microsoft.graph.event entity.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderGetQueryParameters the collection of open extensions defined for the event. Nullable.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderGetQueryParameters
}
// ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByExtensionId provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) ByExtensionId(extensionId string)(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if extensionId != "" {
        urlTplParams["extension%2Did"] = extensionId
    }
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/extensions{?%24count,%24expand,%24filter,%24orderby,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsCountRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) Count()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsCountRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of open extensions defined for the event. Nullable.
// returns a ExtensionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExtensionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExtensionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExtensionCollectionResponseable), nil
}
// Post create new navigation property to extensions for users
// returns a Extensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Extensionable, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Extensionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Extensionable), nil
}
// ToGetRequestInformation the collection of open extensions defined for the event. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to extensions for users
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Extensionable, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
