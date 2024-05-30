package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder provides operations to manage the endUserNotifications property of the microsoft.graph.attackSimulationRoot entity.
type AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderGetQueryParameters represents an end user's notification for an attack simulation training.
type AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderGetQueryParameters
}
// AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderInternal instantiates a new AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder and sets the default values.
func NewAttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) {
    m := &AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/endUserNotifications/{endUserNotification%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder instantiates a new AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder and sets the default values.
func NewAttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property endUserNotifications for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Details provides operations to manage the details property of the microsoft.graph.endUserNotification entity.
// returns a *AttacksimulationEndusernotificationsItemDetailsRequestBuilder when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) Details()(*AttacksimulationEndusernotificationsItemDetailsRequestBuilder) {
    return NewAttacksimulationEndusernotificationsItemDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents an end user's notification for an attack simulation training.
// returns a EndUserNotificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEndUserNotificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationable), nil
}
// Patch update the navigation property endUserNotifications in security
// returns a EndUserNotificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationable, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEndUserNotificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationable), nil
}
// ToDeleteRequestInformation delete navigation property endUserNotifications for security
// returns a *RequestInformation when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents an end user's notification for an attack simulation training.
// returns a *RequestInformation when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property endUserNotifications in security
// returns a *RequestInformation when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationable, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) {
    return NewAttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
