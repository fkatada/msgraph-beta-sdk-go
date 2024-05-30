package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder provides operations to call the forward method.
type ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilderInternal instantiates a new ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder) {
    m := &ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo/forward", pathParameters),
    }
    return m
}
// NewItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder instantiates a new ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilderInternal(urlParams, requestAdapter)
}
// Post forward a post to a recipient. You can specify both the parent conversation and thread in the request, or, you can specify just the parent thread without the parent conversation. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/post-forward?view=graph-rest-beta
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder) Post(ctx context.Context, body ItemConversationsItemThreadsItemPostsItemInreplytoForwardPostRequestBodyable, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation forward a post to a recipient. You can specify both the parent conversation and thread in the request, or, you can specify just the parent thread without the parent conversation. 
// returns a *RequestInformation when successful
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemConversationsItemThreadsItemPostsItemInreplytoForwardPostRequestBodyable, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder when successful
func (m *ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder) WithUrl(rawUrl string)(*ItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInreplytoForwardRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
