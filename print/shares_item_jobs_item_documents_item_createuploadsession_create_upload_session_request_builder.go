package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder provides operations to call the createUploadSession method.
type SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal instantiates a new SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewSharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    m := &SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/shares/{printerShare%2Did}/jobs/{printJob%2Did}/documents/{printDocument%2Did}/createUploadSession", pathParameters),
    }
    return m
}
// NewSharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder instantiates a new SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewSharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an upload session that allows an app to iteratively upload ranges of a binary file linked to the print document. As part of the response, this action returns an upload URL that can be used in subsequent sequential PUT queries. Request headers for each PUT operation can be used to specify the exact range of bytes to be uploaded. This allows transfer to be resumed, in case the network connection is dropped during upload. 
// returns a UploadSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/printdocument-createuploadsession?view=graph-rest-beta
func (m *SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder) Post(ctx context.Context, body SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UploadSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUploadSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UploadSessionable), nil
}
// ToPostRequestInformation create an upload session that allows an app to iteratively upload ranges of a binary file linked to the print document. As part of the response, this action returns an upload URL that can be used in subsequent sequential PUT queries. Request headers for each PUT operation can be used to specify the exact range of bytes to be uploaded. This allows transfer to be resumed, in case the network connection is dropped during upload. 
// returns a *RequestInformation when successful
func (m *SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder) WithUrl(rawUrl string)(*SharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewSharesItemJobsItemDocumentsItemCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
