package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type Sensor struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewSensor instantiates a new Sensor and sets the default values.
func NewSensor()(*Sensor) {
    m := &Sensor{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSensorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSensorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensor(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
// returns a *Time when successful
func (m *Sensor) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeploymentStatus gets the deploymentStatus property value. The deploymentStatus property
// returns a *DeploymentStatus when successful
func (m *Sensor) GetDeploymentStatus()(*DeploymentStatus) {
    val, err := m.GetBackingStore().Get("deploymentStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeploymentStatus)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *Sensor) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomainName gets the domainName property value. The domainName property
// returns a *string when successful
func (m *Sensor) GetDomainName()(*string) {
    val, err := m.GetBackingStore().Get("domainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Sensor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deploymentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentStatus(val.(*DeploymentStatus))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
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
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensorHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*SensorHealthStatus))
        }
        return nil
    }
    res["openHealthIssuesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenHealthIssuesCount(val)
        }
        return nil
    }
    res["sensorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensorType(val.(*SensorType))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSensorSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(SensorSettingsable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetHealthIssues gets the healthIssues property value. The healthIssues property
// returns a []HealthIssueable when successful
func (m *Sensor) GetHealthIssues()([]HealthIssueable) {
    val, err := m.GetBackingStore().Get("healthIssues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HealthIssueable)
    }
    return nil
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *SensorHealthStatus when successful
func (m *Sensor) GetHealthStatus()(*SensorHealthStatus) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SensorHealthStatus)
    }
    return nil
}
// GetOpenHealthIssuesCount gets the openHealthIssuesCount property value. The openHealthIssuesCount property
// returns a *int64 when successful
func (m *Sensor) GetOpenHealthIssuesCount()(*int64) {
    val, err := m.GetBackingStore().Get("openHealthIssuesCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSensorType gets the sensorType property value. The sensorType property
// returns a *SensorType when successful
func (m *Sensor) GetSensorType()(*SensorType) {
    val, err := m.GetBackingStore().Get("sensorType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SensorType)
    }
    return nil
}
// GetSettings gets the settings property value. The settings property
// returns a SensorSettingsable when successful
func (m *Sensor) GetSettings()(SensorSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SensorSettingsable)
    }
    return nil
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *Sensor) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Sensor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentStatus() != nil {
        cast := (*m.GetDeploymentStatus()).String()
        err = writer.WriteStringValue("deploymentStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
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
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("openHealthIssuesCount", m.GetOpenHealthIssuesCount())
        if err != nil {
            return err
        }
    }
    if m.GetSensorType() != nil {
        cast := (*m.GetSensorType()).String()
        err = writer.WriteStringValue("sensorType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *Sensor) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentStatus sets the deploymentStatus property value. The deploymentStatus property
func (m *Sensor) SetDeploymentStatus(value *DeploymentStatus)() {
    err := m.GetBackingStore().Set("deploymentStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Sensor) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainName sets the domainName property value. The domainName property
func (m *Sensor) SetDomainName(value *string)() {
    err := m.GetBackingStore().Set("domainName", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthIssues sets the healthIssues property value. The healthIssues property
func (m *Sensor) SetHealthIssues(value []HealthIssueable)() {
    err := m.GetBackingStore().Set("healthIssues", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *Sensor) SetHealthStatus(value *SensorHealthStatus)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetOpenHealthIssuesCount sets the openHealthIssuesCount property value. The openHealthIssuesCount property
func (m *Sensor) SetOpenHealthIssuesCount(value *int64)() {
    err := m.GetBackingStore().Set("openHealthIssuesCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSensorType sets the sensorType property value. The sensorType property
func (m *Sensor) SetSensorType(value *SensorType)() {
    err := m.GetBackingStore().Set("sensorType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The settings property
func (m *Sensor) SetSettings(value SensorSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. The version property
func (m *Sensor) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
type Sensorable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentStatus()(*DeploymentStatus)
    GetDisplayName()(*string)
    GetDomainName()(*string)
    GetHealthIssues()([]HealthIssueable)
    GetHealthStatus()(*SensorHealthStatus)
    GetOpenHealthIssuesCount()(*int64)
    GetSensorType()(*SensorType)
    GetSettings()(SensorSettingsable)
    GetVersion()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentStatus(value *DeploymentStatus)()
    SetDisplayName(value *string)()
    SetDomainName(value *string)()
    SetHealthIssues(value []HealthIssueable)()
    SetHealthStatus(value *SensorHealthStatus)()
    SetOpenHealthIssuesCount(value *int64)()
    SetSensorType(value *SensorType)()
    SetSettings(value SensorSettingsable)()
    SetVersion(value *string)()
}
