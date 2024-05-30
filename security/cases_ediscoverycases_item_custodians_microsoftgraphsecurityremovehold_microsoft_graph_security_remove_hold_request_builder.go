package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder provides operations to call the removeHold method.
type CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/microsoft.graph.security.removeHold", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action removeHold
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-removehold?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) Post(ctx context.Context, body CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdRemoveHoldPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action removeHold
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdRemoveHoldPostRequestBodyable, requestConfiguration *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
