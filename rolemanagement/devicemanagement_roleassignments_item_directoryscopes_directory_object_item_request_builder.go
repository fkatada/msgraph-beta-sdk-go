package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
type DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderGetQueryParameters read-only collection that references the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
type DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderInternal instantiates a new DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder and sets the default values.
func NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder) {
    m := &DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/deviceManagement/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}/directoryScopes/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder instantiates a new DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder and sets the default values.
func NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only collection that references the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// ToGetRequestInformation read-only collection that references the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// returns a *RequestInformation when successful
func (m *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder when successful
func (m *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder) {
    return NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
