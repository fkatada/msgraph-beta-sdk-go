package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder provides operations to call the retrieveJustInTimeConfiguration method.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilderInternal instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder) {
    m := &ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/retrieveJustInTimeConfiguration", pathParameters),
    }
    return m
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action retrieveJustInTimeConfiguration
// returns a DeviceManagementConfigurationJustInTimeAssignmentPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder) Post(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationJustInTimeAssignmentPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationJustInTimeAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationJustInTimeAssignmentPolicyable), nil
}
// ToPostRequestInformation invoke action retrieveJustInTimeConfiguration
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder) WithUrl(rawUrl string)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievejustintimeconfigurationRetrieveJustInTimeConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
