package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse struct {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponse
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse{
        FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponse: *NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountResponseable interface {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemTasksItemChangesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
