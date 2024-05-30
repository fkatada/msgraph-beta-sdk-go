package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder provides operations to manage the linkedEligibleRoleAssignment property of the microsoft.graph.governanceRoleAssignment entity.
type ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderGetQueryParameters read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
type ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderGetQueryParameters
}
// NewItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderInternal instantiates a new ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder and sets the default values.
func NewItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder) {
    m := &ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}/roleAssignments/{governanceRoleAssignment%2Did}/linkedEligibleRoleAssignment{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder instantiates a new ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder and sets the default values.
func NewItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
// returns a GovernanceRoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable), nil
}
// ToGetRequestInformation read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
// returns a *RequestInformation when successful
func (m *ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder when successful
func (m *ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder) WithUrl(rawUrl string)(*ItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder) {
    return NewItemResourcesItemRoleassignmentsItemLinkedeligibleroleassignmentLinkedEligibleRoleAssignmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
