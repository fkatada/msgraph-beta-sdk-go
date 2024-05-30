package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderGetQueryParameters represents the aggregation of task execution data for tasks within a workflow object.
type LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents the aggregation of task execution data for tasks within a workflow object.
// returns a TaskReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskReportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateTaskReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskReportable), nil
}
// Task provides operations to manage the task property of the microsoft.graph.identityGovernance.taskReport entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) Task()(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskDefinition provides operations to manage the taskDefinition property of the microsoft.graph.identityGovernance.taskReport entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) TaskDefinition()(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskProcessingResults provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.taskReport entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) TaskProcessingResults()(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation represents the aggregation of task execution data for tasks within a workflow object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
