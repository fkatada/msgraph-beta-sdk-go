package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
type ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
type ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters
}
// ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal instantiates a new ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    m := &ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions/{accessReviewInstanceDecisionItem%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder instantiates a new ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property decisions for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
// returns a AccessReviewInstanceDecisionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable), nil
}
// Insights provides operations to manage the insights property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
// returns a *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Insights()(*ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property decisions in users
// returns a AccessReviewInstanceDecisionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable), nil
}
// ToDeleteRequestInformation delete navigation property decisions for users
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property decisions in users
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) WithUrl(rawUrl string)(*ItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
