package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDevicesItemCommandsItemResponsepayloadRequestBuilder provides operations to manage the responsepayload property of the microsoft.graph.command entity.
type ItemDevicesItemCommandsItemResponsepayloadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDevicesItemCommandsItemResponsepayloadRequestBuilderGetQueryParameters get responsepayload from users
type ItemDevicesItemCommandsItemResponsepayloadRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDevicesItemCommandsItemResponsepayloadRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesItemCommandsItemResponsepayloadRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDevicesItemCommandsItemResponsepayloadRequestBuilderGetQueryParameters
}
// NewItemDevicesItemCommandsItemResponsepayloadRequestBuilderInternal instantiates a new ItemDevicesItemCommandsItemResponsepayloadRequestBuilder and sets the default values.
func NewItemDevicesItemCommandsItemResponsepayloadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemCommandsItemResponsepayloadRequestBuilder) {
    m := &ItemDevicesItemCommandsItemResponsepayloadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/commands/{command%2Did}/responsepayload{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemDevicesItemCommandsItemResponsepayloadRequestBuilder instantiates a new ItemDevicesItemCommandsItemResponsepayloadRequestBuilder and sets the default values.
func NewItemDevicesItemCommandsItemResponsepayloadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemCommandsItemResponsepayloadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemCommandsItemResponsepayloadRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get responsepayload from users
// returns a PayloadResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDevicesItemCommandsItemResponsepayloadRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDevicesItemCommandsItemResponsepayloadRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PayloadResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePayloadResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PayloadResponseable), nil
}
// ToGetRequestInformation get responsepayload from users
// returns a *RequestInformation when successful
func (m *ItemDevicesItemCommandsItemResponsepayloadRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDevicesItemCommandsItemResponsepayloadRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDevicesItemCommandsItemResponsepayloadRequestBuilder when successful
func (m *ItemDevicesItemCommandsItemResponsepayloadRequestBuilder) WithUrl(rawUrl string)(*ItemDevicesItemCommandsItemResponsepayloadRequestBuilder) {
    return NewItemDevicesItemCommandsItemResponsepayloadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
