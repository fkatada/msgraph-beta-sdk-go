package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder provides operations to manage the dependentResources property of the microsoft.graph.educationAssignmentResource entity.
type MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderGetQueryParameters get dependentResources from education
type MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderGetQueryParameters struct {
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
// MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderGetQueryParameters
}
// MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEducationAssignmentResourceId1 provides operations to manage the dependentResources property of the microsoft.graph.educationAssignmentResource entity.
// returns a *MeAssignmentsItemResourcesItemDependentresourcesEducationAssignmentResourceItemRequestBuilder when successful
func (m *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) ByEducationAssignmentResourceId1(educationAssignmentResourceId1 string)(*MeAssignmentsItemResourcesItemDependentresourcesEducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if educationAssignmentResourceId1 != "" {
        urlTplParams["educationAssignmentResource%2Did1"] = educationAssignmentResourceId1
    }
    return NewMeAssignmentsItemResourcesItemDependentresourcesEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderInternal instantiates a new MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder and sets the default values.
func NewMeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) {
    m := &MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/resources/{educationAssignmentResource%2Did}/dependentResources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder instantiates a new MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder and sets the default values.
func NewMeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MeAssignmentsItemResourcesItemDependentresourcesCountRequestBuilder when successful
func (m *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) Count()(*MeAssignmentsItemResourcesItemDependentresourcesCountRequestBuilder) {
    return NewMeAssignmentsItemResourcesItemDependentresourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get dependentResources from education
// returns a EducationAssignmentResourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentResourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentResourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentResourceCollectionResponseable), nil
}
// Post create new navigation property to dependentResources for education
// returns a EducationAssignmentResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentResourceable, requestConfiguration *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentResourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentResourceable), nil
}
// ToGetRequestInformation get dependentResources from education
// returns a *RequestInformation when successful
func (m *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dependentResources for education
// returns a *RequestInformation when successful
func (m *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentResourceable, requestConfiguration *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder when successful
func (m *MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) WithUrl(rawUrl string)(*MeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder) {
    return NewMeAssignmentsItemResourcesItemDependentresourcesDependentResourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
