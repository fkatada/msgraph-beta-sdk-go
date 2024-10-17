package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MailboxProtectionUnitsBulkAdditionJobCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewMailboxProtectionUnitsBulkAdditionJobCollectionResponse instantiates a new MailboxProtectionUnitsBulkAdditionJobCollectionResponse and sets the default values.
func NewMailboxProtectionUnitsBulkAdditionJobCollectionResponse()(*MailboxProtectionUnitsBulkAdditionJobCollectionResponse) {
    m := &MailboxProtectionUnitsBulkAdditionJobCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMailboxProtectionUnitsBulkAdditionJobCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMailboxProtectionUnitsBulkAdditionJobCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxProtectionUnitsBulkAdditionJobCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MailboxProtectionUnitsBulkAdditionJobCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMailboxProtectionUnitsBulkAdditionJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailboxProtectionUnitsBulkAdditionJobable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MailboxProtectionUnitsBulkAdditionJobable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []MailboxProtectionUnitsBulkAdditionJobable when successful
func (m *MailboxProtectionUnitsBulkAdditionJobCollectionResponse) GetValue()([]MailboxProtectionUnitsBulkAdditionJobable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MailboxProtectionUnitsBulkAdditionJobable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MailboxProtectionUnitsBulkAdditionJobCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MailboxProtectionUnitsBulkAdditionJobCollectionResponse) SetValue(value []MailboxProtectionUnitsBulkAdditionJobable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type MailboxProtectionUnitsBulkAdditionJobCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]MailboxProtectionUnitsBulkAdditionJobable)
    SetValue(value []MailboxProtectionUnitsBulkAdditionJobable)()
}
