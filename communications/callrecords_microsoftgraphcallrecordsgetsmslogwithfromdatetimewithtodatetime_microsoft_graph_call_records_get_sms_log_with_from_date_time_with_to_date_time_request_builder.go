package communications

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder provides operations to call the getSmsLog method.
type CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters get the log of a sent/received SMS as a collection of smsLogRow entries.
type CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters
}
// NewCallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderInternal instantiates a new CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) {
    m := &CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/microsoft.graph.callRecords.getSmsLog(fromDateTime={fromDateTime},toDateTime={toDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if fromDateTime != nil {
        m.BaseRequestBuilder.PathParameters["fromDateTime"] = (*fromDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if toDateTime != nil {
        m.BaseRequestBuilder.PathParameters["toDateTime"] = (*toDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewCallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder instantiates a new CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get get the log of a sent/received SMS as a collection of smsLogRow entries.
// Deprecated: This method is obsolete. Use GetAsGetSmsLogWithFromDateTimeWithToDateTimeGetResponse instead.
// returns a CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeGetSmsLogWithFromDateTimeWithToDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeGetSmsLogWithFromDateTimeWithToDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeGetSmsLogWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeGetSmsLogWithFromDateTimeWithToDateTimeResponseable), nil
}
// GetAsGetSmsLogWithFromDateTimeWithToDateTimeGetResponse get the log of a sent/received SMS as a collection of smsLogRow entries.
// returns a CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeGetSmsLogWithFromDateTimeWithToDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) GetAsGetSmsLogWithFromDateTimeWithToDateTimeGetResponse(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeGetSmsLogWithFromDateTimeWithToDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeGetSmsLogWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeGetSmsLogWithFromDateTimeWithToDateTimeGetResponseable), nil
}
// ToGetRequestInformation get the log of a sent/received SMS as a collection of smsLogRow entries.
// returns a *RequestInformation when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) WithUrl(rawUrl string)(*CallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallrecordsMicrosoftgraphcallrecordsgetsmslogwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetSmsLogWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
