package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder provides operations to call the publish method.
type FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveListContenttypesItemPublishRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesItemPublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder) {
    m := &FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}/publish", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder instantiates a new FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListContenttypesItemPublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action publish
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-publish?view=graph-rest-beta
func (m *FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action publish
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
