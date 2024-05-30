package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderGetQueryParameters device compliance policy states for this device.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderGetQueryParameters struct {
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
// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderGetQueryParameters
}
// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceCompliancePolicyStateId provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) ByDeviceCompliancePolicyStateId(deviceCompliancePolicyStateId string)(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceCompliancePolicyStateId != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = deviceCompliancePolicyStateId
    }
    return NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderInternal instantiates a new ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder and sets the default values.
func NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) {
    m := &ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceCompliancePolicyStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder instantiates a new ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder and sets the default values.
func NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) Count()(*ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder) {
    return NewManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get device compliance policy states for this device.
// returns a DeviceCompliancePolicyStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateCollectionResponseable), nil
}
// Post create new navigation property to deviceCompliancePolicyStates for deviceManagement
// returns a DeviceCompliancePolicyStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable), nil
}
// ToGetRequestInformation device compliance policy states for this device.
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceCompliancePolicyStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) {
    return NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
