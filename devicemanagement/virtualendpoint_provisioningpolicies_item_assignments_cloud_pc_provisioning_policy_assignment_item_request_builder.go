package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.cloudPcProvisioningPolicy entity.
type VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderGetQueryParameters a defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments relationship, see Get cloudPcProvisioningPolicy.
type VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderGetQueryParameters
}
// VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedUsers provides operations to manage the assignedUsers property of the microsoft.graph.cloudPcProvisioningPolicyAssignment entity.
// returns a *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) AssignedUsers()(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderInternal instantiates a new VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder and sets the default values.
func NewVirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) {
    m := &VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy%2Did}/assignments/{cloudPcProvisioningPolicyAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder instantiates a new VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder and sets the default values.
func NewVirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments relationship, see Get cloudPcProvisioningPolicy.
// returns a CloudPcProvisioningPolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcProvisioningPolicyAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcProvisioningPolicyAssignmentable), nil
}
// Patch update the navigation property assignments in deviceManagement
// returns a CloudPcProvisioningPolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcProvisioningPolicyAssignmentable, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcProvisioningPolicyAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcProvisioningPolicyAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments relationship, see Get cloudPcProvisioningPolicy.
// returns a *RequestInformation when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcProvisioningPolicyAssignmentable, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsCloudPcProvisioningPolicyAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
