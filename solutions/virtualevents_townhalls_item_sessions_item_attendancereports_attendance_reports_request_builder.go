package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
type VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetQueryParameters the attendance reports of an online meeting. Read-only.
type VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetQueryParameters struct {
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
// VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetQueryParameters
}
// VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMeetingAttendanceReportId provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
// returns a *VirtualeventsTownhallsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder when successful
func (m *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) ByMeetingAttendanceReportId(meetingAttendanceReportId string)(*VirtualeventsTownhallsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if meetingAttendanceReportId != "" {
        urlTplParams["meetingAttendanceReport%2Did"] = meetingAttendanceReportId
    }
    return NewVirtualeventsTownhallsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderInternal instantiates a new VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder and sets the default values.
func NewVirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) {
    m := &VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder instantiates a new VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder and sets the default values.
func NewVirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualeventsTownhallsItemSessionsItemAttendancereportsCountRequestBuilder when successful
func (m *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) Count()(*VirtualeventsTownhallsItemSessionsItemAttendancereportsCountRequestBuilder) {
    return NewVirtualeventsTownhallsItemSessionsItemAttendancereportsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the attendance reports of an online meeting. Read-only.
// returns a MeetingAttendanceReportCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingAttendanceReportCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportCollectionResponseable), nil
}
// Post create new navigation property to attendanceReports for solutions
// returns a MeetingAttendanceReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, requestConfiguration *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable), nil
}
// ToGetRequestInformation the attendance reports of an online meeting. Read-only.
// returns a *RequestInformation when successful
func (m *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attendanceReports for solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, requestConfiguration *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder when successful
func (m *VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) {
    return NewVirtualeventsTownhallsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
