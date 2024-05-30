package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder provides operations to call the requestRemoteAssistance method.
type ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal instantiates a new ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder and sets the default values.
func NewManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    m := &ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/requestRemoteAssistance", pathParameters),
    }
    return m
}
// NewManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder instantiates a new ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder and sets the default values.
func NewManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post request remote assistance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) Post(ctx context.Context, requestConfiguration *ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation request remote assistance
// returns a *RequestInformation when successful
func (m *ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder when successful
func (m *ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    return NewManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
