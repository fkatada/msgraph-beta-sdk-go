package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInvalidateAllRefreshTokensRequestBuilder provides operations to call the invalidateAllRefreshTokens method.
type ItemInvalidateAllRefreshTokensRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInvalidateAllRefreshTokensRequestBuilderInternal instantiates a new ItemInvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewItemInvalidateAllRefreshTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvalidateAllRefreshTokensRequestBuilder) {
    m := &ItemInvalidateAllRefreshTokensRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/invalidateAllRefreshTokens", pathParameters),
    }
    return m
}
// NewItemInvalidateAllRefreshTokensRequestBuilder instantiates a new ItemInvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewItemInvalidateAllRefreshTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvalidateAllRefreshTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInvalidateAllRefreshTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invalidates all of the user's refresh tokens issued to applications and session cookies in a user's browser, by resetting the refreshTokensValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device. This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application receives an error. If this happens, the application needs to acquire a new refresh token by making a request to the OAuth 2.0 /authorize endpoint, which forces the user to sign in.
// Deprecated: This method is obsolete. Use PostAsInvalidateAllRefreshTokensPostResponse instead.
// returns a ItemInvalidateAllRefreshTokensResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-invalidateallrefreshtokens?view=graph-rest-beta
func (m *ItemInvalidateAllRefreshTokensRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(ItemInvalidateAllRefreshTokensResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemInvalidateAllRefreshTokensResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInvalidateAllRefreshTokensResponseable), nil
}
// PostAsInvalidateAllRefreshTokensPostResponse invalidates all of the user's refresh tokens issued to applications and session cookies in a user's browser, by resetting the refreshTokensValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device. This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application receives an error. If this happens, the application needs to acquire a new refresh token by making a request to the OAuth 2.0 /authorize endpoint, which forces the user to sign in.
// returns a ItemInvalidateAllRefreshTokensPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-invalidateallrefreshtokens?view=graph-rest-beta
func (m *ItemInvalidateAllRefreshTokensRequestBuilder) PostAsInvalidateAllRefreshTokensPostResponse(ctx context.Context, requestConfiguration *ItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(ItemInvalidateAllRefreshTokensPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemInvalidateAllRefreshTokensPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInvalidateAllRefreshTokensPostResponseable), nil
}
// ToPostRequestInformation invalidates all of the user's refresh tokens issued to applications and session cookies in a user's browser, by resetting the refreshTokensValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device. This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application receives an error. If this happens, the application needs to acquire a new refresh token by making a request to the OAuth 2.0 /authorize endpoint, which forces the user to sign in.
// returns a *RequestInformation when successful
func (m *ItemInvalidateAllRefreshTokensRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInvalidateAllRefreshTokensRequestBuilder when successful
func (m *ItemInvalidateAllRefreshTokensRequestBuilder) WithUrl(rawUrl string)(*ItemInvalidateAllRefreshTokensRequestBuilder) {
    return NewItemInvalidateAllRefreshTokensRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
