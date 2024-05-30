package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetQueryParameters the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
type ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
// returns a *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Attachments()(*ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderInternal instantiates a new ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) {
    m := &ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder instantiates a new ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.post entity.
// returns a *ItemThreadsItemPostsItemInreplytoExtensionsRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Extensions()(*ItemThreadsItemPostsItemInreplytoExtensionsRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemThreadsItemPostsItemInreplytoForwardRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Forward()(*ItemThreadsItemPostsItemInreplytoForwardRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
// returns a Postable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable), nil
}
// Mentions provides operations to manage the mentions property of the microsoft.graph.post entity.
// returns a *ItemThreadsItemPostsItemInreplytoMentionsRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Mentions()(*ItemThreadsItemPostsItemInreplytoMentionsRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoMentionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reply provides operations to call the reply method.
// returns a *ItemThreadsItemPostsItemInreplytoReplyRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) Reply()(*ItemThreadsItemPostsItemInreplytoReplyRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) WithUrl(rawUrl string)(*ItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoInReplyToRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
