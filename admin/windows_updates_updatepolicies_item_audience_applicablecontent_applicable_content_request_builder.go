package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder provides operations to manage the applicableContent property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderGetQueryParameters content eligible to deploy to devices in the audience. Not nullable. Read-only.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderGetQueryParameters
}
// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByApplicableContentCatalogEntryId provides operations to manage the applicableContent property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) ByApplicableContentCatalogEntryId(applicableContentCatalogEntryId string)(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if applicableContentCatalogEntryId != "" {
        urlTplParams["applicableContent%2DcatalogEntryId"] = applicableContentCatalogEntryId
    }
    return NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderInternal instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) {
    m := &WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/applicableContent{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentCountRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) Count()(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentCountRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get content eligible to deploy to devices in the audience. Not nullable. Read-only.
// returns a ApplicableContentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateApplicableContentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentCollectionResponseable), nil
}
// Post create new navigation property to applicableContent for admin
// returns a ApplicableContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) Post(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateApplicableContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable), nil
}
// ToGetRequestInformation content eligible to deploy to devices in the audience. Not nullable. Read-only.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to applicableContent for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) ToPostRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
