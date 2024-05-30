package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters get roleEligibilityScheduleRequests from roleManagement
type EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleEligibilityScheduleRequestId provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) ByUnifiedRoleEligibilityScheduleRequestId(unifiedRoleEligibilityScheduleRequestId string)(*EntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleEligibilityScheduleRequestId != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest%2Did"] = unifiedRoleEligibilityScheduleRequestId
    }
    return NewEntitlementmanagementRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderInternal instantiates a new EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) {
    m := &EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder instantiates a new EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsCountRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) Count()(*EntitlementmanagementRoleeligibilityschedulerequestsCountRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*EntitlementmanagementRoleeligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get roleEligibilityScheduleRequests from roleManagement
// returns a UnifiedRoleEligibilityScheduleRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestCollectionResponseable), nil
}
// Post create new navigation property to roleEligibilityScheduleRequests for roleManagement
// returns a UnifiedRoleEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable), nil
}
// ToGetRequestInformation get roleEligibilityScheduleRequests from roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleEligibilityScheduleRequests for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
