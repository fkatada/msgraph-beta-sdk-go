package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder provides operations to manage the roleAssignmentRequests property of the microsoft.graph.privilegedAccess entity.
type ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters a collection of role assignment requests for the provider.
type ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters
}
// ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
// returns a *ItemRoleassignmentrequestsItemCancelRequestBuilder when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Cancel()(*ItemRoleassignmentrequestsItemCancelRequestBuilder) {
    return NewItemRoleassignmentrequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal instantiates a new ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    m := &ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/roleAssignmentRequests/{governanceRoleAssignmentRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder instantiates a new ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentRequests for privilegedAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of role assignment requests for the provider.
// returns a GovernanceRoleAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// Patch update the navigation property roleAssignmentRequests in privilegedAccess
// returns a GovernanceRoleAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// Resource provides operations to manage the resource property of the microsoft.graph.governanceRoleAssignmentRequest entity.
// returns a *ItemRoleassignmentrequestsItemResourceRequestBuilder when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Resource()(*ItemRoleassignmentrequestsItemResourceRequestBuilder) {
    return NewItemRoleassignmentrequestsItemResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.governanceRoleAssignmentRequest entity.
// returns a *ItemRoleassignmentrequestsItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) RoleDefinition()(*ItemRoleassignmentrequestsItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewItemRoleassignmentrequestsItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subject provides operations to manage the subject property of the microsoft.graph.governanceRoleAssignmentRequest entity.
// returns a *ItemRoleassignmentrequestsItemSubjectRequestBuilder when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Subject()(*ItemRoleassignmentrequestsItemSubjectRequestBuilder) {
    return NewItemRoleassignmentrequestsItemSubjectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentRequests for privilegedAccess
// returns a *RequestInformation when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of role assignment requests for the provider.
// returns a *RequestInformation when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignmentRequests in privilegedAccess
// returns a *RequestInformation when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateRequest provides operations to call the updateRequest method.
// returns a *ItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) UpdateRequest()(*ItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder) {
    return NewItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder when successful
func (m *ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) WithUrl(rawUrl string)(*ItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    return NewItemRoleassignmentrequestsGovernanceRoleAssignmentRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
