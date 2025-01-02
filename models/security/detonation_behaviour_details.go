package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DetonationBehaviourDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDetonationBehaviourDetails instantiates a new DetonationBehaviourDetails and sets the default values.
func NewDetonationBehaviourDetails()(*DetonationBehaviourDetails) {
    m := &DetonationBehaviourDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDetonationBehaviourDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDetonationBehaviourDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDetonationBehaviourDetails(), nil
}
// GetActionStatus gets the actionStatus property value. The actionStatus property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetActionStatus()(*string) {
    val, err := m.GetBackingStore().Get("actionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DetonationBehaviourDetails) GetAdditionalData()(map[string]any) {
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
func (m *DetonationBehaviourDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBehaviourCapability gets the behaviourCapability property value. The behaviourCapability property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetBehaviourCapability()(*string) {
    val, err := m.GetBackingStore().Get("behaviourCapability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBehaviourGroup gets the behaviourGroup property value. The behaviourGroup property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetBehaviourGroup()(*string) {
    val, err := m.GetBackingStore().Get("behaviourGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDetails gets the details property value. The details property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetDetails()(*string) {
    val, err := m.GetBackingStore().Get("details")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEventDateTime gets the eventDateTime property value. The eventDateTime property
// returns a *Time when successful
func (m *DetonationBehaviourDetails) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("eventDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DetonationBehaviourDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionStatus(val)
        }
        return nil
    }
    res["behaviourCapability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBehaviourCapability(val)
        }
        return nil
    }
    res["behaviourGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBehaviourGroup(val)
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val)
        }
        return nil
    }
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
    res["operation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperation(val)
        }
        return nil
    }
    res["processId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessId(val)
        }
        return nil
    }
    res["processName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessName(val)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperation gets the operation property value. The operation property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetOperation()(*string) {
    val, err := m.GetBackingStore().Get("operation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProcessId gets the processId property value. The processId property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetProcessId()(*string) {
    val, err := m.GetBackingStore().Get("processId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProcessName gets the processName property value. The processName property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetProcessName()(*string) {
    val, err := m.GetBackingStore().Get("processName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTarget gets the target property value. The target property
// returns a *string when successful
func (m *DetonationBehaviourDetails) GetTarget()(*string) {
    val, err := m.GetBackingStore().Get("target")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DetonationBehaviourDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionStatus", m.GetActionStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("behaviourCapability", m.GetBehaviourCapability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("behaviourGroup", m.GetBehaviourGroup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
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
        err := writer.WriteStringValue("operation", m.GetOperation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("processId", m.GetProcessId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("processName", m.GetProcessName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
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
// SetActionStatus sets the actionStatus property value. The actionStatus property
func (m *DetonationBehaviourDetails) SetActionStatus(value *string)() {
    err := m.GetBackingStore().Set("actionStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetonationBehaviourDetails) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DetonationBehaviourDetails) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBehaviourCapability sets the behaviourCapability property value. The behaviourCapability property
func (m *DetonationBehaviourDetails) SetBehaviourCapability(value *string)() {
    err := m.GetBackingStore().Set("behaviourCapability", value)
    if err != nil {
        panic(err)
    }
}
// SetBehaviourGroup sets the behaviourGroup property value. The behaviourGroup property
func (m *DetonationBehaviourDetails) SetBehaviourGroup(value *string)() {
    err := m.GetBackingStore().Set("behaviourGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetDetails sets the details property value. The details property
func (m *DetonationBehaviourDetails) SetDetails(value *string)() {
    err := m.GetBackingStore().Set("details", value)
    if err != nil {
        panic(err)
    }
}
// SetEventDateTime sets the eventDateTime property value. The eventDateTime property
func (m *DetonationBehaviourDetails) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("eventDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DetonationBehaviourDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperation sets the operation property value. The operation property
func (m *DetonationBehaviourDetails) SetOperation(value *string)() {
    err := m.GetBackingStore().Set("operation", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessId sets the processId property value. The processId property
func (m *DetonationBehaviourDetails) SetProcessId(value *string)() {
    err := m.GetBackingStore().Set("processId", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessName sets the processName property value. The processName property
func (m *DetonationBehaviourDetails) SetProcessName(value *string)() {
    err := m.GetBackingStore().Set("processName", value)
    if err != nil {
        panic(err)
    }
}
// SetTarget sets the target property value. The target property
func (m *DetonationBehaviourDetails) SetTarget(value *string)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
type DetonationBehaviourDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionStatus()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBehaviourCapability()(*string)
    GetBehaviourGroup()(*string)
    GetDetails()(*string)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetOperation()(*string)
    GetProcessId()(*string)
    GetProcessName()(*string)
    GetTarget()(*string)
    SetActionStatus(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBehaviourCapability(value *string)()
    SetBehaviourGroup(value *string)()
    SetDetails(value *string)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetOperation(value *string)()
    SetProcessId(value *string)()
    SetProcessName(value *string)()
    SetTarget(value *string)()
}
