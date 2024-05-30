package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder provides operations to call the getTeamsDeviceUsageUserDetail method.
type GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal instantiates a new GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    m := &GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getTeamsDeviceUsageUserDetail(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder instantiates a new GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getTeamsDeviceUsageUserDetail
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getTeamsDeviceUsageUserDetail
// returns a *RequestInformation when successful
func (m *GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder when successful
func (m *GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewGetteamsdeviceusageuserdetailwithperiodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
