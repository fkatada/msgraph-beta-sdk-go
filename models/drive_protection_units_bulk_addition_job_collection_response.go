package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DriveProtectionUnitsBulkAdditionJobCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewDriveProtectionUnitsBulkAdditionJobCollectionResponse instantiates a new DriveProtectionUnitsBulkAdditionJobCollectionResponse and sets the default values.
func NewDriveProtectionUnitsBulkAdditionJobCollectionResponse()(*DriveProtectionUnitsBulkAdditionJobCollectionResponse) {
    m := &DriveProtectionUnitsBulkAdditionJobCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDriveProtectionUnitsBulkAdditionJobCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDriveProtectionUnitsBulkAdditionJobCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveProtectionUnitsBulkAdditionJobCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DriveProtectionUnitsBulkAdditionJobCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveProtectionUnitsBulkAdditionJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveProtectionUnitsBulkAdditionJobable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DriveProtectionUnitsBulkAdditionJobable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []DriveProtectionUnitsBulkAdditionJobable when successful
func (m *DriveProtectionUnitsBulkAdditionJobCollectionResponse) GetValue()([]DriveProtectionUnitsBulkAdditionJobable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DriveProtectionUnitsBulkAdditionJobable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DriveProtectionUnitsBulkAdditionJobCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DriveProtectionUnitsBulkAdditionJobCollectionResponse) SetValue(value []DriveProtectionUnitsBulkAdditionJobable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type DriveProtectionUnitsBulkAdditionJobCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]DriveProtectionUnitsBulkAdditionJobable)
    SetValue(value []DriveProtectionUnitsBulkAdditionJobable)()
}
