package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder provides operations to call the invalidateAllRefreshTokens method.
type ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderInternal instantiates a new ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder) {
    m := &ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/invalidateAllRefreshTokens", pathParameters),
    }
    return m
}
// NewItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder instantiates a new ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invalidates all of the user's refresh tokens issued to applications (as well as session cookies in a user's browser), by resetting the refreshTokensValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device.  This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application will get an error. If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint, which will force the user to sign in.
// Deprecated: This method is obsolete. Use PostAsInvalidateAllRefreshTokensPostResponse instead.
// returns a ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-invalidateallrefreshtokens?view=graph-rest-beta
func (m *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemInvalidateallrefreshtokensInvalidateAllRefreshTokensResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensResponseable), nil
}
// PostAsInvalidateAllRefreshTokensPostResponse invalidates all of the user's refresh tokens issued to applications (as well as session cookies in a user's browser), by resetting the refreshTokensValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device.  This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application will get an error. If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint, which will force the user to sign in.
// returns a ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-invalidateallrefreshtokens?view=graph-rest-beta
func (m *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder) PostAsInvalidateAllRefreshTokensPostResponse(ctx context.Context, requestConfiguration *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemInvalidateallrefreshtokensInvalidateAllRefreshTokensPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensPostResponseable), nil
}
// ToPostRequestInformation invalidates all of the user's refresh tokens issued to applications (as well as session cookies in a user's browser), by resetting the refreshTokensValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device.  This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application will get an error. If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint, which will force the user to sign in.
// returns a *RequestInformation when successful
func (m *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder when successful
func (m *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder) WithUrl(rawUrl string)(*ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder) {
    return NewItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
