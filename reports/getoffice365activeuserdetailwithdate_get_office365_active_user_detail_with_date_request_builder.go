package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder provides operations to call the getOffice365ActiveUserDetail method.
type Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderGetQueryParameters invoke function getOffice365ActiveUserDetail
type Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderGetQueryParameters struct {
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
// Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderGetQueryParameters
}
// NewGetoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderInternal instantiates a new Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder and sets the default values.
func NewGetoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder) {
    m := &Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365ActiveUserDetail(date={date}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder instantiates a new Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder and sets the default values.
func NewGetoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365ActiveUserDetail
// Deprecated: This method is obsolete. Use GetAsGetOffice365ActiveUserDetailWithDateGetResponse instead.
// returns a Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderGetRequestConfiguration)(Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateResponseable), nil
}
// GetAsGetOffice365ActiveUserDetailWithDateGetResponse invoke function getOffice365ActiveUserDetail
// returns a Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder) GetAsGetOffice365ActiveUserDetailWithDateGetResponse(ctx context.Context, requestConfiguration *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderGetRequestConfiguration)(Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365ActiveUserDetail
// returns a *RequestInformation when successful
func (m *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder when successful
func (m *Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*Getoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return NewGetoffice365activeuserdetailwithdateGetOffice365ActiveUserDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
