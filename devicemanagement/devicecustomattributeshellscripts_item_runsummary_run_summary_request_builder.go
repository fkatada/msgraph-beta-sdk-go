package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder provides operations to manage the runSummary property of the microsoft.graph.deviceCustomAttributeShellScript entity.
type DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderGetQueryParameters run summary for device management script.
type DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderGetQueryParameters
}
// NewDevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderInternal instantiates a new DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder) {
    m := &DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}/runSummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder instantiates a new DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get run summary for device management script.
// returns a DeviceManagementScriptRunSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptRunSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptRunSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptRunSummaryable), nil
}
// ToGetRequestInformation run summary for device management script.
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder) WithUrl(rawUrl string)(*DevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder) {
    return NewDevicecustomattributeshellscriptsItemRunsummaryRunSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
