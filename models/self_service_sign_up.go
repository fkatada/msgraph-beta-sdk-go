package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SelfServiceSignUp struct {
    Entity
}
// NewSelfServiceSignUp instantiates a new SelfServiceSignUp and sets the default values.
func NewSelfServiceSignUp()(*SelfServiceSignUp) {
    m := &SelfServiceSignUp{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSelfServiceSignUpFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSelfServiceSignUpFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSelfServiceSignUp(), nil
}
// GetAppDisplayName gets the appDisplayName property value. The appDisplayName property
// returns a *string when successful
func (m *SelfServiceSignUp) GetAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("appDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppId gets the appId property value. The appId property
// returns a *string when successful
func (m *SelfServiceSignUp) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppliedEventListeners gets the appliedEventListeners property value. The appliedEventListeners property
// returns a []AppliedAuthenticationEventListenerable when successful
func (m *SelfServiceSignUp) GetAppliedEventListeners()([]AppliedAuthenticationEventListenerable) {
    val, err := m.GetBackingStore().Get("appliedEventListeners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppliedAuthenticationEventListenerable)
    }
    return nil
}
// GetCorrelationId gets the correlationId property value. The correlationId property
// returns a *string when successful
func (m *SelfServiceSignUp) GetCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("correlationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
// returns a *Time when successful
func (m *SelfServiceSignUp) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
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
func (m *SelfServiceSignUp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["appliedEventListeners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppliedAuthenticationEventListenerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppliedAuthenticationEventListenerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppliedAuthenticationEventListenerable)
                }
            }
            m.SetAppliedEventListeners(res)
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
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
    res["signUpIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignUpIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignUpIdentity(val.(SignUpIdentityable))
        }
        return nil
    }
    res["signUpIdentityProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignUpIdentityProvider(val)
        }
        return nil
    }
    res["signUpStage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignUpStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignUpStage(val.(*SignUpStage))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignUpStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(SignUpStatusable))
        }
        return nil
    }
    res["userSnapshot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCiamUserSnapshotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSnapshot(val.(CiamUserSnapshotable))
        }
        return nil
    }
    return res
}
// GetSignUpIdentity gets the signUpIdentity property value. The signUpIdentity property
// returns a SignUpIdentityable when successful
func (m *SelfServiceSignUp) GetSignUpIdentity()(SignUpIdentityable) {
    val, err := m.GetBackingStore().Get("signUpIdentity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignUpIdentityable)
    }
    return nil
}
// GetSignUpIdentityProvider gets the signUpIdentityProvider property value. The signUpIdentityProvider property
// returns a *string when successful
func (m *SelfServiceSignUp) GetSignUpIdentityProvider()(*string) {
    val, err := m.GetBackingStore().Get("signUpIdentityProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSignUpStage gets the signUpStage property value. The signUpStage property
// returns a *SignUpStage when successful
func (m *SelfServiceSignUp) GetSignUpStage()(*SignUpStage) {
    val, err := m.GetBackingStore().Get("signUpStage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SignUpStage)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a SignUpStatusable when successful
func (m *SelfServiceSignUp) GetStatus()(SignUpStatusable) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignUpStatusable)
    }
    return nil
}
// GetUserSnapshot gets the userSnapshot property value. The userSnapshot property
// returns a CiamUserSnapshotable when successful
func (m *SelfServiceSignUp) GetUserSnapshot()(CiamUserSnapshotable) {
    val, err := m.GetBackingStore().Get("userSnapshot")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CiamUserSnapshotable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SelfServiceSignUp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetAppliedEventListeners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliedEventListeners()))
        for i, v := range m.GetAppliedEventListeners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appliedEventListeners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("signUpIdentity", m.GetSignUpIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signUpIdentityProvider", m.GetSignUpIdentityProvider())
        if err != nil {
            return err
        }
    }
    if m.GetSignUpStage() != nil {
        cast := (*m.GetSignUpStage()).String()
        err = writer.WriteStringValue("signUpStage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userSnapshot", m.GetUserSnapshot())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. The appDisplayName property
func (m *SelfServiceSignUp) SetAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("appDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppId sets the appId property value. The appId property
func (m *SelfServiceSignUp) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliedEventListeners sets the appliedEventListeners property value. The appliedEventListeners property
func (m *SelfServiceSignUp) SetAppliedEventListeners(value []AppliedAuthenticationEventListenerable)() {
    err := m.GetBackingStore().Set("appliedEventListeners", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationId sets the correlationId property value. The correlationId property
func (m *SelfServiceSignUp) SetCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("correlationId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *SelfServiceSignUp) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSignUpIdentity sets the signUpIdentity property value. The signUpIdentity property
func (m *SelfServiceSignUp) SetSignUpIdentity(value SignUpIdentityable)() {
    err := m.GetBackingStore().Set("signUpIdentity", value)
    if err != nil {
        panic(err)
    }
}
// SetSignUpIdentityProvider sets the signUpIdentityProvider property value. The signUpIdentityProvider property
func (m *SelfServiceSignUp) SetSignUpIdentityProvider(value *string)() {
    err := m.GetBackingStore().Set("signUpIdentityProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetSignUpStage sets the signUpStage property value. The signUpStage property
func (m *SelfServiceSignUp) SetSignUpStage(value *SignUpStage)() {
    err := m.GetBackingStore().Set("signUpStage", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *SelfServiceSignUp) SetStatus(value SignUpStatusable)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUserSnapshot sets the userSnapshot property value. The userSnapshot property
func (m *SelfServiceSignUp) SetUserSnapshot(value CiamUserSnapshotable)() {
    err := m.GetBackingStore().Set("userSnapshot", value)
    if err != nil {
        panic(err)
    }
}
type SelfServiceSignUpable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetAppliedEventListeners()([]AppliedAuthenticationEventListenerable)
    GetCorrelationId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSignUpIdentity()(SignUpIdentityable)
    GetSignUpIdentityProvider()(*string)
    GetSignUpStage()(*SignUpStage)
    GetStatus()(SignUpStatusable)
    GetUserSnapshot()(CiamUserSnapshotable)
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetAppliedEventListeners(value []AppliedAuthenticationEventListenerable)()
    SetCorrelationId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSignUpIdentity(value SignUpIdentityable)()
    SetSignUpIdentityProvider(value *string)()
    SetSignUpStage(value *SignUpStage)()
    SetStatus(value SignUpStatusable)()
    SetUserSnapshot(value CiamUserSnapshotable)()
}
