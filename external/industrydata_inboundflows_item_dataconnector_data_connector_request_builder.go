package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder provides operations to manage the dataConnector property of the microsoft.graph.industryData.inboundFlow entity.
type IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderGetQueryParameters the data connector in the context of which this flow pulls in data from a source system.
type IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderGetQueryParameters
}
// NewIndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderInternal instantiates a new IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder and sets the default values.
func NewIndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder) {
    m := &IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/inboundFlows/{inboundFlow%2Did}/dataConnector{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder instantiates a new IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder and sets the default values.
func NewIndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the data connector in the context of which this flow pulls in data from a source system.
// returns a IndustryDataConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataConnectorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataConnectorable), nil
}
// ToGetRequestInformation the data connector in the context of which this flow pulls in data from a source system.
// returns a *RequestInformation when successful
func (m *IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder when successful
func (m *IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder) WithUrl(rawUrl string)(*IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder) {
    return NewIndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
