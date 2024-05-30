package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder provides operations to call the refresh method.
type EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageResourceScope/accessPackageResource/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource/refresh", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, refresh the accessPackageResource object to fetch the latest details for displayName, description, and resourceType from the origin system. For the AadApplication originSystem, this operation also updates the displayName and description for the accessPackageResourceRole. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageresource-refresh?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation in Microsoft Entra entitlement management, refresh the accessPackageResource object to fetch the latest details for displayName, description, and resourceType from the origin system. For the AadApplication originSystem, this operation also updates the displayName and description for the accessPackageResourceRole. 
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesItemAccesspackageresourcescopeAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
