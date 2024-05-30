package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder provides operations to manage the customExtension property of the microsoft.graph.customExtensionStageSetting entity.
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters indicates the custom workflow extension that is executed at this stage. Nullable. Supports $expand.
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionStageSettings/{customExtensionStageSetting%2Did}/customExtension{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get indicates the custom workflow extension that is executed at this stage. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a CustomCalloutExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomCalloutExtensionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomCalloutExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomCalloutExtensionable), nil
}
// ToGetRequestInformation indicates the custom workflow extension that is executed at this stage. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionstagesettingsItemCustomextensionCustomExtensionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
