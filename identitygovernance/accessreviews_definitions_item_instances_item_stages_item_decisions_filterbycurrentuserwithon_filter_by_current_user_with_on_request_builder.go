package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder provides operations to call the filterByCurrentUser method.
type AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters retrieve the accessReviewInstanceDecisionItem objects for a specific accessReviewInstance. A list of zero or more accessReviewInstanceDecisionItem objects are returned, including all of their nested properties.
type AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters struct {
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
// AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters
}
// NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, on *string)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/filterByCurrentUser(on='{on}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if on != nil {
        m.BaseRequestBuilder.PathParameters["on"] = *on
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get retrieve the accessReviewInstanceDecisionItem objects for a specific accessReviewInstance. A list of zero or more accessReviewInstanceDecisionItem objects are returned, including all of their nested properties.
// Deprecated: This method is obsolete. Use GetAsFilterByCurrentUserWithOnGetResponse instead.
// returns a AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-list-decisions?view=graph-rest-beta
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable), nil
}
// GetAsFilterByCurrentUserWithOnGetResponse retrieve the accessReviewInstanceDecisionItem objects for a specific accessReviewInstance. A list of zero or more accessReviewInstanceDecisionItem objects are returned, including all of their nested properties.
// returns a AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-list-decisions?view=graph-rest-beta
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) GetAsFilterByCurrentUserWithOnGetResponse(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable), nil
}
// ToGetRequestInformation retrieve the accessReviewInstanceDecisionItem objects for a specific accessReviewInstance. A list of zero or more accessReviewInstanceDecisionItem objects are returned, including all of their nested properties.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
