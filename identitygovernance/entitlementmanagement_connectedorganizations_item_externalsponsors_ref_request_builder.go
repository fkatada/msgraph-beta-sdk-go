package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder provides operations to manage the collection of identityGovernance entities.
type EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderDeleteQueryParameters remove a user or a group from the connected organization's external sponsors. The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
type EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderDeleteQueryParameters struct {
    // The delete Uri
    Id *string `uriparametername:"%40id"`
}
// EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderDeleteQueryParameters
}
// EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderGetQueryParameters retrieve a list of a connectedOrganization's external sponsors.  The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
type EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderGetQueryParameters
}
// EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderInternal instantiates a new EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder and sets the default values.
func NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) {
    m := &EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/externalSponsors/$ref?@id={%40id}{&%24count,%24filter,%24orderby,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder instantiates a new EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder and sets the default values.
func NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a user or a group from the connected organization's external sponsors. The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectedorganization-delete-externalsponsors?view=graph-rest-beta
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve a list of a connectedOrganization's external sponsors.  The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a StringCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectedorganization-list-externalsponsors?view=graph-rest-beta
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.StringCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateStringCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.StringCollectionResponseable), nil
}
// Post add a user or a group to the connected organization's external sponsors. The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectedorganization-post-externalsponsors?view=graph-rest-beta
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceCreateable, requestConfiguration *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation remove a user or a group from the connected organization's external sponsors. The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/externalSponsors/$ref?@id={%40id}", m.BaseRequestBuilder.PathParameters)
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
// ToGetRequestInformation retrieve a list of a connectedOrganization's external sponsors.  The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/externalSponsors/$ref{?%24count,%24filter,%24orderby,%24search,%24skip,%24top}", m.BaseRequestBuilder.PathParameters)
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
// ToPostRequestInformation add a user or a group to the connected organization's external sponsors. The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceCreateable, requestConfiguration *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/externalSponsors/$ref", m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder when successful
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder) {
    return NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
