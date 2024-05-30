package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder provides operations to manage the knownIssues property of the microsoft.graph.windowsUpdates.product entity.
type WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderGetQueryParameters represents a known issue related to a Windows product.
type WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderInternal instantiates a new WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) {
    m := &WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/products/{product%2Did}/knownIssues/{knownIssue%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder instantiates a new WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property knownIssues for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents a known issue related to a Windows product.
// returns a KnownIssueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateKnownIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable), nil
}
// OriginatingKnowledgeBaseArticle provides operations to manage the originatingKnowledgeBaseArticle property of the microsoft.graph.windowsUpdates.knownIssue entity.
// returns a *WindowsUpdatesProductsItemKnownissuesItemOriginatingknowledgebasearticleOriginatingKnowledgeBaseArticleRequestBuilder when successful
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) OriginatingKnowledgeBaseArticle()(*WindowsUpdatesProductsItemKnownissuesItemOriginatingknowledgebasearticleOriginatingKnowledgeBaseArticleRequestBuilder) {
    return NewWindowsUpdatesProductsItemKnownissuesItemOriginatingknowledgebasearticleOriginatingKnowledgeBaseArticleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property knownIssues in admin
// returns a KnownIssueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable, requestConfiguration *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateKnownIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable), nil
}
// ResolvingKnowledgeBaseArticle provides operations to manage the resolvingKnowledgeBaseArticle property of the microsoft.graph.windowsUpdates.knownIssue entity.
// returns a *WindowsUpdatesProductsItemKnownissuesItemResolvingknowledgebasearticleResolvingKnowledgeBaseArticleRequestBuilder when successful
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) ResolvingKnowledgeBaseArticle()(*WindowsUpdatesProductsItemKnownissuesItemResolvingknowledgebasearticleResolvingKnowledgeBaseArticleRequestBuilder) {
    return NewWindowsUpdatesProductsItemKnownissuesItemResolvingknowledgebasearticleResolvingKnowledgeBaseArticleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property knownIssues for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents a known issue related to a Windows product.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property knownIssues in admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable, requestConfiguration *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder when successful
func (m *WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder) {
    return NewWindowsUpdatesProductsItemKnownissuesKnownIssueItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
