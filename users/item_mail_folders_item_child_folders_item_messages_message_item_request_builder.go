package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Attachments()(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal instantiates a new ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder instantiates a new ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Content()(*ItemMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemCopyRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Copy()(*ItemMailFoldersItemChildFoldersItemMessagesItemCopyRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateForward provides operations to call the createForward method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemCreateForwardRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateForward()(*ItemMailFoldersItemChildFoldersItemMessagesItemCreateForwardRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemCreateForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateReply provides operations to call the createReply method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateReply()(*ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateReplyAll provides operations to call the createReplyAll method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateReplyAll()(*ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property messages for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Extensions()(*ItemMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemForwardRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Forward()(*ItemMailFoldersItemChildFoldersItemMessagesItemForwardRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of messages in the mailFolder.
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
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
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsJunkRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MarkAsJunk()(*ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsJunkRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMarkAsJunkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MarkAsNotJunk provides operations to call the markAsNotJunk method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MarkAsNotJunk()(*ItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMarkAsNotJunkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mentions provides operations to manage the mentions property of the microsoft.graph.message entity.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemMentionsRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Mentions()(*ItemMailFoldersItemChildFoldersItemMessagesItemMentionsRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMentionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Move provides operations to call the move method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Move()(*ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property messages in users
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
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
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemPermanentDeleteRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) PermanentDelete()(*ItemMailFoldersItemChildFoldersItemMessagesItemPermanentDeleteRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reply provides operations to call the reply method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemReplyRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Reply()(*ItemMailFoldersItemChildFoldersItemMessagesItemReplyRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReplyAll provides operations to call the replyAll method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemReplyAllRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ReplyAll()(*ItemMailFoldersItemChildFoldersItemMessagesItemReplyAllRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemReplyAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Send provides operations to call the send method.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemSendRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Send()(*ItemMailFoldersItemChildFoldersItemMessagesItemSendRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property messages for users
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailFoldersItemChildFoldersItemMessagesItemUnsubscribeRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Unsubscribe()(*ItemMailFoldersItemChildFoldersItemMessagesItemUnsubscribeRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemUnsubscribeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
