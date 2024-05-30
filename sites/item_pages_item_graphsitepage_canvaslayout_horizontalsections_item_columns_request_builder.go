package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderGetQueryParameters the set of vertical columns in this section.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderGetQueryParameters struct {
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
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderGetQueryParameters
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHorizontalSectionColumnId provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) ByHorizontalSectionColumnId(horizontalSectionColumnId string)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if horizontalSectionColumnId != "" {
        urlTplParams["horizontalSectionColumn%2Did"] = horizontalSectionColumnId
    }
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderInternal instantiates a new ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) {
    m := &ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder instantiates a new ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsCountRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) Count()(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsCountRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the set of vertical columns in this section.
// returns a HorizontalSectionColumnCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHorizontalSectionColumnCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnCollectionResponseable), nil
}
// Post create new navigation property to columns for sites
// returns a HorizontalSectionColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable), nil
}
// ToGetRequestInformation the set of vertical columns in this section.
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to columns for sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
