package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder provides operations to call the createEnrollmentNotificationConfiguration method.
type ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderInternal instantiates a new ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder and sets the default values.
func NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) {
    m := &ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/deviceEnrollmentConfigurations/createEnrollmentNotificationConfiguration", pathParameters),
    }
    return m
}
// NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder instantiates a new ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder and sets the default values.
func NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createEnrollmentNotificationConfiguration
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) Post(ctx context.Context, body ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyable, requestConfiguration *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action createEnrollmentNotificationConfiguration
// returns a *RequestInformation when successful
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyable, requestConfiguration *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder when successful
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) WithUrl(rawUrl string)(*ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) {
    return NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
