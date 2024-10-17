package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder provides operations to call the rotateBitLockerKeys method.
type ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemRotateBitLockerKeysRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemRotateBitLockerKeysRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal instantiates a new ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder and sets the default values.
func NewItemManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    m := &ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/rotateBitLockerKeys", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemRotateBitLockerKeysRequestBuilder instantiates a new ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder and sets the default values.
func NewItemManagedDevicesItemRotateBitLockerKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Post rotate BitLockerKeys
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManagedDevicesItemRotateBitLockerKeysRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation rotate BitLockerKeys
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemRotateBitLockerKeysRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder when successful
func (m *ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewItemManagedDevicesItemRotateBitLockerKeysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
