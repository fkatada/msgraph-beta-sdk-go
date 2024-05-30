package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder provides operations to manage the userRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
type DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderGetQueryParameters list of run states for this script across all users.
type DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderGetQueryParameters
}
// DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderInternal instantiates a new DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) {
    m := &DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}/userRunStates/{deviceManagementScriptUserState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder instantiates a new DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userRunStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScriptUserState entity.
// returns a *DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) DeviceRunStates()(*DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    return NewDevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of run states for this script across all users.
// returns a DeviceManagementScriptUserStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptUserStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable), nil
}
// Patch update the navigation property userRunStates in deviceManagement
// returns a DeviceManagementScriptUserStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptUserStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable), nil
}
// ToDeleteRequestInformation delete navigation property userRunStates for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of run states for this script across all users.
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userRunStates in deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) WithUrl(rawUrl string)(*DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) {
    return NewDevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
