package monitoring

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
)

// MonitoringRequestBuilder provides operations to manage the monitoring singleton.
type MonitoringRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MonitoringRequestBuilderGetQueryParameters get monitoring
type MonitoringRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MonitoringRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MonitoringRequestBuilderGetQueryParameters
}
// MonitoringRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertRecords provides operations to manage the alertRecords property of the microsoft.graph.deviceManagement.monitoring entity.
// returns a *AlertrecordsAlertRecordsRequestBuilder when successful
func (m *MonitoringRequestBuilder) AlertRecords()(*AlertrecordsAlertRecordsRequestBuilder) {
    return NewAlertrecordsAlertRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AlertRules provides operations to manage the alertRules property of the microsoft.graph.deviceManagement.monitoring entity.
// returns a *AlertrulesAlertRulesRequestBuilder when successful
func (m *MonitoringRequestBuilder) AlertRules()(*AlertrulesAlertRulesRequestBuilder) {
    return NewAlertrulesAlertRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMonitoringRequestBuilderInternal instantiates a new MonitoringRequestBuilder and sets the default values.
func NewMonitoringRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringRequestBuilder) {
    m := &MonitoringRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/monitoring{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMonitoringRequestBuilder instantiates a new MonitoringRequestBuilder and sets the default values.
func NewMonitoringRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonitoringRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get monitoring
// returns a Monitoringable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MonitoringRequestBuilder) Get(ctx context.Context, requestConfiguration *MonitoringRequestBuilderGetRequestConfiguration)(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.CreateMonitoringFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable), nil
}
// Patch update monitoring
// returns a Monitoringable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MonitoringRequestBuilder) Patch(ctx context.Context, body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, requestConfiguration *MonitoringRequestBuilderPatchRequestConfiguration)(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.CreateMonitoringFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable), nil
}
// ToGetRequestInformation get monitoring
// returns a *RequestInformation when successful
func (m *MonitoringRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MonitoringRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update monitoring
// returns a *RequestInformation when successful
func (m *MonitoringRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.Monitoringable, requestConfiguration *MonitoringRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MonitoringRequestBuilder when successful
func (m *MonitoringRequestBuilder) WithUrl(rawUrl string)(*MonitoringRequestBuilder) {
    return NewMonitoringRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
