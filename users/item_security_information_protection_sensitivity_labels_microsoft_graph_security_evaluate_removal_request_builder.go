package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal instantiates a new MicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    m := &ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateRemoval", pathParameters),
    }
    return m
}
// NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder instantiates a new MicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-1.0
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
