package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder provides operations to call the delta method.
type ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters invoke function delta
type ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters struct {
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
// ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetQueryParameters
}
// NewItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal instantiates a new ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, token *string)(*ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    m := &ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/delta(token='{token}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if token != nil {
        m.BaseRequestBuilder.PathParameters["token"] = *token
    }
    return m
}
// NewItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder instantiates a new ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder and sets the default values.
func NewItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function delta
// Deprecated: This method is obsolete. Use GetAsDeltaWithTokenGetResponse instead.
// returns a ItemListsItemItemsDeltawithtokenDeltaWithTokenResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(ItemListsItemItemsDeltawithtokenDeltaWithTokenResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListsItemItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListsItemItemsDeltawithtokenDeltaWithTokenResponseable), nil
}
// GetAsDeltaWithTokenGetResponse invoke function delta
// returns a ItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) GetAsDeltaWithTokenGetResponse(ctx context.Context, requestConfiguration *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(ItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable), nil
}
// ToGetRequestInformation invoke function delta
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder when successful
func (m *ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder) {
    return NewItemListsItemItemsDeltawithtokenDeltaWithTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
