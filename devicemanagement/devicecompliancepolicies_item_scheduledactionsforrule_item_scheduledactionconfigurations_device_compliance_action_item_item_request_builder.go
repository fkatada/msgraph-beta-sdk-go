package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceComplianceScheduledActionForRule entity.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderGetQueryParameters the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderGetQueryParameters
}
// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderInternal instantiates a new DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) {
    m := &DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule%2Did}/scheduledActionConfigurations/{deviceComplianceActionItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder instantiates a new DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property scheduledActionConfigurations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
// returns a DeviceComplianceActionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable), nil
}
// Patch update the navigation property scheduledActionConfigurations in deviceManagement
// returns a DeviceComplianceActionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable), nil
}
// ToDeleteRequestInformation delete navigation property scheduledActionConfigurations for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property scheduledActionConfigurations in deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) {
    return NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
