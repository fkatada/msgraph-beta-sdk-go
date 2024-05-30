package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder provides operations to call the updateIndex method.
type EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderInternal instantiates a new EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) {
    m := &EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/noncustodialDataSources/{noncustodialDataSource%2Did}/microsoft.graph.ediscovery.updateIndex", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder instantiates a new EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateIndex
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action updateIndex
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
