package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder provides operations to manage the advancedThreatProtectionOnboardingDeviceSettingStates property of the microsoft.graph.advancedThreatProtectionOnboardingStateSummary entity.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters get advancedThreatProtectionOnboardingDeviceSettingStates from deviceManagement
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters struct {
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
// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAdvancedThreatProtectionOnboardingDeviceSettingStateId provides operations to manage the advancedThreatProtectionOnboardingDeviceSettingStates property of the microsoft.graph.advancedThreatProtectionOnboardingStateSummary entity.
// returns a *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) ByAdvancedThreatProtectionOnboardingDeviceSettingStateId(advancedThreatProtectionOnboardingDeviceSettingStateId string)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if advancedThreatProtectionOnboardingDeviceSettingStateId != "" {
        urlTplParams["advancedThreatProtectionOnboardingDeviceSettingState%2Did"] = advancedThreatProtectionOnboardingDeviceSettingStateId
    }
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal instantiates a new AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder and sets the default values.
func NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    m := &AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder instantiates a new AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder and sets the default values.
func NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesCountRequestBuilder when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) Count()(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesCountRequestBuilder) {
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get advancedThreatProtectionOnboardingDeviceSettingStates from deviceManagement
// returns a AdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponseable), nil
}
// Post create new navigation property to advancedThreatProtectionOnboardingDeviceSettingStates for deviceManagement
// returns a AdvancedThreatProtectionOnboardingDeviceSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get advancedThreatProtectionOnboardingDeviceSettingStates from deviceManagement
// returns a *RequestInformation when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to advancedThreatProtectionOnboardingDeviceSettingStates for deviceManagement
// returns a *RequestInformation when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) WithUrl(rawUrl string)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
