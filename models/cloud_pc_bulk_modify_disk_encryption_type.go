package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcBulkModifyDiskEncryptionType 
type CloudPcBulkModifyDiskEncryptionType struct {
    CloudPcBulkAction
}
// NewCloudPcBulkModifyDiskEncryptionType instantiates a new cloudPcBulkModifyDiskEncryptionType and sets the default values.
func NewCloudPcBulkModifyDiskEncryptionType()(*CloudPcBulkModifyDiskEncryptionType) {
    m := &CloudPcBulkModifyDiskEncryptionType{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkModifyDiskEncryptionType"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkModifyDiskEncryptionTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcBulkModifyDiskEncryptionTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkModifyDiskEncryptionType(), nil
}
// GetDiskEncryptionType gets the diskEncryptionType property value. The diskEncryptionType property
func (m *CloudPcBulkModifyDiskEncryptionType) GetDiskEncryptionType()(*CloudPcDiskEncryptionType) {
    val, err := m.GetBackingStore().Get("diskEncryptionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcDiskEncryptionType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcBulkModifyDiskEncryptionType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    res["diskEncryptionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDiskEncryptionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskEncryptionType(val.(*CloudPcDiskEncryptionType))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CloudPcBulkModifyDiskEncryptionType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDiskEncryptionType() != nil {
        cast := (*m.GetDiskEncryptionType()).String()
        err = writer.WriteStringValue("diskEncryptionType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDiskEncryptionType sets the diskEncryptionType property value. The diskEncryptionType property
func (m *CloudPcBulkModifyDiskEncryptionType) SetDiskEncryptionType(value *CloudPcDiskEncryptionType)() {
    err := m.GetBackingStore().Set("diskEncryptionType", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcBulkModifyDiskEncryptionTypeable 
type CloudPcBulkModifyDiskEncryptionTypeable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDiskEncryptionType()(*CloudPcDiskEncryptionType)
    SetDiskEncryptionType(value *CloudPcDiskEncryptionType)()
}
