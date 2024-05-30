package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder provides operations to manage the frontLineServicePlans property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderGetQueryParameters get a list of the cloudPcFrontLineServicePlan objects and their properties.
type VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderGetQueryParameters struct {
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
// VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderGetQueryParameters
}
// VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudPcFrontLineServicePlanId provides operations to manage the frontLineServicePlans property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointFrontlineserviceplansCloudPcFrontLineServicePlanItemRequestBuilder when successful
func (m *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) ByCloudPcFrontLineServicePlanId(cloudPcFrontLineServicePlanId string)(*VirtualendpointFrontlineserviceplansCloudPcFrontLineServicePlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudPcFrontLineServicePlanId != "" {
        urlTplParams["cloudPcFrontLineServicePlan%2Did"] = cloudPcFrontLineServicePlanId
    }
    return NewVirtualendpointFrontlineserviceplansCloudPcFrontLineServicePlanItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderInternal instantiates a new VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder and sets the default values.
func NewVirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) {
    m := &VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/frontLineServicePlans{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder instantiates a new VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder and sets the default values.
func NewVirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualendpointFrontlineserviceplansCountRequestBuilder when successful
func (m *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) Count()(*VirtualendpointFrontlineserviceplansCountRequestBuilder) {
    return NewVirtualendpointFrontlineserviceplansCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the cloudPcFrontLineServicePlan objects and their properties.
// returns a CloudPcFrontLineServicePlanCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-list-frontlineserviceplans?view=graph-rest-beta
func (m *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcFrontLineServicePlanCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanCollectionResponseable), nil
}
// Post create new navigation property to frontLineServicePlans for deviceManagement
// returns a CloudPcFrontLineServicePlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, requestConfiguration *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcFrontLineServicePlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable), nil
}
// ToGetRequestInformation get a list of the cloudPcFrontLineServicePlan objects and their properties.
// returns a *RequestInformation when successful
func (m *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to frontLineServicePlans for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, requestConfiguration *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder when successful
func (m *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) {
    return NewVirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
