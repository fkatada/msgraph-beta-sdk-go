package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageResourceEnvironment entity.
type EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderGetQueryParameters read-only. Required.
type EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder) {
    m := &EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/{accessPackageResourceEnvironment%2Did}/accessPackageResources/{accessPackageResource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder instantiates a new EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only. Required.
// returns a AccessPackageResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
// ToGetRequestInformation read-only. Required.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
