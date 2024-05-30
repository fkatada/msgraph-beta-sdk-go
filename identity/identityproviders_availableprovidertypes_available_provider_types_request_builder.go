package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder provides operations to call the availableProviderTypes method.
type IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters get all identity providers supported in a directory.
type IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters
}
// NewIdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal instantiates a new IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder and sets the default values.
func NewIdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    m := &IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/identityProviders/availableProviderTypes(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder instantiates a new IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder and sets the default values.
func NewIdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all identity providers supported in a directory.
// Deprecated: This method is obsolete. Use GetAsAvailableProviderTypesGetResponse instead.
// returns a IdentityprovidersAvailableprovidertypesAvailableProviderTypesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityproviderbase-availableprovidertypes?view=graph-rest-beta
func (m *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(IdentityprovidersAvailableprovidertypesAvailableProviderTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIdentityprovidersAvailableprovidertypesAvailableProviderTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IdentityprovidersAvailableprovidertypesAvailableProviderTypesResponseable), nil
}
// GetAsAvailableProviderTypesGetResponse get all identity providers supported in a directory.
// returns a IdentityprovidersAvailableprovidertypesAvailableProviderTypesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityproviderbase-availableprovidertypes?view=graph-rest-beta
func (m *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder) GetAsAvailableProviderTypesGetResponse(ctx context.Context, requestConfiguration *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(IdentityprovidersAvailableprovidertypesAvailableProviderTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIdentityprovidersAvailableprovidertypesAvailableProviderTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IdentityprovidersAvailableprovidertypesAvailableProviderTypesGetResponseable), nil
}
// ToGetRequestInformation get all identity providers supported in a directory.
// returns a *RequestInformation when successful
func (m *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder when successful
func (m *IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder) WithUrl(rawUrl string)(*IdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    return NewIdentityprovidersAvailableprovidertypesAvailableProviderTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
