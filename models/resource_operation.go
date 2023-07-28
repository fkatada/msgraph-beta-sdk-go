package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceOperation describes the resourceOperation resource (entity) of the Microsoft Graph API (REST), which supports Intune workflows related to role-based access control (RBAC).
type ResourceOperation struct {
    Entity
}
// NewResourceOperation instantiates a new resourceOperation and sets the default values.
func NewResourceOperation()(*ResourceOperation) {
    m := &ResourceOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateResourceOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceOperation(), nil
}
// GetActionName gets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) GetActionName()(*string) {
    val, err := m.GetBackingStore().Get("actionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnabledForScopeValidation gets the enabledForScopeValidation property value. Determines whether the Permission is validated for Scopes defined per Role Assignment. This property is read-only.
func (m *ResourceOperation) GetEnabledForScopeValidation()(*bool) {
    val, err := m.GetBackingStore().Get("enabledForScopeValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["enabledForScopeValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabledForScopeValidation(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    res["resourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. Resource category to which this Operation belongs. This property is read-only.
func (m *ResourceOperation) GetResource()(*string) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceName gets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) GetResourceName()(*string) {
    val, err := m.GetBackingStore().Get("resourceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ResourceOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionName sets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) SetActionName(value *string)() {
    err := m.GetBackingStore().Set("actionName", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabledForScopeValidation sets the enabledForScopeValidation property value. Determines whether the Permission is validated for Scopes defined per Role Assignment. This property is read-only.
func (m *ResourceOperation) SetEnabledForScopeValidation(value *bool)() {
    err := m.GetBackingStore().Set("enabledForScopeValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetResource sets the resource property value. Resource category to which this Operation belongs. This property is read-only.
func (m *ResourceOperation) SetResource(value *string)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceName sets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) SetResourceName(value *string)() {
    err := m.GetBackingStore().Set("resourceName", value)
    if err != nil {
        panic(err)
    }
}
// ResourceOperationable 
type ResourceOperationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionName()(*string)
    GetDescription()(*string)
    GetEnabledForScopeValidation()(*bool)
    GetResource()(*string)
    GetResourceName()(*string)
    SetActionName(value *string)()
    SetDescription(value *string)()
    SetEnabledForScopeValidation(value *bool)()
    SetResource(value *string)()
    SetResourceName(value *string)()
}
