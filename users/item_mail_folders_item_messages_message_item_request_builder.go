package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type ItemMailFoldersItemMessagesMessageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMailFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type ItemMailFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMailFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// ItemMailFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
// returns a *ItemMailFoldersItemMessagesItemAttachmentsRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Attachments()(*ItemMailFoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailFoldersItemMessagesMessageItemRequestBuilderInternal instantiates a new ItemMailFoldersItemMessagesMessageItemRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessagesMessageItemRequestBuilder) {
    m := &ItemMailFoldersItemMessagesMessageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemMessagesMessageItemRequestBuilder instantiates a new ItemMailFoldersItemMessagesMessageItemRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
// returns a *ItemMailFoldersItemMessagesItemValueContentRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Content()(*ItemMailFoldersItemMessagesItemValueContentRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *ItemMailFoldersItemMessagesItemCopyRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Copy()(*ItemMailFoldersItemMessagesItemCopyRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateForward provides operations to call the createForward method.
// returns a *ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) CreateForward()(*ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemCreateForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateReply provides operations to call the createReply method.
// returns a *ItemMailFoldersItemMessagesItemCreateReplyRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) CreateReply()(*ItemMailFoldersItemMessagesItemCreateReplyRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemCreateReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateReplyAll provides operations to call the createReplyAll method.
// returns a *ItemMailFoldersItemMessagesItemCreateReplyAllRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) CreateReplyAll()(*ItemMailFoldersItemMessagesItemCreateReplyAllRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemCreateReplyAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property messages for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.message entity.
// returns a *ItemMailFoldersItemMessagesItemExtensionsRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Extensions()(*ItemMailFoldersItemMessagesItemExtensionsRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemMailFoldersItemMessagesItemForwardRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Forward()(*ItemMailFoldersItemMessagesItemForwardRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of messages in the mailFolder.
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// MarkAsJunk provides operations to call the markAsJunk method.
// returns a *ItemMailFoldersItemMessagesItemMarkAsJunkRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) MarkAsJunk()(*ItemMailFoldersItemMessagesItemMarkAsJunkRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemMarkAsJunkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MarkAsNotJunk provides operations to call the markAsNotJunk method.
// returns a *ItemMailFoldersItemMessagesItemMarkAsNotJunkRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) MarkAsNotJunk()(*ItemMailFoldersItemMessagesItemMarkAsNotJunkRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemMarkAsNotJunkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mentions provides operations to manage the mentions property of the microsoft.graph.message entity.
// returns a *ItemMailFoldersItemMessagesItemMentionsRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Mentions()(*ItemMailFoldersItemMessagesItemMentionsRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemMentionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Move provides operations to call the move method.
// returns a *ItemMailFoldersItemMessagesItemMoveRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Move()(*ItemMailFoldersItemMessagesItemMoveRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property messages in users
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *ItemMailFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemMailFoldersItemMessagesItemPermanentDeleteRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) PermanentDelete()(*ItemMailFoldersItemMessagesItemPermanentDeleteRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reply provides operations to call the reply method.
// returns a *ItemMailFoldersItemMessagesItemReplyRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Reply()(*ItemMailFoldersItemMessagesItemReplyRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReplyAll provides operations to call the replyAll method.
// returns a *ItemMailFoldersItemMessagesItemReplyAllRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) ReplyAll()(*ItemMailFoldersItemMessagesItemReplyAllRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemReplyAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Send provides operations to call the send method.
// returns a *ItemMailFoldersItemMessagesItemSendRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Send()(*ItemMailFoldersItemMessagesItemSendRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property messages for users
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of messages in the mailFolder.
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property messages in users
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *ItemMailFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// Unsubscribe provides operations to call the unsubscribe method.
// returns a *ItemMailFoldersItemMessagesItemUnsubscribeRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) Unsubscribe()(*ItemMailFoldersItemMessagesItemUnsubscribeRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemUnsubscribeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemMailFoldersItemMessagesMessageItemRequestBuilder when successful
func (m *ItemMailFoldersItemMessagesMessageItemRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemMessagesMessageItemRequestBuilder) {
    return NewItemMailFoldersItemMessagesMessageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
