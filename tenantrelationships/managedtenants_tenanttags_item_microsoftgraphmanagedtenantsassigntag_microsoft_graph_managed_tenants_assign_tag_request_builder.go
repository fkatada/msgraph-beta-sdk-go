package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder provides operations to call the assignTag method.
type ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilderInternal instantiates a new ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder and sets the default values.
func NewManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder) {
    m := &ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenantTags/{tenantTag%2Did}/microsoft.graph.managedTenants.assignTag", pathParameters),
    }
    return m
}
// NewManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder instantiates a new ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder and sets the default values.
func NewManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Post assign the tenant tag to the specified managed tenants.
// returns a TenantTagable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-tenanttag-assigntag?view=graph-rest-beta
func (m *ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder) Post(ctx context.Context, body ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagAssignTagPostRequestBodyable, requestConfiguration *ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantTagable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantTagable), nil
}
// ToPostRequestInformation assign the tenant tag to the specified managed tenants.
// returns a *RequestInformation when successful
func (m *ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagAssignTagPostRequestBodyable, requestConfiguration *ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder when successful
func (m *ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder) {
    return NewManagedtenantsTenanttagsItemMicrosoftgraphmanagedtenantsassigntagMicrosoftGraphManagedTenantsAssignTagRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
