package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder provides operations to call the confirmCompromised method.
type RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyusersConfirmcompromisedConfirmCompromisedRequestBuilderInternal instantiates a new RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder and sets the default values.
func NewRiskyusersConfirmcompromisedConfirmCompromisedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder) {
    m := &RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/riskyUsers/confirmCompromised", pathParameters),
    }
    return m
}
// NewRiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder instantiates a new RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder and sets the default values.
func NewRiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyusersConfirmcompromisedConfirmCompromisedRequestBuilderInternal(urlParams, requestAdapter)
}
// Post confirm one or more riskyUser objects as compromised. This action sets the targeted user's risk level to high.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/riskyusers-confirmcompromised?view=graph-rest-beta
func (m *RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder) Post(ctx context.Context, body RiskyusersConfirmcompromisedConfirmCompromisedPostRequestBodyable, requestConfiguration *RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation confirm one or more riskyUser objects as compromised. This action sets the targeted user's risk level to high.
// returns a *RequestInformation when successful
func (m *RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder) ToPostRequestInformation(ctx context.Context, body RiskyusersConfirmcompromisedConfirmCompromisedPostRequestBodyable, requestConfiguration *RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder when successful
func (m *RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder) WithUrl(rawUrl string)(*RiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder) {
    return NewRiskyusersConfirmcompromisedConfirmCompromisedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
