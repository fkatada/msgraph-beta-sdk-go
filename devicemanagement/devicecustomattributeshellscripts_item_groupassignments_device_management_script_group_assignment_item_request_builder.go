package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder provides operations to manage the groupAssignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
type DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderGetQueryParameters the list of group assignments for the device management script.
type DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderGetQueryParameters
}
// DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal instantiates a new DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    m := &DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}/groupAssignments/{deviceManagementScriptGroupAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder instantiates a new DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupAssignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of group assignments for the device management script.
// returns a DeviceManagementScriptGroupAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable), nil
}
// Patch update the navigation property groupAssignments in deviceManagement
// returns a DeviceManagementScriptGroupAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable, requestConfiguration *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property groupAssignments for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of group assignments for the device management script.
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupAssignments in deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable, requestConfiguration *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*DevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    return NewDevicecustomattributeshellscriptsItemGroupassignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
