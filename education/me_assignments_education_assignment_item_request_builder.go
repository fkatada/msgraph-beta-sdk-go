package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAssignmentsEducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
type MeAssignmentsEducationAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters list of assignments for the user. Nullable.
type MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters
}
// MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate provides operations to call the activate method.
// returns a *MeAssignmentsItemActivateRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Activate()(*MeAssignmentsItemActivateRequestBuilder) {
    return NewMeAssignmentsItemActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.educationAssignment entity.
// returns a *MeAssignmentsItemCategoriesRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Categories()(*MeAssignmentsItemCategoriesRequestBuilder) {
    return NewMeAssignmentsItemCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal instantiates a new MeAssignmentsEducationAssignmentItemRequestBuilder and sets the default values.
func NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsEducationAssignmentItemRequestBuilder) {
    m := &MeAssignmentsEducationAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMeAssignmentsEducationAssignmentItemRequestBuilder instantiates a new MeAssignmentsEducationAssignmentItemRequestBuilder and sets the default values.
func NewMeAssignmentsEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsEducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Deactivate provides operations to call the deactivate method.
// returns a *MeAssignmentsItemDeactivateRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Deactivate()(*MeAssignmentsItemDeactivateRequestBuilder) {
    return NewMeAssignmentsItemDeactivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property assignments for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of assignments for the user. Nullable.
// returns a EducationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// GradingCategory provides operations to manage the gradingCategory property of the microsoft.graph.educationAssignment entity.
// returns a *MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) GradingCategory()(*MeAssignmentsItemGradingcategoryGradingCategoryRequestBuilder) {
    return NewMeAssignmentsItemGradingcategoryGradingCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GradingScheme provides operations to manage the gradingScheme property of the microsoft.graph.educationAssignment entity.
// returns a *MeAssignmentsItemGradingschemeGradingSchemeRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) GradingScheme()(*MeAssignmentsItemGradingschemeGradingSchemeRequestBuilder) {
    return NewMeAssignmentsItemGradingschemeGradingSchemeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignments in education
// returns a EducationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// Publish provides operations to call the publish method.
// returns a *MeAssignmentsItemPublishRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Publish()(*MeAssignmentsItemPublishRequestBuilder) {
    return NewMeAssignmentsItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resources provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
// returns a *MeAssignmentsItemResourcesRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Resources()(*MeAssignmentsItemResourcesRequestBuilder) {
    return NewMeAssignmentsItemResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rubric provides operations to manage the rubric property of the microsoft.graph.educationAssignment entity.
// returns a *MeAssignmentsItemRubricRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Rubric()(*MeAssignmentsItemRubricRequestBuilder) {
    return NewMeAssignmentsItemRubricRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetUpFeedbackResourcesFolder provides operations to call the setUpFeedbackResourcesFolder method.
// returns a *MeAssignmentsItemSetupfeedbackresourcesfolderSetUpFeedbackResourcesFolderRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) SetUpFeedbackResourcesFolder()(*MeAssignmentsItemSetupfeedbackresourcesfolderSetUpFeedbackResourcesFolderRequestBuilder) {
    return NewMeAssignmentsItemSetupfeedbackresourcesfolderSetUpFeedbackResourcesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
// returns a *MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) SetUpResourcesFolder()(*MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) {
    return NewMeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Submissions provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
// returns a *MeAssignmentsItemSubmissionsRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Submissions()(*MeAssignmentsItemSubmissionsRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignments for education
// returns a *RequestInformation when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of assignments for the user. Nullable.
// returns a *RequestInformation when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in education
// returns a *RequestInformation when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MeAssignmentsEducationAssignmentItemRequestBuilder when successful
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*MeAssignmentsEducationAssignmentItemRequestBuilder) {
    return NewMeAssignmentsEducationAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
