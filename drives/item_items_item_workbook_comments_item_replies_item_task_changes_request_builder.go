package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder provides operations to manage the changes property of the microsoft.graph.workbookDocumentTask entity.
type ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderGetQueryParameters a collection of task change histories.
type ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderGetQueryParameters struct {
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
// ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWorkbookDocumentTaskChangeId provides operations to manage the changes property of the microsoft.graph.workbookDocumentTask entity.
// returns a *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder when successful
func (m *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) ByWorkbookDocumentTaskChangeId(workbookDocumentTaskChangeId string)(*ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workbookDocumentTaskChangeId != "" {
        urlTplParams["workbookDocumentTaskChange%2Did"] = workbookDocumentTaskChangeId
    }
    return NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderInternal instantiates a new ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) {
    m := &ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/comments/{workbookComment%2Did}/replies/{workbookCommentReply%2Did}/task/changes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder instantiates a new ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to call the count method.
// returns a *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder when successful
func (m *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) Count()(*ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder) {
    return NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of task change histories.
// returns a WorkbookDocumentTaskChangeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookDocumentTaskChangeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeCollectionResponseable), nil
}
// ItemAtWithIndex provides operations to call the itemAt method.
// returns a *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) ItemAtWithIndex(index *int32)(*ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesItematwithindexItemAtWithIndexRequestBuilder) {
    return NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesItematwithindexItemAtWithIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, index)
}
// Post create new navigation property to changes for drives
// returns a WorkbookDocumentTaskChangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, requestConfiguration *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookDocumentTaskChangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable), nil
}
// ToGetRequestInformation a collection of task change histories.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to changes for drives
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, requestConfiguration *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder when successful
func (m *ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder) {
    return NewItemItemsItemWorkbookCommentsItemRepliesItemTaskChangesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
