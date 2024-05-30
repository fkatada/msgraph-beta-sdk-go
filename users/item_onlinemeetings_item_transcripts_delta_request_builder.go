package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder provides operations to call the delta method.
type ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderGetQueryParameters invoke function delta
type ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderGetQueryParameters struct {
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
// ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderGetQueryParameters
}
// NewItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderInternal instantiates a new ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder) {
    m := &ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/transcripts/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder instantiates a new ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function delta
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ItemOnlinemeetingsItemTranscriptsDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderGetRequestConfiguration)(ItemOnlinemeetingsItemTranscriptsDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlinemeetingsItemTranscriptsDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlinemeetingsItemTranscriptsDeltaResponseable), nil
}
// GetAsDeltaGetResponse invoke function delta
// returns a ItemOnlinemeetingsItemTranscriptsDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderGetRequestConfiguration)(ItemOnlinemeetingsItemTranscriptsDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlinemeetingsItemTranscriptsDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlinemeetingsItemTranscriptsDeltaGetResponseable), nil
}
// ToGetRequestInformation invoke function delta
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder when successful
func (m *ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder) {
    return NewItemOnlinemeetingsItemTranscriptsDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
