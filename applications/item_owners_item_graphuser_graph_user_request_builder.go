package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOwnersItemGraphuserGraphUserRequestBuilder casts the previous resource to user.
type ItemOwnersItemGraphuserGraphUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOwnersItemGraphuserGraphUserRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.user
type ItemOwnersItemGraphuserGraphUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOwnersItemGraphuserGraphUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOwnersItemGraphuserGraphUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOwnersItemGraphuserGraphUserRequestBuilderGetQueryParameters
}
// NewItemOwnersItemGraphuserGraphUserRequestBuilderInternal instantiates a new ItemOwnersItemGraphuserGraphUserRequestBuilder and sets the default values.
func NewItemOwnersItemGraphuserGraphUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnersItemGraphuserGraphUserRequestBuilder) {
    m := &ItemOwnersItemGraphuserGraphUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/owners/{directoryObject%2Did}/graph.user{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOwnersItemGraphuserGraphUserRequestBuilder instantiates a new ItemOwnersItemGraphuserGraphUserRequestBuilder and sets the default values.
func NewItemOwnersItemGraphuserGraphUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnersItemGraphuserGraphUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOwnersItemGraphuserGraphUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOwnersItemGraphuserGraphUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOwnersItemGraphuserGraphUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.user
// returns a *RequestInformation when successful
func (m *ItemOwnersItemGraphuserGraphUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOwnersItemGraphuserGraphUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOwnersItemGraphuserGraphUserRequestBuilder when successful
func (m *ItemOwnersItemGraphuserGraphUserRequestBuilder) WithUrl(rawUrl string)(*ItemOwnersItemGraphuserGraphUserRequestBuilder) {
    return NewItemOwnersItemGraphuserGraphUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
