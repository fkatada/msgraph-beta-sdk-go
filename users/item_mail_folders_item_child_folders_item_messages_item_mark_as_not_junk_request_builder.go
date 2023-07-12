package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder provides operations to call the markAsNotJunk method.
type ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilderInternal instantiates a new MarkAsNotJunkRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/markAsNotJunk", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder instantiates a new MarkAsNotJunkRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark a message as not junk. This API removes the sender from the list of blocked senders and moves the message to the **Inbox** folder, when **moveToInbox** is `true`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/message-markasnotjunk?view=graph-rest-1.0
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder) Post(ctx context.Context, body ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkPostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// ToPostRequestInformation mark a message as not junk. This API removes the sender from the list of blocked senders and moves the message to the **Inbox** folder, when **moveToInbox** is `true`.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkPostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
