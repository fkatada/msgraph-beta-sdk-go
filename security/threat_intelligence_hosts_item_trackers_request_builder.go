package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligenceHostsItemTrackersRequestBuilder provides operations to manage the trackers property of the microsoft.graph.security.host entity.
type ThreatIntelligenceHostsItemTrackersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceHostsItemTrackersRequestBuilderGetQueryParameters get a list of hostTracker resources.
type ThreatIntelligenceHostsItemTrackersRequestBuilderGetQueryParameters struct {
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
// ThreatIntelligenceHostsItemTrackersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostsItemTrackersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceHostsItemTrackersRequestBuilderGetQueryParameters
}
// ByHostTrackerId provides operations to manage the trackers property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemTrackersHostTrackerItemRequestBuilder when successful
func (m *ThreatIntelligenceHostsItemTrackersRequestBuilder) ByHostTrackerId(hostTrackerId string)(*ThreatIntelligenceHostsItemTrackersHostTrackerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hostTrackerId != "" {
        urlTplParams["hostTracker%2Did"] = hostTrackerId
    }
    return NewThreatIntelligenceHostsItemTrackersHostTrackerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatIntelligenceHostsItemTrackersRequestBuilderInternal instantiates a new ThreatIntelligenceHostsItemTrackersRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemTrackersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemTrackersRequestBuilder) {
    m := &ThreatIntelligenceHostsItemTrackersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/trackers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceHostsItemTrackersRequestBuilder instantiates a new ThreatIntelligenceHostsItemTrackersRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemTrackersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemTrackersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceHostsItemTrackersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatIntelligenceHostsItemTrackersCountRequestBuilder when successful
func (m *ThreatIntelligenceHostsItemTrackersRequestBuilder) Count()(*ThreatIntelligenceHostsItemTrackersCountRequestBuilder) {
    return NewThreatIntelligenceHostsItemTrackersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of hostTracker resources.
// returns a HostTrackerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-host-list-trackers?view=graph-rest-beta
func (m *ThreatIntelligenceHostsItemTrackersRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemTrackersRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostTrackerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateHostTrackerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostTrackerCollectionResponseable), nil
}
// ToGetRequestInformation get a list of hostTracker resources.
// returns a *RequestInformation when successful
func (m *ThreatIntelligenceHostsItemTrackersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemTrackersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatIntelligenceHostsItemTrackersRequestBuilder when successful
func (m *ThreatIntelligenceHostsItemTrackersRequestBuilder) WithUrl(rawUrl string)(*ThreatIntelligenceHostsItemTrackersRequestBuilder) {
    return NewThreatIntelligenceHostsItemTrackersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
