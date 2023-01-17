package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody 
type WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ids property
    ids []string
    // The memberEntityType property
    memberEntityType *string
}
// NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody instantiates a new WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody()(*WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) {
    m := &WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIds(res)
        }
        return nil
    }
    res["memberEntityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEntityType(val)
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) GetIds()([]string) {
    return m.ids
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) GetMemberEntityType()(*string) {
    return m.memberEntityType
}
// Serialize serializes information the current object
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memberEntityType", m.GetMemberEntityType())
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
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
