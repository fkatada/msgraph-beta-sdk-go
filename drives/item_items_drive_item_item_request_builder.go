package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type ItemItemsDriveItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type ItemItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsDriveItemItemRequestBuilderGetQueryParameters
}
// ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemActivitiesRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Activities()(*ItemItemsItemActivitiesRequestBuilder) {
    return NewItemItemsItemActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemAnalyticsRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Analytics()(*ItemItemsItemAnalyticsRequestBuilder) {
    return NewItemItemsItemAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
// returns a *ItemItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) AssignSensitivityLabel()(*ItemItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilder) {
    return NewItemItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkin provides operations to call the checkin method.
// returns a *ItemItemsItemCheckinRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Checkin()(*ItemItemsItemCheckinRequestBuilder) {
    return NewItemItemsItemCheckinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkout provides operations to call the checkout method.
// returns a *ItemItemsItemCheckoutRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Checkout()(*ItemItemsItemCheckoutRequestBuilder) {
    return NewItemItemsItemCheckoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemChildrenRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Children()(*ItemItemsItemChildrenRequestBuilder) {
    return NewItemItemsItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsDriveItemItemRequestBuilderInternal instantiates a new ItemItemsDriveItemItemRequestBuilder and sets the default values.
func NewItemItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsDriveItemItemRequestBuilder) {
    m := &ItemItemsDriveItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsDriveItemItemRequestBuilder instantiates a new ItemItemsDriveItemItemRequestBuilder and sets the default values.
func NewItemItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the drive entity.
// returns a *ItemItemsItemContentRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Content()(*ItemItemsItemContentRequestBuilder) {
    return NewItemItemsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ContentStream provides operations to manage the media for the drive entity.
// returns a *ItemItemsItemContentstreamContentStreamRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) ContentStream()(*ItemItemsItemContentstreamContentStreamRequestBuilder) {
    return NewItemItemsItemContentstreamContentStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *ItemItemsItemCopyRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Copy()(*ItemItemsItemCopyRequestBuilder) {
    return NewItemItemsItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemItemsItemCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) CreatedByUser()(*ItemItemsItemCreatedbyuserCreatedByUserRequestBuilder) {
    return NewItemItemsItemCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateLink provides operations to call the createLink method.
// returns a *ItemItemsItemCreatelinkCreateLinkRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) CreateLink()(*ItemItemsItemCreatelinkCreateLinkRequestBuilder) {
    return NewItemItemsItemCreatelinkCreateLinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *ItemItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) CreateUploadSession()(*ItemItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewItemItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property items for drives
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Delta provides operations to call the delta method.
// returns a *ItemItemsItemDeltaRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Delta()(*ItemItemsItemDeltaRequestBuilder) {
    return NewItemItemsItemDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeltaWithToken provides operations to call the delta method.
// returns a *ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) DeltaWithToken(token *string)(*ItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    return NewItemItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, token)
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
// returns a *ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) ExtractSensitivityLabels()(*ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    return NewItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Follow provides operations to call the follow method.
// returns a *ItemItemsItemFollowRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Follow()(*ItemItemsItemFollowRequestBuilder) {
    return NewItemItemsItemFollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get all items contained in the drive. Read-only. Nullable.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// returns a *ItemItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// Invite provides operations to call the invite method.
// returns a *ItemItemsItemInviteRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Invite()(*ItemItemsItemInviteRequestBuilder) {
    return NewItemItemsItemInviteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) LastModifiedByUser()(*ItemItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewItemItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemListitemListItemRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) ListItem()(*ItemItemsItemListitemListItemRequestBuilder) {
    return NewItemItemsItemListitemListItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property items in drives
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemItemsItemPermanentdeletePermanentDeleteRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) PermanentDelete()(*ItemItemsItemPermanentdeletePermanentDeleteRequestBuilder) {
    return NewItemItemsItemPermanentdeletePermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemPermissionsRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Permissions()(*ItemItemsItemPermissionsRequestBuilder) {
    return NewItemItemsItemPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Preview provides operations to call the preview method.
// returns a *ItemItemsItemPreviewRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Preview()(*ItemItemsItemPreviewRequestBuilder) {
    return NewItemItemsItemPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *ItemItemsItemRestoreRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Restore()(*ItemItemsItemRestoreRequestBuilder) {
    return NewItemItemsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetentionLabel provides operations to manage the retentionLabel property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemRetentionlabelRetentionLabelRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) RetentionLabel()(*ItemItemsItemRetentionlabelRetentionLabelRequestBuilder) {
    return NewItemItemsItemRetentionlabelRetentionLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchWithQ provides operations to call the search method.
// returns a *ItemItemsItemSearchwithqSearchWithQRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) SearchWithQ(q *string)(*ItemItemsItemSearchwithqSearchWithQRequestBuilder) {
    return NewItemItemsItemSearchwithqSearchWithQRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, q)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemSubscriptionsRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Subscriptions()(*ItemItemsItemSubscriptionsRequestBuilder) {
    return NewItemItemsItemSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemThumbnailsRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Thumbnails()(*ItemItemsItemThumbnailsRequestBuilder) {
    return NewItemItemsItemThumbnailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property items for drives
// returns a *RequestInformation when successful
func (m *ItemItemsDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation all items contained in the drive. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemItemsDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property items in drives
// returns a *RequestInformation when successful
func (m *ItemItemsDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unfollow provides operations to call the unfollow method.
// returns a *ItemItemsItemUnfollowRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Unfollow()(*ItemItemsItemUnfollowRequestBuilder) {
    return NewItemItemsItemUnfollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ValidatePermission provides operations to call the validatePermission method.
// returns a *ItemItemsItemValidatepermissionValidatePermissionRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) ValidatePermission()(*ItemItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    return NewItemItemsItemValidatepermissionValidatePermissionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemVersionsRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Versions()(*ItemItemsItemVersionsRequestBuilder) {
    return NewItemItemsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsDriveItemItemRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemsDriveItemItemRequestBuilder) {
    return NewItemItemsDriveItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Workbook provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
// returns a *ItemItemsItemWorkbookRequestBuilder when successful
func (m *ItemItemsDriveItemItemRequestBuilder) Workbook()(*ItemItemsItemWorkbookRequestBuilder) {
    return NewItemItemsItemWorkbookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
