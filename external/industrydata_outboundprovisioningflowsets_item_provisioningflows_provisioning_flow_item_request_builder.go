package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder provides operations to manage the provisioningFlows property of the microsoft.graph.industryData.outboundProvisioningFlowSet entity.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderGetQueryParameters a flow that provisions relevant records of a given entity type in the Microsoft 365 tenant.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderGetQueryParameters
}
// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderInternal instantiates a new IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) {
    m := &IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/outboundProvisioningFlowSets/{outboundProvisioningFlowSet%2Did}/provisioningFlows/{provisioningFlow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder instantiates a new IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property provisioningFlows for external
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get a flow that provisions relevant records of a given entity type in the Microsoft 365 tenant.
// returns a ProvisioningFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateProvisioningFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable), nil
}
// MicrosoftGraphIndustryDataReset provides operations to call the reset method.
// returns a *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) MicrosoftGraphIndustryDataReset()(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property provisioningFlows in external
// returns a ProvisioningFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) Patch(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderPatchRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateProvisioningFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable), nil
}
// ToDeleteRequestInformation delete navigation property provisioningFlows for external
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a flow that provisions relevant records of a given entity type in the Microsoft 365 tenant.
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property provisioningFlows in external
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) WithUrl(rawUrl string)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
