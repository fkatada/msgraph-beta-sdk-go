package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppContentFile contains properties for a single installer file that is associated with a given mobileAppContent version.
type MobileAppContentFile struct {
    Entity
}
// NewMobileAppContentFile instantiates a new MobileAppContentFile and sets the default values.
func NewMobileAppContentFile()(*MobileAppContentFile) {
    m := &MobileAppContentFile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppContentFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileAppContentFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppContentFile(), nil
}
// GetAzureStorageUri gets the azureStorageUri property value. Indicates the Azure Storage URI that the file is uploaded to. Created by the service upon receiving a valid mobileAppContentFile. Read-only. This property is read-only.
// returns a *string when successful
func (m *MobileAppContentFile) GetAzureStorageUri()(*string) {
    val, err := m.GetBackingStore().Get("azureStorageUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureStorageUriExpirationDateTime gets the azureStorageUriExpirationDateTime property value. Indicates the date and time when the Azure storage URI expires, in ISO 8601 format. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Read-only. This property is read-only.
// returns a *Time when successful
func (m *MobileAppContentFile) GetAzureStorageUriExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("azureStorageUriExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Indicates created date and time associated with app content file, in ISO 8601 format. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Read-only. This property is read-only.
// returns a *Time when successful
func (m *MobileAppContentFile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *MobileAppContentFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureStorageUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageUri(val)
        }
        return nil
    }
    res["azureStorageUriExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageUriExpirationDateTime(val)
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
    res["isCommitted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCommitted(val)
        }
        return nil
    }
    res["isDependency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDependency(val)
        }
        return nil
    }
    res["isFrameworkFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFrameworkFile(val)
        }
        return nil
    }
    res["manifest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifest(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["sizeEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeEncrypted(val)
        }
        return nil
    }
    res["sizeEncryptedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeEncryptedInBytes(val)
        }
        return nil
    }
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    res["uploadState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppContentFileUploadState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadState(val.(*MobileAppContentFileUploadState))
        }
        return nil
    }
    return res
}
// GetIsCommitted gets the isCommitted property value. A value indicating whether the file is committed. A committed app content file has been fully uploaded and validated by the Intune service. TRUE means that app content file is committed, FALSE means that app content file is not committed. Defaults to FALSE. Read-only. This property is read-only.
// returns a *bool when successful
func (m *MobileAppContentFile) GetIsCommitted()(*bool) {
    val, err := m.GetBackingStore().Get("isCommitted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsDependency gets the isDependency property value. Indicates whether this content file is a dependency for the main content file. TRUE means that the content file is a dependency, FALSE means that the content file is not a dependency and is the main content file. Defaults to FALSE.
// returns a *bool when successful
func (m *MobileAppContentFile) GetIsDependency()(*bool) {
    val, err := m.GetBackingStore().Get("isDependency")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsFrameworkFile gets the isFrameworkFile property value. A value indicating whether the file is a framework file. To be deprecated.
// returns a *bool when successful
func (m *MobileAppContentFile) GetIsFrameworkFile()(*bool) {
    val, err := m.GetBackingStore().Get("isFrameworkFile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManifest gets the manifest property value. Indicates the manifest information, containing file metadata.
// returns a []byte when successful
func (m *MobileAppContentFile) GetManifest()([]byte) {
    val, err := m.GetBackingStore().Get("manifest")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetName gets the name property value. Indicates the name of the file.
// returns a *string when successful
func (m *MobileAppContentFile) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSize gets the size property value. Indicates the original size of the file, in bytes.
// returns a *int64 when successful
func (m *MobileAppContentFile) GetSize()(*int64) {
    val, err := m.GetBackingStore().Get("size")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSizeEncrypted gets the sizeEncrypted property value. Indicates the size of the file after encryption, in bytes.
// returns a *int64 when successful
func (m *MobileAppContentFile) GetSizeEncrypted()(*int64) {
    val, err := m.GetBackingStore().Get("sizeEncrypted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSizeEncryptedInBytes gets the sizeEncryptedInBytes property value. Indicates the size of the file after encryption, in bytes. To be deprecated in February 2025, please use SizeEncrypted property instead. Valid values 0 to 9.22337203685478E+18
// returns a *int64 when successful
func (m *MobileAppContentFile) GetSizeEncryptedInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("sizeEncryptedInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSizeInBytes gets the sizeInBytes property value. Indicates the original size of the file, in bytes. To be deprecated in February 2025, please use Size property instead. Valid values 0 to 9.22337203685478E+18
// returns a *int64 when successful
func (m *MobileAppContentFile) GetSizeInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("sizeInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUploadState gets the uploadState property value. Contains properties for upload request states.
// returns a *MobileAppContentFileUploadState when successful
func (m *MobileAppContentFile) GetUploadState()(*MobileAppContentFileUploadState) {
    val, err := m.GetBackingStore().Get("uploadState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MobileAppContentFileUploadState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppContentFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDependency", m.GetIsDependency())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFrameworkFile", m.GetIsFrameworkFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("manifest", m.GetManifest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeEncrypted", m.GetSizeEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeEncryptedInBytes", m.GetSizeEncryptedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    if m.GetUploadState() != nil {
        cast := (*m.GetUploadState()).String()
        err = writer.WriteStringValue("uploadState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureStorageUri sets the azureStorageUri property value. Indicates the Azure Storage URI that the file is uploaded to. Created by the service upon receiving a valid mobileAppContentFile. Read-only. This property is read-only.
func (m *MobileAppContentFile) SetAzureStorageUri(value *string)() {
    err := m.GetBackingStore().Set("azureStorageUri", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureStorageUriExpirationDateTime sets the azureStorageUriExpirationDateTime property value. Indicates the date and time when the Azure storage URI expires, in ISO 8601 format. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Read-only. This property is read-only.
func (m *MobileAppContentFile) SetAzureStorageUriExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("azureStorageUriExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Indicates created date and time associated with app content file, in ISO 8601 format. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Read-only. This property is read-only.
func (m *MobileAppContentFile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCommitted sets the isCommitted property value. A value indicating whether the file is committed. A committed app content file has been fully uploaded and validated by the Intune service. TRUE means that app content file is committed, FALSE means that app content file is not committed. Defaults to FALSE. Read-only. This property is read-only.
func (m *MobileAppContentFile) SetIsCommitted(value *bool)() {
    err := m.GetBackingStore().Set("isCommitted", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDependency sets the isDependency property value. Indicates whether this content file is a dependency for the main content file. TRUE means that the content file is a dependency, FALSE means that the content file is not a dependency and is the main content file. Defaults to FALSE.
func (m *MobileAppContentFile) SetIsDependency(value *bool)() {
    err := m.GetBackingStore().Set("isDependency", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFrameworkFile sets the isFrameworkFile property value. A value indicating whether the file is a framework file. To be deprecated.
func (m *MobileAppContentFile) SetIsFrameworkFile(value *bool)() {
    err := m.GetBackingStore().Set("isFrameworkFile", value)
    if err != nil {
        panic(err)
    }
}
// SetManifest sets the manifest property value. Indicates the manifest information, containing file metadata.
func (m *MobileAppContentFile) SetManifest(value []byte)() {
    err := m.GetBackingStore().Set("manifest", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Indicates the name of the file.
func (m *MobileAppContentFile) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetSize sets the size property value. Indicates the original size of the file, in bytes.
func (m *MobileAppContentFile) SetSize(value *int64)() {
    err := m.GetBackingStore().Set("size", value)
    if err != nil {
        panic(err)
    }
}
// SetSizeEncrypted sets the sizeEncrypted property value. Indicates the size of the file after encryption, in bytes.
func (m *MobileAppContentFile) SetSizeEncrypted(value *int64)() {
    err := m.GetBackingStore().Set("sizeEncrypted", value)
    if err != nil {
        panic(err)
    }
}
// SetSizeEncryptedInBytes sets the sizeEncryptedInBytes property value. Indicates the size of the file after encryption, in bytes. To be deprecated in February 2025, please use SizeEncrypted property instead. Valid values 0 to 9.22337203685478E+18
func (m *MobileAppContentFile) SetSizeEncryptedInBytes(value *int64)() {
    err := m.GetBackingStore().Set("sizeEncryptedInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetSizeInBytes sets the sizeInBytes property value. Indicates the original size of the file, in bytes. To be deprecated in February 2025, please use Size property instead. Valid values 0 to 9.22337203685478E+18
func (m *MobileAppContentFile) SetSizeInBytes(value *int64)() {
    err := m.GetBackingStore().Set("sizeInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetUploadState sets the uploadState property value. Contains properties for upload request states.
func (m *MobileAppContentFile) SetUploadState(value *MobileAppContentFileUploadState)() {
    err := m.GetBackingStore().Set("uploadState", value)
    if err != nil {
        panic(err)
    }
}
type MobileAppContentFileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureStorageUri()(*string)
    GetAzureStorageUriExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsCommitted()(*bool)
    GetIsDependency()(*bool)
    GetIsFrameworkFile()(*bool)
    GetManifest()([]byte)
    GetName()(*string)
    GetSize()(*int64)
    GetSizeEncrypted()(*int64)
    GetSizeEncryptedInBytes()(*int64)
    GetSizeInBytes()(*int64)
    GetUploadState()(*MobileAppContentFileUploadState)
    SetAzureStorageUri(value *string)()
    SetAzureStorageUriExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsCommitted(value *bool)()
    SetIsDependency(value *bool)()
    SetIsFrameworkFile(value *bool)()
    SetManifest(value []byte)()
    SetName(value *string)()
    SetSize(value *int64)()
    SetSizeEncrypted(value *int64)()
    SetSizeEncryptedInBytes(value *int64)()
    SetSizeInBytes(value *int64)()
    SetUploadState(value *MobileAppContentFileUploadState)()
}
