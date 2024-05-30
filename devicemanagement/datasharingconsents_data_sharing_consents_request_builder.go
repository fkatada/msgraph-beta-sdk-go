package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DatasharingconsentsDataSharingConsentsRequestBuilder provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
type DatasharingconsentsDataSharingConsentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DatasharingconsentsDataSharingConsentsRequestBuilderGetQueryParameters data sharing consents.
type DatasharingconsentsDataSharingConsentsRequestBuilderGetQueryParameters struct {
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
// DatasharingconsentsDataSharingConsentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasharingconsentsDataSharingConsentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DatasharingconsentsDataSharingConsentsRequestBuilderGetQueryParameters
}
// DatasharingconsentsDataSharingConsentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DatasharingconsentsDataSharingConsentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDataSharingConsentId provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
// returns a *DatasharingconsentsDataSharingConsentItemRequestBuilder when successful
func (m *DatasharingconsentsDataSharingConsentsRequestBuilder) ByDataSharingConsentId(dataSharingConsentId string)(*DatasharingconsentsDataSharingConsentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if dataSharingConsentId != "" {
        urlTplParams["dataSharingConsent%2Did"] = dataSharingConsentId
    }
    return NewDatasharingconsentsDataSharingConsentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDatasharingconsentsDataSharingConsentsRequestBuilderInternal instantiates a new DatasharingconsentsDataSharingConsentsRequestBuilder and sets the default values.
func NewDatasharingconsentsDataSharingConsentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasharingconsentsDataSharingConsentsRequestBuilder) {
    m := &DatasharingconsentsDataSharingConsentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/dataSharingConsents{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDatasharingconsentsDataSharingConsentsRequestBuilder instantiates a new DatasharingconsentsDataSharingConsentsRequestBuilder and sets the default values.
func NewDatasharingconsentsDataSharingConsentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DatasharingconsentsDataSharingConsentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDatasharingconsentsDataSharingConsentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DatasharingconsentsCountRequestBuilder when successful
func (m *DatasharingconsentsDataSharingConsentsRequestBuilder) Count()(*DatasharingconsentsCountRequestBuilder) {
    return NewDatasharingconsentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get data sharing consents.
// returns a DataSharingConsentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DatasharingconsentsDataSharingConsentsRequestBuilder) Get(ctx context.Context, requestConfiguration *DatasharingconsentsDataSharingConsentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataSharingConsentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataSharingConsentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataSharingConsentCollectionResponseable), nil
}
// Post create new navigation property to dataSharingConsents for deviceManagement
// returns a DataSharingConsentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DatasharingconsentsDataSharingConsentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataSharingConsentable, requestConfiguration *DatasharingconsentsDataSharingConsentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataSharingConsentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataSharingConsentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataSharingConsentable), nil
}
// ToGetRequestInformation data sharing consents.
// returns a *RequestInformation when successful
func (m *DatasharingconsentsDataSharingConsentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DatasharingconsentsDataSharingConsentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dataSharingConsents for deviceManagement
// returns a *RequestInformation when successful
func (m *DatasharingconsentsDataSharingConsentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataSharingConsentable, requestConfiguration *DatasharingconsentsDataSharingConsentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DatasharingconsentsDataSharingConsentsRequestBuilder when successful
func (m *DatasharingconsentsDataSharingConsentsRequestBuilder) WithUrl(rawUrl string)(*DatasharingconsentsDataSharingConsentsRequestBuilder) {
    return NewDatasharingconsentsDataSharingConsentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
