package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder provides operations to call the addMembersById method.
type WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    m := &WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/members/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.addMembersById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder instantiates a new WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add members of the same type to an updatableAssetGroup. You can also use the method addMembers to add members.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-addmembersbyid?view=graph-rest-beta
func (m *WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidAddMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation add members of the same type to an updatableAssetGroup. You can also use the method addMembers to add members.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidAddMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder when successful
func (m *WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentaudiencesItemMembersItemMicrosoftgraphwindowsupdatesaddmembersbyidMicrosoftGraphWindowsUpdatesAddMembersByIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
