package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/{conversationMember%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a collection of team members who have access to the shared channel.
// returns a ConversationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConversationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberable), nil
}
// ToGetRequestInformation a collection of team members who have access to the shared channel.
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) WithUrl(rawUrl string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
