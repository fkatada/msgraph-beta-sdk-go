package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder provides operations to call the permanentDelete method.
type ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilderInternal instantiates a new ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/permanentDelete", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder instantiates a new ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post permanently delete a calendar folder and the events that it contains and remove them from the mailbox. For more information about item retention, see Configure Deleted Item retention and Recoverable Items quotas.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/calendar-permanentdelete?view=graph-rest-beta
func (m *ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation permanently delete a calendar folder and the events that it contains and remove them from the mailbox. For more information about item retention, see Configure Deleted Item retention and Recoverable Items quotas.
// returns a *RequestInformation when successful
func (m *ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemPermanentDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
