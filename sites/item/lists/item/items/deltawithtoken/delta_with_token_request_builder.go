package deltawithtoken

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeltaWithTokenRequestBuilder provides operations to call the delta method.
type DeltaWithTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeltaWithTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeltaWithTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeltaWithTokenRequestBuilderInternal instantiates a new DeltaWithTokenRequestBuilder and sets the default values.
func NewDeltaWithTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, token *string)(*DeltaWithTokenRequestBuilder) {
    m := &DeltaWithTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/microsoft.graph.delta(token='{token}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if token != nil {
        urlTplParams["token"] = *token
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeltaWithTokenRequestBuilder instantiates a new DeltaWithTokenRequestBuilder and sets the default values.
func NewDeltaWithTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeltaWithTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeltaWithTokenRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function delta
func (m *DeltaWithTokenRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function delta
func (m *DeltaWithTokenRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeltaWithTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function delta
func (m *DeltaWithTokenRequestBuilder) Get()(DeltaWithTokenResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function delta
func (m *DeltaWithTokenRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeltaWithTokenRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(DeltaWithTokenResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateDeltaWithTokenResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(DeltaWithTokenResponseable), nil
}
