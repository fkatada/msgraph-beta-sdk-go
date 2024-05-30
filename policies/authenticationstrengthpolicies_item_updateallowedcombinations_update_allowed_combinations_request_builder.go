package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder provides operations to call the updateAllowedCombinations method.
type AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderInternal instantiates a new AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder and sets the default values.
func NewAuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) {
    m := &AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationStrengthPolicies/{authenticationStrengthPolicy%2Did}/updateAllowedCombinations", pathParameters),
    }
    return m
}
// NewAuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder instantiates a new AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder and sets the default values.
func NewAuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
// returns a UpdateAllowedCombinationsResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-updateallowedcombinations?view=graph-rest-beta
func (m *AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) Post(ctx context.Context, body AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsPostRequestBodyable, requestConfiguration *AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateAllowedCombinationsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUpdateAllowedCombinationsResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UpdateAllowedCombinationsResultable), nil
}
// ToPostRequestInformation update the allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
// returns a *RequestInformation when successful
func (m *AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsPostRequestBodyable, requestConfiguration *AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder when successful
func (m *AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) {
    return NewAuthenticationstrengthpoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
