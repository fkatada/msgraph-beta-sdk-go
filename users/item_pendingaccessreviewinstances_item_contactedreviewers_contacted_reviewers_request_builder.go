package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
type ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderGetQueryParameters returns the collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers, contactedReviewers returns their individual identities. Supports $select. Read-only.
type ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderGetQueryParameters struct {
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
// ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderGetQueryParameters
}
// ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessReviewReviewerId provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingaccessreviewinstancesItemContactedreviewersAccessReviewReviewerItemRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) ByAccessReviewReviewerId(accessReviewReviewerId string)(*ItemPendingaccessreviewinstancesItemContactedreviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessReviewReviewerId != "" {
        urlTplParams["accessReviewReviewer%2Did"] = accessReviewReviewerId
    }
    return NewItemPendingaccessreviewinstancesItemContactedreviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderInternal instantiates a new ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) {
    m := &ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/contactedReviewers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder instantiates a new ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPendingaccessreviewinstancesItemContactedreviewersCountRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) Count()(*ItemPendingaccessreviewinstancesItemContactedreviewersCountRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemContactedreviewersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns the collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers, contactedReviewers returns their individual identities. Supports $select. Read-only.
// returns a AccessReviewReviewerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewReviewerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerCollectionResponseable), nil
}
// Post create new navigation property to contactedReviewers for users
// returns a AccessReviewReviewerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerable, requestConfiguration *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewReviewerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerable), nil
}
// ToGetRequestInformation returns the collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers, contactedReviewers returns their individual identities. Supports $select. Read-only.
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to contactedReviewers for users
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerable, requestConfiguration *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) WithUrl(rawUrl string)(*ItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemContactedreviewersContactedReviewersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
