package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder provides operations to manage the media for the user entity.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderInternal instantiates a new ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) {
    m := &ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/content", pathParameters),
    }
    return m
}
// NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder instantiates a new ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the page's HTML content.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the page's HTML content.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put the page's HTML content.
// returns a OnenotePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable), nil
}
// ToDeleteRequestInformation the page's HTML content.
// returns a *RequestInformation when successful
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the page's HTML content.
// returns a *RequestInformation when successful
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the page's HTML content.
// returns a *RequestInformation when successful
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder when successful
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) {
    return NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
