package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedteamsItemChannelsItemArchiveRequestBuilder provides operations to call the archive method.
type DeletedteamsItemChannelsItemArchiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsItemChannelsItemArchiveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemArchiveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedteamsItemChannelsItemArchiveRequestBuilderInternal instantiates a new DeletedteamsItemChannelsItemArchiveRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemArchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemArchiveRequestBuilder) {
    m := &DeletedteamsItemChannelsItemArchiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/archive", pathParameters),
    }
    return m
}
// NewDeletedteamsItemChannelsItemArchiveRequestBuilder instantiates a new DeletedteamsItemChannelsItemArchiveRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemArchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemArchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsItemChannelsItemArchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post archive a channel in a team. When a channel is archived, users can't send new messages or react to existing messages in the channel, edit the channel settings, or make other changes to the channel. You can delete an archived channel, or add and remove members from it. If you archive a team, its channels are archived for you. Archiving is asynchronous; a channel is archived after the asynchronous archiving operation completes successfully, which might occur after the response returns. A channel without an owner, or that belongs to a group that has no owner, can't be archived. To restore a channel from its archived state, use the unarchive method. A channel can’t be archived or unarchived if its team is archived.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-archive?view=graph-rest-beta
func (m *DeletedteamsItemChannelsItemArchiveRequestBuilder) Post(ctx context.Context, body DeletedteamsItemChannelsItemArchivePostRequestBodyable, requestConfiguration *DeletedteamsItemChannelsItemArchiveRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation archive a channel in a team. When a channel is archived, users can't send new messages or react to existing messages in the channel, edit the channel settings, or make other changes to the channel. You can delete an archived channel, or add and remove members from it. If you archive a team, its channels are archived for you. Archiving is asynchronous; a channel is archived after the asynchronous archiving operation completes successfully, which might occur after the response returns. A channel without an owner, or that belongs to a group that has no owner, can't be archived. To restore a channel from its archived state, use the unarchive method. A channel can’t be archived or unarchived if its team is archived.
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemArchiveRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeletedteamsItemChannelsItemArchivePostRequestBodyable, requestConfiguration *DeletedteamsItemChannelsItemArchiveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedteamsItemChannelsItemArchiveRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemArchiveRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsItemChannelsItemArchiveRequestBuilder) {
    return NewDeletedteamsItemChannelsItemArchiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
