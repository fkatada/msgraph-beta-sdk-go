package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder provides operations to manage the agentGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
type ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderGetQueryParameters list of existing onPremisesAgentGroup objects. Read-only. Nullable.
type ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderGetQueryParameters
}
// ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Agents provides operations to manage the agents property of the microsoft.graph.onPremisesAgentGroup entity.
// returns a *ItemAgentgroupsItemAgentsRequestBuilder when successful
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) Agents()(*ItemAgentgroupsItemAgentsRequestBuilder) {
    return NewItemAgentgroupsItemAgentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property agentGroups for onPremisesPublishingProfiles
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of existing onPremisesAgentGroup objects. Read-only. Nullable.
// returns a OnPremisesAgentGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesAgentGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable), nil
}
// Patch update the navigation property agentGroups in onPremisesPublishingProfiles
// returns a OnPremisesAgentGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable, requestConfiguration *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesAgentGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable), nil
}
// PublishedResources provides operations to manage the publishedResources property of the microsoft.graph.onPremisesAgentGroup entity.
// returns a *ItemAgentgroupsItemPublishedresourcesPublishedResourcesRequestBuilder when successful
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) PublishedResources()(*ItemAgentgroupsItemPublishedresourcesPublishedResourcesRequestBuilder) {
    return NewItemAgentgroupsItemPublishedresourcesPublishedResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property agentGroups for onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of existing onPremisesAgentGroup objects. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property agentGroups in onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable, requestConfiguration *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder when successful
func (m *ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) WithUrl(rawUrl string)(*ItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    return NewItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
