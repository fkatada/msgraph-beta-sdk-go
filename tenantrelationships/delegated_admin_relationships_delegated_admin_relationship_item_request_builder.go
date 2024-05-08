package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder provides operations to manage the delegatedAdminRelationships property of the microsoft.graph.tenantRelationship entity.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters read the properties of a delegatedAdminRelationship object.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters
}
// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessAssignments provides operations to manage the accessAssignments property of the microsoft.graph.delegatedAdminRelationship entity.
// returns a *DelegatedAdminRelationshipsItemAccessAssignmentsRequestBuilder when successful
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) AccessAssignments()(*DelegatedAdminRelationshipsItemAccessAssignmentsRequestBuilder) {
    return NewDelegatedAdminRelationshipsItemAccessAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal instantiates a new DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    m := &DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder instantiates a new DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a delegatedAdminRelationship object. A relationship can only be deleted if it's in the 'created' status. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationship-delete?view=graph-rest-beta
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties of a delegatedAdminRelationship object.
// returns a DelegatedAdminRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationship-get?view=graph-rest-beta
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable), nil
}
// Operations provides operations to manage the operations property of the microsoft.graph.delegatedAdminRelationship entity.
// returns a *DelegatedAdminRelationshipsItemOperationsRequestBuilder when successful
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Operations()(*DelegatedAdminRelationshipsItemOperationsRequestBuilder) {
    return NewDelegatedAdminRelationshipsItemOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a delegatedAdminRelationship object.  The following restrictions apply:- You can update this relationship when its status property is created.- You can update the autoExtendDuration property when status is either created or active.- You can only remove the Microsoft Entra Global Administrator role when the status property is active, which indicates a long-running operation.
// returns a DelegatedAdminRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationship-update?view=graph-rest-beta
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable), nil
}
// Requests provides operations to manage the requests property of the microsoft.graph.delegatedAdminRelationship entity.
// returns a *DelegatedAdminRelationshipsItemRequestsRequestBuilder when successful
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Requests()(*DelegatedAdminRelationshipsItemRequestsRequestBuilder) {
    return NewDelegatedAdminRelationshipsItemRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a delegatedAdminRelationship object. A relationship can only be deleted if it's in the 'created' status. 
// returns a *RequestInformation when successful
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties of a delegatedAdminRelationship object.
// returns a *RequestInformation when successful
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a delegatedAdminRelationship object.  The following restrictions apply:- You can update this relationship when its status property is created.- You can update the autoExtendDuration property when status is either created or active.- You can only remove the Microsoft Entra Global Administrator role when the status property is active, which indicates a long-running operation.
// returns a *RequestInformation when successful
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder when successful
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) WithUrl(rawUrl string)(*DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    return NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
