package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder provides operations to manage the complianceChanges property of the microsoft.graph.windowsUpdates.updatePolicy entity.
type WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderGetQueryParameters get a list of the complianceChange objects and their properties.
type WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderGetQueryParameters
}
// WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByComplianceChangeId provides operations to manage the complianceChanges property of the microsoft.graph.windowsUpdates.updatePolicy entity.
// returns a *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangeItemRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) ByComplianceChangeId(complianceChangeId string)(*WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if complianceChangeId != "" {
        urlTplParams["complianceChange%2Did"] = complianceChangeId
    }
    return NewWindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderInternal instantiates a new WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) {
    m := &WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/complianceChanges{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder instantiates a new WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsUpdatesUpdatepoliciesItemCompliancechangesCountRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) Count()(*WindowsUpdatesUpdatepoliciesItemCompliancechangesCountRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemCompliancechangesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the complianceChange objects and their properties.
// returns a ComplianceChangeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatepolicy-list-compliancechanges?view=graph-rest-beta
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ComplianceChangeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateComplianceChangeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ComplianceChangeCollectionResponseable), nil
}
// Post create a new contentApproval object.
// returns a ComplianceChangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatepolicy-post-compliancechanges-contentapproval?view=graph-rest-beta
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) Post(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ComplianceChangeable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ComplianceChangeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateComplianceChangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ComplianceChangeable), nil
}
// ToGetRequestInformation get a list of the complianceChange objects and their properties.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new contentApproval object.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ComplianceChangeable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemCompliancechangesComplianceChangesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
