package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
type MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderGetQueryParameters list of relationships for this mobile app.
type MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderGetQueryParameters struct {
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
// MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileAppRelationshipId provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemGraphManagedMobileLobAppRelationshipsMobileAppRelationshipItemRequestBuilder when successful
func (m *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) ByMobileAppRelationshipId(mobileAppRelationshipId string)(*MobileAppsItemGraphManagedMobileLobAppRelationshipsMobileAppRelationshipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileAppRelationshipId != "" {
        urlTplParams["mobileAppRelationship%2Did"] = mobileAppRelationshipId
    }
    return NewMobileAppsItemGraphManagedMobileLobAppRelationshipsMobileAppRelationshipItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderInternal instantiates a new MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder and sets the default values.
func NewMobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) {
    m := &MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedMobileLobApp/relationships{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder instantiates a new MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder and sets the default values.
func NewMobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileAppsItemGraphManagedMobileLobAppRelationshipsCountRequestBuilder when successful
func (m *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) Count()(*MobileAppsItemGraphManagedMobileLobAppRelationshipsCountRequestBuilder) {
    return NewMobileAppsItemGraphManagedMobileLobAppRelationshipsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of relationships for this mobile app.
// returns a MobileAppRelationshipCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppRelationshipCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipCollectionResponseable), nil
}
// Post create new navigation property to relationships for deviceAppManagement
// returns a MobileAppRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable), nil
}
// ToGetRequestInformation list of relationships for this mobile app.
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to relationships for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder when successful
func (m *MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder) {
    return NewMobileAppsItemGraphManagedMobileLobAppRelationshipsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
