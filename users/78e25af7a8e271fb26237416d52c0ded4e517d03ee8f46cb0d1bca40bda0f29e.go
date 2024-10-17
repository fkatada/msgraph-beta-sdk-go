package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal instantiates a new ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    m := &ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateApplication", pathParameters),
    }
    return m
}
// NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder instantiates a new ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// Deprecated: This method is obsolete. Use PostAsEvaluateApplicationPostResponse instead.
// returns a ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateapplication?view=graph-rest-beta
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationResponseable), nil
}
// PostAsEvaluateApplicationPostResponse compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// returns a ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateapplication?view=graph-rest-beta
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) PostAsEvaluateApplicationPostResponse(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostResponseable), nil
}
// ToPostRequestInformation compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// returns a *RequestInformation when successful
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder when successful
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
