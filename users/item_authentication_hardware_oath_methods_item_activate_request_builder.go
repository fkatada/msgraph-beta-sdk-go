package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder provides operations to call the activate method.
type ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationHardwareOathMethodsItemActivateRequestBuilderInternal instantiates a new ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsItemActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder) {
    m := &ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/hardwareOathMethods/{hardwareOathAuthenticationMethod%2Did}/activate", pathParameters),
    }
    return m
}
// NewItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder instantiates a new ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationHardwareOathMethodsItemActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post activate a hardware OATH token that is already assigned to a user. A user can self-activate their token or an admin can activate for a user.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/hardwareoathauthenticationmethod-activate?view=graph-rest-beta
func (m *ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder) Post(ctx context.Context, body ItemAuthenticationHardwareOathMethodsItemActivatePostRequestBodyable, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation activate a hardware OATH token that is already assigned to a user. A user can self-activate their token or an admin can activate for a user.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAuthenticationHardwareOathMethodsItemActivatePostRequestBodyable, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
