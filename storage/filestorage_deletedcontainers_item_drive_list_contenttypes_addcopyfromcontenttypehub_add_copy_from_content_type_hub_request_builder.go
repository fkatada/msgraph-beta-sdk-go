package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder provides operations to call the addCopyFromContentTypeHub method.
type FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/list/contentTypes/addCopyFromContentTypeHub", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add or sync a copy of a published content type from the content type hub to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see getCompatibleHubContentTypes and the blog post Syntex Product Updates – August 2021.
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-addcopyfromcontenttypehub?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) Post(ctx context.Context, body FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubPostRequestBodyable, requestConfiguration *FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// ToPostRequestInformation add or sync a copy of a published content type from the content type hub to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see getCompatibleHubContentTypes and the blog post Syntex Product Updates – August 2021.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubPostRequestBodyable, requestConfiguration *FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
