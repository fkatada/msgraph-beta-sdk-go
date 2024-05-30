package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSharedwithmeSharedWithMeRequestBuilder provides operations to call the sharedWithMe method.
type ItemSharedwithmeSharedWithMeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSharedwithmeSharedWithMeRequestBuilderGetQueryParameters get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they're items from a different drive.
type ItemSharedwithmeSharedWithMeRequestBuilderGetQueryParameters struct {
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
// ItemSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSharedwithmeSharedWithMeRequestBuilderGetQueryParameters
}
// NewItemSharedwithmeSharedWithMeRequestBuilderInternal instantiates a new ItemSharedwithmeSharedWithMeRequestBuilder and sets the default values.
func NewItemSharedwithmeSharedWithMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSharedwithmeSharedWithMeRequestBuilder) {
    m := &ItemSharedwithmeSharedWithMeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/sharedWithMe(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSharedwithmeSharedWithMeRequestBuilder instantiates a new ItemSharedwithmeSharedWithMeRequestBuilder and sets the default values.
func NewItemSharedwithmeSharedWithMeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSharedwithmeSharedWithMeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSharedwithmeSharedWithMeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they're items from a different drive.
// Deprecated: This method is obsolete. Use GetAsSharedWithMeGetResponse instead.
// returns a ItemSharedwithmeSharedWithMeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/drive-sharedwithme?view=graph-rest-beta
func (m *ItemSharedwithmeSharedWithMeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(ItemSharedwithmeSharedWithMeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSharedwithmeSharedWithMeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSharedwithmeSharedWithMeResponseable), nil
}
// GetAsSharedWithMeGetResponse get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they're items from a different drive.
// returns a ItemSharedwithmeSharedWithMeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/drive-sharedwithme?view=graph-rest-beta
func (m *ItemSharedwithmeSharedWithMeRequestBuilder) GetAsSharedWithMeGetResponse(ctx context.Context, requestConfiguration *ItemSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(ItemSharedwithmeSharedWithMeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSharedwithmeSharedWithMeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSharedwithmeSharedWithMeGetResponseable), nil
}
// ToGetRequestInformation get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they're items from a different drive.
// returns a *RequestInformation when successful
func (m *ItemSharedwithmeSharedWithMeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSharedwithmeSharedWithMeRequestBuilder when successful
func (m *ItemSharedwithmeSharedWithMeRequestBuilder) WithUrl(rawUrl string)(*ItemSharedwithmeSharedWithMeRequestBuilder) {
    return NewItemSharedwithmeSharedWithMeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
