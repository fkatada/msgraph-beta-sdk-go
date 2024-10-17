package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPlannerAllDeltaRequestBuilder provides operations to call the delta method.
type ItemPlannerAllDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPlannerAllDeltaRequestBuilderGetQueryParameters retrieves changes to objects that the user is subscribed to. This method allows your application to track changes to objects that the user can access from within Planner over time. The return value of this method might contain heterogeneous types of objects from Planner. For more information about tracking changes in Microsoft Graph data, see Use delta query to track changes in Microsoft Graph data.
type ItemPlannerAllDeltaRequestBuilderGetQueryParameters struct {
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
// ItemPlannerAllDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerAllDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPlannerAllDeltaRequestBuilderGetQueryParameters
}
// NewItemPlannerAllDeltaRequestBuilderInternal instantiates a new ItemPlannerAllDeltaRequestBuilder and sets the default values.
func NewItemPlannerAllDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerAllDeltaRequestBuilder) {
    m := &ItemPlannerAllDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/planner/all/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPlannerAllDeltaRequestBuilder instantiates a new ItemPlannerAllDeltaRequestBuilder and sets the default values.
func NewItemPlannerAllDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerAllDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPlannerAllDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves changes to objects that the user is subscribed to. This method allows your application to track changes to objects that the user can access from within Planner over time. The return value of this method might contain heterogeneous types of objects from Planner. For more information about tracking changes in Microsoft Graph data, see Use delta query to track changes in Microsoft Graph data.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ItemPlannerAllDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/planneruser-list-delta?view=graph-rest-beta
func (m *ItemPlannerAllDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPlannerAllDeltaRequestBuilderGetRequestConfiguration)(ItemPlannerAllDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemPlannerAllDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemPlannerAllDeltaResponseable), nil
}
// GetAsDeltaGetResponse retrieves changes to objects that the user is subscribed to. This method allows your application to track changes to objects that the user can access from within Planner over time. The return value of this method might contain heterogeneous types of objects from Planner. For more information about tracking changes in Microsoft Graph data, see Use delta query to track changes in Microsoft Graph data.
// returns a ItemPlannerAllDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/planneruser-list-delta?view=graph-rest-beta
func (m *ItemPlannerAllDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemPlannerAllDeltaRequestBuilderGetRequestConfiguration)(ItemPlannerAllDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemPlannerAllDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemPlannerAllDeltaGetResponseable), nil
}
// ToGetRequestInformation retrieves changes to objects that the user is subscribed to. This method allows your application to track changes to objects that the user can access from within Planner over time. The return value of this method might contain heterogeneous types of objects from Planner. For more information about tracking changes in Microsoft Graph data, see Use delta query to track changes in Microsoft Graph data.
// returns a *RequestInformation when successful
func (m *ItemPlannerAllDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPlannerAllDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPlannerAllDeltaRequestBuilder when successful
func (m *ItemPlannerAllDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemPlannerAllDeltaRequestBuilder) {
    return NewItemPlannerAllDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
