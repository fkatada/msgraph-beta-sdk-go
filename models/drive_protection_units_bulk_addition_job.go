package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DriveProtectionUnitsBulkAdditionJob struct {
    ProtectionUnitsBulkJobBase
}
// NewDriveProtectionUnitsBulkAdditionJob instantiates a new DriveProtectionUnitsBulkAdditionJob and sets the default values.
func NewDriveProtectionUnitsBulkAdditionJob()(*DriveProtectionUnitsBulkAdditionJob) {
    m := &DriveProtectionUnitsBulkAdditionJob{
        ProtectionUnitsBulkJobBase: *NewProtectionUnitsBulkJobBase(),
    }
    odataTypeValue := "#microsoft.graph.driveProtectionUnitsBulkAdditionJob"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDriveProtectionUnitsBulkAdditionJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDriveProtectionUnitsBulkAdditionJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveProtectionUnitsBulkAdditionJob(), nil
}
// GetDirectoryObjectIds gets the directoryObjectIds property value. The list of OneDrive directoryObjectIds to add to the OneDrive protection policy.
// returns a []string when successful
func (m *DriveProtectionUnitsBulkAdditionJob) GetDirectoryObjectIds()([]string) {
    val, err := m.GetBackingStore().Get("directoryObjectIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDrives gets the drives property value. The list of email addresses to add to the OneDrive protection policy.
// returns a []string when successful
func (m *DriveProtectionUnitsBulkAdditionJob) GetDrives()([]string) {
    val, err := m.GetBackingStore().Get("drives")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DriveProtectionUnitsBulkAdditionJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectionUnitsBulkJobBase.GetFieldDeserializers()
    res["directoryObjectIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDirectoryObjectIds(res)
        }
        return nil
    }
    res["drives"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDrives(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DriveProtectionUnitsBulkAdditionJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectionUnitsBulkJobBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDirectoryObjectIds() != nil {
        err = writer.WriteCollectionOfStringValues("directoryObjectIds", m.GetDirectoryObjectIds())
        if err != nil {
            return err
        }
    }
    if m.GetDrives() != nil {
        err = writer.WriteCollectionOfStringValues("drives", m.GetDrives())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectoryObjectIds sets the directoryObjectIds property value. The list of OneDrive directoryObjectIds to add to the OneDrive protection policy.
func (m *DriveProtectionUnitsBulkAdditionJob) SetDirectoryObjectIds(value []string)() {
    err := m.GetBackingStore().Set("directoryObjectIds", value)
    if err != nil {
        panic(err)
    }
}
// SetDrives sets the drives property value. The list of email addresses to add to the OneDrive protection policy.
func (m *DriveProtectionUnitsBulkAdditionJob) SetDrives(value []string)() {
    err := m.GetBackingStore().Set("drives", value)
    if err != nil {
        panic(err)
    }
}
type DriveProtectionUnitsBulkAdditionJobable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectionUnitsBulkJobBaseable
    GetDirectoryObjectIds()([]string)
    GetDrives()([]string)
    SetDirectoryObjectIds(value []string)()
    SetDrives(value []string)()
}
