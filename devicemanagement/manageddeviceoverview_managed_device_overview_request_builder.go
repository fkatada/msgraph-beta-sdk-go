package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddeviceoverviewManagedDeviceOverviewRequestBuilder provides operations to manage the managedDeviceOverview property of the microsoft.graph.deviceManagement entity.
type ManageddeviceoverviewManagedDeviceOverviewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddeviceoverviewManagedDeviceOverviewRequestBuilderGetQueryParameters device overview
type ManageddeviceoverviewManagedDeviceOverviewRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManageddeviceoverviewManagedDeviceOverviewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddeviceoverviewManagedDeviceOverviewRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddeviceoverviewManagedDeviceOverviewRequestBuilderGetQueryParameters
}
// NewManageddeviceoverviewManagedDeviceOverviewRequestBuilderInternal instantiates a new ManageddeviceoverviewManagedDeviceOverviewRequestBuilder and sets the default values.
func NewManageddeviceoverviewManagedDeviceOverviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddeviceoverviewManagedDeviceOverviewRequestBuilder) {
    m := &ManageddeviceoverviewManagedDeviceOverviewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDeviceOverview{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManageddeviceoverviewManagedDeviceOverviewRequestBuilder instantiates a new ManageddeviceoverviewManagedDeviceOverviewRequestBuilder and sets the default values.
func NewManageddeviceoverviewManagedDeviceOverviewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddeviceoverviewManagedDeviceOverviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddeviceoverviewManagedDeviceOverviewRequestBuilderInternal(urlParams, requestAdapter)
}
// Get device overview
// returns a ManagedDeviceOverviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddeviceoverviewManagedDeviceOverviewRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddeviceoverviewManagedDeviceOverviewRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceOverviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceOverviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceOverviewable), nil
}
// ToGetRequestInformation device overview
// returns a *RequestInformation when successful
func (m *ManageddeviceoverviewManagedDeviceOverviewRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddeviceoverviewManagedDeviceOverviewRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddeviceoverviewManagedDeviceOverviewRequestBuilder when successful
func (m *ManageddeviceoverviewManagedDeviceOverviewRequestBuilder) WithUrl(rawUrl string)(*ManageddeviceoverviewManagedDeviceOverviewRequestBuilder) {
    return NewManageddeviceoverviewManagedDeviceOverviewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
