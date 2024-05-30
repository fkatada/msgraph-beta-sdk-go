package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder casts the previous resource to managedAndroidLobApp.
type MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.managedAndroidLobApp
type MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmanagedandroidlobappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) Assignments()(*MobileappsItemGraphmanagedandroidlobappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphmanagedandroidlobappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmanagedandroidlobappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) Categories()(*MobileappsItemGraphmanagedandroidlobappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphmanagedandroidlobappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderInternal instantiates a new MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) {
    m := &MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedAndroidLobApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder instantiates a new MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentVersions provides operations to manage the contentVersions property of the microsoft.graph.managedMobileLobApp entity.
// returns a *MobileappsItemGraphmanagedandroidlobappContentversionsContentVersionsRequestBuilder when successful
func (m *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) ContentVersions()(*MobileappsItemGraphmanagedandroidlobappContentversionsContentVersionsRequestBuilder) {
    return NewMobileappsItemGraphmanagedandroidlobappContentversionsContentVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.managedAndroidLobApp
// returns a ManagedAndroidLobAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAndroidLobAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAndroidLobAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAndroidLobAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmanagedandroidlobappRelationshipsRequestBuilder when successful
func (m *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) Relationships()(*MobileappsItemGraphmanagedandroidlobappRelationshipsRequestBuilder) {
    return NewMobileappsItemGraphmanagedandroidlobappRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.managedAndroidLobApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder when successful
func (m *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) {
    return NewMobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
