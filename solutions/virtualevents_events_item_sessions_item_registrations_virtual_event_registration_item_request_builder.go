package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
type VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderGetQueryParameters get registrations from solutions
type VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderGetQueryParameters
}
// NewVirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderInternal instantiates a new VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder and sets the default values.
func NewVirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder) {
    m := &VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/sessions/{virtualEventSession%2Did}/registrations/{virtualEventRegistration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder instantiates a new VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder and sets the default values.
func NewVirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get registrations from solutions
// returns a VirtualEventRegistrationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationable), nil
}
// ToGetRequestInformation get registrations from solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder when successful
func (m *VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder) {
    return NewVirtualeventsEventsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
