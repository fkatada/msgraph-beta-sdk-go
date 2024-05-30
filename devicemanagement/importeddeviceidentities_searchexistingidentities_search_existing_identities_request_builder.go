package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder provides operations to call the searchExistingIdentities method.
type ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderInternal instantiates a new ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder and sets the default values.
func NewImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder) {
    m := &ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/importedDeviceIdentities/searchExistingIdentities", pathParameters),
    }
    return m
}
// NewImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder instantiates a new ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder and sets the default values.
func NewImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action searchExistingIdentities
// Deprecated: This method is obsolete. Use PostAsSearchExistingIdentitiesPostResponse instead.
// returns a ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder) Post(ctx context.Context, body ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesPostRequestBodyable, requestConfiguration *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration)(ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesResponseable), nil
}
// PostAsSearchExistingIdentitiesPostResponse invoke action searchExistingIdentities
// returns a ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder) PostAsSearchExistingIdentitiesPostResponse(ctx context.Context, body ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesPostRequestBodyable, requestConfiguration *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration)(ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesPostResponseable), nil
}
// ToPostRequestInformation invoke action searchExistingIdentities
// returns a *RequestInformation when successful
func (m *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesPostRequestBodyable, requestConfiguration *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder when successful
func (m *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder) WithUrl(rawUrl string)(*ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder) {
    return NewImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
