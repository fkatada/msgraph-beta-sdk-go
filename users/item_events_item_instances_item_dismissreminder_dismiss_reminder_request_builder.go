package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder provides operations to call the dismissReminder method.
type ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal instantiates a new ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    m := &ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/events/{event%2Did}/instances/{event%2Did1}/dismissReminder", pathParameters),
    }
    return m
}
// NewItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder instantiates a new ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post dismiss a reminder that has been triggered for an event in a user calendar.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-dismissreminder?view=graph-rest-beta
func (m *ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation dismiss a reminder that has been triggered for an event in a user calendar.
// returns a *RequestInformation when successful
func (m *ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder) WithUrl(rawUrl string)(*ItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
