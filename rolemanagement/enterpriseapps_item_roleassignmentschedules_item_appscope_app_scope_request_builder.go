package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
type EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetQueryParameters read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
type EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetQueryParameters
}
// NewEnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderInternal instantiates a new EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder and sets the default values.
func NewEnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) {
    m := &EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleAssignmentSchedules/{unifiedRoleAssignmentSchedule%2Did}/appScope{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder instantiates a new EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder and sets the default values.
func NewEnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
// returns a AppScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppScopeable), nil
}
// ToGetRequestInformation read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
// returns a *RequestInformation when successful
func (m *EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder when successful
func (m *EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) WithUrl(rawUrl string)(*EnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) {
    return NewEnterpriseappsItemRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
