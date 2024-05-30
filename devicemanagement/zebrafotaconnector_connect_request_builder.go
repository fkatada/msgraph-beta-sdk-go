package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebrafotaconnectorConnectRequestBuilder provides operations to call the connect method.
type ZebrafotaconnectorConnectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebrafotaconnectorConnectRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotaconnectorConnectRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebrafotaconnectorConnectRequestBuilderInternal instantiates a new ZebrafotaconnectorConnectRequestBuilder and sets the default values.
func NewZebrafotaconnectorConnectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotaconnectorConnectRequestBuilder) {
    m := &ZebrafotaconnectorConnectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaConnector/connect", pathParameters),
    }
    return m
}
// NewZebrafotaconnectorConnectRequestBuilder instantiates a new ZebrafotaconnectorConnectRequestBuilder and sets the default values.
func NewZebrafotaconnectorConnectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotaconnectorConnectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebrafotaconnectorConnectRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action connect
// Deprecated: This method is obsolete. Use PostAsConnectPostResponse instead.
// returns a ZebrafotaconnectorConnectResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorConnectRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebrafotaconnectorConnectRequestBuilderPostRequestConfiguration)(ZebrafotaconnectorConnectResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebrafotaconnectorConnectResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebrafotaconnectorConnectResponseable), nil
}
// PostAsConnectPostResponse invoke action connect
// returns a ZebrafotaconnectorConnectPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorConnectRequestBuilder) PostAsConnectPostResponse(ctx context.Context, requestConfiguration *ZebrafotaconnectorConnectRequestBuilderPostRequestConfiguration)(ZebrafotaconnectorConnectPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebrafotaconnectorConnectPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebrafotaconnectorConnectPostResponseable), nil
}
// ToPostRequestInformation invoke action connect
// returns a *RequestInformation when successful
func (m *ZebrafotaconnectorConnectRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebrafotaconnectorConnectRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ZebrafotaconnectorConnectRequestBuilder when successful
func (m *ZebrafotaconnectorConnectRequestBuilder) WithUrl(rawUrl string)(*ZebrafotaconnectorConnectRequestBuilder) {
    return NewZebrafotaconnectorConnectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
