package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder provides operations to call the setUpResourcesFolder method.
type MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderInternal instantiates a new MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder and sets the default values.
func NewMeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) {
    m := &MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/setUpResourcesFolder", pathParameters),
    }
    return m
}
// NewMeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder instantiates a new MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder and sets the default values.
func NewMeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a SharePoint folder to upload files for a given educationAssignment. Only teachers can perform this operation. The teacher determines the resources to upload in the assignment's folder.
// returns a EducationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationassignment-setupresourcesfolder?view=graph-rest-beta
func (m *MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) Post(ctx context.Context, requestConfiguration *MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation create a SharePoint folder to upload files for a given educationAssignment. Only teachers can perform this operation. The teacher determines the resources to upload in the assignment's folder.
// returns a *RequestInformation when successful
func (m *MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder when successful
func (m *MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) WithUrl(rawUrl string)(*MeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) {
    return NewMeAssignmentsItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
