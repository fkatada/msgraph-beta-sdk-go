package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder provides operations to call the move method.
type ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal instantiates a new ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/move", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder instantiates a new ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post move a message to another folder within the specified user's mailbox. This creates a new copy of the message in the destination folder and removes the original message.
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/message-move?view=graph-rest-beta
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) Post(ctx context.Context, body ItemMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation move a message to another folder within the specified user's mailbox. This creates a new copy of the message in the destination folder and removes the original message.
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
