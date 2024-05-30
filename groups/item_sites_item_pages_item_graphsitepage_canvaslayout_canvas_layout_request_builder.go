package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder provides operations to manage the canvasLayout property of the microsoft.graph.sitePage entity.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderGetQueryParameters indicates the layout of the content in a given SharePoint page, including horizontal sections and vertical sections.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderGetQueryParameters
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderInternal instantiates a new ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) {
    m := &ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder instantiates a new ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property canvasLayout for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get indicates the layout of the content in a given SharePoint page, including horizontal sections and vertical sections.
// returns a CanvasLayoutable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCanvasLayoutFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable), nil
}
// HorizontalSections provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) HorizontalSections()(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property canvasLayout in groups
// returns a CanvasLayoutable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCanvasLayoutFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable), nil
}
// ToDeleteRequestInformation delete navigation property canvasLayout for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation indicates the layout of the content in a given SharePoint page, including horizontal sections and vertical sections.
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property canvasLayout in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// VerticalSection provides operations to manage the verticalSection property of the microsoft.graph.canvasLayout entity.
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) VerticalSection()(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
