package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder provides operations to call the refresh method.
type EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceRole/accessPackageResource/refresh", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, refresh the accessPackageResource object to fetch the latest details for displayName, description, and resourceType from the origin system. For the AadApplication originSystem, this operation also updates the displayName and description for the accessPackageResourceRole. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageresource-refresh?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilderPostRequestConfiguration)(error) {
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
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageresourcerolescopesItemAccesspackageresourceroleAccesspackageresourceRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
