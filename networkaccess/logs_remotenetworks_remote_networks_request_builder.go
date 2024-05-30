package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// LogsRemotenetworksRemoteNetworksRequestBuilder provides operations to manage the remoteNetworks property of the microsoft.graph.networkaccess.logs entity.
type LogsRemotenetworksRemoteNetworksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LogsRemotenetworksRemoteNetworksRequestBuilderGetQueryParameters retrieve a list of remote network health status microsoft.graph.networkaccess.remoteNetworkHealthStatusEvent events, providing insights into the health and status of remote networks.
type LogsRemotenetworksRemoteNetworksRequestBuilderGetQueryParameters struct {
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
// LogsRemotenetworksRemoteNetworksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LogsRemotenetworksRemoteNetworksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LogsRemotenetworksRemoteNetworksRequestBuilderGetQueryParameters
}
// LogsRemotenetworksRemoteNetworksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LogsRemotenetworksRemoteNetworksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRemoteNetworkHealthEventId provides operations to manage the remoteNetworks property of the microsoft.graph.networkaccess.logs entity.
// returns a *LogsRemotenetworksRemoteNetworkHealthEventItemRequestBuilder when successful
func (m *LogsRemotenetworksRemoteNetworksRequestBuilder) ByRemoteNetworkHealthEventId(remoteNetworkHealthEventId string)(*LogsRemotenetworksRemoteNetworkHealthEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if remoteNetworkHealthEventId != "" {
        urlTplParams["remoteNetworkHealthEvent%2Did"] = remoteNetworkHealthEventId
    }
    return NewLogsRemotenetworksRemoteNetworkHealthEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLogsRemotenetworksRemoteNetworksRequestBuilderInternal instantiates a new LogsRemotenetworksRemoteNetworksRequestBuilder and sets the default values.
func NewLogsRemotenetworksRemoteNetworksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogsRemotenetworksRemoteNetworksRequestBuilder) {
    m := &LogsRemotenetworksRemoteNetworksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/logs/remoteNetworks{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLogsRemotenetworksRemoteNetworksRequestBuilder instantiates a new LogsRemotenetworksRemoteNetworksRequestBuilder and sets the default values.
func NewLogsRemotenetworksRemoteNetworksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogsRemotenetworksRemoteNetworksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLogsRemotenetworksRemoteNetworksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LogsRemotenetworksCountRequestBuilder when successful
func (m *LogsRemotenetworksRemoteNetworksRequestBuilder) Count()(*LogsRemotenetworksCountRequestBuilder) {
    return NewLogsRemotenetworksCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of remote network health status microsoft.graph.networkaccess.remoteNetworkHealthStatusEvent events, providing insights into the health and status of remote networks.
// returns a RemoteNetworkHealthEventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/networkaccess-remotenetworkhealthstatusevent-list?view=graph-rest-beta
func (m *LogsRemotenetworksRemoteNetworksRequestBuilder) Get(ctx context.Context, requestConfiguration *LogsRemotenetworksRemoteNetworksRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkHealthEventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateRemoteNetworkHealthEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkHealthEventCollectionResponseable), nil
}
// Post create new navigation property to remoteNetworks for networkAccess
// returns a RemoteNetworkHealthEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LogsRemotenetworksRemoteNetworksRequestBuilder) Post(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkHealthEventable, requestConfiguration *LogsRemotenetworksRemoteNetworksRequestBuilderPostRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkHealthEventable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateRemoteNetworkHealthEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkHealthEventable), nil
}
// ToGetRequestInformation retrieve a list of remote network health status microsoft.graph.networkaccess.remoteNetworkHealthStatusEvent events, providing insights into the health and status of remote networks.
// returns a *RequestInformation when successful
func (m *LogsRemotenetworksRemoteNetworksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LogsRemotenetworksRemoteNetworksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to remoteNetworks for networkAccess
// returns a *RequestInformation when successful
func (m *LogsRemotenetworksRemoteNetworksRequestBuilder) ToPostRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkHealthEventable, requestConfiguration *LogsRemotenetworksRemoteNetworksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LogsRemotenetworksRemoteNetworksRequestBuilder when successful
func (m *LogsRemotenetworksRemoteNetworksRequestBuilder) WithUrl(rawUrl string)(*LogsRemotenetworksRemoteNetworksRequestBuilder) {
    return NewLogsRemotenetworksRemoteNetworksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
