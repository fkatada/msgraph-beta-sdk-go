package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
type ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters device configuration states for this device.
type ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetQueryParameters
}
// ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal instantiates a new ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder and sets the default values.
func NewComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    m := &ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/deviceConfigurationStates/{deviceConfigurationState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder instantiates a new ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder and sets the default values.
func NewComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceConfigurationStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get device configuration states for this device.
// returns a DeviceConfigurationStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationStateable), nil
}
// Patch update the navigation property deviceConfigurationStates in deviceManagement
// returns a DeviceConfigurationStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationStateable, requestConfiguration *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationStateable), nil
}
// ToDeleteRequestInformation delete navigation property deviceConfigurationStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation device configuration states for this device.
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceConfigurationStates in deviceManagement
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationStateable, requestConfiguration *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder when successful
func (m *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder) {
    return NewComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
