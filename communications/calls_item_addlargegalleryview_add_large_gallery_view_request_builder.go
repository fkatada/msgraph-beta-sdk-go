package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder provides operations to call the addLargeGalleryView method.
type CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilderInternal instantiates a new CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder and sets the default values.
func NewCallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder) {
    m := &CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/addLargeGalleryView", pathParameters),
    }
    return m
}
// NewCallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder instantiates a new CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder and sets the default values.
func NewCallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add the large gallery view to a call.  For details about how to identify a large gallery view participant in a roster so that you can retrieve the relevant data to subscribe to the video feed, see Identify large gallery view participants in a roster.
// returns a AddLargeGalleryViewOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/call-addlargegalleryview?view=graph-rest-beta
func (m *CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder) Post(ctx context.Context, body CallsItemAddlargegalleryviewAddLargeGalleryViewPostRequestBodyable, requestConfiguration *CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AddLargeGalleryViewOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAddLargeGalleryViewOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AddLargeGalleryViewOperationable), nil
}
// ToPostRequestInformation add the large gallery view to a call.  For details about how to identify a large gallery view participant in a roster so that you can retrieve the relevant data to subscribe to the video feed, see Identify large gallery view participants in a roster.
// returns a *RequestInformation when successful
func (m *CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemAddlargegalleryviewAddLargeGalleryViewPostRequestBodyable, requestConfiguration *CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder when successful
func (m *CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder) WithUrl(rawUrl string)(*CallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder) {
    return NewCallsItemAddlargegalleryviewAddLargeGalleryViewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
