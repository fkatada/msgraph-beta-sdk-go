package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder provides operations to manage the activities property of the microsoft.graph.itemActivityStat entity.
type ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderGetQueryParameters exposes the itemActivities represented in this itemActivityStat resource.
type ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderGetQueryParameters struct {
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
// ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderGetQueryParameters
}
// ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByItemActivityId provides operations to manage the activities property of the microsoft.graph.itemActivityStat entity.
// returns a *ItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder when successful
func (m *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) ByItemActivityId(itemActivityId string)(*ItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if itemActivityId != "" {
        urlTplParams["itemActivity%2Did"] = itemActivityId
    }
    return NewItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderInternal instantiates a new ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder and sets the default values.
func NewItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) {
    m := &ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/analytics/itemActivityStats/{itemActivityStat%2Did}/activities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder instantiates a new ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder and sets the default values.
func NewItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAnalyticsItemactivitystatsItemActivitiesCountRequestBuilder when successful
func (m *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) Count()(*ItemAnalyticsItemactivitystatsItemActivitiesCountRequestBuilder) {
    return NewItemAnalyticsItemactivitystatsItemActivitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get exposes the itemActivities represented in this itemActivityStat resource.
// returns a ItemActivityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemActivityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityCollectionResponseable), nil
}
// Post create new navigation property to activities for sites
// returns a ItemActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityable, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityable), nil
}
// ToGetRequestInformation exposes the itemActivities represented in this itemActivityStat resource.
// returns a *RequestInformation when successful
func (m *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to activities for sites
// returns a *RequestInformation when successful
func (m *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityable, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder when successful
func (m *ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) WithUrl(rawUrl string)(*ItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder) {
    return NewItemAnalyticsItemactivitystatsItemActivitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
