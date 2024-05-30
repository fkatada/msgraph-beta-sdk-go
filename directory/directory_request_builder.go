package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryRequestBuilder provides operations to manage the directory singleton.
type DirectoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRequestBuilderGetQueryParameters get directory
type DirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRequestBuilderGetQueryParameters
}
// DirectoryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdministrativeUnits provides operations to manage the administrativeUnits property of the microsoft.graph.directory entity.
// returns a *AdministrativeunitsAdministrativeUnitsRequestBuilder when successful
func (m *DirectoryRequestBuilder) AdministrativeUnits()(*AdministrativeunitsAdministrativeUnitsRequestBuilder) {
    return NewAdministrativeunitsAdministrativeUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttributeSets provides operations to manage the attributeSets property of the microsoft.graph.directory entity.
// returns a *AttributesetsAttributeSetsRequestBuilder when successful
func (m *DirectoryRequestBuilder) AttributeSets()(*AttributesetsAttributeSetsRequestBuilder) {
    return NewAttributesetsAttributeSetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CertificateAuthorities provides operations to manage the certificateAuthorities property of the microsoft.graph.directory entity.
// returns a *CertificateauthoritiesCertificateAuthoritiesRequestBuilder when successful
func (m *DirectoryRequestBuilder) CertificateAuthorities()(*CertificateauthoritiesCertificateAuthoritiesRequestBuilder) {
    return NewCertificateauthoritiesCertificateAuthoritiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomSecurityAttributeDefinitions provides operations to manage the customSecurityAttributeDefinitions property of the microsoft.graph.directory entity.
// returns a *CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder when successful
func (m *DirectoryRequestBuilder) CustomSecurityAttributeDefinitions()(*CustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilder) {
    return NewCustomsecurityattributedefinitionsCustomSecurityAttributeDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeletedItems provides operations to manage the deletedItems property of the microsoft.graph.directory entity.
// returns a *DeleteditemsDeletedItemsRequestBuilder when successful
func (m *DirectoryRequestBuilder) DeletedItems()(*DeleteditemsDeletedItemsRequestBuilder) {
    return NewDeleteditemsDeletedItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceLocalCredentials provides operations to manage the deviceLocalCredentials property of the microsoft.graph.directory entity.
// returns a *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder when successful
func (m *DirectoryRequestBuilder) DeviceLocalCredentials()(*DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) {
    return NewDevicelocalcredentialsDeviceLocalCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalUserProfiles provides operations to manage the externalUserProfiles property of the microsoft.graph.directory entity.
// returns a *ExternaluserprofilesExternalUserProfilesRequestBuilder when successful
func (m *DirectoryRequestBuilder) ExternalUserProfiles()(*ExternaluserprofilesExternalUserProfilesRequestBuilder) {
    return NewExternaluserprofilesExternalUserProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FeatureRolloutPolicies provides operations to manage the featureRolloutPolicies property of the microsoft.graph.directory entity.
// returns a *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder when successful
func (m *DirectoryRequestBuilder) FeatureRolloutPolicies()(*FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) {
    return NewFeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FederationConfigurations provides operations to manage the federationConfigurations property of the microsoft.graph.directory entity.
// returns a *FederationconfigurationsFederationConfigurationsRequestBuilder when successful
func (m *DirectoryRequestBuilder) FederationConfigurations()(*FederationconfigurationsFederationConfigurationsRequestBuilder) {
    return NewFederationconfigurationsFederationConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get directory
// returns a Directoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable), nil
}
// ImpactedResources provides operations to manage the impactedResources property of the microsoft.graph.directory entity.
// returns a *ImpactedresourcesImpactedResourcesRequestBuilder when successful
func (m *DirectoryRequestBuilder) ImpactedResources()(*ImpactedresourcesImpactedResourcesRequestBuilder) {
    return NewImpactedresourcesImpactedResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InboundSharedUserProfiles provides operations to manage the inboundSharedUserProfiles property of the microsoft.graph.directory entity.
// returns a *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder when successful
func (m *DirectoryRequestBuilder) InboundSharedUserProfiles()(*InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) {
    return NewInboundshareduserprofilesInboundSharedUserProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OnPremisesSynchronization provides operations to manage the onPremisesSynchronization property of the microsoft.graph.directory entity.
// returns a *OnpremisessynchronizationOnPremisesSynchronizationRequestBuilder when successful
func (m *DirectoryRequestBuilder) OnPremisesSynchronization()(*OnpremisessynchronizationOnPremisesSynchronizationRequestBuilder) {
    return NewOnpremisessynchronizationOnPremisesSynchronizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OutboundSharedUserProfiles provides operations to manage the outboundSharedUserProfiles property of the microsoft.graph.directory entity.
// returns a *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder when successful
func (m *DirectoryRequestBuilder) OutboundSharedUserProfiles()(*OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) {
    return NewOutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update directory
// returns a Directoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable, requestConfiguration *DirectoryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable), nil
}
// PendingExternalUserProfiles provides operations to manage the pendingExternalUserProfiles property of the microsoft.graph.directory entity.
// returns a *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder when successful
func (m *DirectoryRequestBuilder) PendingExternalUserProfiles()(*PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) {
    return NewPendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Recommendations provides operations to manage the recommendations property of the microsoft.graph.directory entity.
// returns a *RecommendationsRequestBuilder when successful
func (m *DirectoryRequestBuilder) Recommendations()(*RecommendationsRequestBuilder) {
    return NewRecommendationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedEmailDomains provides operations to manage the sharedEmailDomains property of the microsoft.graph.directory entity.
// returns a *SharedemaildomainsSharedEmailDomainsRequestBuilder when successful
func (m *DirectoryRequestBuilder) SharedEmailDomains()(*SharedemaildomainsSharedEmailDomainsRequestBuilder) {
    return NewSharedemaildomainsSharedEmailDomainsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.directory entity.
// returns a *SubscriptionsRequestBuilder when successful
func (m *DirectoryRequestBuilder) Subscriptions()(*SubscriptionsRequestBuilder) {
    return NewSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SubscriptionsWithCommerceSubscriptionId provides operations to manage the subscriptions property of the microsoft.graph.directory entity.
// returns a *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder when successful
func (m *DirectoryRequestBuilder) SubscriptionsWithCommerceSubscriptionId(commerceSubscriptionId *string)(*SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) {
    return NewSubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, commerceSubscriptionId)
}
// SubscriptionsWithOcpSubscriptionId provides operations to manage the subscriptions property of the microsoft.graph.directory entity.
// returns a *SubscriptionswithocpsubscriptionidSubscriptionsWithOcpSubscriptionIdRequestBuilder when successful
func (m *DirectoryRequestBuilder) SubscriptionsWithOcpSubscriptionId(ocpSubscriptionId *string)(*SubscriptionswithocpsubscriptionidSubscriptionsWithOcpSubscriptionIdRequestBuilder) {
    return NewSubscriptionswithocpsubscriptionidSubscriptionsWithOcpSubscriptionIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, ocpSubscriptionId)
}
// ToGetRequestInformation get directory
// returns a *RequestInformation when successful
func (m *DirectoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update directory
// returns a *RequestInformation when successful
func (m *DirectoryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable, requestConfiguration *DirectoryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRequestBuilder when successful
func (m *DirectoryRequestBuilder) WithUrl(rawUrl string)(*DirectoryRequestBuilder) {
    return NewDirectoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
