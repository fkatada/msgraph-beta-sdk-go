package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder provides operations to call the copyToDefaultContentLocation method.
type ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal instantiates a new ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder and sets the default values.
func NewItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    m := &ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/contentTypes/{contentType%2Did}/copyToDefaultContentLocation", pathParameters),
    }
    return m
}
// NewItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder instantiates a new ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder and sets the default values.
func NewItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action copyToDefaultContentLocation
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-copytodefaultcontentlocation?view=graph-rest-beta
func (m *ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) Post(ctx context.Context, body ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBodyable, requestConfiguration *ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action copyToDefaultContentLocation
// returns a *RequestInformation when successful
func (m *ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBodyable, requestConfiguration *ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder when successful
func (m *ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) WithUrl(rawUrl string)(*ItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    return NewItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
