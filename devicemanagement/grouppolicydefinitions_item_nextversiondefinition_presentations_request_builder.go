package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
type GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderGetQueryParameters the group policy presentations associated with the definition.
type GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderGetQueryParameters struct {
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
// GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderGetQueryParameters
}
// GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGroupPolicyPresentationId provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GrouppolicydefinitionsItemNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) ByGroupPolicyPresentationId(groupPolicyPresentationId string)(*GrouppolicydefinitionsItemNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyPresentationId != "" {
        urlTplParams["groupPolicyPresentation%2Did"] = groupPolicyPresentationId
    }
    return NewGrouppolicydefinitionsItemNextversiondefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderInternal instantiates a new GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder and sets the default values.
func NewGrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) {
    m := &GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/nextVersionDefinition/presentations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder instantiates a new GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder and sets the default values.
func NewGrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GrouppolicydefinitionsItemNextversiondefinitionPresentationsCountRequestBuilder when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) Count()(*GrouppolicydefinitionsItemNextversiondefinitionPresentationsCountRequestBuilder) {
    return NewGrouppolicydefinitionsItemNextversiondefinitionPresentationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the group policy presentations associated with the definition.
// returns a GroupPolicyPresentationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationCollectionResponseable), nil
}
// Post create new navigation property to presentations for deviceManagement
// returns a GroupPolicyPresentationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable), nil
}
// ToGetRequestInformation the group policy presentations associated with the definition.
// returns a *RequestInformation when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to presentations for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder when successful
func (m *GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) WithUrl(rawUrl string)(*GrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder) {
    return NewGrouppolicydefinitionsItemNextversiondefinitionPresentationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
