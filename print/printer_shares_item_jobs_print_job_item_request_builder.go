package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrinterSharesItemJobsPrintJobItemRequestBuilder provides operations to manage the jobs property of the microsoft.graph.printerBase entity.
type PrinterSharesItemJobsPrintJobItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrinterSharesItemJobsPrintJobItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrinterSharesItemJobsPrintJobItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrinterSharesItemJobsPrintJobItemRequestBuilderGetQueryParameters the list of jobs that are queued for printing by the printer/printerShare.
type PrinterSharesItemJobsPrintJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrinterSharesItemJobsPrintJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrinterSharesItemJobsPrintJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrinterSharesItemJobsPrintJobItemRequestBuilderGetQueryParameters
}
// PrinterSharesItemJobsPrintJobItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrinterSharesItemJobsPrintJobItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Abort provides operations to call the abort method.
// returns a *PrinterSharesItemJobsItemAbortRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Abort()(*PrinterSharesItemJobsItemAbortRequestBuilder) {
    return NewPrinterSharesItemJobsItemAbortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *PrinterSharesItemJobsItemCancelRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Cancel()(*PrinterSharesItemJobsItemCancelRequestBuilder) {
    return NewPrinterSharesItemJobsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CancelPrintJob provides operations to call the cancelPrintJob method.
// returns a *PrinterSharesItemJobsItemCancelPrintJobRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) CancelPrintJob()(*PrinterSharesItemJobsItemCancelPrintJobRequestBuilder) {
    return NewPrinterSharesItemJobsItemCancelPrintJobRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrinterSharesItemJobsPrintJobItemRequestBuilderInternal instantiates a new PrinterSharesItemJobsPrintJobItemRequestBuilder and sets the default values.
func NewPrinterSharesItemJobsPrintJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemJobsPrintJobItemRequestBuilder) {
    m := &PrinterSharesItemJobsPrintJobItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/jobs/{printJob%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrinterSharesItemJobsPrintJobItemRequestBuilder instantiates a new PrinterSharesItemJobsPrintJobItemRequestBuilder and sets the default values.
func NewPrinterSharesItemJobsPrintJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemJobsPrintJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterSharesItemJobsPrintJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property jobs for print
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrinterSharesItemJobsPrintJobItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Documents provides operations to manage the documents property of the microsoft.graph.printJob entity.
// returns a *PrinterSharesItemJobsItemDocumentsRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Documents()(*PrinterSharesItemJobsItemDocumentsRequestBuilder) {
    return NewPrinterSharesItemJobsItemDocumentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of jobs that are queued for printing by the printer/printerShare.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrinterSharesItemJobsPrintJobItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable), nil
}
// Patch update the navigation property jobs in print
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable, requestConfiguration *PrinterSharesItemJobsPrintJobItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable), nil
}
// Redirect provides operations to call the redirect method.
// returns a *PrinterSharesItemJobsItemRedirectRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Redirect()(*PrinterSharesItemJobsItemRedirectRequestBuilder) {
    return NewPrinterSharesItemJobsItemRedirectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start provides operations to call the start method.
// returns a *PrinterSharesItemJobsItemStartRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Start()(*PrinterSharesItemJobsItemStartRequestBuilder) {
    return NewPrinterSharesItemJobsItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StartPrintJob provides operations to call the startPrintJob method.
// returns a *PrinterSharesItemJobsItemStartPrintJobRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) StartPrintJob()(*PrinterSharesItemJobsItemStartPrintJobRequestBuilder) {
    return NewPrinterSharesItemJobsItemStartPrintJobRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.printJob entity.
// returns a *PrinterSharesItemJobsItemTasksRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) Tasks()(*PrinterSharesItemJobsItemTasksRequestBuilder) {
    return NewPrinterSharesItemJobsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property jobs for print
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrinterSharesItemJobsPrintJobItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of jobs that are queued for printing by the printer/printerShare.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrinterSharesItemJobsPrintJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property jobs in print
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintJobable, requestConfiguration *PrinterSharesItemJobsPrintJobItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *PrinterSharesItemJobsPrintJobItemRequestBuilder when successful
func (m *PrinterSharesItemJobsPrintJobItemRequestBuilder) WithUrl(rawUrl string)(*PrinterSharesItemJobsPrintJobItemRequestBuilder) {
    return NewPrinterSharesItemJobsPrintJobItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
