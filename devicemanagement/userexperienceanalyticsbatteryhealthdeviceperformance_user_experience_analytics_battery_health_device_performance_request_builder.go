package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderGetQueryParameters user Experience Analytics Battery Health Device Performance
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsBatteryHealthDevicePerformanceId provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) ByUserExperienceAnalyticsBatteryHealthDevicePerformanceId(userExperienceAnalyticsBatteryHealthDevicePerformanceId string)(*UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsBatteryHealthDevicePerformanceId != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthDevicePerformance%2Did"] = userExperienceAnalyticsBatteryHealthDevicePerformanceId
    }
    return NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderInternal instantiates a new UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) {
    m := &UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder instantiates a new UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsbatteryhealthdeviceperformanceCountRequestBuilder when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) Count()(*UserexperienceanalyticsbatteryhealthdeviceperformanceCountRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthdeviceperformanceCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user Experience Analytics Battery Health Device Performance
// returns a UserExperienceAnalyticsBatteryHealthDevicePerformanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsBatteryHealthDevicePerformance for deviceManagement
// returns a UserExperienceAnalyticsBatteryHealthDevicePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable), nil
}
// ToGetRequestInformation user Experience Analytics Battery Health Device Performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsBatteryHealthDevicePerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDevicePerformanceable, requestConfiguration *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder when successful
func (m *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
