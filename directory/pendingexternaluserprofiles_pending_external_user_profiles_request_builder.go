package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder provides operations to manage the pendingExternalUserProfiles property of the microsoft.graph.directory entity.
type PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderGetQueryParameters retrieve the properties of all pendingExternalUserProfiles.
type PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderGetQueryParameters
}
// PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPendingExternalUserProfileId provides operations to manage the pendingExternalUserProfiles property of the microsoft.graph.directory entity.
// returns a *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) ByPendingExternalUserProfileId(pendingExternalUserProfileId string)(*PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if pendingExternalUserProfileId != "" {
        urlTplParams["pendingExternalUserProfile%2Did"] = pendingExternalUserProfileId
    }
    return NewPendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderInternal instantiates a new PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder and sets the default values.
func NewPendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) {
    m := &PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/pendingExternalUserProfiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder instantiates a new PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder and sets the default values.
func NewPendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PendingexternaluserprofilesCountRequestBuilder when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) Count()(*PendingexternaluserprofilesCountRequestBuilder) {
    return NewPendingexternaluserprofilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties of all pendingExternalUserProfiles.
// returns a PendingExternalUserProfileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-list-pendingexternaluserprofile?view=graph-rest-beta
func (m *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePendingExternalUserProfileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileCollectionResponseable), nil
}
// Post create a new pendingExternalUserProfile object.
// returns a PendingExternalUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-post-pendingexternaluserprofile?view=graph-rest-beta
func (m *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePendingExternalUserProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable), nil
}
// ToGetRequestInformation retrieve the properties of all pendingExternalUserProfiles.
// returns a *RequestInformation when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new pendingExternalUserProfile object.
// returns a *RequestInformation when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) WithUrl(rawUrl string)(*PendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder) {
    return NewPendingexternaluserprofilesPendingExternalUserProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
