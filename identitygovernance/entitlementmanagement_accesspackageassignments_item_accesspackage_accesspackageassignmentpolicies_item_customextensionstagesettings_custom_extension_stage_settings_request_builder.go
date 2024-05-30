package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder provides operations to manage the customExtensionStageSettings property of the microsoft.graph.accessPackageAssignmentPolicy entity.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetQueryParameters the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomExtensionStageSettingId provides operations to manage the customExtensionStageSettings property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) ByCustomExtensionStageSettingId(customExtensionStageSettingId string)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customExtensionStageSettingId != "" {
        urlTplParams["customExtensionStageSetting%2Did"] = customExtensionStageSettingId
    }
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionStageSettings{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) Count()(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a CustomExtensionStageSettingCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionStageSettingCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomExtensionStageSettingCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionStageSettingCollectionResponseable), nil
}
// Post create new navigation property to customExtensionStageSettings for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a CustomExtensionStageSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionStageSettingable, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionStageSettingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomExtensionStageSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionStageSettingable), nil
}
// ToGetRequestInformation the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to customExtensionStageSettings for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionStageSettingable, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageassignmentpoliciesItemCustomextensionstagesettingsCustomExtensionStageSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
