package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeletedItemContainer 
type DeletedItemContainer struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewDeletedItemContainer instantiates a new deletedItemContainer and sets the default values.
func NewDeletedItemContainer()(*DeletedItemContainer) {
    m := &DeletedItemContainer{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeletedItemContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletedItemContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedItemContainer(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeletedItemContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeletedItemContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// DeletedItemContainerable 
type DeletedItemContainerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
