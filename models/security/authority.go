package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Authority 
type Authority struct {
    FilePlanDescriptorBase
    // The OdataType property
    OdataType *string
}
// NewAuthority instantiates a new authority and sets the default values.
func NewAuthority()(*Authority) {
    m := &Authority{
        FilePlanDescriptorBase: *NewFilePlanDescriptorBase(),
    }
    return m
}
// CreateAuthorityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthorityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthority(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Authority) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FilePlanDescriptorBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Authority) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FilePlanDescriptorBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// Authorityable 
type Authorityable interface {
    FilePlanDescriptorBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
