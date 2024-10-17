package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder provides operations to call the getCloudPcConnectivityHistory method.
type ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters get the connectivity history of a specific Cloud PC.
type ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters
}
// NewItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal instantiates a new ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder and sets the default values.
func NewItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    m := &ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/getCloudPcConnectivityHistory(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder instantiates a new ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder and sets the default values.
func NewItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the connectivity history of a specific Cloud PC.
// Deprecated: This method is obsolete. Use GetAsGetCloudPcConnectivityHistoryGetResponse instead.
// returns a ItemCloudPCsItemGetCloudPcConnectivityHistoryResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getcloudpcconnectivityhistory?view=graph-rest-beta
func (m *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(ItemCloudPCsItemGetCloudPcConnectivityHistoryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsItemGetCloudPcConnectivityHistoryResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsItemGetCloudPcConnectivityHistoryResponseable), nil
}
// GetAsGetCloudPcConnectivityHistoryGetResponse get the connectivity history of a specific Cloud PC.
// returns a ItemCloudPCsItemGetCloudPcConnectivityHistoryGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getcloudpcconnectivityhistory?view=graph-rest-beta
func (m *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) GetAsGetCloudPcConnectivityHistoryGetResponse(ctx context.Context, requestConfiguration *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(ItemCloudPCsItemGetCloudPcConnectivityHistoryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsItemGetCloudPcConnectivityHistoryGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsItemGetCloudPcConnectivityHistoryGetResponseable), nil
}
// ToGetRequestInformation get the connectivity history of a specific Cloud PC.
// returns a *RequestInformation when successful
func (m *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder when successful
func (m *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
