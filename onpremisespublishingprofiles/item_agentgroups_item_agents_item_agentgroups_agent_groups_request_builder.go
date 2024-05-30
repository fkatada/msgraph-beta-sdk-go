package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder provides operations to manage the agentGroups property of the microsoft.graph.onPremisesAgent entity.
type ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderGetQueryParameters list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
type ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderGetQueryParameters struct {
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
// ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderGetQueryParameters
}
// ByOnPremisesAgentGroupId1 gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agentGroups.item.agents.item.agentGroups.item collection
// returns a *ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder when successful
func (m *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) ByOnPremisesAgentGroupId1(onPremisesAgentGroupId1 string)(*ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if onPremisesAgentGroupId1 != "" {
        urlTplParams["onPremisesAgentGroup%2Did1"] = onPremisesAgentGroupId1
    }
    return NewItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderInternal instantiates a new ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder and sets the default values.
func NewItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) {
    m := &ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/agents/{onPremisesAgent%2Did}/agentGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder instantiates a new ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder and sets the default values.
func NewItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAgentgroupsItemAgentsItemAgentgroupsCountRequestBuilder when successful
func (m *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) Count()(*ItemAgentgroupsItemAgentsItemAgentgroupsCountRequestBuilder) {
    return NewItemAgentgroupsItemAgentsItemAgentgroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
// returns a OnPremisesAgentGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesAgentGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupCollectionResponseable), nil
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemAgentgroupsItemAgentsItemAgentgroupsRefRequestBuilder when successful
func (m *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) Ref()(*ItemAgentgroupsItemAgentsItemAgentgroupsRefRequestBuilder) {
    return NewItemAgentgroupsItemAgentsItemAgentgroupsRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder when successful
func (m *ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder) {
    return NewItemAgentgroupsItemAgentsItemAgentgroupsAgentGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
