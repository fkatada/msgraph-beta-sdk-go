package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder provides operations to manage the deviceAppManagementTasks property of the microsoft.graph.deviceAppManagement entity.
type DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderGetQueryParameters device app management tasks.
type DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderGetQueryParameters
}
// DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderInternal instantiates a new DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder and sets the default values.
func NewDeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) {
    m := &DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/deviceAppManagementTasks/{deviceAppManagementTask%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder instantiates a new DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder and sets the default values.
func NewDeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceAppManagementTasks for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get device app management tasks.
// returns a DeviceAppManagementTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAppManagementTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementTaskable), nil
}
// Patch update the navigation property deviceAppManagementTasks in deviceAppManagement
// returns a DeviceAppManagementTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementTaskable, requestConfiguration *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementTaskable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAppManagementTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementTaskable), nil
}
// ToDeleteRequestInformation delete navigation property deviceAppManagementTasks for deviceAppManagement
// returns a *RequestInformation when successful
func (m *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation device app management tasks.
// returns a *RequestInformation when successful
func (m *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceAppManagementTasks in deviceAppManagement
// returns a *RequestInformation when successful
func (m *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementTaskable, requestConfiguration *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateStatus provides operations to call the updateStatus method.
// returns a *DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder when successful
func (m *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) UpdateStatus()(*DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder) {
    return NewDeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder when successful
func (m *DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) WithUrl(rawUrl string)(*DeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder) {
    return NewDeviceappmanagementtasksDeviceAppManagementTaskItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
