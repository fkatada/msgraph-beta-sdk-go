package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters struct {
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
// ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters
}
// ByConversationMemberId provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
// returns a *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder when successful
func (m *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) ByConversationMemberId(conversationMemberId string)(*ItemTeamChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if conversationMemberId != "" {
        urlTplParams["conversationMember%2Did"] = conversationMemberId
    }
    return NewItemTeamChannelsItemSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal instantiates a new ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder and sets the default values.
func NewItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    m := &ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder instantiates a new ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder and sets the default values.
func NewItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersCountRequestBuilder when successful
func (m *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) Count()(*ItemTeamChannelsItemSharedwithteamsItemAllowedmembersCountRequestBuilder) {
    return NewItemTeamChannelsItemSharedwithteamsItemAllowedmembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of team members who have access to the shared channel.
// returns a ConversationMemberCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConversationMemberCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberCollectionResponseable), nil
}
// ToGetRequestInformation a collection of team members who have access to the shared channel.
// returns a *RequestInformation when successful
func (m *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder when successful
func (m *ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) WithUrl(rawUrl string)(*ItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    return NewItemTeamChannelsItemSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
