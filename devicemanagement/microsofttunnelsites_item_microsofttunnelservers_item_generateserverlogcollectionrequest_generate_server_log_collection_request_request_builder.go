package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder provides operations to call the generateServerLogCollectionRequest method.
type MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilderInternal instantiates a new MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder) {
    m := &MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers/{microsoftTunnelServer%2Did}/generateServerLogCollectionRequest", pathParameters),
    }
    return m
}
// NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder instantiates a new MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action generateServerLogCollectionRequest
// returns a MicrosoftTunnelServerLogCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder) Post(ctx context.Context, body MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestPostRequestBodyable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable), nil
}
// ToPostRequestInformation invoke action generateServerLogCollectionRequest
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestPostRequestBodyable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
