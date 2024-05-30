package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeysetsItemGetactivekeyGetActiveKeyRequestBuilder provides operations to call the getActiveKey method.
type KeysetsItemGetactivekeyGetActiveKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// KeysetsItemGetactivekeyGetActiveKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeysetsItemGetactivekeyGetActiveKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewKeysetsItemGetactivekeyGetActiveKeyRequestBuilderInternal instantiates a new KeysetsItemGetactivekeyGetActiveKeyRequestBuilder and sets the default values.
func NewKeysetsItemGetactivekeyGetActiveKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsItemGetactivekeyGetActiveKeyRequestBuilder) {
    m := &KeysetsItemGetactivekeyGetActiveKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}/getActiveKey()", pathParameters),
    }
    return m
}
// NewKeysetsItemGetactivekeyGetActiveKeyRequestBuilder instantiates a new KeysetsItemGetactivekeyGetActiveKeyRequestBuilder and sets the default values.
func NewKeysetsItemGetactivekeyGetActiveKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsItemGetactivekeyGetActiveKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeysetsItemGetactivekeyGetActiveKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the currently active trustFrameworkKey in a trustFrameworkKeySet. Only one key is active in the keyset at a time.
// returns a TrustFrameworkKeyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/trustframeworkkeyset-getactivekey?view=graph-rest-beta
func (m *KeysetsItemGetactivekeyGetActiveKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *KeysetsItemGetactivekeyGetActiveKeyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable), nil
}
// ToGetRequestInformation get the currently active trustFrameworkKey in a trustFrameworkKeySet. Only one key is active in the keyset at a time.
// returns a *RequestInformation when successful
func (m *KeysetsItemGetactivekeyGetActiveKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *KeysetsItemGetactivekeyGetActiveKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *KeysetsItemGetactivekeyGetActiveKeyRequestBuilder when successful
func (m *KeysetsItemGetactivekeyGetActiveKeyRequestBuilder) WithUrl(rawUrl string)(*KeysetsItemGetactivekeyGetActiveKeyRequestBuilder) {
    return NewKeysetsItemGetactivekeyGetActiveKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
