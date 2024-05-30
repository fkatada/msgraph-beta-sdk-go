package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder provides operations to manage the deviceRunStates property of the microsoft.graph.hardwareConfiguration entity.
type HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters list of run states for the hardware configuration across all devices. Read-Only.
type HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters struct {
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
// HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters
}
// HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHardwareConfigurationDeviceStateId provides operations to manage the deviceRunStates property of the microsoft.graph.hardwareConfiguration entity.
// returns a *HardwareconfigurationsItemDevicerunstatesHardwareConfigurationDeviceStateItemRequestBuilder when successful
func (m *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) ByHardwareConfigurationDeviceStateId(hardwareConfigurationDeviceStateId string)(*HardwareconfigurationsItemDevicerunstatesHardwareConfigurationDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hardwareConfigurationDeviceStateId != "" {
        urlTplParams["hardwareConfigurationDeviceState%2Did"] = hardwareConfigurationDeviceStateId
    }
    return NewHardwareconfigurationsItemDevicerunstatesHardwareConfigurationDeviceStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewHardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal instantiates a new HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder and sets the default values.
func NewHardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    m := &HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/hardwareConfigurations/{hardwareConfiguration%2Did}/deviceRunStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewHardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder instantiates a new HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder and sets the default values.
func NewHardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *HardwareconfigurationsItemDevicerunstatesCountRequestBuilder when successful
func (m *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) Count()(*HardwareconfigurationsItemDevicerunstatesCountRequestBuilder) {
    return NewHardwareconfigurationsItemDevicerunstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of run states for the hardware configuration across all devices. Read-Only.
// returns a HardwareConfigurationDeviceStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationDeviceStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationDeviceStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationDeviceStateCollectionResponseable), nil
}
// Post create new navigation property to deviceRunStates for deviceManagement
// returns a HardwareConfigurationDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationDeviceStateable, requestConfiguration *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationDeviceStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationDeviceStateable), nil
}
// ToGetRequestInformation list of run states for the hardware configuration across all devices. Read-Only.
// returns a *RequestInformation when successful
func (m *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceRunStates for deviceManagement
// returns a *RequestInformation when successful
func (m *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationDeviceStateable, requestConfiguration *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder when successful
func (m *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) WithUrl(rawUrl string)(*HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    return NewHardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
