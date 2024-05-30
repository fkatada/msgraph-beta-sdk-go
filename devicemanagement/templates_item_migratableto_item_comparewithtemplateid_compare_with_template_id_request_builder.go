package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder provides operations to call the compare method.
type TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetQueryParameters invoke function compare
type TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetQueryParameters
}
// NewTemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderInternal instantiates a new TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder and sets the default values.
func NewTemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, templateId *string)(*TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) {
    m := &TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}/compare(templateId='{templateId}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if templateId != nil {
        m.BaseRequestBuilder.PathParameters["templateId"] = *templateId
    }
    return m
}
// NewTemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder instantiates a new TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder and sets the default values.
func NewTemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function compare
// Deprecated: This method is obsolete. Use GetAsCompareWithTemplateIdGetResponse instead.
// returns a TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdResponseable), nil
}
// GetAsCompareWithTemplateIdGetResponse invoke function compare
// returns a TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) GetAsCompareWithTemplateIdGetResponse(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdGetResponseable), nil
}
// ToGetRequestInformation invoke function compare
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder when successful
func (m *TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) WithUrl(rawUrl string)(*TemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) {
    return NewTemplatesItemMigratabletoItemComparewithtemplateidCompareWithTemplateIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
