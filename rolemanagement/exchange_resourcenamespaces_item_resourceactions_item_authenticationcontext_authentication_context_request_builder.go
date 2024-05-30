package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder provides operations to manage the authenticationContext property of the microsoft.graph.unifiedRbacResourceAction entity.
type ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderGetQueryParameters get authenticationContext from roleManagement
type ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderGetQueryParameters
}
// NewExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderInternal instantiates a new ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder and sets the default values.
func NewExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder) {
    m := &ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/exchange/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}/resourceActions/{unifiedRbacResourceAction%2Did}/authenticationContext{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder instantiates a new ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder and sets the default values.
func NewExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get authenticationContext from roleManagement
// returns a AuthenticationContextClassReferenceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationContextClassReferenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable), nil
}
// ToGetRequestInformation get authenticationContext from roleManagement
// returns a *RequestInformation when successful
func (m *ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder when successful
func (m *ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder) WithUrl(rawUrl string)(*ExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder) {
    return NewExchangeResourcenamespacesItemResourceactionsItemAuthenticationcontextAuthenticationContextRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
