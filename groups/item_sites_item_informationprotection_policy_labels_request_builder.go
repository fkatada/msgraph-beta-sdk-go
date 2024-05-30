package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder provides operations to manage the labels property of the microsoft.graph.informationProtectionPolicy entity.
type ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderGetQueryParameters get labels from groups
type ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByInformationProtectionLabelId provides operations to manage the labels property of the microsoft.graph.informationProtectionPolicy entity.
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *ItemSitesItemInformationprotectionPolicyLabelsInformationProtectionLabelItemRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) ByInformationProtectionLabelId(informationProtectionLabelId string)(*ItemSitesItemInformationprotectionPolicyLabelsInformationProtectionLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if informationProtectionLabelId != "" {
        urlTplParams["informationProtectionLabel%2Did"] = informationProtectionLabelId
    }
    return NewItemSitesItemInformationprotectionPolicyLabelsInformationProtectionLabelItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemInformationprotectionPolicyLabelsRequestBuilderInternal instantiates a new ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionPolicyLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) {
    m := &ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/policy/labels{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationprotectionPolicyLabelsRequestBuilder instantiates a new ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionPolicyLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationprotectionPolicyLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemInformationprotectionPolicyLabelsCountRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) Count()(*ItemSitesItemInformationprotectionPolicyLabelsCountRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyLabelsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EvaluateApplication provides operations to call the evaluateApplication method.
// returns a *ItemSitesItemInformationprotectionPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) EvaluateApplication()(*ItemSitesItemInformationprotectionPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EvaluateClassificationResults provides operations to call the evaluateClassificationResults method.
// returns a *ItemSitesItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) EvaluateClassificationResults()(*ItemSitesItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EvaluateRemoval provides operations to call the evaluateRemoval method.
// returns a *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) EvaluateRemoval()(*ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExtractLabel provides operations to call the extractLabel method.
// returns a *ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) ExtractLabel()(*ItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyLabelsExtractlabelExtractLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get labels from groups
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a InformationProtectionLabelCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionLabelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelCollectionResponseable), nil
}
// Post create new navigation property to labels for groups
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a InformationProtectionLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelable, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelable), nil
}
// ToGetRequestInformation get labels from groups
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to labels for groups
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionLabelable, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationprotectionPolicyLabelsRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
