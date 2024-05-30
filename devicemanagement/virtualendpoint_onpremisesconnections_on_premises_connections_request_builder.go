package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderGetQueryParameters get a list of the cloudPcOnPremisesConnection objects and their properties.
type VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderGetQueryParameters struct {
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
// VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderGetQueryParameters
}
// VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudPcOnPremisesConnectionId provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder when successful
func (m *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) ByCloudPcOnPremisesConnectionId(cloudPcOnPremisesConnectionId string)(*VirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudPcOnPremisesConnectionId != "" {
        urlTplParams["cloudPcOnPremisesConnection%2Did"] = cloudPcOnPremisesConnectionId
    }
    return NewVirtualendpointOnpremisesconnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderInternal instantiates a new VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder and sets the default values.
func NewVirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) {
    m := &VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder instantiates a new VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder and sets the default values.
func NewVirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualendpointOnpremisesconnectionsCountRequestBuilder when successful
func (m *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) Count()(*VirtualendpointOnpremisesconnectionsCountRequestBuilder) {
    return NewVirtualendpointOnpremisesconnectionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the cloudPcOnPremisesConnection objects and their properties.
// returns a CloudPcOnPremisesConnectionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-list-onpremisesconnections?view=graph-rest-beta
func (m *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcOnPremisesConnectionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionCollectionResponseable), nil
}
// Post create a new cloudPcOnPremisesConnection object for provisioning Cloud PCs.
// returns a CloudPcOnPremisesConnectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-post-onpremisesconnections?view=graph-rest-beta
func (m *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, requestConfiguration *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcOnPremisesConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable), nil
}
// ToGetRequestInformation get a list of the cloudPcOnPremisesConnection objects and their properties.
// returns a *RequestInformation when successful
func (m *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new cloudPcOnPremisesConnection object for provisioning Cloud PCs.
// returns a *RequestInformation when successful
func (m *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, requestConfiguration *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder when successful
func (m *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) {
    return NewVirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
