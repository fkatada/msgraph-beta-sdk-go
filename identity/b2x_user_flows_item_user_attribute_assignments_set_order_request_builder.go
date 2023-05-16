package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder provides operations to call the setOrder method.
type B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderInternal instantiates a new SetOrderRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder) {
    m := &B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userAttributeAssignments/setOrder", pathParameters),
    }
    return m
}
// NewB2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder instantiates a new SetOrderRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the order of identityUserFlowAttributeAssignments being collected within a user flow.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/identityuserflowattributeassignment-setorder?view=graph-rest-1.0
func (m *B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder) Post(ctx context.Context, body B2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyable, requestConfiguration *B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation set the order of identityUserFlowAttributeAssignments being collected within a user flow.
func (m *B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder) ToPostRequestInformation(ctx context.Context, body B2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyable, requestConfiguration *B2xUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
