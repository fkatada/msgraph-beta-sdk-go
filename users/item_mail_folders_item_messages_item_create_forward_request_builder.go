package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder provides operations to call the createForward method.
type ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemMessagesItemCreateForwardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessagesItemCreateForwardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemMessagesItemCreateForwardRequestBuilderInternal instantiates a new CreateForwardRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessagesItemCreateForwardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder) {
    m := &ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/createForward", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemMessagesItemCreateForwardRequestBuilder instantiates a new CreateForwardRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessagesItemCreateForwardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemMessagesItemCreateForwardRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a draft to forward an existing message, in either JSON or MIME format. When using JSON format, you can: - Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- Specify either the `toRecipients` parameter or the **toRecipients** property of the `message` parameter. Specifying both or specifying neither will return an HTTP 400 Bad Request error.- Update the draft later to add content to the **body** or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. Send the draft message in a subsequent operation. Alternatively, forward a message in a single operation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/message-createforward?view=graph-rest-1.0
func (m *ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder) Post(ctx context.Context, body ItemMailFoldersItemMessagesItemCreateForwardPostRequestBodyable, requestConfiguration *ItemMailFoldersItemMessagesItemCreateForwardRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
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
// ToPostRequestInformation create a draft to forward an existing message, in either JSON or MIME format. When using JSON format, you can: - Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- Specify either the `toRecipients` parameter or the **toRecipients** property of the `message` parameter. Specifying both or specifying neither will return an HTTP 400 Bad Request error.- Update the draft later to add content to the **body** or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. Send the draft message in a subsequent operation. Alternatively, forward a message in a single operation.
func (m *ItemMailFoldersItemMessagesItemCreateForwardRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailFoldersItemMessagesItemCreateForwardPostRequestBodyable, requestConfiguration *ItemMailFoldersItemMessagesItemCreateForwardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
