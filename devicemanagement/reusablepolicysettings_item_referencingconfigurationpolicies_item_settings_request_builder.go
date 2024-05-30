package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderGetQueryParameters policy settings
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderGetQueryParameters struct {
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
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderGetQueryParameters
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementConfigurationSettingId provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) ByDeviceManagementConfigurationSettingId(deviceManagementConfigurationSettingId string)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementConfigurationSettingId != "" {
        urlTplParams["deviceManagementConfigurationSetting%2Did"] = deviceManagementConfigurationSettingId
    }
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderInternal instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) {
    m := &ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/settings{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsCountRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) Count()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsCountRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get policy settings
// returns a DeviceManagementConfigurationSettingCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingCollectionResponseable), nil
}
// Post create new navigation property to settings for deviceManagement
// returns a DeviceManagementConfigurationSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable), nil
}
// ToGetRequestInformation policy settings
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to settings for deviceManagement
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) WithUrl(rawUrl string)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
