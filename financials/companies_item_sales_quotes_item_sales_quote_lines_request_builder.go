package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
type CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderGetQueryParameters get salesQuoteLines from financials
type CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderGetQueryParameters struct {
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
// CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderGetQueryParameters
}
// BySalesQuoteLineId provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder) BySalesQuoteLineId(salesQuoteLineId string)(*CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if salesQuoteLineId != "" {
        urlTplParams["salesQuoteLine%2Did"] = salesQuoteLineId
    }
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderInternal instantiates a new SalesQuoteLinesRequestBuilder and sets the default values.
func NewCompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder) {
    m := &CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesQuotes/{salesQuote%2Did}/salesQuoteLines{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder instantiates a new SalesQuoteLinesRequestBuilder and sets the default values.
func NewCompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder) Count()(*CompaniesItemSalesQuotesItemSalesQuoteLinesCountRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesQuoteLines from financials
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteLineCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineCollectionResponseable), nil
}
// ToGetRequestInformation get salesQuoteLines from financials
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
