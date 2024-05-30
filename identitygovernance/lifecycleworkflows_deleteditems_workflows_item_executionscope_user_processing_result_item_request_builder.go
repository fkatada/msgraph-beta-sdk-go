package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderGetQueryParameters the unique identifier of the Microsoft Entra identity that last modified the workflow object.
type LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/executionScope/{userProcessingResult%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the unique identifier of the Microsoft Entra identity that last modified the workflow object.
// returns a UserProcessingResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateUserProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable), nil
}
// ToGetRequestInformation the unique identifier of the Microsoft Entra identity that last modified the workflow object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
