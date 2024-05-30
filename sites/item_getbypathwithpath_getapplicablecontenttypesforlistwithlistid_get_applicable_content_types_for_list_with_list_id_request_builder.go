package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder provides operations to call the getApplicableContentTypesForList method.
type ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetQueryParameters get site contentTypes that can be added to a list.
type ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetQueryParameters struct {
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
// ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetQueryParameters
}
// NewItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderInternal instantiates a new ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder and sets the default values.
func NewItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, listId *string)(*ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) {
    m := &ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/getByPath(path='{path}')/getApplicableContentTypesForList(listId='{listId}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if listId != nil {
        m.BaseRequestBuilder.PathParameters["listId"] = *listId
    }
    return m
}
// NewItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder instantiates a new ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder and sets the default values.
func NewItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get site contentTypes that can be added to a list.
// Deprecated: This method is obsolete. Use GetAsGetApplicableContentTypesForListWithListIdGetResponse instead.
// returns a ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-getapplicablecontenttypesforlist?view=graph-rest-beta
func (m *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration)(ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseable), nil
}
// GetAsGetApplicableContentTypesForListWithListIdGetResponse get site contentTypes that can be added to a list.
// returns a ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-getapplicablecontenttypesforlist?view=graph-rest-beta
func (m *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) GetAsGetApplicableContentTypesForListWithListIdGetResponse(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration)(ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable), nil
}
// ToGetRequestInformation get site contentTypes that can be added to a list.
// returns a *RequestInformation when successful
func (m *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder when successful
func (m *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) WithUrl(rawUrl string)(*ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
