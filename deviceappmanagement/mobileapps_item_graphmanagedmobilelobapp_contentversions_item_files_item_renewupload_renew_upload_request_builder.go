package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder provides operations to call the renewUpload method.
type MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal instantiates a new MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    m := &MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedMobileLobApp/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}/renewUpload", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder instantiates a new MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post renews the SAS URI for an application file upload.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) Post(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation renews the SAS URI for an application file upload.
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder when successful
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    return NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
