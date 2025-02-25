package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.mailboxFolder entity.
type ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderGetQueryParameters read the properties and relationships of a mailboxItem object.
type ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderGetQueryParameters
}
// NewExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderInternal instantiates a new ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder) {
    m := &ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/mailboxes/{mailbox%2Did}/folders/{mailboxFolder%2Did}/items/{mailboxItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder instantiates a new ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a mailboxItem object.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport on 2021-08-19 and will be removed 2021-11-15
// returns a MailboxItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailboxitem-get?view=graph-rest-beta
func (m *ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxItemable), nil
}
// ToGetRequestInformation read the properties and relationships of a mailboxItem object.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport on 2021-08-19 and will be removed 2021-11-15
// returns a *RequestInformation when successful
func (m *ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport on 2021-08-19 and will be removed 2021-11-15
// returns a *ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder) WithUrl(rawUrl string)(*ExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder) {
    return NewExchangeMailboxesItemFoldersItemItemsMailboxItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
