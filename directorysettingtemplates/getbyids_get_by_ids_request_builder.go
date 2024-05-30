package directorysettingtemplates

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetbyidsGetByIdsRequestBuilder provides operations to call the getByIds method.
type GetbyidsGetByIdsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetbyidsGetByIdsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetbyidsGetByIdsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetbyidsGetByIdsRequestBuilderInternal instantiates a new GetbyidsGetByIdsRequestBuilder and sets the default values.
func NewGetbyidsGetByIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetbyidsGetByIdsRequestBuilder) {
    m := &GetbyidsGetByIdsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directorySettingTemplates/getByIds", pathParameters),
    }
    return m
}
// NewGetbyidsGetByIdsRequestBuilder instantiates a new GetbyidsGetByIdsRequestBuilder and sets the default values.
func NewGetbyidsGetByIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetbyidsGetByIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetbyidsGetByIdsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return the directory objects specified in a list of IDs. Some common uses for this function are to:
// Deprecated: This method is obsolete. Use PostAsGetByIdsPostResponse instead.
// returns a GetbyidsGetByIdsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getbyids?view=graph-rest-beta
func (m *GetbyidsGetByIdsRequestBuilder) Post(ctx context.Context, body GetbyidsGetByIdsPostRequestBodyable, requestConfiguration *GetbyidsGetByIdsRequestBuilderPostRequestConfiguration)(GetbyidsGetByIdsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetbyidsGetByIdsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetbyidsGetByIdsResponseable), nil
}
// PostAsGetByIdsPostResponse return the directory objects specified in a list of IDs. Some common uses for this function are to:
// returns a GetbyidsGetByIdsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getbyids?view=graph-rest-beta
func (m *GetbyidsGetByIdsRequestBuilder) PostAsGetByIdsPostResponse(ctx context.Context, body GetbyidsGetByIdsPostRequestBodyable, requestConfiguration *GetbyidsGetByIdsRequestBuilderPostRequestConfiguration)(GetbyidsGetByIdsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetbyidsGetByIdsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetbyidsGetByIdsPostResponseable), nil
}
// ToPostRequestInformation return the directory objects specified in a list of IDs. Some common uses for this function are to:
// returns a *RequestInformation when successful
func (m *GetbyidsGetByIdsRequestBuilder) ToPostRequestInformation(ctx context.Context, body GetbyidsGetByIdsPostRequestBodyable, requestConfiguration *GetbyidsGetByIdsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetbyidsGetByIdsRequestBuilder when successful
func (m *GetbyidsGetByIdsRequestBuilder) WithUrl(rawUrl string)(*GetbyidsGetByIdsRequestBuilder) {
    return NewGetbyidsGetByIdsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
