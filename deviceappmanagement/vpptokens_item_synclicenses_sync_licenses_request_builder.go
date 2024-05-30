package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VpptokensItemSynclicensesSyncLicensesRequestBuilder provides operations to call the syncLicenses method.
type VpptokensItemSynclicensesSyncLicensesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VpptokensItemSynclicensesSyncLicensesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VpptokensItemSynclicensesSyncLicensesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVpptokensItemSynclicensesSyncLicensesRequestBuilderInternal instantiates a new VpptokensItemSynclicensesSyncLicensesRequestBuilder and sets the default values.
func NewVpptokensItemSynclicensesSyncLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VpptokensItemSynclicensesSyncLicensesRequestBuilder) {
    m := &VpptokensItemSynclicensesSyncLicensesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/vppTokens/{vppToken%2Did}/syncLicenses", pathParameters),
    }
    return m
}
// NewVpptokensItemSynclicensesSyncLicensesRequestBuilder instantiates a new VpptokensItemSynclicensesSyncLicensesRequestBuilder and sets the default values.
func NewVpptokensItemSynclicensesSyncLicensesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VpptokensItemSynclicensesSyncLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVpptokensItemSynclicensesSyncLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post syncs licenses associated with a specific appleVolumePurchaseProgramToken
// returns a VppTokenable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VpptokensItemSynclicensesSyncLicensesRequestBuilder) Post(ctx context.Context, requestConfiguration *VpptokensItemSynclicensesSyncLicensesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VppTokenable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVppTokenFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VppTokenable), nil
}
// ToPostRequestInformation syncs licenses associated with a specific appleVolumePurchaseProgramToken
// returns a *RequestInformation when successful
func (m *VpptokensItemSynclicensesSyncLicensesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VpptokensItemSynclicensesSyncLicensesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VpptokensItemSynclicensesSyncLicensesRequestBuilder when successful
func (m *VpptokensItemSynclicensesSyncLicensesRequestBuilder) WithUrl(rawUrl string)(*VpptokensItemSynclicensesSyncLicensesRequestBuilder) {
    return NewVpptokensItemSynclicensesSyncLicensesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
