package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BookingbusinessesItemAppointmentsItemCancelRequestBuilder provides operations to call the cancel method.
type BookingbusinessesItemAppointmentsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingbusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingbusinessesItemAppointmentsItemCancelRequestBuilderInternal instantiates a new BookingbusinessesItemAppointmentsItemCancelRequestBuilder and sets the default values.
func NewBookingbusinessesItemAppointmentsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemAppointmentsItemCancelRequestBuilder) {
    m := &BookingbusinessesItemAppointmentsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/appointments/{bookingAppointment%2Did}/cancel", pathParameters),
    }
    return m
}
// NewBookingbusinessesItemAppointmentsItemCancelRequestBuilder instantiates a new BookingbusinessesItemAppointmentsItemCancelRequestBuilder and sets the default values.
func NewBookingbusinessesItemAppointmentsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemAppointmentsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingbusinessesItemAppointmentsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel the specified bookingAppointment in the specified bookingBusiness, and send a message to the involved customer and staff members.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingappointment-cancel?view=graph-rest-beta
func (m *BookingbusinessesItemAppointmentsItemCancelRequestBuilder) Post(ctx context.Context, body BookingbusinessesItemAppointmentsItemCancelPostRequestBodyable, requestConfiguration *BookingbusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation cancel the specified bookingAppointment in the specified bookingBusiness, and send a message to the involved customer and staff members.
// returns a *RequestInformation when successful
func (m *BookingbusinessesItemAppointmentsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, body BookingbusinessesItemAppointmentsItemCancelPostRequestBodyable, requestConfiguration *BookingbusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BookingbusinessesItemAppointmentsItemCancelRequestBuilder when successful
func (m *BookingbusinessesItemAppointmentsItemCancelRequestBuilder) WithUrl(rawUrl string)(*BookingbusinessesItemAppointmentsItemCancelRequestBuilder) {
    return NewBookingbusinessesItemAppointmentsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
