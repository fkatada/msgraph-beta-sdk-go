package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder provides operations to manage the childTags property of the microsoft.graph.ediscovery.tag entity.
type EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderGetQueryParameters returns the tags that are a child of a tag.
type EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderGetQueryParameters struct {
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
// EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderGetQueryParameters
}
// ByTagId1 provides operations to manage the childTags property of the microsoft.graph.ediscovery.tag entity.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemTagsItemChildtagsTagItemRequestBuilder when successful
func (m *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder) ByTagId1(tagId1 string)(*EdiscoveryCasesItemTagsItemChildtagsTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tagId1 != "" {
        urlTplParams["tag%2Did1"] = tagId1
    }
    return NewEdiscoveryCasesItemTagsItemChildtagsTagItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderInternal instantiates a new EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder) {
    m := &EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/tags/{tag%2Did}/childTags{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder instantiates a new EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EdiscoveryCasesItemTagsItemChildtagsCountRequestBuilder when successful
func (m *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder) Count()(*EdiscoveryCasesItemTagsItemChildtagsCountRequestBuilder) {
    return NewEdiscoveryCasesItemTagsItemChildtagsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns the tags that are a child of a tag.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a TagCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.TagCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateTagCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.TagCollectionResponseable), nil
}
// ToGetRequestInformation returns the tags that are a child of a tag.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder when successful
func (m *EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder) {
    return NewEdiscoveryCasesItemTagsItemChildtagsChildTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
