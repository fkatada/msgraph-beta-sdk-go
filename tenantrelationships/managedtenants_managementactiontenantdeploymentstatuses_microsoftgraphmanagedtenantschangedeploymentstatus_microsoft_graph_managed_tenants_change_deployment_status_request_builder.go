package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder provides operations to call the changeDeploymentStatus method.
type ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal instantiates a new ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) {
    m := &ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementActionTenantDeploymentStatuses/microsoft.graph.managedTenants.changeDeploymentStatus", pathParameters),
    }
    return m
}
// NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder instantiates a new ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post changes the tenant level deployment status for the management action. This information is used to provide insights into what management actions are in a specific state. As example there might be a plan to apply the require multi-factor authentication for admins, so it would be ideal to change the status to planned to reflect the appropriate status.
// returns a ManagementActionDeploymentStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) Post(ctx context.Context, body ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionDeploymentStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementActionDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionDeploymentStatusable), nil
}
// ToPostRequestInformation changes the tenant level deployment status for the management action. This information is used to provide insights into what management actions are in a specific state. As example there might be a plan to apply the require multi-factor authentication for admins, so it would be ideal to change the status to planned to reflect the appropriate status.
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) {
    return NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
