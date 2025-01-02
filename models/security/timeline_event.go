package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TimelineEvent struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTimelineEvent instantiates a new TimelineEvent and sets the default values.
func NewTimelineEvent()(*TimelineEvent) {
    m := &TimelineEvent{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimelineEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimelineEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimelineEvent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TimelineEvent) GetAdditionalData()(map[string]any) {
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
func (m *TimelineEvent) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEventDateTime gets the eventDateTime property value. The eventDateTime property
// returns a *Time when successful
func (m *TimelineEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("eventDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEventDetails gets the eventDetails property value. The eventDetails property
// returns a *string when successful
func (m *TimelineEvent) GetEventDetails()(*string) {
    val, err := m.GetBackingStore().Get("eventDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEventResult gets the eventResult property value. The eventResult property
// returns a *string when successful
func (m *TimelineEvent) GetEventResult()(*string) {
    val, err := m.GetBackingStore().Get("eventResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEventSource gets the eventSource property value. The eventSource property
// returns a *EventSource when successful
func (m *TimelineEvent) GetEventSource()(*EventSource) {
    val, err := m.GetBackingStore().Get("eventSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EventSource)
    }
    return nil
}
// GetEventThreats gets the eventThreats property value. The eventThreats property
// returns a []string when successful
func (m *TimelineEvent) GetEventThreats()([]string) {
    val, err := m.GetBackingStore().Get("eventThreats")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetEventType gets the eventType property value. The eventType property
// returns a *TimelineEventType when successful
func (m *TimelineEvent) GetEventType()(*TimelineEventType) {
    val, err := m.GetBackingStore().Get("eventType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TimelineEventType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimelineEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["eventDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDetails(val)
        }
        return nil
    }
    res["eventResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventResult(val)
        }
        return nil
    }
    res["eventSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEventSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventSource(val.(*EventSource))
        }
        return nil
    }
    res["eventThreats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetEventThreats(res)
        }
        return nil
    }
    res["eventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTimelineEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventType(val.(*TimelineEventType))
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
// returns a *string when successful
func (m *TimelineEvent) GetOdataType()(*string) {
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
func (m *TimelineEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventDetails", m.GetEventDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventResult", m.GetEventResult())
        if err != nil {
            return err
        }
    }
    if m.GetEventSource() != nil {
        cast := (*m.GetEventSource()).String()
        err := writer.WriteStringValue("eventSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEventThreats() != nil {
        err := writer.WriteCollectionOfStringValues("eventThreats", m.GetEventThreats())
        if err != nil {
            return err
        }
    }
    if m.GetEventType() != nil {
        cast := (*m.GetEventType()).String()
        err := writer.WriteStringValue("eventType", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimelineEvent) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TimelineEvent) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEventDateTime sets the eventDateTime property value. The eventDateTime property
func (m *TimelineEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("eventDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEventDetails sets the eventDetails property value. The eventDetails property
func (m *TimelineEvent) SetEventDetails(value *string)() {
    err := m.GetBackingStore().Set("eventDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetEventResult sets the eventResult property value. The eventResult property
func (m *TimelineEvent) SetEventResult(value *string)() {
    err := m.GetBackingStore().Set("eventResult", value)
    if err != nil {
        panic(err)
    }
}
// SetEventSource sets the eventSource property value. The eventSource property
func (m *TimelineEvent) SetEventSource(value *EventSource)() {
    err := m.GetBackingStore().Set("eventSource", value)
    if err != nil {
        panic(err)
    }
}
// SetEventThreats sets the eventThreats property value. The eventThreats property
func (m *TimelineEvent) SetEventThreats(value []string)() {
    err := m.GetBackingStore().Set("eventThreats", value)
    if err != nil {
        panic(err)
    }
}
// SetEventType sets the eventType property value. The eventType property
func (m *TimelineEvent) SetEventType(value *TimelineEventType)() {
    err := m.GetBackingStore().Set("eventType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TimelineEvent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type TimelineEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEventDetails()(*string)
    GetEventResult()(*string)
    GetEventSource()(*EventSource)
    GetEventThreats()([]string)
    GetEventType()(*TimelineEventType)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEventDetails(value *string)()
    SetEventResult(value *string)()
    SetEventSource(value *EventSource)()
    SetEventThreats(value []string)()
    SetEventType(value *TimelineEventType)()
    SetOdataType(value *string)()
}
