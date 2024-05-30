package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder provides operations to count the resources in the collection.
type ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderGetQueryParameters get the number of the resource
type ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderGetQueryParameters
}
// NewItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderInternal instantiates a new ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder and sets the default values.
func NewItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder) {
    m := &ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder instantiates a new ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder and sets the default values.
func NewItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder when successful
func (m *ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder) {
    return NewItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
