package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder provides operations to manage the itemCategory property of the microsoft.graph.item entity.
type CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderGetQueryParameters get itemCategory from financials
type CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderGetQueryParameters
}
// CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderInternal instantiates a new CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder and sets the default values.
func NewCompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) {
    m := &CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemoLines/{salesCreditMemoLine%2Did}/item/itemCategory{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder instantiates a new CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder and sets the default values.
func NewCompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property itemCategory for financials
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get itemCategory from financials
// returns a ItemCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable), nil
}
// Patch update the navigation property itemCategory in financials
// returns a ItemCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable, requestConfiguration *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable), nil
}
// ToDeleteRequestInformation delete navigation property itemCategory for financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get itemCategory from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property itemCategory in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable, requestConfiguration *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder when successful
func (m *CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder) {
    return NewCompaniesItemSalescreditmemolinesItemItemItemcategoryItemCategoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
