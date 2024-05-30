package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetQueryParameters get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetQueryParameters struct {
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
// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetQueryParameters
}
// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByBusinessFlowId provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
// returns a *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) ByBusinessFlowId(businessFlowId string)(*ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if businessFlowId != "" {
        urlTplParams["businessFlow%2Did"] = businessFlowId
    }
    return NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal instantiates a new ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder and sets the default values.
func NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    m := &ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}/businessFlowsWithRequestsAwaitingMyDecision{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder instantiates a new ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder and sets the default values.
func NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemBusinessflowswithrequestsawaitingmydecisionCountRequestBuilder when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) Count()(*ItemBusinessflowswithrequestsawaitingmydecisionCountRequestBuilder) {
    return NewItemBusinessflowswithrequestsawaitingmydecisionCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
// returns a BusinessFlowCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessFlowCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowCollectionResponseable), nil
}
// Post create new navigation property to businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
// returns a BusinessFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable), nil
}
// ToGetRequestInformation get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
// returns a *RequestInformation when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
// returns a *RequestInformation when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) WithUrl(rawUrl string)(*ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    return NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
