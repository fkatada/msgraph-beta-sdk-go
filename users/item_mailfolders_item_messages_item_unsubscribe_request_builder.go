package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder provides operations to call the unsubscribe method.
type ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailfoldersItemMessagesItemUnsubscribeRequestBuilderInternal instantiates a new ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder and sets the default values.
func NewItemMailfoldersItemMessagesItemUnsubscribeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder) {
    m := &ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/unsubscribe", pathParameters),
    }
    return m
}
// NewItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder instantiates a new ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder and sets the default values.
func NewItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersItemMessagesItemUnsubscribeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post submits an email request on behalf of the signed-in user to unsubscribe from an email distribution list. Uses the information in the List-Unsubscribe header. Message senders can use mailing lists in a user-friendly way by including an option for recipients to opt out. They can do so by specifying the List-Unsubscribe header in each message following RFC-2369. Note In particular, for the unsubscribe action to work, the sender must specify mailto: and not URL-based unsubscribe information. Setting that header would also set the unsubscribeEnabled property of the message instance to true, and the unsubscribeData property to the header data. If the unsubscribeEnabled property of a message is true, you can use the unsubscribe action to unsubscribe the user from similar future messages as managed by the message sender. A successful unsubscribe action moves the message to the Deleted Items folder. The actual exclusion of the user from future mail distribution is managed by the sender.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/message-unsubscribe?view=graph-rest-beta
func (m *ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation submits an email request on behalf of the signed-in user to unsubscribe from an email distribution list. Uses the information in the List-Unsubscribe header. Message senders can use mailing lists in a user-friendly way by including an option for recipients to opt out. They can do so by specifying the List-Unsubscribe header in each message following RFC-2369. Note In particular, for the unsubscribe action to work, the sender must specify mailto: and not URL-based unsubscribe information. Setting that header would also set the unsubscribeEnabled property of the message instance to true, and the unsubscribeData property to the header data. If the unsubscribeEnabled property of a message is true, you can use the unsubscribe action to unsubscribe the user from similar future messages as managed by the message sender. A successful unsubscribe action moves the message to the Deleted Items folder. The actual exclusion of the user from future mail distribution is managed by the sender.
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder when successful
func (m *ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder) {
    return NewItemMailfoldersItemMessagesItemUnsubscribeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
