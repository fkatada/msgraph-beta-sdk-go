package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder provides operations to call the extractLabel method.
type ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilderInternal instantiates a new ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder) {
    m := &ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/policy/labels/extractLabel", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder instantiates a new ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post using the metadata that exists on an already-labeled piece of information, resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a InformationProtectionContentLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-extractlabel?view=graph-rest-beta
func (m *ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder) Post(ctx context.Context, body ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelPostRequestBodyable, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionContentLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionContentLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionContentLabelable), nil
}
// ToPostRequestInformation using the metadata that exists on an already-labeled piece of information, resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelPostRequestBodyable, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
