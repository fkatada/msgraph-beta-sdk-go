package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder provides operations to call the endBreak method.
type ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemScheduleTimecardsItemEndbreakEndBreakRequestBuilderInternal instantiates a new ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder and sets the default values.
func NewItemScheduleTimecardsItemEndbreakEndBreakRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder) {
    m := &ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/schedule/timeCards/{timeCard%2Did}/endBreak", pathParameters),
    }
    return m
}
// NewItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder instantiates a new ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder and sets the default values.
func NewItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemScheduleTimecardsItemEndbreakEndBreakRequestBuilderInternal(urlParams, requestAdapter)
}
// Post end the open break in a specific timeCard.
// returns a TimeCardable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/timecard-endbreak?view=graph-rest-beta
func (m *ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder) Post(ctx context.Context, body ItemScheduleTimecardsItemEndbreakEndBreakPostRequestBodyable, requestConfiguration *ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeCardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable), nil
}
// ToPostRequestInformation end the open break in a specific timeCard.
// returns a *RequestInformation when successful
func (m *ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemScheduleTimecardsItemEndbreakEndBreakPostRequestBodyable, requestConfiguration *ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder when successful
func (m *ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder) WithUrl(rawUrl string)(*ItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder) {
    return NewItemScheduleTimecardsItemEndbreakEndBreakRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
