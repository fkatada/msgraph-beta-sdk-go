package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder provides operations to call the undoSoftDelete method.
type ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal instantiates a new ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder and sets the default values.
func NewItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    m := &ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/messages/{chatMessage%2Did}/undoSoftDelete", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder instantiates a new ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder and sets the default values.
func NewItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post undo soft deletion of a single chatMessage or a chat message reply in a channel or a chat.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-undosoftdelete?view=graph-rest-beta
func (m *ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation undo soft deletion of a single chatMessage or a chat message reply in a channel or a chat.
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder when successful
func (m *ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder) {
    return NewItemTeamdefinitionPrimarychannelMessagesItemUndosoftdeleteUndoSoftDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
