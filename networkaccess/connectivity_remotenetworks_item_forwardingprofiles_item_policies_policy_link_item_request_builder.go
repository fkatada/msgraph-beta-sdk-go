package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder provides operations to manage the policies property of the microsoft.graph.networkaccess.profile entity.
type ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderGetQueryParameters traffic forwarding policies associated with this profile.
type ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderGetQueryParameters
}
// ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderInternal instantiates a new ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) {
    m := &ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/forwardingProfiles/{forwardingProfile%2Did}/policies/{policyLink%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder instantiates a new ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property policies for networkAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get traffic forwarding policies associated with this profile.
// returns a PolicyLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreatePolicyLinkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable), nil
}
// Patch update the navigation property policies in networkAccess
// returns a PolicyLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreatePolicyLinkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.networkaccess.policyLink entity.
// returns a *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesItemPolicyRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) Policy()(*ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesItemPolicyRequestBuilder) {
    return NewConnectivityRemotenetworksItemForwardingprofilesItemPoliciesItemPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property policies for networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation traffic forwarding policies associated with this profile.
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property policies in networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) WithUrl(rawUrl string)(*ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder) {
    return NewConnectivityRemotenetworksItemForwardingprofilesItemPoliciesPolicyLinkItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
