package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder provides operations to manage the deviceStatuses property of the microsoft.graph.deviceConfiguration entity.
type DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters device configuration installation status by device.
type DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters struct {
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
// DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters
}
// DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceConfigurationDeviceStatusId provides operations to manage the deviceStatuses property of the microsoft.graph.deviceConfiguration entity.
// returns a *DeviceconfigurationsItemDevicestatusesDeviceConfigurationDeviceStatusItemRequestBuilder when successful
func (m *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) ByDeviceConfigurationDeviceStatusId(deviceConfigurationDeviceStatusId string)(*DeviceconfigurationsItemDevicestatusesDeviceConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceConfigurationDeviceStatusId != "" {
        urlTplParams["deviceConfigurationDeviceStatus%2Did"] = deviceConfigurationDeviceStatusId
    }
    return NewDeviceconfigurationsItemDevicestatusesDeviceConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderInternal instantiates a new DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) {
    m := &DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/deviceStatuses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder instantiates a new DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DeviceconfigurationsItemDevicestatusesCountRequestBuilder when successful
func (m *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) Count()(*DeviceconfigurationsItemDevicestatusesCountRequestBuilder) {
    return NewDeviceconfigurationsItemDevicestatusesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get device configuration installation status by device.
// returns a DeviceConfigurationDeviceStatusCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationDeviceStatusCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationDeviceStatusCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationDeviceStatusCollectionResponseable), nil
}
// Post create new navigation property to deviceStatuses for deviceManagement
// returns a DeviceConfigurationDeviceStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationDeviceStatusable, requestConfiguration *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationDeviceStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationDeviceStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationDeviceStatusable), nil
}
// ToGetRequestInformation device configuration installation status by device.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceStatuses for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationDeviceStatusable, requestConfiguration *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder when successful
func (m *DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) {
    return NewDeviceconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
