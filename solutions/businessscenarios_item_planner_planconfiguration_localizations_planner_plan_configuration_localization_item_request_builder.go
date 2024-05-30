package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder provides operations to manage the localizations property of the microsoft.graph.plannerPlanConfiguration entity.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetQueryParameters localized names for the plan configuration.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetQueryParameters
}
// BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderInternal instantiates a new BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) {
    m := &BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner/planConfiguration/localizations/{plannerPlanConfigurationLocalization%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder instantiates a new BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property localizations for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get localized names for the plan configuration.
// returns a PlannerPlanConfigurationLocalizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable), nil
}
// Patch update the navigation property localizations in solutions
// returns a PlannerPlanConfigurationLocalizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable), nil
}
// ToDeleteRequestInformation delete navigation property localizations for solutions
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation localized names for the plan configuration.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property localizations in solutions
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationLocalizationable, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) WithUrl(rawUrl string)(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder) {
    return NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsPlannerPlanConfigurationLocalizationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
