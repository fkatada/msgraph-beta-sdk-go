package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminWindowsUpdates 
type AdminWindowsUpdates struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewAdminWindowsUpdates instantiates a new adminWindowsUpdates and sets the default values.
func NewAdminWindowsUpdates()(*AdminWindowsUpdates) {
    m := &AdminWindowsUpdates{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminWindowsUpdatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminWindowsUpdatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminWindowsUpdates(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminWindowsUpdates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AdminWindowsUpdates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AdminWindowsUpdatesable 
type AdminWindowsUpdatesable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
