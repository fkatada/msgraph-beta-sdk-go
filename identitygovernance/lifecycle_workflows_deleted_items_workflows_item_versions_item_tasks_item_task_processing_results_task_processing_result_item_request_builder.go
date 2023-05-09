package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.task entity.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters the result of processing the task.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal instantiates a new TaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}/tasks/{task%2Did}/taskProcessingResults/{taskProcessingResult%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder instantiates a new TaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the result of processing the task.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskProcessingResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateTaskProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskProcessingResultable), nil
}
// MicrosoftGraphIdentityGovernanceResume provides operations to call the resume method.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) MicrosoftGraphIdentityGovernanceResume()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subject provides operations to manage the subject property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Subject()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Task provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Task()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemTaskRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsItemTaskRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the result of processing the task.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
