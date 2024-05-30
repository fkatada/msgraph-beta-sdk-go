package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder provides operations to call the addMembers method.
type WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    m := &WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/exclusions/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.addMembers", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder instantiates a new WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add members to an updatableAssetGroup. You can add azureADDevice resources as members, but may not add updatableAssetGroup resources as members. Adding a Microsoft Entra device as a member of an updatable asset group automatically creates an azureADDevice object, if it does not already exist. You can also use the method addMembersById to add members.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-addmembers?view=graph-rest-beta
func (m *WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersAddMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation add members to an updatableAssetGroup. You can add azureADDevice resources as members, but may not add updatableAssetGroup resources as members. Adding a Microsoft Entra device as a member of an updatable asset group automatically creates an azureADDevice object, if it does not already exist. You can also use the method addMembersById to add members.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersAddMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder when successful
func (m *WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentaudiencesItemExclusionsItemMicrosoftgraphwindowsupdatesaddmembersMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
