package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder provides operations to call the provisionEmail method.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilderInternal instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/provisionEmail", pathParameters),
    }
    return m
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post provision an email address for a channel. Microsoft Teams doesn't automatically provision an email address for a channel by default. To have Teams provision an email address, you can call provisionEmail, or through the Teams user interface, select Get email address, which triggers Teams to generate an email address if it didn't provisioned one. To remove the email address of a channel, use the removeEmail method.
// returns a ProvisionChannelEmailResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-provisionemail?view=graph-rest-beta
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder) Post(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProvisionChannelEmailResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProvisionChannelEmailResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProvisionChannelEmailResultable), nil
}
// ToPostRequestInformation provision an email address for a channel. Microsoft Teams doesn't automatically provision an email address for a channel by default. To have Teams provision an email address, you can call provisionEmail, or through the Teams user interface, select Get email address, which triggers Teams to generate an email address if it didn't provisioned one. To remove the email address of a channel, use the removeEmail method.
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder) WithUrl(rawUrl string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsItemProvisionemailProvisionEmailRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
