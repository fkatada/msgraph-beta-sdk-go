package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentfiltersValidatefilterValidateFilterRequestBuilder provides operations to call the validateFilter method.
type AssignmentfiltersValidatefilterValidateFilterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentfiltersValidatefilterValidateFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersValidatefilterValidateFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAssignmentfiltersValidatefilterValidateFilterRequestBuilderInternal instantiates a new AssignmentfiltersValidatefilterValidateFilterRequestBuilder and sets the default values.
func NewAssignmentfiltersValidatefilterValidateFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersValidatefilterValidateFilterRequestBuilder) {
    m := &AssignmentfiltersValidatefilterValidateFilterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/assignmentFilters/validateFilter", pathParameters),
    }
    return m
}
// NewAssignmentfiltersValidatefilterValidateFilterRequestBuilder instantiates a new AssignmentfiltersValidatefilterValidateFilterRequestBuilder and sets the default values.
func NewAssignmentfiltersValidatefilterValidateFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersValidatefilterValidateFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentfiltersValidatefilterValidateFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validateFilter
// returns a AssignmentFilterValidationResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersValidatefilterValidateFilterRequestBuilder) Post(ctx context.Context, body AssignmentfiltersValidatefilterValidateFilterPostRequestBodyable, requestConfiguration *AssignmentfiltersValidatefilterValidateFilterRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterValidationResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentFilterValidationResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterValidationResultable), nil
}
// ToPostRequestInformation invoke action validateFilter
// returns a *RequestInformation when successful
func (m *AssignmentfiltersValidatefilterValidateFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body AssignmentfiltersValidatefilterValidateFilterPostRequestBodyable, requestConfiguration *AssignmentfiltersValidatefilterValidateFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AssignmentfiltersValidatefilterValidateFilterRequestBuilder when successful
func (m *AssignmentfiltersValidatefilterValidateFilterRequestBuilder) WithUrl(rawUrl string)(*AssignmentfiltersValidatefilterValidateFilterRequestBuilder) {
    return NewAssignmentfiltersValidatefilterValidateFilterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
