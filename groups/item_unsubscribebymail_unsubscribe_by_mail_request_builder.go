package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemUnsubscribebymailUnsubscribeByMailRequestBuilder provides operations to call the unsubscribeByMail method.
type ItemUnsubscribebymailUnsubscribeByMailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemUnsubscribebymailUnsubscribeByMailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemUnsubscribebymailUnsubscribeByMailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemUnsubscribebymailUnsubscribeByMailRequestBuilderInternal instantiates a new ItemUnsubscribebymailUnsubscribeByMailRequestBuilder and sets the default values.
func NewItemUnsubscribebymailUnsubscribeByMailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnsubscribebymailUnsubscribeByMailRequestBuilder) {
    m := &ItemUnsubscribebymailUnsubscribeByMailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/unsubscribeByMail", pathParameters),
    }
    return m
}
// NewItemUnsubscribebymailUnsubscribeByMailRequestBuilder instantiates a new ItemUnsubscribebymailUnsubscribeByMailRequestBuilder and sets the default values.
func NewItemUnsubscribebymailUnsubscribeByMailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnsubscribebymailUnsubscribeByMailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUnsubscribebymailUnsubscribeByMailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post calling this method disables the current user to receive email notifications for this group about new posts, events, and files in that group. Supported for Microsoft 365 groups only.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-unsubscribebymail?view=graph-rest-beta
func (m *ItemUnsubscribebymailUnsubscribeByMailRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemUnsubscribebymailUnsubscribeByMailRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation calling this method disables the current user to receive email notifications for this group about new posts, events, and files in that group. Supported for Microsoft 365 groups only.
// returns a *RequestInformation when successful
func (m *ItemUnsubscribebymailUnsubscribeByMailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemUnsubscribebymailUnsubscribeByMailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemUnsubscribebymailUnsubscribeByMailRequestBuilder when successful
func (m *ItemUnsubscribebymailUnsubscribeByMailRequestBuilder) WithUrl(rawUrl string)(*ItemUnsubscribebymailUnsubscribeByMailRequestBuilder) {
    return NewItemUnsubscribebymailUnsubscribeByMailRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
