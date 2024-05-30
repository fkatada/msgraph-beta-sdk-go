package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
type ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters collection of resource access settings associated with account.
type ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters
}
// ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *ResourceaccessprofilesItemAssignRequestBuilder when successful
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Assign()(*ResourceaccessprofilesItemAssignRequestBuilder) {
    return NewResourceaccessprofilesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementResourceAccessProfileBase entity.
// returns a *ResourceaccessprofilesItemAssignmentsRequestBuilder when successful
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Assignments()(*ResourceaccessprofilesItemAssignmentsRequestBuilder) {
    return NewResourceaccessprofilesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal instantiates a new ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    m := &ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/resourceAccessProfiles/{deviceManagementResourceAccessProfileBase%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder instantiates a new ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resourceAccessProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of resource access settings associated with account.
// returns a DeviceManagementResourceAccessProfileBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable), nil
}
// Patch update the navigation property resourceAccessProfiles in deviceManagement
// returns a DeviceManagementResourceAccessProfileBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable), nil
}
// ToDeleteRequestInformation delete navigation property resourceAccessProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of resource access settings associated with account.
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resourceAccessProfiles in deviceManagement
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder when successful
func (m *ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) WithUrl(rawUrl string)(*ResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    return NewResourceaccessprofilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
