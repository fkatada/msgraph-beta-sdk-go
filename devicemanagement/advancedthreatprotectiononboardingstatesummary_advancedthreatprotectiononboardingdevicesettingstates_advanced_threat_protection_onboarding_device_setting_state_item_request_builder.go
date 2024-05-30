package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder provides operations to manage the advancedThreatProtectionOnboardingDeviceSettingStates property of the microsoft.graph.advancedThreatProtectionOnboardingStateSummary entity.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetQueryParameters get advancedThreatProtectionOnboardingDeviceSettingStates from deviceManagement
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetQueryParameters
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal instantiates a new AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder and sets the default values.
func NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    m := &AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates/{advancedThreatProtectionOnboardingDeviceSettingState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder instantiates a new AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder and sets the default values.
func NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property advancedThreatProtectionOnboardingDeviceSettingStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get advancedThreatProtectionOnboardingDeviceSettingStates from deviceManagement
// returns a AdvancedThreatProtectionOnboardingDeviceSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable), nil
}
// Patch update the navigation property advancedThreatProtectionOnboardingDeviceSettingStates in deviceManagement
// returns a AdvancedThreatProtectionOnboardingDeviceSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable), nil
}
// ToDeleteRequestInformation delete navigation property advancedThreatProtectionOnboardingDeviceSettingStates for deviceManagement
// returns a *RequestInformation when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get advancedThreatProtectionOnboardingDeviceSettingStates from deviceManagement
// returns a *RequestInformation when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property advancedThreatProtectionOnboardingDeviceSettingStates in deviceManagement
// returns a *RequestInformation when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) WithUrl(rawUrl string)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
