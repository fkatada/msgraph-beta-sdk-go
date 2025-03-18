package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder provides operations to call the replyWithQuote method.
type ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilderInternal instantiates a new ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder) {
    m := &ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/messages/replyWithQuote", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder instantiates a new ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reply with quote to a single chat message or multiple chat messages in a chat.
// returns a ChatMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-replywithquote?view=graph-rest-beta
func (m *ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder) Post(ctx context.Context, body ItemTeamDefinitionChannelsItemMessagesReplyWithQuotePostRequestBodyable, requestConfiguration *ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable), nil
}
// ToPostRequestInformation reply with quote to a single chat message or multiple chat messages in a chat.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamDefinitionChannelsItemMessagesReplyWithQuotePostRequestBodyable, requestConfiguration *ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMessagesReplyWithQuoteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
