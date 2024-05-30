package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder provides operations to manage the windowsInformationProtectionWipeActions property of the microsoft.graph.deviceAppManagement entity.
type WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderGetQueryParameters windows information protection wipe actions.
type WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderGetQueryParameters
}
// WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderInternal instantiates a new WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) {
    m := &WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionWipeActions/{windowsInformationProtectionWipeAction%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder instantiates a new WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsInformationProtectionWipeActions for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get windows information protection wipe actions.
// returns a WindowsInformationProtectionWipeActionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionWipeActionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionWipeActionable), nil
}
// Patch update the navigation property windowsInformationProtectionWipeActions in deviceAppManagement
// returns a WindowsInformationProtectionWipeActionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionWipeActionable, requestConfiguration *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionWipeActionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionWipeActionable), nil
}
// ToDeleteRequestInformation delete navigation property windowsInformationProtectionWipeActions for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation windows information protection wipe actions.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsInformationProtectionWipeActions in deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionWipeActionable, requestConfiguration *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder when successful
func (m *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder) {
    return NewWindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
