package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder provides operations to call the createUploadSession method.
type ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal instantiates a new ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    m := &ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/attachments/createUploadSession", pathParameters),
    }
    return m
}
// NewItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder instantiates a new ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an upload session that allows an app to iteratively upload ranges of a file, so as to attach the file to an Outlook item. The item can be a message or event. Use this approach to attach a file if the file size is between 3 MB and 150 MB. To attach a file that's smaller than 3 MB, do a POST operation on the attachments navigation property of the Outlook item; see how to do this for a message or for an event. As part of the response, this action returns an upload URL that you can use in subsequent sequential PUT queries. Request headers for each PUT operation let you specify the exact range of bytes to be uploaded. This allows transfer to be resumed, in case the network connection is dropped during upload. The following are the steps to attach a file to an Outlook item using an upload session: See attach large files to Outlook messages or events for an example.
// returns a UploadSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/attachment-createuploadsession?view=graph-rest-beta
func (m *ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) Post(ctx context.Context, body ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UploadSessionable, error) {
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
// ToPostRequestInformation create an upload session that allows an app to iteratively upload ranges of a file, so as to attach the file to an Outlook item. The item can be a message or event. Use this approach to attach a file if the file size is between 3 MB and 150 MB. To attach a file that's smaller than 3 MB, do a POST operation on the attachments navigation property of the Outlook item; see how to do this for a message or for an event. As part of the response, this action returns an upload URL that you can use in subsequent sequential PUT queries. Request headers for each PUT operation let you specify the exact range of bytes to be uploaded. This allows transfer to be resumed, in case the network connection is dropped during upload. The following are the steps to attach a file to an Outlook item using an upload session: See attach large files to Outlook messages or events for an example.
// returns a *RequestInformation when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
