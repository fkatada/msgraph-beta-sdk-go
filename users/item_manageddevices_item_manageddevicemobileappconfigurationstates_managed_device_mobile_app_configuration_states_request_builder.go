package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderGetQueryParameters managed device mobile app configuration states for this device.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderGetQueryParameters struct {
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
// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderGetQueryParameters
}
// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedDeviceMobileAppConfigurationStateId provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) ByManagedDeviceMobileAppConfigurationStateId(managedDeviceMobileAppConfigurationStateId string)(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceMobileAppConfigurationStateId != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = managedDeviceMobileAppConfigurationStateId
    }
    return NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal instantiates a new ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder and sets the default values.
func NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    m := &ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/managedDeviceMobileAppConfigurationStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder instantiates a new ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder and sets the default values.
func NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemManageddevicesItemManageddevicemobileappconfigurationstatesCountRequestBuilder when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) Count()(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesCountRequestBuilder) {
    return NewItemManageddevicesItemManageddevicemobileappconfigurationstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get managed device mobile app configuration states for this device.
// returns a ManagedDeviceMobileAppConfigurationStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateCollectionResponseable), nil
}
// Post create new navigation property to managedDeviceMobileAppConfigurationStates for users
// returns a ManagedDeviceMobileAppConfigurationStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable), nil
}
// ToGetRequestInformation managed device mobile app configuration states for this device.
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managedDeviceMobileAppConfigurationStates for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
