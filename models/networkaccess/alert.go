package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Alert 
type Alert struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAlert instantiates a new alert and sets the default values.
func NewAlert()(*Alert) {
    m := &Alert{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlert(), nil
}
// GetActions gets the actions property value. The actions property
func (m *Alert) GetActions()([]AlertActionable) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AlertActionable)
    }
    return nil
}
// GetAlertType gets the alertType property value. The alertType property
func (m *Alert) GetAlertType()(*AlertType) {
    val, err := m.GetBackingStore().Get("alertType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertType)
    }
    return nil
}
// GetCreationDateTime gets the creationDateTime property value. The creationDateTime property
func (m *Alert) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("creationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *Alert) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Alert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AlertActionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    res["alertType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertType(val.(*AlertType))
        }
        return nil
    }
    res["creationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
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
    res["firstImpactedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstImpactedDateTime(val)
        }
        return nil
    }
    return res
}
// GetFirstImpactedDateTime gets the firstImpactedDateTime property value. The firstImpactedDateTime property
func (m *Alert) GetFirstImpactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("firstImpactedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Alert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlertType() != nil {
        cast := (*m.GetAlertType()).String()
        err = writer.WriteStringValue("alertType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
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
        err = writer.WriteTimeValue("firstImpactedDateTime", m.GetFirstImpactedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The actions property
func (m *Alert) SetActions(value []AlertActionable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertType sets the alertType property value. The alertType property
func (m *Alert) SetAlertType(value *AlertType)() {
    err := m.GetBackingStore().Set("alertType", value)
    if err != nil {
        panic(err)
    }
}
// SetCreationDateTime sets the creationDateTime property value. The creationDateTime property
func (m *Alert) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("creationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *Alert) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstImpactedDateTime sets the firstImpactedDateTime property value. The firstImpactedDateTime property
func (m *Alert) SetFirstImpactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("firstImpactedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// Alertable 
type Alertable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]AlertActionable)
    GetAlertType()(*AlertType)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetFirstImpactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActions(value []AlertActionable)()
    SetAlertType(value *AlertType)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetFirstImpactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
