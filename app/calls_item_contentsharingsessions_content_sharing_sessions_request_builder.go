package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
type CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetQueryParameters get contentSharingSessions from app
type CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetQueryParameters struct {
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
// CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetQueryParameters
}
// CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByContentSharingSessionId provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
// returns a *CallsItemContentsharingsessionsContentSharingSessionItemRequestBuilder when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) ByContentSharingSessionId(contentSharingSessionId string)(*CallsItemContentsharingsessionsContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if contentSharingSessionId != "" {
        urlTplParams["contentSharingSession%2Did"] = contentSharingSessionId
    }
    return NewCallsItemContentsharingsessionsContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilderInternal instantiates a new CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder and sets the default values.
func NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) {
    m := &CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/calls/{call%2Did}/contentSharingSessions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilder instantiates a new CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder and sets the default values.
func NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CallsItemContentsharingsessionsCountRequestBuilder when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) Count()(*CallsItemContentsharingsessionsCountRequestBuilder) {
    return NewCallsItemContentsharingsessionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get contentSharingSessions from app
// returns a ContentSharingSessionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) Get(ctx context.Context, requestConfiguration *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentSharingSessionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentSharingSessionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentSharingSessionCollectionResponseable), nil
}
// Post create new navigation property to contentSharingSessions for app
// returns a ContentSharingSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentSharingSessionable, requestConfiguration *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentSharingSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentSharingSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentSharingSessionable), nil
}
// ToGetRequestInformation get contentSharingSessions from app
// returns a *RequestInformation when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to contentSharingSessions for app
// returns a *RequestInformation when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentSharingSessionable, requestConfiguration *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) WithUrl(rawUrl string)(*CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) {
    return NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
