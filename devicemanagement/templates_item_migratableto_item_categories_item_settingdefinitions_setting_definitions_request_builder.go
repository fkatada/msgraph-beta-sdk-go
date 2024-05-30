package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementSettingCategory entity.
type TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderGetQueryParameters the setting definitions this category contains
type TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderGetQueryParameters struct {
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
// TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderGetQueryParameters
}
// TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementSettingDefinitionId provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementSettingCategory entity.
// returns a *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder when successful
func (m *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) ByDeviceManagementSettingDefinitionId(deviceManagementSettingDefinitionId string)(*TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementSettingDefinitionId != "" {
        urlTplParams["deviceManagementSettingDefinition%2Did"] = deviceManagementSettingDefinitionId
    }
    return NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderInternal instantiates a new TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder and sets the default values.
func NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) {
    m := &TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}/categories/{deviceManagementTemplateSettingCategory%2Did}/settingDefinitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder instantiates a new TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder and sets the default values.
func NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsCountRequestBuilder when successful
func (m *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) Count()(*TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsCountRequestBuilder) {
    return NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the setting definitions this category contains
// returns a DeviceManagementSettingDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementSettingDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionCollectionResponseable), nil
}
// Post create new navigation property to settingDefinitions for deviceManagement
// returns a DeviceManagementSettingDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, requestConfiguration *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementSettingDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable), nil
}
// ToGetRequestInformation the setting definitions this category contains
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to settingDefinitions for deviceManagement
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, requestConfiguration *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder when successful
func (m *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) WithUrl(rawUrl string)(*TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) {
    return NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
