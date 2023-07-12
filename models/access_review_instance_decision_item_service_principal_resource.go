package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItemServicePrincipalResource 
type AccessReviewInstanceDecisionItemServicePrincipalResource struct {
    AccessReviewInstanceDecisionItemResource
}
// NewAccessReviewInstanceDecisionItemServicePrincipalResource instantiates a new accessReviewInstanceDecisionItemServicePrincipalResource and sets the default values.
func NewAccessReviewInstanceDecisionItemServicePrincipalResource()(*AccessReviewInstanceDecisionItemServicePrincipalResource) {
    m := &AccessReviewInstanceDecisionItemServicePrincipalResource{
        AccessReviewInstanceDecisionItemResource: *NewAccessReviewInstanceDecisionItemResource(),
    }
    odataTypeValue := "#microsoft.graph.accessReviewInstanceDecisionItemServicePrincipalResource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessReviewInstanceDecisionItemServicePrincipalResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceDecisionItemServicePrincipalResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewInstanceDecisionItemServicePrincipalResource(), nil
}
// GetAppId gets the appId property value. The globally unique identifier of the application to which access has been granted.
func (m *AccessReviewInstanceDecisionItemServicePrincipalResource) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstanceDecisionItemServicePrincipalResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessReviewInstanceDecisionItemResource.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewInstanceDecisionItemServicePrincipalResource) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessReviewInstanceDecisionItemServicePrincipalResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessReviewInstanceDecisionItemResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The globally unique identifier of the application to which access has been granted.
func (m *AccessReviewInstanceDecisionItemServicePrincipalResource) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewInstanceDecisionItemServicePrincipalResource) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AccessReviewInstanceDecisionItemServicePrincipalResourceable 
type AccessReviewInstanceDecisionItemServicePrincipalResourceable interface {
    AccessReviewInstanceDecisionItemResourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetOdataType()(*string)
    SetAppId(value *string)()
    SetOdataType(value *string)()
}
