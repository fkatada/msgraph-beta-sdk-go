package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder provides operations to call the getCredentialUserRegistrationCount method.
type GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderGetQueryParameters report the current state of how many users in your organization are registered for self-service password reset and multifactor authentication (MFA) capabilities.
type GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderGetQueryParameters struct {
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
// GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderGetQueryParameters
}
// NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderInternal instantiates a new GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder and sets the default values.
func NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder) {
    m := &GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getCredentialUserRegistrationCount(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder instantiates a new GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder and sets the default values.
func NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get report the current state of how many users in your organization are registered for self-service password reset and multifactor authentication (MFA) capabilities.
// Deprecated: This method is obsolete. Use GetAsGetCredentialUserRegistrationCountGetResponse instead.
// returns a GetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getcredentialuserregistrationcount?view=graph-rest-beta
func (m *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder) Get(ctx context.Context, requestConfiguration *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration)(GetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponseable), nil
}
// GetAsGetCredentialUserRegistrationCountGetResponse report the current state of how many users in your organization are registered for self-service password reset and multifactor authentication (MFA) capabilities.
// returns a GetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getcredentialuserregistrationcount?view=graph-rest-beta
func (m *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder) GetAsGetCredentialUserRegistrationCountGetResponse(ctx context.Context, requestConfiguration *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration)(GetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponseable), nil
}
// ToGetRequestInformation report the current state of how many users in your organization are registered for self-service password reset and multifactor authentication (MFA) capabilities.
// returns a *RequestInformation when successful
func (m *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder when successful
func (m *GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder) WithUrl(rawUrl string)(*GetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder) {
    return NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
