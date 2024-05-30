package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder provides operations to manage the teamsAppDefinition property of the microsoft.graph.teamsAppInstallation entity.
type ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetQueryParameters the details of this version of the app.
type ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetQueryParameters
}
// NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderInternal instantiates a new ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder and sets the default values.
func NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) {
    m := &ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/installedApps/{teamsAppInstallation%2Did}/teamsAppDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder instantiates a new ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder and sets the default values.
func NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the details of this version of the app.
// returns a TeamsAppDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamsAppDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable), nil
}
// ToGetRequestInformation the details of this version of the app.
// returns a *RequestInformation when successful
func (m *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder when successful
func (m *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) WithUrl(rawUrl string)(*ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) {
    return NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
