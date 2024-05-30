package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
type GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal instantiates a new GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder and sets the default values.
func NewGetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    m := &GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getMailboxUsageQuotaStatusMailboxCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder instantiates a new GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder and sets the default values.
func NewGetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getMailboxUsageQuotaStatusMailboxCounts
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getMailboxUsageQuotaStatusMailboxCounts
// returns a *RequestInformation when successful
func (m *GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder when successful
func (m *GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return NewGetmailboxusagequotastatusmailboxcountswithperiodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
