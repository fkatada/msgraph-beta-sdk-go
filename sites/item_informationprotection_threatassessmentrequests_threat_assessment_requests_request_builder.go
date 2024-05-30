package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
type ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetQueryParameters get threatAssessmentRequests from sites
type ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetQueryParameters struct {
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
// ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetQueryParameters
}
// ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByThreatAssessmentRequestId provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
// returns a *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder when successful
func (m *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) ByThreatAssessmentRequestId(threatAssessmentRequestId string)(*ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if threatAssessmentRequestId != "" {
        urlTplParams["threatAssessmentRequest%2Did"] = threatAssessmentRequestId
    }
    return NewItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderInternal instantiates a new ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder and sets the default values.
func NewItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) {
    m := &ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/informationProtection/threatAssessmentRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder instantiates a new ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder and sets the default values.
func NewItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemInformationprotectionThreatassessmentrequestsCountRequestBuilder when successful
func (m *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) Count()(*ItemInformationprotectionThreatassessmentrequestsCountRequestBuilder) {
    return NewItemInformationprotectionThreatassessmentrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get threatAssessmentRequests from sites
// returns a ThreatAssessmentRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateThreatAssessmentRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentRequestCollectionResponseable), nil
}
// Post create new navigation property to threatAssessmentRequests for sites
// returns a ThreatAssessmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentRequestable, requestConfiguration *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateThreatAssessmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentRequestable), nil
}
// ToGetRequestInformation get threatAssessmentRequests from sites
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to threatAssessmentRequests for sites
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentRequestable, requestConfiguration *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder when successful
func (m *ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) WithUrl(rawUrl string)(*ItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) {
    return NewItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
