package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody instantiates a new ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody and sets the default values.
func NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody()(*ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) {
    m := &ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementActionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementActionId(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["managementTemplateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateVersion(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["tenantGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantGroupId(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetManagementActionId gets the managementActionId property value. The managementActionId property
// returns a *string when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetManagementActionId()(*string) {
    val, err := m.GetBackingStore().Get("managementActionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
// returns a *string when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetManagementTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateVersion gets the managementTemplateVersion property value. The managementTemplateVersion property
// returns a *int32 when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetManagementTemplateVersion()(*int32) {
    val, err := m.GetBackingStore().Get("managementTemplateVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantGroupId gets the tenantGroupId property value. The tenantGroupId property
// returns a *string when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetTenantGroupId()(*string) {
    val, err := m.GetBackingStore().Get("tenantGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
// returns a *string when successful
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementActionId", m.GetManagementActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("managementTemplateVersion", m.GetManagementTemplateVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantGroupId", m.GetTenantGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetManagementActionId sets the managementActionId property value. The managementActionId property
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) SetManagementActionId(value *string)() {
    err := m.GetBackingStore().Set("managementActionId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) SetManagementTemplateId(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateVersion sets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) SetManagementTemplateVersion(value *int32)() {
    err := m.GetBackingStore().Set("managementTemplateVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantGroupId sets the tenantGroupId property value. The tenantGroupId property
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) SetTenantGroupId(value *string)() {
    err := m.GetBackingStore().Set("tenantGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBody) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
type ManagedtenantsManagementactiontenantdeploymentstatusesMicrosoftgraphmanagedtenantschangedeploymentstatusChangeDeploymentStatusPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetManagementActionId()(*string)
    GetManagementTemplateId()(*string)
    GetManagementTemplateVersion()(*int32)
    GetStatus()(*string)
    GetTenantGroupId()(*string)
    GetTenantId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetManagementActionId(value *string)()
    SetManagementTemplateId(value *string)()
    SetManagementTemplateVersion(value *int32)()
    SetStatus(value *string)()
    SetTenantGroupId(value *string)()
    SetTenantId(value *string)()
}
