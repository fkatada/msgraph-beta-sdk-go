package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder provides operations to call the findByMethodMode method.
type ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
type ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters
}
// NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal instantiates a new ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, authenticationMethodModes *string)(*ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    m := &ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrengths/policies/findByMethodMode(authenticationMethodModes={authenticationMethodModes}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if authenticationMethodModes != nil {
        m.BaseRequestBuilder.PathParameters["authenticationMethodModes"] = *authenticationMethodModes
    }
    return m
}
// NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder instantiates a new ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: This method is obsolete. Use GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse instead.
// returns a ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-findbymethodmode?view=graph-rest-beta
func (m *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable), nil
}
// GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
// returns a ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-findbymethodmode?view=graph-rest-beta
func (m *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable), nil
}
// ToGetRequestInformation get a list of the authenticationStrengthPolicy objects and their properties filtered to only include policies that include the authentication method mode specified in the request.
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
// returns a *RequestInformation when successful
func (m *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The &apos;authenticationStrengths&apos; segment is deprecated. Please use &apos;authenticationStrength&apos; instead. as of 2023-02/AuthenticationStrengthsRemove on 2023-02-01 and will be removed 2023-03-31
// returns a *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder when successful
func (m *ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) WithUrl(rawUrl string)(*ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
