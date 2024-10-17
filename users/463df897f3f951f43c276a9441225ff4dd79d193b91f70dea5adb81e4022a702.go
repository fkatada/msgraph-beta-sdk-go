package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder provides operations to count the resources in the collection.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderGetQueryParameters
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderInternal instantiates a new ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/attachments/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder instantiates a new ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
