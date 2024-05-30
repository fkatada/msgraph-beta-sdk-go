package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
type OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetQueryParameters list of attendance records of an attendance report. Read-only.
type OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetQueryParameters
}
// OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAttendanceRecordId provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
// returns a *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordItemRequestBuilder when successful
func (m *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) ByAttendanceRecordId(attendanceRecordId string)(*OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attendanceRecordId != "" {
        urlTplParams["attendanceRecord%2Did"] = attendanceRecordId
    }
    return NewOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderInternal instantiates a new OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder and sets the default values.
func NewOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) {
    m := &OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/onlineMeetings/{onlineMeeting%2Did}/attendanceReports/{meetingAttendanceReport%2Did}/attendanceRecords{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder instantiates a new OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder and sets the default values.
func NewOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *OnlinemeetingsItemAttendancereportsItemAttendancerecordsCountRequestBuilder when successful
func (m *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) Count()(*OnlinemeetingsItemAttendancereportsItemAttendancerecordsCountRequestBuilder) {
    return NewOnlinemeetingsItemAttendancereportsItemAttendancerecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of attendance records of an attendance report. Read-only.
// returns a AttendanceRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendanceRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordCollectionResponseable), nil
}
// Post create new navigation property to attendanceRecords for app
// returns a AttendanceRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, requestConfiguration *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendanceRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable), nil
}
// ToGetRequestInformation list of attendance records of an attendance report. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attendanceRecords for app
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, requestConfiguration *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder when successful
func (m *OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) {
    return NewOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
