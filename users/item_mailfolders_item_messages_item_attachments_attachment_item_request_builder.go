package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder provides operations to manage the attachments property of the microsoft.graph.message entity.
type ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters the fileAttachment and itemAttachment attachments for the message.
type ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters
}
// NewItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal instantiates a new ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder and sets the default values.
func NewItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    m := &ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/attachments/{attachment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder instantiates a new ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder and sets the default values.
func NewItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property attachments for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the fileAttachment and itemAttachment attachments for the message.
// returns a Attachmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable), nil
}
// ToDeleteRequestInformation delete navigation property attachments for users
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the fileAttachment and itemAttachment attachments for the message.
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder when successful
func (m *ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    return NewItemMailfoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
