package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RiskyserviceprincipalsDismissRequestBuilder provides operations to call the dismiss method.
type RiskyserviceprincipalsDismissRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RiskyserviceprincipalsDismissRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyserviceprincipalsDismissRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyserviceprincipalsDismissRequestBuilderInternal instantiates a new RiskyserviceprincipalsDismissRequestBuilder and sets the default values.
func NewRiskyserviceprincipalsDismissRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyserviceprincipalsDismissRequestBuilder) {
    m := &RiskyserviceprincipalsDismissRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/riskyServicePrincipals/dismiss", pathParameters),
    }
    return m
}
// NewRiskyserviceprincipalsDismissRequestBuilder instantiates a new RiskyserviceprincipalsDismissRequestBuilder and sets the default values.
func NewRiskyserviceprincipalsDismissRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyserviceprincipalsDismissRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyserviceprincipalsDismissRequestBuilderInternal(urlParams, requestAdapter)
}
// Post dismiss the risk of one or more riskyServicePrincipal objects. This action sets the targeted service principal account's risk level to none. You can dismiss up to 60 service principal accounts in one request.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/riskyserviceprincipal-dismiss?view=graph-rest-beta
func (m *RiskyserviceprincipalsDismissRequestBuilder) Post(ctx context.Context, body RiskyserviceprincipalsDismissPostRequestBodyable, requestConfiguration *RiskyserviceprincipalsDismissRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation dismiss the risk of one or more riskyServicePrincipal objects. This action sets the targeted service principal account's risk level to none. You can dismiss up to 60 service principal accounts in one request.
// returns a *RequestInformation when successful
func (m *RiskyserviceprincipalsDismissRequestBuilder) ToPostRequestInformation(ctx context.Context, body RiskyserviceprincipalsDismissPostRequestBodyable, requestConfiguration *RiskyserviceprincipalsDismissRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RiskyserviceprincipalsDismissRequestBuilder when successful
func (m *RiskyserviceprincipalsDismissRequestBuilder) WithUrl(rawUrl string)(*RiskyserviceprincipalsDismissRequestBuilder) {
    return NewRiskyserviceprincipalsDismissRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
