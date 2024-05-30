package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder provides operations to manage the approval property of the microsoft.graph.userConsentRequest entity.
type ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderGetQueryParameters approval decisions associated with a request.
type ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderGetQueryParameters
}
// ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderInternal instantiates a new ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder and sets the default values.
func NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) {
    m := &ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/appConsentRequestsForApproval/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}/approval{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder instantiates a new ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder and sets the default values.
func NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property approval for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get approval decisions associated with a request.
// returns a Approvalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable), nil
}
// Patch update the navigation property approval in users
// returns a Approvalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable, requestConfiguration *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable), nil
}
// Steps provides operations to manage the steps property of the microsoft.graph.approval entity.
// returns a *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsRequestBuilder when successful
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) Steps()(*ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsRequestBuilder) {
    return NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property approval for users
// returns a *RequestInformation when successful
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation approval decisions associated with a request.
// returns a *RequestInformation when successful
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property approval in users
// returns a *RequestInformation when successful
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Approvalable, requestConfiguration *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder when successful
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) WithUrl(rawUrl string)(*ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder) {
    return NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
