package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderGetQueryParameters user experience analytics devices without cloud identity.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsDeviceWithoutCloudIdentityId provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) ByUserExperienceAnalyticsDeviceWithoutCloudIdentityId(userExperienceAnalyticsDeviceWithoutCloudIdentityId string)(*UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsDeviceWithoutCloudIdentityId != "" {
        urlTplParams["userExperienceAnalyticsDeviceWithoutCloudIdentity%2Did"] = userExperienceAnalyticsDeviceWithoutCloudIdentityId
    }
    return NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderInternal instantiates a new UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) {
    m := &UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder instantiates a new UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsdeviceswithoutcloudidentityCountRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) Count()(*UserexperienceanalyticsdeviceswithoutcloudidentityCountRequestBuilder) {
    return NewUserexperienceanalyticsdeviceswithoutcloudidentityCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics devices without cloud identity.
// returns a UserExperienceAnalyticsDeviceWithoutCloudIdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsDevicesWithoutCloudIdentity for deviceManagement
// returns a UserExperienceAnalyticsDeviceWithoutCloudIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable), nil
}
// ToGetRequestInformation user experience analytics devices without cloud identity.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsDevicesWithoutCloudIdentity for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceWithoutCloudIdentityable, requestConfiguration *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder when successful
func (m *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) {
    return NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
