package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
type CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderGetQueryParameters get salesQuoteLines from financials
type CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderGetQueryParameters struct {
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
// CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderGetQueryParameters
}
// BySalesQuoteLineId provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalesquotelinesSalesQuoteLineItemRequestBuilder when successful
func (m *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) BySalesQuoteLineId(salesQuoteLineId string)(*CompaniesItemSalesquotelinesSalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if salesQuoteLineId != "" {
        urlTplParams["salesQuoteLine%2Did"] = salesQuoteLineId
    }
    return NewCompaniesItemSalesquotelinesSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderInternal instantiates a new CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder and sets the default values.
func NewCompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) {
    m := &CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesQuoteLines{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder instantiates a new CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder and sets the default values.
func NewCompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemSalesquotelinesCountRequestBuilder when successful
func (m *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) Count()(*CompaniesItemSalesquotelinesCountRequestBuilder) {
    return NewCompaniesItemSalesquotelinesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesQuoteLines from financials
// returns a SalesQuoteLineCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder when successful
func (m *CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder) {
    return NewCompaniesItemSalesquotelinesSalesQuoteLinesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
