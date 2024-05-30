package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder provides operations to call the extractContentLabel method.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilderInternal instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder) {
    m := &ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/microsoft.graph.security.extractContentLabel", pathParameters),
    }
    return m
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post use the metadata that exists on an already-labeled piece of information to resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
// returns a ContentLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-extractcontentlabel?view=graph-rest-beta
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelExtractContentLabelPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateContentLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentLabelable), nil
}
// ToPostRequestInformation use the metadata that exists on an already-labeled piece of information to resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
// returns a *RequestInformation when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelExtractContentLabelPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder) {
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityextractcontentlabelMicrosoftGraphSecurityExtractContentLabelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
