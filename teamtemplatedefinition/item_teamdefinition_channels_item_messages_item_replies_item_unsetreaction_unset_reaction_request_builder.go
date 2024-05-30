package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder provides operations to call the unsetReaction method.
type ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal instantiates a new ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder and sets the default values.
func NewItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    m := &ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/unsetReaction", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder instantiates a new ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder and sets the default values.
func NewItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unsetReaction
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) Post(ctx context.Context, body ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionPostRequestBodyable, requestConfiguration *ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action unsetReaction
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionPostRequestBodyable, requestConfiguration *ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder when successful
func (m *ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    return NewItemTeamdefinitionChannelsItemMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
