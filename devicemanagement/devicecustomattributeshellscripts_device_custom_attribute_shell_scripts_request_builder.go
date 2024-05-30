package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
type DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderGetQueryParameters the list of device custom attribute shell scripts associated with the tenant.
type DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderGetQueryParameters struct {
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
// DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderGetQueryParameters
}
// DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceCustomAttributeShellScriptId provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptItemRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) ByDeviceCustomAttributeShellScriptId(deviceCustomAttributeShellScriptId string)(*DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceCustomAttributeShellScriptId != "" {
        urlTplParams["deviceCustomAttributeShellScript%2Did"] = deviceCustomAttributeShellScriptId
    }
    return NewDevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderInternal instantiates a new DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) {
    m := &DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder instantiates a new DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicecustomattributeshellscriptsCountRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) Count()(*DevicecustomattributeshellscriptsCountRequestBuilder) {
    return NewDevicecustomattributeshellscriptsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device custom attribute shell scripts associated with the tenant.
// returns a DeviceCustomAttributeShellScriptCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCustomAttributeShellScriptCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptCollectionResponseable), nil
}
// Post create new navigation property to deviceCustomAttributeShellScripts for deviceManagement
// returns a DeviceCustomAttributeShellScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable), nil
}
// ToGetRequestInformation the list of device custom attribute shell scripts associated with the tenant.
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceCustomAttributeShellScripts for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) WithUrl(rawUrl string)(*DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) {
    return NewDevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
