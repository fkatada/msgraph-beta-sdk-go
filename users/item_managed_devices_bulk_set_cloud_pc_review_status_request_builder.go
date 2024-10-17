package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder provides operations to call the bulkSetCloudPcReviewStatus method.
type ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal instantiates a new ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    m := &ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/bulkSetCloudPcReviewStatus", pathParameters),
    }
    return m
}
// NewItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder instantiates a new ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder and sets the default values.
func NewItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action bulkSetCloudPcReviewStatus
// Deprecated: The BulkSetCloudPcReviewStatus action is deprecated and will stop supporting on September 30, 2024. Please use bulk action entity api. as of 2024-05/BulkSetCloudPcReviewStatus
// returns a CloudPcBulkRemoteActionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-bulksetcloudpcreviewstatus?view=graph-rest-beta
func (m *ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) Post(ctx context.Context, body ItemManagedDevicesBulkSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcBulkRemoteActionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcBulkRemoteActionResultable), nil
}
// ToPostRequestInformation invoke action bulkSetCloudPcReviewStatus
// Deprecated: The BulkSetCloudPcReviewStatus action is deprecated and will stop supporting on September 30, 2024. Please use bulk action entity api. as of 2024-05/BulkSetCloudPcReviewStatus
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManagedDevicesBulkSetCloudPcReviewStatusPostRequestBodyable, requestConfiguration *ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The BulkSetCloudPcReviewStatus action is deprecated and will stop supporting on September 30, 2024. Please use bulk action entity api. as of 2024-05/BulkSetCloudPcReviewStatus
// returns a *ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder when successful
func (m *ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    return NewItemManagedDevicesBulkSetCloudPcReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
