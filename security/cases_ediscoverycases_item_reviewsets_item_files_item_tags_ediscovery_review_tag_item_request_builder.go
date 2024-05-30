package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder provides operations to manage the tags property of the microsoft.graph.security.ediscoveryFile entity.
type CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters tags associated with the file.
type CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters
}
// NewCasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/files/{ediscoveryFile%2Did}/tags/{ediscoveryReviewTag%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder instantiates a new CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get tags associated with the file.
// returns a EdiscoveryReviewTagable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable), nil
}
// ToGetRequestInformation tags associated with the file.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemFilesItemTagsEdiscoveryReviewTagItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
