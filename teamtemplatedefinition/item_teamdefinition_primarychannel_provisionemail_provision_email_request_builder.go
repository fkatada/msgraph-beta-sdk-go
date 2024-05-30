package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder provides operations to call the provisionEmail method.
type ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilderInternal instantiates a new ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder and sets the default values.
func NewItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder) {
    m := &ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/provisionEmail", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder instantiates a new ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder and sets the default values.
func NewItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post provision an email address for a channel. Microsoft Teams doesn't automatically provision an email address for a channel by default. To have Teams provision an email address, you can call provisionEmail, or through the Teams user interface, select Get email address, which triggers Teams to generate an email address if it didn't provisioned one. To remove the email address of a channel, use the removeEmail method.
// returns a ProvisionChannelEmailResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-provisionemail?view=graph-rest-beta
func (m *ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProvisionChannelEmailResultable, error) {
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
func (m *ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder when successful
func (m *ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder) {
    return NewItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
