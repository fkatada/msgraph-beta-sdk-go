package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder provides operations to manage the policy property of the microsoft.graph.networkaccess.policyLink entity.
type ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetQueryParameters policy.
type ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetQueryParameters
}
// NewConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderInternal instantiates a new ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder and sets the default values.
func NewConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder) {
    m := &ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/branches/{branchSite%2Did}/forwardingProfiles/{forwardingProfile%2Did}/policies/{policyLink%2Did}/policy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder instantiates a new ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder and sets the default values.
func NewConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get policy.
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a Policyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Policyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreatePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Policyable), nil
}
// ToGetRequestInformation policy.
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a *RequestInformation when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a *ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder) WithUrl(rawUrl string)(*ConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder) {
    return NewConnectivityBranchesItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
