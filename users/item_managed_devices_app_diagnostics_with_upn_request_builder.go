package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder provides operations to call the appDiagnostics method.
type ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderGetQueryParameters invoke function appDiagnostics
type ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderGetQueryParameters
}
// NewItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderInternal instantiates a new ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder and sets the default values.
func NewItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, upn *string)(*ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder) {
    m := &ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/appDiagnostics(upn='{upn}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if upn != nil {
        m.BaseRequestBuilder.PathParameters["upn"] = *upn
    }
    return m
}
// NewItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder instantiates a new ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder and sets the default values.
func NewItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function appDiagnostics
// Deprecated: This method is obsolete. Use GetAsAppDiagnosticsWithUpnGetResponse instead.
// returns a ItemManagedDevicesAppDiagnosticsWithUpnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(ItemManagedDevicesAppDiagnosticsWithUpnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesAppDiagnosticsWithUpnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesAppDiagnosticsWithUpnResponseable), nil
}
// GetAsAppDiagnosticsWithUpnGetResponse invoke function appDiagnostics
// returns a ItemManagedDevicesAppDiagnosticsWithUpnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder) GetAsAppDiagnosticsWithUpnGetResponse(ctx context.Context, requestConfiguration *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(ItemManagedDevicesAppDiagnosticsWithUpnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesAppDiagnosticsWithUpnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesAppDiagnosticsWithUpnGetResponseable), nil
}
// ToGetRequestInformation invoke function appDiagnostics
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder when successful
func (m *ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder) {
    return NewItemManagedDevicesAppDiagnosticsWithUpnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
