package contacts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDirectreportsDirectReportsRequestBuilder provides operations to manage the directReports property of the microsoft.graph.orgContact entity.
type ItemDirectreportsDirectReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDirectreportsDirectReportsRequestBuilderGetQueryParameters get the contact's direct reports.
type ItemDirectreportsDirectReportsRequestBuilderGetQueryParameters struct {
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
// ItemDirectreportsDirectReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDirectreportsDirectReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDirectreportsDirectReportsRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the directReports property of the microsoft.graph.orgContact entity.
// returns a *ItemDirectreportsDirectoryObjectItemRequestBuilder when successful
func (m *ItemDirectreportsDirectReportsRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemDirectreportsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemDirectreportsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDirectreportsDirectReportsRequestBuilderInternal instantiates a new ItemDirectreportsDirectReportsRequestBuilder and sets the default values.
func NewItemDirectreportsDirectReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDirectreportsDirectReportsRequestBuilder) {
    m := &ItemDirectreportsDirectReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/contacts/{orgContact%2Did}/directReports{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemDirectreportsDirectReportsRequestBuilder instantiates a new ItemDirectreportsDirectReportsRequestBuilder and sets the default values.
func NewItemDirectreportsDirectReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDirectreportsDirectReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDirectreportsDirectReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemDirectreportsCountRequestBuilder when successful
func (m *ItemDirectreportsDirectReportsRequestBuilder) Count()(*ItemDirectreportsCountRequestBuilder) {
    return NewItemDirectreportsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the contact's direct reports.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/orgcontact-list-directreports?view=graph-rest-beta
func (m *ItemDirectreportsDirectReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDirectreportsDirectReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *ItemDirectreportsGraphorgcontactGraphOrgContactRequestBuilder when successful
func (m *ItemDirectreportsDirectReportsRequestBuilder) GraphOrgContact()(*ItemDirectreportsGraphorgcontactGraphOrgContactRequestBuilder) {
    return NewItemDirectreportsGraphorgcontactGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemDirectreportsGraphuserGraphUserRequestBuilder when successful
func (m *ItemDirectreportsDirectReportsRequestBuilder) GraphUser()(*ItemDirectreportsGraphuserGraphUserRequestBuilder) {
    return NewItemDirectreportsGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the contact's direct reports.
// returns a *RequestInformation when successful
func (m *ItemDirectreportsDirectReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDirectreportsDirectReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDirectreportsDirectReportsRequestBuilder when successful
func (m *ItemDirectreportsDirectReportsRequestBuilder) WithUrl(rawUrl string)(*ItemDirectreportsDirectReportsRequestBuilder) {
    return NewItemDirectreportsDirectReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
