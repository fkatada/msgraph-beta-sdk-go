package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponse struct {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponse
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponse instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponse()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponse) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponse{
        FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponse: *NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponseable interface {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
