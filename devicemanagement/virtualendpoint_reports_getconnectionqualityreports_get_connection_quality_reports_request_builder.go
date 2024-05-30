package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder provides operations to call the getConnectionQualityReports method.
type VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilderInternal instantiates a new VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder) {
    m := &VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getConnectionQualityReports", pathParameters),
    }
    return m
}
// NewVirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder instantiates a new VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the overall connection quality reports for all devices within a current tenant during a given time period, including metrics like the average round trip time (P50), average available bandwidth, and UDP connection percentage. Get also other real-time metrics such as last connection round trip time, last connection client IP, last connection gateway, and last connection protocol.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getconnectionqualityreports?view=graph-rest-beta
func (m *VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder) Post(ctx context.Context, body VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation get the overall connection quality reports for all devices within a current tenant during a given time period, including metrics like the average round trip time (P50), average available bandwidth, and UDP connection percentage. Get also other real-time metrics such as last connection round trip time, last connection client IP, last connection gateway, and last connection protocol.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder when successful
func (m *VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder) {
    return NewVirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
