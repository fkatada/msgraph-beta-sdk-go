package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphsitepageGraphSitePageRequestBuilder casts the previous resource to sitePage.
type ItemPagesItemGraphsitepageGraphSitePageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphsitepageGraphSitePageRequestBuilderGetQueryParameters returns the metadata for a sitePage in the site pages list in a site.
type ItemPagesItemGraphsitepageGraphSitePageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPagesItemGraphsitepageGraphSitePageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageGraphSitePageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphsitepageGraphSitePageRequestBuilderGetQueryParameters
}
// CanvasLayout provides operations to manage the canvasLayout property of the microsoft.graph.sitePage entity.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) CanvasLayout()(*ItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutCanvasLayoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPagesItemGraphsitepageGraphSitePageRequestBuilderInternal instantiates a new ItemPagesItemGraphsitepageGraphSitePageRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageGraphSitePageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) {
    m := &ItemPagesItemGraphsitepageGraphSitePageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphsitepageGraphSitePageRequestBuilder instantiates a new ItemPagesItemGraphsitepageGraphSitePageRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageGraphSitePageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphsitepageGraphSitePageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemPagesItemGraphsitepageCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) CreatedByUser()(*ItemPagesItemGraphsitepageCreatedbyuserCreatedByUserRequestBuilder) {
    return NewItemPagesItemGraphsitepageCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns the metadata for a sitePage in the site pages list in a site.
// returns a SitePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/sitepage-get?view=graph-rest-beta
func (m *ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageGraphSitePageRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SitePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSitePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SitePageable), nil
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemPagesItemGraphsitepageLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) LastModifiedByUser()(*ItemPagesItemGraphsitepageLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewItemPagesItemGraphsitepageLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation returns the metadata for a sitePage in the site pages list in a site.
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageGraphSitePageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WebParts provides operations to manage the webParts property of the microsoft.graph.sitePage entity.
// returns a *ItemPagesItemGraphsitepageWebpartsWebPartsRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) WebParts()(*ItemPagesItemGraphsitepageWebpartsWebPartsRequestBuilder) {
    return NewItemPagesItemGraphsitepageWebpartsWebPartsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPagesItemGraphsitepageGraphSitePageRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphsitepageGraphSitePageRequestBuilder) {
    return NewItemPagesItemGraphsitepageGraphSitePageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
