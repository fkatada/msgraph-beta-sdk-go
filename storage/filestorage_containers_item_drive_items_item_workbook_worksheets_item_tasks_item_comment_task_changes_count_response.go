package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponseable instead.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse struct {
    FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponse
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse{
        FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponse: *NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponseable instead.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountResponseable interface {
    FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemCommentTaskChangesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
