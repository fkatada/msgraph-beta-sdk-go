package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder provides operations to call the moveToCatalog method.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/moveToCatalog", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, this action moves the accessPackage to a specified target accessPackageCatalog. The resources in the access package must be present in the target catalog.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-movetocatalog?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder) Post(ctx context.Context, body EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogPostRequestBodyable, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation in Microsoft Entra entitlement management, this action moves the accessPackage to a specified target accessPackageCatalog. The resources in the access package must be present in the target catalog.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder) ToPostRequestInformation(ctx context.Context, body EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogPostRequestBodyable, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageMovetocatalogMoveToCatalogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
