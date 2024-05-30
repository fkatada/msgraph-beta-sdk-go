package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder provides operations to manage the authenticationEventsFlows property of the microsoft.graph.identityContainer entity.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a specific authenticationEventsFlow object by ID. Only externalUsersSelfServiceSignupEventsFlow object types are available.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Conditions the conditions property
// returns a *AuthenticationeventsflowsItemConditionsRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) Conditions()(*AuthenticationeventsflowsItemConditionsRequestBuilder) {
    return NewAuthenticationeventsflowsItemConditionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderInternal instantiates a new AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) {
    m := &AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder instantiates a new AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a specific authenticationEventsFlow resource by ID. Only externalUsersSelfServiceSignupEventsFlow object types are available.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationeventsflow-delete?view=graph-rest-beta
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a specific authenticationEventsFlow object by ID. Only externalUsersSelfServiceSignupEventsFlow object types are available.
// returns a AuthenticationEventsFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationeventsflow-get?view=graph-rest-beta
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationEventsFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationEventsFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationEventsFlowable), nil
}
// GraphExternalUsersSelfServiceSignUpEventsFlow casts the previous resource to externalUsersSelfServiceSignUpEventsFlow.
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) GraphExternalUsersSelfServiceSignUpEventsFlow()(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an authenticationEventsFlow object. Only the externalUsersSelfServiceSignupEventsFlow object type is supported.
// returns a AuthenticationEventsFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationeventsflow-update?view=graph-rest-beta
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationEventsFlowable, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationEventsFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationEventsFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationEventsFlowable), nil
}
// ToDeleteRequestInformation delete a specific authenticationEventsFlow resource by ID. Only externalUsersSelfServiceSignupEventsFlow object types are available.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a specific authenticationEventsFlow object by ID. Only externalUsersSelfServiceSignupEventsFlow object types are available.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an authenticationEventsFlow object. Only the externalUsersSelfServiceSignupEventsFlow object type is supported.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationEventsFlowable, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) {
    return NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
