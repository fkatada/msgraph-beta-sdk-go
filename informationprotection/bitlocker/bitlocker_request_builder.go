package bitlocker

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3aff3a656fbbde11cf121aa411c961df9053d91b2360949922f9c52f0a45d2c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/bitlocker/recoverykeys"
    if2ebcd3e1f6994611d75fef91c6a8af89427d1b5902b51756526761c63ff4f79 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/bitlocker/recoverykeys/item"
)

// BitlockerRequestBuilder provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
type BitlockerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BitlockerRequestBuilderGetQueryParameters get bitlocker from informationProtection
type BitlockerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BitlockerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BitlockerRequestBuilderGetQueryParameters
}
// NewBitlockerRequestBuilderInternal instantiates a new BitlockerRequestBuilder and sets the default values.
func NewBitlockerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRequestBuilder) {
    m := &BitlockerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/bitlocker{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBitlockerRequestBuilder instantiates a new BitlockerRequestBuilder and sets the default values.
func NewBitlockerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBitlockerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get bitlocker from informationProtection
func (m *BitlockerRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *BitlockerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get bitlocker from informationProtection
func (m *BitlockerRequestBuilder) Get(ctx context.Context, requestConfiguration *BitlockerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBitlockerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable), nil
}
// RecoveryKeys the recoveryKeys property
func (m *BitlockerRequestBuilder) RecoveryKeys()(*i3aff3a656fbbde11cf121aa411c961df9053d91b2360949922f9c52f0a45d2c5.RecoveryKeysRequestBuilder) {
    return i3aff3a656fbbde11cf121aa411c961df9053d91b2360949922f9c52f0a45d2c5.NewRecoveryKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoveryKeysById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.informationProtection.bitlocker.recoveryKeys.item collection
func (m *BitlockerRequestBuilder) RecoveryKeysById(id string)(*if2ebcd3e1f6994611d75fef91c6a8af89427d1b5902b51756526761c63ff4f79.BitlockerRecoveryKeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bitlockerRecoveryKey%2Did"] = id
    }
    return if2ebcd3e1f6994611d75fef91c6a8af89427d1b5902b51756526761c63ff4f79.NewBitlockerRecoveryKeyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
