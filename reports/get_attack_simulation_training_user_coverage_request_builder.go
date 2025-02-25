package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetAttackSimulationTrainingUserCoverageRequestBuilder provides operations to call the getAttackSimulationTrainingUserCoverage method.
type GetAttackSimulationTrainingUserCoverageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters list training coverage for each user of a tenant in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
type GetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters struct {
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
// GetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters
}
// NewGetAttackSimulationTrainingUserCoverageRequestBuilderInternal instantiates a new GetAttackSimulationTrainingUserCoverageRequestBuilder and sets the default values.
func NewGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAttackSimulationTrainingUserCoverageRequestBuilder) {
    m := &GetAttackSimulationTrainingUserCoverageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getAttackSimulationTrainingUserCoverage(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGetAttackSimulationTrainingUserCoverageRequestBuilder instantiates a new GetAttackSimulationTrainingUserCoverageRequestBuilder and sets the default values.
func NewGetAttackSimulationTrainingUserCoverageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAttackSimulationTrainingUserCoverageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list training coverage for each user of a tenant in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// Deprecated: This method is obsolete. Use GetAsGetAttackSimulationTrainingUserCoverageGetResponse instead.
// returns a GetAttackSimulationTrainingUserCoverageResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getattacksimulationtrainingusercoverage?view=graph-rest-beta
func (m *GetAttackSimulationTrainingUserCoverageRequestBuilder) Get(ctx context.Context, requestConfiguration *GetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration)(GetAttackSimulationTrainingUserCoverageResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAttackSimulationTrainingUserCoverageResponseable), nil
}
// GetAsGetAttackSimulationTrainingUserCoverageGetResponse list training coverage for each user of a tenant in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// Deprecated: This report function api is deprecated and will stop returning data on August 20, 2022. Api is now moved to /reports/security. Please use the new API. as of 2022-05/Tasks_And_Plans on 2022-05-24 and will be removed 2022-08-20
// returns a GetAttackSimulationTrainingUserCoverageGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getattacksimulationtrainingusercoverage?view=graph-rest-beta
func (m *GetAttackSimulationTrainingUserCoverageRequestBuilder) GetAsGetAttackSimulationTrainingUserCoverageGetResponse(ctx context.Context, requestConfiguration *GetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration)(GetAttackSimulationTrainingUserCoverageGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetAttackSimulationTrainingUserCoverageGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAttackSimulationTrainingUserCoverageGetResponseable), nil
}
// ToGetRequestInformation list training coverage for each user of a tenant in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// Deprecated: This report function api is deprecated and will stop returning data on August 20, 2022. Api is now moved to /reports/security. Please use the new API. as of 2022-05/Tasks_And_Plans on 2022-05-24 and will be removed 2022-08-20
// returns a *RequestInformation when successful
func (m *GetAttackSimulationTrainingUserCoverageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: This report function api is deprecated and will stop returning data on August 20, 2022. Api is now moved to /reports/security. Please use the new API. as of 2022-05/Tasks_And_Plans on 2022-05-24 and will be removed 2022-08-20
// returns a *GetAttackSimulationTrainingUserCoverageRequestBuilder when successful
func (m *GetAttackSimulationTrainingUserCoverageRequestBuilder) WithUrl(rawUrl string)(*GetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return NewGetAttackSimulationTrainingUserCoverageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
