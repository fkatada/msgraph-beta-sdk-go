package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderInternal instantiates a new PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder and sets the default values.
func NewPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) {
    m := &PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/policy/labels/evaluateRemoval", pathParameters),
    }
    return m
}
// NewPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder instantiates a new PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder and sets the default values.
func NewPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// Deprecated: This method is obsolete. Use PostAsEvaluateRemovalPostResponse instead.
// returns a PolicyLabelsEvaluateremovalEvaluateRemovalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateremoval?view=graph-rest-beta
func (m *PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) Post(ctx context.Context, body PolicyLabelsEvaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration)(PolicyLabelsEvaluateremovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicyLabelsEvaluateremovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicyLabelsEvaluateremovalEvaluateRemovalResponseable), nil
}
// PostAsEvaluateRemovalPostResponse indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a PolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateremoval?view=graph-rest-beta
func (m *PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) PostAsEvaluateRemovalPostResponse(ctx context.Context, body PolicyLabelsEvaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration)(PolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicyLabelsEvaluateremovalEvaluateRemovalPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body PolicyLabelsEvaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder when successful
func (m *PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) WithUrl(rawUrl string)(*PolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) {
    return NewPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
