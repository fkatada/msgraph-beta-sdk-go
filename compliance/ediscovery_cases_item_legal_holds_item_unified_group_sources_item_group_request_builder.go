package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.ediscovery.unifiedGroupSource entity.
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderGetQueryParameters the group associated with the unifiedGroupSource.
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderInternal instantiates a new EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder) {
    m := &EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/legalHolds/{legalHold%2Did}/unifiedGroupSources/{unifiedGroupSource%2Did}/group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder instantiates a new EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the group associated with the unifiedGroupSource.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// returns a Groupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder) ServiceProvisioningErrors()(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the group associated with the unifiedGroupSource.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// returns a *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
