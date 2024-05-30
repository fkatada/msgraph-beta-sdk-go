package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder provides operations to manage the mobileAppIntentAndStates property of the microsoft.graph.user entity.
type ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderGetQueryParameters the list of troubleshooting events for this user.
type ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderGetQueryParameters struct {
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
// ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderGetQueryParameters
}
// ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileAppIntentAndStateId provides operations to manage the mobileAppIntentAndStates property of the microsoft.graph.user entity.
// returns a *ItemMobileappintentandstatesMobileAppIntentAndStateItemRequestBuilder when successful
func (m *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) ByMobileAppIntentAndStateId(mobileAppIntentAndStateId string)(*ItemMobileappintentandstatesMobileAppIntentAndStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileAppIntentAndStateId != "" {
        urlTplParams["mobileAppIntentAndState%2Did"] = mobileAppIntentAndStateId
    }
    return NewItemMobileappintentandstatesMobileAppIntentAndStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderInternal instantiates a new ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder and sets the default values.
func NewItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) {
    m := &ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mobileAppIntentAndStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder instantiates a new ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder and sets the default values.
func NewItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemMobileappintentandstatesCountRequestBuilder when successful
func (m *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) Count()(*ItemMobileappintentandstatesCountRequestBuilder) {
    return NewItemMobileappintentandstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of troubleshooting events for this user.
// returns a MobileAppIntentAndStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppIntentAndStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppIntentAndStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppIntentAndStateCollectionResponseable), nil
}
// Post create new navigation property to mobileAppIntentAndStates for users
// returns a MobileAppIntentAndStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppIntentAndStateable, requestConfiguration *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppIntentAndStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppIntentAndStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppIntentAndStateable), nil
}
// ToGetRequestInformation the list of troubleshooting events for this user.
// returns a *RequestInformation when successful
func (m *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to mobileAppIntentAndStates for users
// returns a *RequestInformation when successful
func (m *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppIntentAndStateable, requestConfiguration *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder when successful
func (m *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) WithUrl(rawUrl string)(*ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) {
    return NewItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
