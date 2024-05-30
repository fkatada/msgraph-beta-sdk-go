package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder provides operations to manage the customerPayments property of the microsoft.graph.customerPaymentJournal entity.
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderGetQueryParameters get customerPayments from financials
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderGetQueryParameters struct {
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
// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderGetQueryParameters
}
// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomerPaymentId provides operations to manage the customerPayments property of the microsoft.graph.customerPaymentJournal entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentItemRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) ByCustomerPaymentId(customerPaymentId string)(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customerPaymentId != "" {
        urlTplParams["customerPayment%2Did"] = customerPaymentId
    }
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByCustomerPaymentIdGuid provides operations to manage the customerPayments property of the microsoft.graph.customerPaymentJournal entity.
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentItemRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) ByCustomerPaymentIdGuid(customerPaymentId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["customerPayment%2Did"] = customerPaymentId.String()
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderInternal instantiates a new CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder and sets the default values.
func NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) {
    m := &CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}/customerPayments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder instantiates a new CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder and sets the default values.
func NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCountRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) Count()(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCountRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get customerPayments from financials
// returns a CustomerPaymentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerPaymentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentCollectionResponseable), nil
}
// Post create new navigation property to customerPayments for financials
// returns a CustomerPaymentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerPaymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable), nil
}
// ToGetRequestInformation get customerPayments from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to customerPayments for financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsCustomerPaymentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
