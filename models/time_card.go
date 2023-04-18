package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeCard 
type TimeCard struct {
    ChangeTrackedEntity
}
// NewTimeCard instantiates a new timeCard and sets the default values.
func NewTimeCard()(*TimeCard) {
    m := &TimeCard{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.timeCard"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTimeCardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeCardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeCard(), nil
}
// GetBreaks gets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCard) GetBreaks()([]TimeCardBreakable) {
    val, err := m.GetBackingStore().Get("breaks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TimeCardBreakable)
    }
    return nil
}
// GetClockInEvent gets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCard) GetClockInEvent()(TimeCardEventable) {
    val, err := m.GetBackingStore().Get("clockInEvent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TimeCardEventable)
    }
    return nil
}
// GetClockOutEvent gets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCard) GetClockOutEvent()(TimeCardEventable) {
    val, err := m.GetBackingStore().Get("clockOutEvent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TimeCardEventable)
    }
    return nil
}
// GetConfirmedBy gets the confirmedBy property value. Indicate if this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
func (m *TimeCard) GetConfirmedBy()(*ConfirmedBy) {
    val, err := m.GetBackingStore().Get("confirmedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConfirmedBy)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCard) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["breaks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeCardBreakFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeCardBreakable, len(val))
            for i, v := range val {
                res[i] = v.(TimeCardBreakable)
            }
            m.SetBreaks(res)
        }
        return nil
    }
    res["clockInEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockInEvent(val.(TimeCardEventable))
        }
        return nil
    }
    res["clockOutEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockOutEvent(val.(TimeCardEventable))
        }
        return nil
    }
    res["confirmedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfirmedBy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfirmedBy(val.(*ConfirmedBy))
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ItemBodyable))
        }
        return nil
    }
    res["originalEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalEntry(val.(TimeCardEntryable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTimeCardState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*TimeCardState))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. Notes about the timeCard.
func (m *TimeCard) GetNotes()(ItemBodyable) {
    val, err := m.GetBackingStore().Get("notes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemBodyable)
    }
    return nil
}
// GetOriginalEntry gets the originalEntry property value. The original timeCardEntry of the timeCard, before user edits.
func (m *TimeCard) GetOriginalEntry()(TimeCardEntryable) {
    val, err := m.GetBackingStore().Get("originalEntry")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TimeCardEntryable)
    }
    return nil
}
// GetState gets the state property value. The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut, unknownFutureValue.
func (m *TimeCard) GetState()(*TimeCardState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TimeCardState)
    }
    return nil
}
// GetUserId gets the userId property value. User ID to which  the timeCard belongs.
func (m *TimeCard) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TimeCard) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBreaks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBreaks()))
        for i, v := range m.GetBreaks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("breaks", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clockInEvent", m.GetClockInEvent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clockOutEvent", m.GetClockOutEvent())
        if err != nil {
            return err
        }
    }
    if m.GetConfirmedBy() != nil {
        cast := (*m.GetConfirmedBy()).String()
        err = writer.WriteStringValue("confirmedBy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("originalEntry", m.GetOriginalEntry())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBreaks sets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCard) SetBreaks(value []TimeCardBreakable)() {
    err := m.GetBackingStore().Set("breaks", value)
    if err != nil {
        panic(err)
    }
}
// SetClockInEvent sets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCard) SetClockInEvent(value TimeCardEventable)() {
    err := m.GetBackingStore().Set("clockInEvent", value)
    if err != nil {
        panic(err)
    }
}
// SetClockOutEvent sets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCard) SetClockOutEvent(value TimeCardEventable)() {
    err := m.GetBackingStore().Set("clockOutEvent", value)
    if err != nil {
        panic(err)
    }
}
// SetConfirmedBy sets the confirmedBy property value. Indicate if this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
func (m *TimeCard) SetConfirmedBy(value *ConfirmedBy)() {
    err := m.GetBackingStore().Set("confirmedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetNotes sets the notes property value. Notes about the timeCard.
func (m *TimeCard) SetNotes(value ItemBodyable)() {
    err := m.GetBackingStore().Set("notes", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginalEntry sets the originalEntry property value. The original timeCardEntry of the timeCard, before user edits.
func (m *TimeCard) SetOriginalEntry(value TimeCardEntryable)() {
    err := m.GetBackingStore().Set("originalEntry", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut, unknownFutureValue.
func (m *TimeCard) SetState(value *TimeCardState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. User ID to which  the timeCard belongs.
func (m *TimeCard) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// TimeCardable 
type TimeCardable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBreaks()([]TimeCardBreakable)
    GetClockInEvent()(TimeCardEventable)
    GetClockOutEvent()(TimeCardEventable)
    GetConfirmedBy()(*ConfirmedBy)
    GetNotes()(ItemBodyable)
    GetOriginalEntry()(TimeCardEntryable)
    GetState()(*TimeCardState)
    GetUserId()(*string)
    SetBreaks(value []TimeCardBreakable)()
    SetClockInEvent(value TimeCardEventable)()
    SetClockOutEvent(value TimeCardEventable)()
    SetConfirmedBy(value *ConfirmedBy)()
    SetNotes(value ItemBodyable)()
    SetOriginalEntry(value TimeCardEntryable)()
    SetState(value *TimeCardState)()
    SetUserId(value *string)()
}
