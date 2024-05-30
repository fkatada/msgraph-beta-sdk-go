package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
type VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderGetQueryParameters get registrations from solutions
type VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderGetQueryParameters struct {
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
// VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderGetQueryParameters
}
// ByVirtualEventRegistrationId provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualeventsTownhallsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder when successful
func (m *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder) ByVirtualEventRegistrationId(virtualEventRegistrationId string)(*VirtualeventsTownhallsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if virtualEventRegistrationId != "" {
        urlTplParams["virtualEventRegistration%2Did"] = virtualEventRegistrationId
    }
    return NewVirtualeventsTownhallsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderInternal instantiates a new VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder and sets the default values.
func NewVirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder) {
    m := &VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/sessions/{virtualEventSession%2Did}/registrations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder instantiates a new VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder and sets the default values.
func NewVirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualeventsTownhallsItemSessionsItemRegistrationsCountRequestBuilder when successful
func (m *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder) Count()(*VirtualeventsTownhallsItemSessionsItemRegistrationsCountRequestBuilder) {
    return NewVirtualeventsTownhallsItemSessionsItemRegistrationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get registrations from solutions
// returns a VirtualEventRegistrationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventRegistrationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationCollectionResponseable), nil
}
// ToGetRequestInformation get registrations from solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder when successful
func (m *VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder) {
    return NewVirtualeventsTownhallsItemSessionsItemRegistrationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
