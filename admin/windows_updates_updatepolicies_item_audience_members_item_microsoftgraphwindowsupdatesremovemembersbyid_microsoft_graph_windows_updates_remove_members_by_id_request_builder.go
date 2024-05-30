package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder provides operations to call the removeMembersById method.
type WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderInternal instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) {
    m := &WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/members/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.removeMembersById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove members of the same type from an updatableAssetGroup. You can also use the method removeMembers to remove members.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-removemembersbyid?view=graph-rest-beta
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidRemoveMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove members of the same type from an updatableAssetGroup. You can also use the method removeMembers to remove members.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidRemoveMembersByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemAudienceMembersItemMicrosoftgraphwindowsupdatesremovemembersbyidMicrosoftGraphWindowsUpdatesRemoveMembersByIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
