package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder provides operations to call the removeEmail method.
type DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderInternal instantiates a new DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder) {
    m := &DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/removeEmail", pathParameters),
    }
    return m
}
// NewDeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder instantiates a new DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove the email address of a channel. You can remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-removeemail?view=graph-rest-beta
func (m *DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder) Post(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove the email address of a channel. You can remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder) {
    return NewDeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
