package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImpactedresourcesItemReactivateRequestBuilder provides operations to call the reactivate method.
type ImpactedresourcesItemReactivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImpactedresourcesItemReactivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImpactedresourcesItemReactivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImpactedresourcesItemReactivateRequestBuilderInternal instantiates a new ImpactedresourcesItemReactivateRequestBuilder and sets the default values.
func NewImpactedresourcesItemReactivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImpactedresourcesItemReactivateRequestBuilder) {
    m := &ImpactedresourcesItemReactivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/impactedResources/{impactedResource%2Did}/reactivate", pathParameters),
    }
    return m
}
// NewImpactedresourcesItemReactivateRequestBuilder instantiates a new ImpactedresourcesItemReactivateRequestBuilder and sets the default values.
func NewImpactedresourcesItemReactivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImpactedresourcesItemReactivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImpactedresourcesItemReactivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reactivate an accidentally dismissed, completed, or postponed impactedResource object. This action updates the status of the resource to active. This method is relevant only if the status of the impactedResource object is dismissed, postponed, or completedByUser. 
// returns a ImpactedResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/impactedresource-reactivate?view=graph-rest-beta
func (m *ImpactedresourcesItemReactivateRequestBuilder) Post(ctx context.Context, requestConfiguration *ImpactedresourcesItemReactivateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImpactedResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable), nil
}
// ToPostRequestInformation reactivate an accidentally dismissed, completed, or postponed impactedResource object. This action updates the status of the resource to active. This method is relevant only if the status of the impactedResource object is dismissed, postponed, or completedByUser. 
// returns a *RequestInformation when successful
func (m *ImpactedresourcesItemReactivateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ImpactedresourcesItemReactivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ImpactedresourcesItemReactivateRequestBuilder when successful
func (m *ImpactedresourcesItemReactivateRequestBuilder) WithUrl(rawUrl string)(*ImpactedresourcesItemReactivateRequestBuilder) {
    return NewImpactedresourcesItemReactivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
