package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentfiltersEnableRequestBuilder provides operations to call the enable method.
type AssignmentfiltersEnableRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentfiltersEnableRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersEnableRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAssignmentfiltersEnableRequestBuilderInternal instantiates a new AssignmentfiltersEnableRequestBuilder and sets the default values.
func NewAssignmentfiltersEnableRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersEnableRequestBuilder) {
    m := &AssignmentfiltersEnableRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/assignmentFilters/enable", pathParameters),
    }
    return m
}
// NewAssignmentfiltersEnableRequestBuilder instantiates a new AssignmentfiltersEnableRequestBuilder and sets the default values.
func NewAssignmentfiltersEnableRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersEnableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentfiltersEnableRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action enable
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersEnableRequestBuilder) Post(ctx context.Context, body AssignmentfiltersEnablePostRequestBodyable, requestConfiguration *AssignmentfiltersEnableRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action enable
// returns a *RequestInformation when successful
func (m *AssignmentfiltersEnableRequestBuilder) ToPostRequestInformation(ctx context.Context, body AssignmentfiltersEnablePostRequestBodyable, requestConfiguration *AssignmentfiltersEnableRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AssignmentfiltersEnableRequestBuilder when successful
func (m *AssignmentfiltersEnableRequestBuilder) WithUrl(rawUrl string)(*AssignmentfiltersEnableRequestBuilder) {
    return NewAssignmentfiltersEnableRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
