package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponseable instead.
type FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse struct {
    FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponse
}
// NewFilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse instantiates a new FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse()(*FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse) {
    m := &FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse{
        FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponse: *NewFilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponseable instead.
type FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponseable interface {
    FilestorageDeletedcontainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
