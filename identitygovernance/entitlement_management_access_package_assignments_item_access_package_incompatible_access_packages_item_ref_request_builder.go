package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder provides operations to manage the collection of identityGovernance entities.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/incompatibleAccessPackages/{accessPackage%2Did}/$ref", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove an access package from the list of access packages that have been marked as incompatible on an accessPackage.  
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-delete-incompatibleaccesspackage?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// ToDeleteRequestInformation remove an access package from the list of access packages that have been marked as incompatible on an accessPackage.  
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageIncompatibleAccessPackagesItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
