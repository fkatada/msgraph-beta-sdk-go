package childfolders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ief621387aa2e54a37a0d220123a58940d48e34167bddee38029a63db1c626c70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/count"
    if29039cf371cd6d763f22b3795249e81b13dceb0437e930f09a56ae5f91473b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/delta"
)

// ChildFoldersRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
type ChildFoldersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChildFoldersRequestBuilderGetQueryParameters get the folder collection under the specified folder. You can use the `.../me/mailFolders` shortcut to get the top-levelfolder collection and navigate to another folder. By default, this operation does not return hidden folders. Use a query parameter _includeHiddenFolders_ to include them in the response.
type ChildFoldersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ChildFoldersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChildFoldersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChildFoldersRequestBuilderGetQueryParameters
}
// ChildFoldersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChildFoldersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChildFoldersRequestBuilderInternal instantiates a new ChildFoldersRequestBuilder and sets the default values.
func NewChildFoldersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChildFoldersRequestBuilder) {
    m := &ChildFoldersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChildFoldersRequestBuilder instantiates a new ChildFoldersRequestBuilder and sets the default values.
func NewChildFoldersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChildFoldersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChildFoldersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ChildFoldersRequestBuilder) Count()(*ief621387aa2e54a37a0d220123a58940d48e34167bddee38029a63db1c626c70.CountRequestBuilder) {
    return ief621387aa2e54a37a0d220123a58940d48e34167bddee38029a63db1c626c70.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the folder collection under the specified folder. You can use the `.../me/mailFolders` shortcut to get the top-levelfolder collection and navigate to another folder. By default, this operation does not return hidden folders. Use a query parameter _includeHiddenFolders_ to include them in the response.
func (m *ChildFoldersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ChildFoldersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation use this API to create a new child mailFolder. If you intend a new folder to be hidden, you must set the **isHidden** property to `true` on creation.
func (m *ChildFoldersRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *ChildFoldersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *ChildFoldersRequestBuilder) Delta()(*if29039cf371cd6d763f22b3795249e81b13dceb0437e930f09a56ae5f91473b3.DeltaRequestBuilder) {
    return if29039cf371cd6d763f22b3795249e81b13dceb0437e930f09a56ae5f91473b3.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the folder collection under the specified folder. You can use the `.../me/mailFolders` shortcut to get the top-levelfolder collection and navigate to another folder. By default, this operation does not return hidden folders. Use a query parameter _includeHiddenFolders_ to include them in the response.
func (m *ChildFoldersRequestBuilder) Get(ctx context.Context, requestConfiguration *ChildFoldersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailFolderCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderCollectionResponseable), nil
}
// Post use this API to create a new child mailFolder. If you intend a new folder to be hidden, you must set the **isHidden** property to `true` on creation.
func (m *ChildFoldersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *ChildFoldersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable), nil
}
