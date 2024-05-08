package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type IdentityContainer struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewIdentityContainer instantiates a new IdentityContainer and sets the default values.
func NewIdentityContainer()(*IdentityContainer) {
    m := &IdentityContainer{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateIdentityContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityContainer(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthIssues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHealthIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HealthIssueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HealthIssueable)
                }
            }
            m.SetHealthIssues(res)
        }
        return nil
    }
    res["sensors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Sensorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Sensorable)
                }
            }
            m.SetSensors(res)
        }
        return nil
    }
    return res
}
// GetHealthIssues gets the healthIssues property value. Represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender for Identity identified.
// returns a []HealthIssueable when successful
func (m *IdentityContainer) GetHealthIssues()([]HealthIssueable) {
    val, err := m.GetBackingStore().Get("healthIssues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HealthIssueable)
    }
    return nil
}
// GetSensors gets the sensors property value. The sensors property
// returns a []Sensorable when successful
func (m *IdentityContainer) GetSensors()([]Sensorable) {
    val, err := m.GetBackingStore().Get("sensors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Sensorable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IdentityContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetHealthIssues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHealthIssues()))
        for i, v := range m.GetHealthIssues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("healthIssues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensors()))
        for i, v := range m.GetSensors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sensors", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHealthIssues sets the healthIssues property value. Represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender for Identity identified.
func (m *IdentityContainer) SetHealthIssues(value []HealthIssueable)() {
    err := m.GetBackingStore().Set("healthIssues", value)
    if err != nil {
        panic(err)
    }
}
// SetSensors sets the sensors property value. The sensors property
func (m *IdentityContainer) SetSensors(value []Sensorable)() {
    err := m.GetBackingStore().Set("sensors", value)
    if err != nil {
        panic(err)
    }
}
type IdentityContainerable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHealthIssues()([]HealthIssueable)
    GetSensors()([]Sensorable)
    SetHealthIssues(value []HealthIssueable)()
    SetSensors(value []Sensorable)()
}
