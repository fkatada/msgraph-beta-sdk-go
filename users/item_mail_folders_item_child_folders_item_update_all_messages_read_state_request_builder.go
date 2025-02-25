package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder provides operations to call the updateAllMessagesReadState method.
type ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilderInternal instantiates a new ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/updateAllMessagesReadState", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder instantiates a new ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateAllMessagesReadState
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI on 2024-04-29 and will be removed 2024-06-30
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder) Post(ctx context.Context, body ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action updateAllMessagesReadState
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI on 2024-04-29 and will be removed 2024-06-30
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStatePostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI on 2024-04-29 and will be removed 2024-06-30
// returns a *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemUpdateAllMessagesReadStateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
