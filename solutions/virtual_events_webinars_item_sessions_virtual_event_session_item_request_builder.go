package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder provides operations to manage the sessions property of the microsoft.graph.virtualEvent entity.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters read the properties and relationships of a virtualEventSession object.  Currently, the following virtual event types are supported: virtualEventTownhall and virtualEventWebinar.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters
}
// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
// returns a *VirtualEventsWebinarsItemSessionsItemAttendanceReportsRequestBuilder when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) AttendanceReports()(*VirtualEventsWebinarsItemSessionsItemAttendanceReportsRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemAttendanceReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderInternal instantiates a new VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) {
    m := &VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/sessions/{virtualEventSession%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder instantiates a new VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sessions for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get read the properties and relationships of a virtualEventSession object.  Currently, the following virtual event types are supported: virtualEventTownhall and virtualEventWebinar.
// returns a VirtualEventSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventsession-get?view=graph-rest-beta
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable), nil
}
// Patch update the navigation property sessions in solutions
// returns a VirtualEventSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable), nil
}
// Presenters provides operations to manage the presenters property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsWebinarsItemSessionsItemPresentersRequestBuilder when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Presenters()(*VirtualEventsWebinarsItemSessionsItemPresentersRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemPresentersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Registrations provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsWebinarsItemSessionsItemRegistrationsRequestBuilder when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Registrations()(*VirtualEventsWebinarsItemSessionsItemRegistrationsRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RegistrationsWithEmail provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsWebinarsItemSessionsItemRegistrationsWithEmailRequestBuilder when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) RegistrationsWithEmail(email *string)(*VirtualEventsWebinarsItemSessionsItemRegistrationsWithEmailRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemRegistrationsWithEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, email)
}
// RegistrationsWithUserId provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsWebinarsItemSessionsItemRegistrationsWithUserIdRequestBuilder when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) RegistrationsWithUserId(userId *string)(*VirtualEventsWebinarsItemSessionsItemRegistrationsWithUserIdRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemRegistrationsWithUserIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, userId)
}
// ToDeleteRequestInformation delete navigation property sessions for solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a virtualEventSession object.  Currently, the following virtual event types are supported: virtualEventTownhall and virtualEventWebinar.
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sessions in solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder when successful
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
