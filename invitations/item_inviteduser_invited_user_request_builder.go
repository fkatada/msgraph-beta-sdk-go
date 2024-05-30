package invitations

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInviteduserInvitedUserRequestBuilder provides operations to manage the invitedUser property of the microsoft.graph.invitation entity.
type ItemInviteduserInvitedUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInviteduserInvitedUserRequestBuilderGetQueryParameters the user created as part of the invitation creation. Read-Only
type ItemInviteduserInvitedUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInviteduserInvitedUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInviteduserInvitedUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInviteduserInvitedUserRequestBuilderGetQueryParameters
}
// NewItemInviteduserInvitedUserRequestBuilderInternal instantiates a new ItemInviteduserInvitedUserRequestBuilder and sets the default values.
func NewItemInviteduserInvitedUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInviteduserInvitedUserRequestBuilder) {
    m := &ItemInviteduserInvitedUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/invitations/{invitation%2Did}/invitedUser{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemInviteduserInvitedUserRequestBuilder instantiates a new ItemInviteduserInvitedUserRequestBuilder and sets the default values.
func NewItemInviteduserInvitedUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInviteduserInvitedUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInviteduserInvitedUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the user created as part of the invitation creation. Read-Only
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInviteduserInvitedUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInviteduserInvitedUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *ItemInviteduserMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *ItemInviteduserInvitedUserRequestBuilder) MailboxSettings()(*ItemInviteduserMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewItemInviteduserMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *ItemInviteduserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *ItemInviteduserInvitedUserRequestBuilder) ServiceProvisioningErrors()(*ItemInviteduserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewItemInviteduserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the user created as part of the invitation creation. Read-Only
// returns a *RequestInformation when successful
func (m *ItemInviteduserInvitedUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInviteduserInvitedUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInviteduserInvitedUserRequestBuilder when successful
func (m *ItemInviteduserInvitedUserRequestBuilder) WithUrl(rawUrl string)(*ItemInviteduserInvitedUserRequestBuilder) {
    return NewItemInviteduserInvitedUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
