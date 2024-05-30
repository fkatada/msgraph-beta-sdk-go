package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder provides operations to call the summarizeDevicePerformanceDevices method.
type UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetQueryParameters invoke function summarizeDevicePerformanceDevices
type UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetQueryParameters
}
// NewUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal instantiates a new UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, summarizeBy *string)(*UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    m := &UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDevicePerformance/summarizeDevicePerformanceDevices(summarizeBy='{summarizeBy}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if summarizeBy != nil {
        m.BaseRequestBuilder.PathParameters["summarizeBy"] = *summarizeBy
    }
    return m
}
// NewUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder instantiates a new UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function summarizeDevicePerformanceDevices
// Deprecated: This method is obsolete. Use GetAsSummarizeDevicePerformanceDevicesWithSummarizeByGetResponse instead.
// returns a UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration)(UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByResponseable), nil
}
// GetAsSummarizeDevicePerformanceDevicesWithSummarizeByGetResponse invoke function summarizeDevicePerformanceDevices
// returns a UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) GetAsSummarizeDevicePerformanceDevicesWithSummarizeByGetResponse(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration)(UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByGetResponseable), nil
}
// ToGetRequestInformation invoke function summarizeDevicePerformanceDevices
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    return NewUserexperienceanalyticsdeviceperformanceSummarizedeviceperformancedeviceswithsummarizebySummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
