package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder provides operations to call the getCloudPcReviewStatus method.
type ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal instantiates a new ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    m := &ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/getCloudPcReviewStatus()", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder instantiates a new ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the review status of a specific Cloud PC device.
// Deprecated: The getCloudPcReviewStatus API is deprecated and will stop returning data on Apr 30, 2024. Please use the new retrieveReviewStatus API as of 2024-01/getCloudPcReviewStatus
// returns a CloudPcReviewStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-getcloudpcreviewstatus?view=graph-rest-beta
func (m *ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReviewStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable), nil
}
// ToGetRequestInformation get the review status of a specific Cloud PC device.
// Deprecated: The getCloudPcReviewStatus API is deprecated and will stop returning data on Apr 30, 2024. Please use the new retrieveReviewStatus API as of 2024-01/getCloudPcReviewStatus
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The getCloudPcReviewStatus API is deprecated and will stop returning data on Apr 30, 2024. Please use the new retrieveReviewStatus API as of 2024-01/getCloudPcReviewStatus
// returns a *ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder when successful
func (m *ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
