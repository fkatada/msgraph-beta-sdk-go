package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder provides operations to manage the media for the cloudCommunications entity.
type OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderInternal instantiates a new OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder and sets the default values.
func NewOnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) {
    m := &OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/broadcastRecording", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder instantiates a new OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder and sets the default values.
func NewOnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete broadcastRecording for the navigation property onlineMeetings in communications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get broadcastRecording for the navigation property onlineMeetings from communications
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put update broadcastRecording for the navigation property onlineMeetings in communications
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToDeleteRequestInformation delete broadcastRecording for the navigation property onlineMeetings in communications
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get broadcastRecording for the navigation property onlineMeetings from communications
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update broadcastRecording for the navigation property onlineMeetings in communications
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder when successful
func (m *OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder) {
    return NewOnlinemeetingsItemBroadcastrecordingBroadcastRecordingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
