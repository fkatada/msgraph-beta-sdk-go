package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintJob 
type PrintJob struct {
    Entity
}
// NewPrintJob instantiates a new printJob and sets the default values.
func NewPrintJob()(*PrintJob) {
    m := &PrintJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintJob(), nil
}
// GetAcknowledgedDateTime gets the acknowledgedDateTime property value. The acknowledgedDateTime property
func (m *PrintJob) GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("acknowledgedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCompletedDateTime gets the completedDateTime property value. The completedDateTime property
func (m *PrintJob) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("completedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetConfiguration gets the configuration property value. The configuration property
func (m *PrintJob) GetConfiguration()(PrintJobConfigurationable) {
    val, err := m.GetBackingStore().Get("configuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrintJobConfigurationable)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *PrintJob) GetCreatedBy()(UserIdentityable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserIdentityable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The DateTimeOffset when the job was created. Read-only.
func (m *PrintJob) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the print job.
func (m *PrintJob) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDocuments gets the documents property value. The documents property
func (m *PrintJob) GetDocuments()([]PrintDocumentable) {
    val, err := m.GetBackingStore().Get("documents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintDocumentable)
    }
    return nil
}
// GetErrorCode gets the errorCode property value. The errorCode property
func (m *PrintJob) GetErrorCode()(*int32) {
    val, err := m.GetBackingStore().Get("errorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acknowledgedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcknowledgedDateTime(val)
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintJobConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(PrintJobConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(UserIdentityable))
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
    res["documents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintDocumentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintDocumentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintDocumentable)
                }
            }
            m.SetDocuments(res)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["isFetchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFetchable(val)
        }
        return nil
    }
    res["redirectedFrom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectedFrom(val)
        }
        return nil
    }
    res["redirectedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectedTo(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintJobStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(PrintJobStatusable))
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTaskable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintTaskable)
                }
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetIsFetchable gets the isFetchable property value. If true, document can be fetched by printer.
func (m *PrintJob) GetIsFetchable()(*bool) {
    val, err := m.GetBackingStore().Get("isFetchable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRedirectedFrom gets the redirectedFrom property value. Contains the source job URL, if the job has been redirected from another printer.
func (m *PrintJob) GetRedirectedFrom()(*string) {
    val, err := m.GetBackingStore().Get("redirectedFrom")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRedirectedTo gets the redirectedTo property value. Contains the destination job URL, if the job has been redirected to another printer.
func (m *PrintJob) GetRedirectedTo()(*string) {
    val, err := m.GetBackingStore().Get("redirectedTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *PrintJob) GetStatus()(PrintJobStatusable) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrintJobStatusable)
    }
    return nil
}
// GetTasks gets the tasks property value. A list of printTasks that were triggered by this print job.
func (m *PrintJob) GetTasks()([]PrintTaskable) {
    val, err := m.GetBackingStore().Get("tasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintTaskable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrintJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acknowledgedDateTime", m.GetAcknowledgedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDocuments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDocuments()))
        for i, v := range m.GetDocuments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("documents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFetchable", m.GetIsFetchable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectedFrom", m.GetRedirectedFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectedTo", m.GetRedirectedTo())
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
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcknowledgedDateTime sets the acknowledgedDateTime property value. The acknowledgedDateTime property
func (m *PrintJob) SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("acknowledgedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletedDateTime sets the completedDateTime property value. The completedDateTime property
func (m *PrintJob) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("completedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetConfiguration sets the configuration property value. The configuration property
func (m *PrintJob) SetConfiguration(value PrintJobConfigurationable)() {
    err := m.GetBackingStore().Set("configuration", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *PrintJob) SetCreatedBy(value UserIdentityable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The DateTimeOffset when the job was created. Read-only.
func (m *PrintJob) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the print job.
func (m *PrintJob) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDocuments sets the documents property value. The documents property
func (m *PrintJob) SetDocuments(value []PrintDocumentable)() {
    err := m.GetBackingStore().Set("documents", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCode sets the errorCode property value. The errorCode property
func (m *PrintJob) SetErrorCode(value *int32)() {
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFetchable sets the isFetchable property value. If true, document can be fetched by printer.
func (m *PrintJob) SetIsFetchable(value *bool)() {
    err := m.GetBackingStore().Set("isFetchable", value)
    if err != nil {
        panic(err)
    }
}
// SetRedirectedFrom sets the redirectedFrom property value. Contains the source job URL, if the job has been redirected from another printer.
func (m *PrintJob) SetRedirectedFrom(value *string)() {
    err := m.GetBackingStore().Set("redirectedFrom", value)
    if err != nil {
        panic(err)
    }
}
// SetRedirectedTo sets the redirectedTo property value. Contains the destination job URL, if the job has been redirected to another printer.
func (m *PrintJob) SetRedirectedTo(value *string)() {
    err := m.GetBackingStore().Set("redirectedTo", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *PrintJob) SetStatus(value PrintJobStatusable)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTasks sets the tasks property value. A list of printTasks that were triggered by this print job.
func (m *PrintJob) SetTasks(value []PrintTaskable)() {
    err := m.GetBackingStore().Set("tasks", value)
    if err != nil {
        panic(err)
    }
}
// PrintJobable 
type PrintJobable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConfiguration()(PrintJobConfigurationable)
    GetCreatedBy()(UserIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetDocuments()([]PrintDocumentable)
    GetErrorCode()(*int32)
    GetIsFetchable()(*bool)
    GetRedirectedFrom()(*string)
    GetRedirectedTo()(*string)
    GetStatus()(PrintJobStatusable)
    GetTasks()([]PrintTaskable)
    SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConfiguration(value PrintJobConfigurationable)()
    SetCreatedBy(value UserIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetDocuments(value []PrintDocumentable)()
    SetErrorCode(value *int32)()
    SetIsFetchable(value *bool)()
    SetRedirectedFrom(value *string)()
    SetRedirectedTo(value *string)()
    SetStatus(value PrintJobStatusable)()
    SetTasks(value []PrintTaskable)()
}
