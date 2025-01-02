package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder provides operations to manage the applicationSegments property of the microsoft.graph.onPremisesPublishingProfile entity.
type ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderGetQueryParameters represents the segment configurations that are allowed for an on-premises non-web application published through Microsoft Entra application proxy.
type ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderGetQueryParameters
}
// NewItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderInternal instantiates a new ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder and sets the default values.
func NewItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder) {
    m := &ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/applicationSegments/{ipApplicationSegment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder instantiates a new ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder and sets the default values.
func NewItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents the segment configurations that are allowed for an on-premises non-web application published through Microsoft Entra application proxy.
// returns a IpApplicationSegmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IpApplicationSegmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIpApplicationSegmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IpApplicationSegmentable), nil
}
// ToGetRequestInformation represents the segment configurations that are allowed for an on-premises non-web application published through Microsoft Entra application proxy.
// returns a *RequestInformation when successful
func (m *ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder when successful
func (m *ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder) WithUrl(rawUrl string)(*ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder) {
    return NewItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
