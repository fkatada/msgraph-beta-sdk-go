package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderGetQueryParameters policy assignments
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderGetQueryParameters struct {
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
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderGetQueryParameters
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementConfigurationPolicyAssignmentId provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) ByDeviceManagementConfigurationPolicyAssignmentId(deviceManagementConfigurationPolicyAssignmentId string)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementConfigurationPolicyAssignmentId != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment%2Did"] = deviceManagementConfigurationPolicyAssignmentId
    }
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderInternal instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) {
    m := &ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/assignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsCountRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) Count()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsCountRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get policy assignments
// returns a DeviceManagementConfigurationPolicyAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentCollectionResponseable), nil
}
// Post create new navigation property to assignments for deviceManagement
// returns a DeviceManagementConfigurationPolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable), nil
}
// ToGetRequestInformation policy assignments
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) WithUrl(rawUrl string)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
