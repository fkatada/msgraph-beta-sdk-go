package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAndAppManagementAssignmentFilterCollectionResponse 
type DeviceAndAppManagementAssignmentFilterCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewDeviceAndAppManagementAssignmentFilterCollectionResponse instantiates a new deviceAndAppManagementAssignmentFilterCollectionResponse and sets the default values.
func NewDeviceAndAppManagementAssignmentFilterCollectionResponse()(*DeviceAndAppManagementAssignmentFilterCollectionResponse) {
    m := &DeviceAndAppManagementAssignmentFilterCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceAndAppManagementAssignmentFilterCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAndAppManagementAssignmentFilterCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAndAppManagementAssignmentFilterCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementAssignmentFilterCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAndAppManagementAssignmentFilterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceAndAppManagementAssignmentFilterable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceAndAppManagementAssignmentFilterCollectionResponse) GetValue()([]DeviceAndAppManagementAssignmentFilterable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceAndAppManagementAssignmentFilterable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignmentFilterCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceAndAppManagementAssignmentFilterCollectionResponse) SetValue(value []DeviceAndAppManagementAssignmentFilterable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// DeviceAndAppManagementAssignmentFilterCollectionResponseable 
type DeviceAndAppManagementAssignmentFilterCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]DeviceAndAppManagementAssignmentFilterable)
    SetValue(value []DeviceAndAppManagementAssignmentFilterable)()
}
