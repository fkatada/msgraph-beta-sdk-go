package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder provides operations to call the enableUnlicensedAdminstrators method.
type EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilderInternal instantiates a new EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder and sets the default values.
func NewEnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder) {
    m := &EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/enableUnlicensedAdminstrators", pathParameters),
    }
    return m
}
// NewEnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder instantiates a new EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder and sets the default values.
func NewEnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upon enabling, users assigned as administrators via Role Assignment Memberships will no longer require an assigned Intune license. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators will continue to function as-is in that transitive memberships apply and are not subject to the 350 member limit.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder) Post(ctx context.Context, requestConfiguration *EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation upon enabling, users assigned as administrators via Role Assignment Memberships will no longer require an assigned Intune license. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators will continue to function as-is in that transitive memberships apply and are not subject to the 350 member limit.
// returns a *RequestInformation when successful
func (m *EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder when successful
func (m *EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder) WithUrl(rawUrl string)(*EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder) {
    return NewEnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
