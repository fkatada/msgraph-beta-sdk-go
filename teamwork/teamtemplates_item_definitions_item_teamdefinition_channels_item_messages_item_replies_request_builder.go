package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderGetQueryParameters replies for a specified message. Supports $expand for channel messages.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderGetQueryParameters struct {
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
// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderGetQueryParameters
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByChatMessageId1 provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesChatMessageItemRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) ByChatMessageId1(chatMessageId1 string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if chatMessageId1 != "" {
        urlTplParams["chatMessage%2Did1"] = chatMessageId1
    }
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesChatMessageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderInternal instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesCountRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) Count()(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesCountRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) Delta()(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get replies for a specified message. Supports $expand for channel messages.
// returns a ChatMessageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatMessageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageCollectionResponseable), nil
}
// Post create new navigation property to replies for teamwork
// returns a ChatMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable), nil
}
// ToGetRequestInformation replies for a specified message. Supports $expand for channel messages.
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to replies for teamwork
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) WithUrl(rawUrl string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemMessagesItemRepliesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
