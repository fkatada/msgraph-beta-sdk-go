package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder provides operations to call the copyToSection method.
type ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilderInternal instantiates a new ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder) {
    m := &ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/copyToSection", pathParameters),
    }
    return m
}
// NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder instantiates a new ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post copy a page to a specific section. For copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result.
// returns a OnenoteOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/page-copytosection?view=graph-rest-beta
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder) Post(ctx context.Context, body ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteOperationable), nil
}
// ToPostRequestInformation copy a page to a specific section. For copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result.
// returns a *RequestInformation when successful
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder when successful
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder) {
    return NewItemOnenoteNotebooksItemSectiongroupsItemSectionsItemPagesItemCopytosectionCopyToSectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
