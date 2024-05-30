package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder provides operations to call the restore method.
type LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/microsoft.graph.identityGovernance.restore", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a workflow that has been deleted. You can only restore a workflow that was deleted within the last 30 days before Microsoft Entra ID automatically permanently deletes it.
// returns a Workflowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-restore?view=graph-rest-beta
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) Post(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderPostRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable), nil
}
// ToPostRequestInformation restore a workflow that has been deleted. You can only restore a workflow that was deleted within the last 30 days before Microsoft Entra ID automatically permanently deletes it.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
