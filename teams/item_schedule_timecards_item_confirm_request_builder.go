package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemScheduleTimecardsItemConfirmRequestBuilder provides operations to call the confirm method.
type ItemScheduleTimecardsItemConfirmRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemScheduleTimecardsItemConfirmRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleTimecardsItemConfirmRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemScheduleTimecardsItemConfirmRequestBuilderInternal instantiates a new ItemScheduleTimecardsItemConfirmRequestBuilder and sets the default values.
func NewItemScheduleTimecardsItemConfirmRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleTimecardsItemConfirmRequestBuilder) {
    m := &ItemScheduleTimecardsItemConfirmRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/schedule/timeCards/{timeCard%2Did}/confirm", pathParameters),
    }
    return m
}
// NewItemScheduleTimecardsItemConfirmRequestBuilder instantiates a new ItemScheduleTimecardsItemConfirmRequestBuilder and sets the default values.
func NewItemScheduleTimecardsItemConfirmRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleTimecardsItemConfirmRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemScheduleTimecardsItemConfirmRequestBuilderInternal(urlParams, requestAdapter)
}
// Post confirm a specific timeCard.
// returns a TimeCardable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/timecard-confirm?view=graph-rest-beta
func (m *ItemScheduleTimecardsItemConfirmRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemScheduleTimecardsItemConfirmRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation confirm a specific timeCard.
// returns a *RequestInformation when successful
func (m *ItemScheduleTimecardsItemConfirmRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemScheduleTimecardsItemConfirmRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemScheduleTimecardsItemConfirmRequestBuilder when successful
func (m *ItemScheduleTimecardsItemConfirmRequestBuilder) WithUrl(rawUrl string)(*ItemScheduleTimecardsItemConfirmRequestBuilder) {
    return NewItemScheduleTimecardsItemConfirmRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
