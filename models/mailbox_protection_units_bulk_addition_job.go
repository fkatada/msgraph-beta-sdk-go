package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MailboxProtectionUnitsBulkAdditionJob struct {
    ProtectionUnitsBulkJobBase
}
// NewMailboxProtectionUnitsBulkAdditionJob instantiates a new MailboxProtectionUnitsBulkAdditionJob and sets the default values.
func NewMailboxProtectionUnitsBulkAdditionJob()(*MailboxProtectionUnitsBulkAdditionJob) {
    m := &MailboxProtectionUnitsBulkAdditionJob{
        ProtectionUnitsBulkJobBase: *NewProtectionUnitsBulkJobBase(),
    }
    odataTypeValue := "#microsoft.graph.mailboxProtectionUnitsBulkAdditionJob"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMailboxProtectionUnitsBulkAdditionJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMailboxProtectionUnitsBulkAdditionJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxProtectionUnitsBulkAdditionJob(), nil
}
// GetDirectoryObjectIds gets the directoryObjectIds property value. The list of Exchange directoryObjectIds to add to the Exchange protection policy.
// returns a []string when successful
func (m *MailboxProtectionUnitsBulkAdditionJob) GetDirectoryObjectIds()([]string) {
    val, err := m.GetBackingStore().Get("directoryObjectIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MailboxProtectionUnitsBulkAdditionJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectionUnitsBulkJobBase.GetFieldDeserializers()
    res["directoryObjectIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDirectoryObjectIds(res)
        }
        return nil
    }
    res["mailboxes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetMailboxes(res)
        }
        return nil
    }
    return res
}
// GetMailboxes gets the mailboxes property value. The list of Exchange email addresses to add to the Exchange protection policy.
// returns a []string when successful
func (m *MailboxProtectionUnitsBulkAdditionJob) GetMailboxes()([]string) {
    val, err := m.GetBackingStore().Get("mailboxes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MailboxProtectionUnitsBulkAdditionJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectionUnitsBulkJobBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDirectoryObjectIds() != nil {
        err = writer.WriteCollectionOfStringValues("directoryObjectIds", m.GetDirectoryObjectIds())
        if err != nil {
            return err
        }
    }
    if m.GetMailboxes() != nil {
        err = writer.WriteCollectionOfStringValues("mailboxes", m.GetMailboxes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectoryObjectIds sets the directoryObjectIds property value. The list of Exchange directoryObjectIds to add to the Exchange protection policy.
func (m *MailboxProtectionUnitsBulkAdditionJob) SetDirectoryObjectIds(value []string)() {
    err := m.GetBackingStore().Set("directoryObjectIds", value)
    if err != nil {
        panic(err)
    }
}
// SetMailboxes sets the mailboxes property value. The list of Exchange email addresses to add to the Exchange protection policy.
func (m *MailboxProtectionUnitsBulkAdditionJob) SetMailboxes(value []string)() {
    err := m.GetBackingStore().Set("mailboxes", value)
    if err != nil {
        panic(err)
    }
}
type MailboxProtectionUnitsBulkAdditionJobable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectionUnitsBulkJobBaseable
    GetDirectoryObjectIds()([]string)
    GetMailboxes()([]string)
    SetDirectoryObjectIds(value []string)()
    SetMailboxes(value []string)()
}
