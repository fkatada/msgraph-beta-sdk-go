package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSetting entity.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters list of related Setting Definitions. This property is read-only.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    m := &ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/settings/{deviceManagementConfigurationSetting%2Did}/settingDefinitions/{deviceManagementConfigurationSettingDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list of related Setting Definitions. This property is read-only.
// returns a DeviceManagementConfigurationSettingDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable), nil
}
// ToGetRequestInformation list of related Setting Definitions. This property is read-only.
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsItemSettingdefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
