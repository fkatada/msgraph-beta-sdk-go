package threatsubmission

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatSubmissionRequestBuilder provides operations to manage the threatSubmissionRoot singleton.
type ThreatSubmissionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatSubmissionRequestBuilderGetQueryParameters get threatSubmission
type ThreatSubmissionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatSubmissionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatSubmissionRequestBuilderGetQueryParameters
}
// ThreatSubmissionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatSubmissionRequestBuilderInternal instantiates a new ThreatSubmissionRequestBuilder and sets the default values.
func NewThreatSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionRequestBuilder) {
    m := &ThreatSubmissionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/threatSubmission{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatSubmissionRequestBuilder instantiates a new ThreatSubmissionRequestBuilder and sets the default values.
func NewThreatSubmissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
// EmailThreats provides operations to manage the emailThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *EmailthreatsEmailThreatsRequestBuilder when successful
func (m *ThreatSubmissionRequestBuilder) EmailThreats()(*EmailthreatsEmailThreatsRequestBuilder) {
    return NewEmailthreatsEmailThreatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EmailThreatSubmissionPolicies provides operations to manage the emailThreatSubmissionPolicies property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder when successful
func (m *ThreatSubmissionRequestBuilder) EmailThreatSubmissionPolicies()(*EmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilder) {
    return NewEmailthreatsubmissionpoliciesEmailThreatSubmissionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileThreats provides operations to manage the fileThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *FilethreatsFileThreatsRequestBuilder when successful
func (m *ThreatSubmissionRequestBuilder) FileThreats()(*FilethreatsFileThreatsRequestBuilder) {
    return NewFilethreatsFileThreatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get threatSubmission
// returns a ThreatSubmissionRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatSubmissionRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatSubmissionRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatSubmissionRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable), nil
}
// Patch update threatSubmission
// returns a ThreatSubmissionRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatSubmissionRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatSubmissionRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatSubmissionRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable), nil
}
// ToGetRequestInformation get threatSubmission
// returns a *RequestInformation when successful
func (m *ThreatSubmissionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatSubmissionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update threatSubmission
// returns a *RequestInformation when successful
func (m *ThreatSubmissionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatSubmissionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UrlThreats provides operations to manage the urlThreats property of the microsoft.graph.security.threatSubmissionRoot entity.
// returns a *UrlthreatsUrlThreatsRequestBuilder when successful
func (m *ThreatSubmissionRequestBuilder) UrlThreats()(*UrlthreatsUrlThreatsRequestBuilder) {
    return NewUrlthreatsUrlThreatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ThreatSubmissionRequestBuilder when successful
func (m *ThreatSubmissionRequestBuilder) WithUrl(rawUrl string)(*ThreatSubmissionRequestBuilder) {
    return NewThreatSubmissionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
