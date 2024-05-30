package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder provides operations to manage the dayNotes property of the microsoft.graph.schedule entity.
type ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderGetQueryParameters the day notes in the schedule.
type ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderGetQueryParameters struct {
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
// ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderGetQueryParameters
}
// ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDayNoteId provides operations to manage the dayNotes property of the microsoft.graph.schedule entity.
// returns a *ItemTeamdefinitionScheduleDaynotesDayNoteItemRequestBuilder when successful
func (m *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) ByDayNoteId(dayNoteId string)(*ItemTeamdefinitionScheduleDaynotesDayNoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if dayNoteId != "" {
        urlTplParams["dayNote%2Did"] = dayNoteId
    }
    return NewItemTeamdefinitionScheduleDaynotesDayNoteItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderInternal instantiates a new ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder and sets the default values.
func NewItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) {
    m := &ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/dayNotes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder instantiates a new ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder and sets the default values.
func NewItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamdefinitionScheduleDaynotesCountRequestBuilder when successful
func (m *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) Count()(*ItemTeamdefinitionScheduleDaynotesCountRequestBuilder) {
    return NewItemTeamdefinitionScheduleDaynotesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the day notes in the schedule.
// returns a DayNoteCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DayNoteCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDayNoteCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DayNoteCollectionResponseable), nil
}
// Post create new navigation property to dayNotes for teamTemplateDefinition
// returns a DayNoteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DayNoteable, requestConfiguration *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DayNoteable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDayNoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DayNoteable), nil
}
// ToGetRequestInformation the day notes in the schedule.
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dayNotes for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DayNoteable, requestConfiguration *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder when successful
func (m *ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder) {
    return NewItemTeamdefinitionScheduleDaynotesDayNotesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
