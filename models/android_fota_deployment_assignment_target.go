package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidFotaDeploymentAssignmentTarget the AAD Group we are deploying firmware updates to
type AndroidFotaDeploymentAssignmentTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // AAD Group Id.
    groupId *string
    // The OdataType property
    odataType *string
}
// NewAndroidFotaDeploymentAssignmentTarget instantiates a new androidFotaDeploymentAssignmentTarget and sets the default values.
func NewAndroidFotaDeploymentAssignmentTarget()(*AndroidFotaDeploymentAssignmentTarget) {
    m := &AndroidFotaDeploymentAssignmentTarget{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateAndroidFotaDeploymentAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidFotaDeploymentAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidFotaDeploymentAssignmentTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidFotaDeploymentAssignmentTarget) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidFotaDeploymentAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
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
// GetGroupId gets the groupId property value. AAD Group Id.
func (m *AndroidFotaDeploymentAssignmentTarget) GetGroupId()(*string) {
    return m.groupId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidFotaDeploymentAssignmentTarget) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AndroidFotaDeploymentAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidFotaDeploymentAssignmentTarget) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupId sets the groupId property value. AAD Group Id.
func (m *AndroidFotaDeploymentAssignmentTarget) SetGroupId(value *string)() {
    m.groupId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidFotaDeploymentAssignmentTarget) SetOdataType(value *string)() {
    m.odataType = value
}
