package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder provides operations to manage the customExtensionHandlers property of the microsoft.graph.accessPackageAssignmentPolicy entity.
type EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderGetQueryParameters the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
type EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomExtensionHandlerId provides operations to manage the customExtensionHandlers property of the microsoft.graph.accessPackageAssignmentPolicy entity.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlerItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) ByCustomExtensionHandlerId(customExtensionHandlerId string)(*EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customExtensionHandlerId != "" {
        urlTplParams["customExtensionHandler%2Did"] = customExtensionHandlerId
    }
    return NewEntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionHandlers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) Count()(*EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a CustomExtensionHandlerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionHandlerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomExtensionHandlerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionHandlerCollectionResponseable), nil
}
// Post create new navigation property to customExtensionHandlers for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a CustomExtensionHandlerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionHandlerable, requestConfiguration *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionHandlerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomExtensionHandlerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionHandlerable), nil
}
// ToGetRequestInformation the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to customExtensionHandlers for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionHandlerable, requestConfiguration *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentpoliciesItemCustomextensionhandlersCustomExtensionHandlersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
