package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
type FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters get createdByUser from storage
type FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) {
    m := &FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/createdByUser{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder instantiates a new FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get createdByUser from storage
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *FilestorageContainersItemDriveListCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) MailboxSettings()(*FilestorageContainersItemDriveListCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewFilestorageContainersItemDriveListCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *FilestorageContainersItemDriveListCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) ServiceProvisioningErrors()(*FilestorageContainersItemDriveListCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewFilestorageContainersItemDriveListCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get createdByUser from storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
