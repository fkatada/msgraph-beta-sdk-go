package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder provides operations to call the getRealTimeRemoteConnectionLatency method.
type VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal instantiates a new VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, cloudPcId *string)(*VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    m := &VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getRealTimeRemoteConnectionLatency(cloudPcId='{cloudPcId}')", pathParameters),
    }
    if cloudPcId != nil {
        m.BaseRequestBuilder.PathParameters["cloudPcId"] = *cloudPcId
    }
    return m
}
// NewVirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder instantiates a new VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the real-time connection latency information for a Cloud PC.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getrealtimeremoteconnectionlatency?view=graph-rest-beta
func (m *VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get the real-time connection latency information for a Cloud PC.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder when successful
func (m *VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    return NewVirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
