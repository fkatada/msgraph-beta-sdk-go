package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder provides operations to call the getActivitiesByInterval method.
type ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters invoke function getActivitiesByInterval
type ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters struct {
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
// ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters
}
// NewItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal instantiates a new ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder and sets the default values.
func NewItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *string, interval *string, startDateTime *string)(*ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    m := &ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = *endDateTime
    }
    if interval != nil {
        m.BaseRequestBuilder.PathParameters["interval"] = *interval
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = *startDateTime
    }
    return m
}
// NewItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder instantiates a new ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder and sets the default values.
func NewItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get invoke function getActivitiesByInterval
// Deprecated: This method is obsolete. Use GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse instead.
// returns a ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable), nil
}
// GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse invoke function getActivitiesByInterval
// returns a ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(ctx context.Context, requestConfiguration *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable), nil
}
// ToGetRequestInformation invoke function getActivitiesByInterval
// returns a *RequestInformation when successful
func (m *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) WithUrl(rawUrl string)(*ItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
