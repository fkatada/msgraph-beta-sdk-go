package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder provides operations to call the importResourceActions method.
type CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilderInternal instantiates a new CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder and sets the default values.
func NewCloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder) {
    m := &CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/cloudPC/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}/importResourceActions", pathParameters),
    }
    return m
}
// NewCloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder instantiates a new CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder and sets the default values.
func NewCloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action importResourceActions
// returns a UnifiedRbacResourceNamespaceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder) Post(ctx context.Context, body CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsPostRequestBodyable, requestConfiguration *CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ToPostRequestInformation invoke action importResourceActions
// returns a *RequestInformation when successful
func (m *CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsPostRequestBodyable, requestConfiguration *CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder when successful
func (m *CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder) WithUrl(rawUrl string)(*CloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder) {
    return NewCloudpcResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
