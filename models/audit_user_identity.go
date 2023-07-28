package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditUserIdentity 
type AuditUserIdentity struct {
    UserIdentity
}
// NewAuditUserIdentity instantiates a new auditUserIdentity and sets the default values.
func NewAuditUserIdentity()(*AuditUserIdentity) {
    m := &AuditUserIdentity{
        UserIdentity: *NewUserIdentity(),
    }
    odataTypeValue := "#microsoft.graph.auditUserIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAuditUserIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditUserIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditUserIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditUserIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserIdentity.GetFieldDeserializers()
    res["homeTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantId(val)
        }
        return nil
    }
    res["homeTenantName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantName(val)
        }
        return nil
    }
    return res
}
// GetHomeTenantId gets the homeTenantId property value. For user sign ins, the identifier of the tenant that the user is a member of.
func (m *AuditUserIdentity) GetHomeTenantId()(*string) {
    val, err := m.GetBackingStore().Get("homeTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHomeTenantName gets the homeTenantName property value. For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *AuditUserIdentity) GetHomeTenantName()(*string) {
    val, err := m.GetBackingStore().Get("homeTenantName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuditUserIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("homeTenantId", m.GetHomeTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homeTenantName", m.GetHomeTenantName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHomeTenantId sets the homeTenantId property value. For user sign ins, the identifier of the tenant that the user is a member of.
func (m *AuditUserIdentity) SetHomeTenantId(value *string)() {
    err := m.GetBackingStore().Set("homeTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeTenantName sets the homeTenantName property value. For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *AuditUserIdentity) SetHomeTenantName(value *string)() {
    err := m.GetBackingStore().Set("homeTenantName", value)
    if err != nil {
        panic(err)
    }
}
// AuditUserIdentityable 
type AuditUserIdentityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserIdentityable
    GetHomeTenantId()(*string)
    GetHomeTenantName()(*string)
    SetHomeTenantId(value *string)()
    SetHomeTenantName(value *string)()
}
