package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
type ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderGetQueryParameters get sensitivityLabels from sites
type ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderGetQueryParameters struct {
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
// ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderGetQueryParameters
}
// ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySensitivityLabelId provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
// returns a *ItemInformationprotectionSensitivitylabelsSensitivityLabelItemRequestBuilder when successful
func (m *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) BySensitivityLabelId(sensitivityLabelId string)(*ItemInformationprotectionSensitivitylabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sensitivityLabelId != "" {
        urlTplParams["sensitivityLabel%2Did"] = sensitivityLabelId
    }
    return NewItemInformationprotectionSensitivitylabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderInternal instantiates a new ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder and sets the default values.
func NewItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) {
    m := &ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/informationProtection/sensitivityLabels{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder instantiates a new ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder and sets the default values.
func NewItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemInformationprotectionSensitivitylabelsCountRequestBuilder when successful
func (m *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) Count()(*ItemInformationprotectionSensitivitylabelsCountRequestBuilder) {
    return NewItemInformationprotectionSensitivitylabelsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Evaluate provides operations to call the evaluate method.
// returns a *ItemInformationprotectionSensitivitylabelsEvaluateRequestBuilder when successful
func (m *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) Evaluate()(*ItemInformationprotectionSensitivitylabelsEvaluateRequestBuilder) {
    return NewItemInformationprotectionSensitivitylabelsEvaluateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get sensitivityLabels from sites
// returns a SensitivityLabelCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityLabelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelCollectionResponseable), nil
}
// Post create new navigation property to sensitivityLabels for sites
// returns a SensitivityLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, requestConfiguration *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable), nil
}
// ToGetRequestInformation get sensitivityLabels from sites
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to sensitivityLabels for sites
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, requestConfiguration *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder when successful
func (m *ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) WithUrl(rawUrl string)(*ItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) {
    return NewItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
