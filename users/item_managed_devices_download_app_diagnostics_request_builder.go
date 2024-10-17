package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder provides operations to call the downloadAppDiagnostics method.
type ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesDownloadAppDiagnosticsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesDownloadAppDiagnosticsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesDownloadAppDiagnosticsRequestBuilderInternal instantiates a new ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder and sets the default values.
func NewItemManagedDevicesDownloadAppDiagnosticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder) {
    m := &ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/downloadAppDiagnostics", pathParameters),
    }
    return m
}
// NewItemManagedDevicesDownloadAppDiagnosticsRequestBuilder instantiates a new ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder and sets the default values.
func NewItemManagedDevicesDownloadAppDiagnosticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesDownloadAppDiagnosticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action downloadAppDiagnostics
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder) Post(ctx context.Context, body ItemManagedDevicesDownloadAppDiagnosticsPostRequestBodyable, requestConfiguration *ItemManagedDevicesDownloadAppDiagnosticsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation invoke action downloadAppDiagnostics
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManagedDevicesDownloadAppDiagnosticsPostRequestBodyable, requestConfiguration *ItemManagedDevicesDownloadAppDiagnosticsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder when successful
func (m *ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesDownloadAppDiagnosticsRequestBuilder) {
    return NewItemManagedDevicesDownloadAppDiagnosticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
