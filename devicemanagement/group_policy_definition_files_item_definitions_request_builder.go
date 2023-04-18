package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder provides operations to manage the definitions property of the microsoft.graph.groupPolicyDefinitionFile entity.
type GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyDefinitionFilesItemDefinitionsRequestBuilderGetQueryParameters the group policy definitions associated with the file.
type GroupPolicyDefinitionFilesItemDefinitionsRequestBuilderGetQueryParameters struct {
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
// GroupPolicyDefinitionFilesItemDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionFilesItemDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionFilesItemDefinitionsRequestBuilderGetQueryParameters
}
// ByGroupPolicyDefinitionId provides operations to manage the definitions property of the microsoft.graph.groupPolicyDefinitionFile entity.
func (m *GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder) ByGroupPolicyDefinitionId(groupPolicyDefinitionId string)(*GroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyDefinitionId != "" {
        urlTplParams["groupPolicyDefinition%2Did"] = groupPolicyDefinitionId
    }
    return NewGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupPolicyDefinitionFilesItemDefinitionsRequestBuilderInternal instantiates a new DefinitionsRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionFilesItemDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder) {
    m := &GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitionFiles/{groupPolicyDefinitionFile%2Did}/definitions{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewGroupPolicyDefinitionFilesItemDefinitionsRequestBuilder instantiates a new DefinitionsRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionFilesItemDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionFilesItemDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder) Count()(*GroupPolicyDefinitionFilesItemDefinitionsCountRequestBuilder) {
    return NewGroupPolicyDefinitionFilesItemDefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the group policy definitions associated with the file.
func (m *GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyDefinitionFilesItemDefinitionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionCollectionResponseable), nil
}
// ToGetRequestInformation the group policy definitions associated with the file.
func (m *GroupPolicyDefinitionFilesItemDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionFilesItemDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
