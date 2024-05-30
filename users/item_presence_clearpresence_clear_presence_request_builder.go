package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPresenceClearpresenceClearPresenceRequestBuilder provides operations to call the clearPresence method.
type ItemPresenceClearpresenceClearPresenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPresenceClearpresenceClearPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPresenceClearpresenceClearPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPresenceClearpresenceClearPresenceRequestBuilderInternal instantiates a new ItemPresenceClearpresenceClearPresenceRequestBuilder and sets the default values.
func NewItemPresenceClearpresenceClearPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceClearpresenceClearPresenceRequestBuilder) {
    m := &ItemPresenceClearpresenceClearPresenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/presence/clearPresence", pathParameters),
    }
    return m
}
// NewItemPresenceClearpresenceClearPresenceRequestBuilder instantiates a new ItemPresenceClearpresenceClearPresenceRequestBuilder and sets the default values.
func NewItemPresenceClearpresenceClearPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceClearpresenceClearPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPresenceClearpresenceClearPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear a presence session of an application for a user. If it is the user's only presence session, a successful clearPresence changes the user's presence to Offline/Offline. Read more about presence sessions and their time-out and expiration. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-clearpresence?view=graph-rest-beta
func (m *ItemPresenceClearpresenceClearPresenceRequestBuilder) Post(ctx context.Context, body ItemPresenceClearpresenceClearPresencePostRequestBodyable, requestConfiguration *ItemPresenceClearpresenceClearPresenceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clear a presence session of an application for a user. If it is the user's only presence session, a successful clearPresence changes the user's presence to Offline/Offline. Read more about presence sessions and their time-out and expiration. 
// returns a *RequestInformation when successful
func (m *ItemPresenceClearpresenceClearPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPresenceClearpresenceClearPresencePostRequestBodyable, requestConfiguration *ItemPresenceClearpresenceClearPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPresenceClearpresenceClearPresenceRequestBuilder when successful
func (m *ItemPresenceClearpresenceClearPresenceRequestBuilder) WithUrl(rawUrl string)(*ItemPresenceClearpresenceClearPresenceRequestBuilder) {
    return NewItemPresenceClearpresenceClearPresenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
