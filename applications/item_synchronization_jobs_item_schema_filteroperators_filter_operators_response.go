package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsGetResponseable instead.
type ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponse struct {
    ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsGetResponse
}
// NewItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponse instantiates a new ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponse and sets the default values.
func NewItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponse()(*ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponse) {
    m := &ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponse{
        ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsGetResponse: *NewItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsGetResponse(),
    }
    return m
}
// CreateItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsGetResponseable instead.
type ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsResponseable interface {
    ItemSynchronizationJobsItemSchemaFilteroperatorsFilterOperatorsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
