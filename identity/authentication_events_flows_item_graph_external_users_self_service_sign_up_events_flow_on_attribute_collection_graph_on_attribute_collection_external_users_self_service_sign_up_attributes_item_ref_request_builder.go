package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder provides operations to manage the collection of identityContainer entities.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilderInternal instantiates a new AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder) {
    m := &AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/graph.onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/{identityUserFlowAttribute%2Did}/$ref", pathParameters),
    }
    return m
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder instantiates a new AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove an attribute from an external identities self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object. You can add both custom and built-in attributes to a user flow. After this step, PATCH the user flow to remove the attribute from the attribute collection step.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onattributecollectionexternalusersselfservicesignup-delete-attributes?view=graph-rest-beta
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove an attribute from an external identities self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object. You can add both custom and built-in attributes to a user flow. After this step, PATCH the user flow to remove the attribute from the attribute collection step.
// returns a *RequestInformation when successful
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder when successful
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder) WithUrl(rawUrl string)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
