package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder provides operations to call the usersRegisteredByFeature method.
type AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal instantiates a new AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder and sets the default values.
func NewAuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, includedUserRoles *string, includedUserTypes *string)(*AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    m := &AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/authenticationMethods/usersRegisteredByFeature(includedUserTypes='{includedUserTypes}',includedUserRoles='{includedUserRoles}')", pathParameters),
    }
    if includedUserRoles != nil {
        m.BaseRequestBuilder.PathParameters["includedUserRoles"] = *includedUserRoles
    }
    if includedUserTypes != nil {
        m.BaseRequestBuilder.PathParameters["includedUserTypes"] = *includedUserTypes
    }
    return m
}
// NewAuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder instantiates a new AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder and sets the default values.
func NewAuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function usersRegisteredByFeature
// returns a UserRegistrationFeatureSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationFeatureSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserRegistrationFeatureSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationFeatureSummaryable), nil
}
// ToGetRequestInformation invoke function usersRegisteredByFeature
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder when successful
func (m *AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) WithUrl(rawUrl string)(*AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return NewAuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
