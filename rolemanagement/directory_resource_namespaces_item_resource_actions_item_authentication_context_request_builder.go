package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder provides operations to manage the authenticationContext property of the microsoft.graph.unifiedRbacResourceAction entity.
type DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetQueryParameters get authenticationContext from roleManagement
type DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetQueryParameters
}
// NewDirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderInternal instantiates a new AuthenticationContextRequestBuilder and sets the default values.
func NewDirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) {
    m := &DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}/resourceActions/{unifiedRbacResourceAction%2Did}/authenticationContext{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewDirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder instantiates a new AuthenticationContextRequestBuilder and sets the default values.
func NewDirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get authenticationContext from roleManagement
func (m *DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationContextClassReferenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable), nil
}
// ToGetRequestInformation get authenticationContext from roleManagement
func (m *DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
