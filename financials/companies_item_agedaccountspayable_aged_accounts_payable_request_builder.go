package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
type CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderGetQueryParameters get agedAccountsPayable from financials
type CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderGetQueryParameters struct {
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
// CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderGetQueryParameters
}
// ByAgedAccountsPayableId provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder when successful
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) ByAgedAccountsPayableId(agedAccountsPayableId string)(*CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if agedAccountsPayableId != "" {
        urlTplParams["agedAccountsPayable%2Did"] = agedAccountsPayableId
    }
    return NewCompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAgedAccountsPayableIdGuid provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
// returns a *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder when successful
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) ByAgedAccountsPayableIdGuid(agedAccountsPayableId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["agedAccountsPayable%2Did"] = agedAccountsPayableId.String()
    return NewCompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderInternal instantiates a new CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder and sets the default values.
func NewCompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) {
    m := &CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/agedAccountsPayable{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder instantiates a new CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder and sets the default values.
func NewCompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemAgedaccountspayableCountRequestBuilder when successful
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) Count()(*CompaniesItemAgedaccountspayableCountRequestBuilder) {
    return NewCompaniesItemAgedaccountspayableCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get agedAccountsPayable from financials
// returns a AgedAccountsPayableCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgedAccountsPayableCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgedAccountsPayableCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgedAccountsPayableCollectionResponseable), nil
}
// ToGetRequestInformation get agedAccountsPayable from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder when successful
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder) {
    return NewCompaniesItemAgedaccountspayableAgedAccountsPayableRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
