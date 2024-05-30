package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder provides operations to call the clockIn method.
type ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilderInternal instantiates a new ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder and sets the default values.
func NewItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder) {
    m := &ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/timeCards/clockIn", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder instantiates a new ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder and sets the default values.
func NewItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clock in to start a timeCard.
// returns a TimeCardable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/timecard-clockin?view=graph-rest-beta
func (m *ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder) Post(ctx context.Context, body ItemTeamdefinitionScheduleTimecardsClockinClockInPostRequestBodyable, requestConfiguration *ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable, error) {
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
// ToPostRequestInformation clock in to start a timeCard.
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamdefinitionScheduleTimecardsClockinClockInPostRequestBodyable, requestConfiguration *ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder when successful
func (m *ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder) {
    return NewItemTeamdefinitionScheduleTimecardsClockinClockInRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
