package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderGetQueryParameters the collection of single-value extended properties defined for the event. Read-only. Nullable.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderGetQueryParameters struct {
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
// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderGetQueryParameters
}
// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySingleValueLegacyExtendedPropertyId provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) BySingleValueLegacyExtendedPropertyId(singleValueLegacyExtendedPropertyId string)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if singleValueLegacyExtendedPropertyId != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = singleValueLegacyExtendedPropertyId
    }
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal instantiates a new SingleValueExtendedPropertiesRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    m := &ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/singleValueExtendedProperties{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder instantiates a new SingleValueExtendedPropertiesRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) Count()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesCountRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of single-value extended properties defined for the event. Read-only. Nullable.
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SingleValueLegacyExtendedPropertyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSingleValueLegacyExtendedPropertyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SingleValueLegacyExtendedPropertyCollectionResponseable), nil
}
// Post create new navigation property to singleValueExtendedProperties for groups
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SingleValueLegacyExtendedPropertyable, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SingleValueLegacyExtendedPropertyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SingleValueLegacyExtendedPropertyable), nil
}
// ToGetRequestInformation the collection of single-value extended properties defined for the event. Read-only. Nullable.
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to singleValueExtendedProperties for groups
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SingleValueLegacyExtendedPropertyable, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
