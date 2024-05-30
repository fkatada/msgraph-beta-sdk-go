package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder provides operations to manage the versions property of the microsoft.graph.agreementFileLocalization entity.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderGetQueryParameters read-only. Customized versions of the terms of use agreement in the Microsoft Entra tenant.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderGetQueryParameters
}
// TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAgreementFileVersionId provides operations to manage the versions property of the microsoft.graph.agreementFileLocalization entity.
// returns a *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) ByAgreementFileVersionId(agreementFileVersionId string)(*TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if agreementFileVersionId != "" {
        urlTplParams["agreementFileVersion%2Did"] = agreementFileVersionId
    }
    return NewTermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderInternal instantiates a new TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) {
    m := &TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}/file/localizations/{agreementFileLocalization%2Did}/versions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder instantiates a new TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TermsofuseAgreementsItemFileLocalizationsItemVersionsCountRequestBuilder when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) Count()(*TermsofuseAgreementsItemFileLocalizationsItemVersionsCountRequestBuilder) {
    return NewTermsofuseAgreementsItemFileLocalizationsItemVersionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Customized versions of the terms of use agreement in the Microsoft Entra tenant.
// returns a AgreementFileVersionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileVersionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgreementFileVersionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileVersionCollectionResponseable), nil
}
// Post create new navigation property to versions for identityGovernance
// returns a AgreementFileVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileVersionable, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileVersionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgreementFileVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileVersionable), nil
}
// ToGetRequestInformation read-only. Customized versions of the terms of use agreement in the Microsoft Entra tenant.
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to versions for identityGovernance
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgreementFileVersionable, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) WithUrl(rawUrl string)(*TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) {
    return NewTermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
