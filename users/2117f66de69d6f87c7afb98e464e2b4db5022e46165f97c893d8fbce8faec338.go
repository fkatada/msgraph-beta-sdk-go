package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetQueryParameters the section that contains the page. Read-only.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetQueryParameters
}
// NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal instantiates a new ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) {
    m := &ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/parentSection{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder instantiates a new ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the section that contains the page. Read-only.
// returns a OnenoteSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable), nil
}
// ToGetRequestInformation the section that contains the page. Read-only.
// returns a *RequestInformation when successful
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder when successful
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) {
    return NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
