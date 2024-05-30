package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody instantiates a new AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody()(*AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) {
    m := &AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetApproveAllPermissions gets the approveAllPermissions property value. The approveAllPermissions property
// returns a *bool when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) GetApproveAllPermissions()(*bool) {
    val, err := m.GetBackingStore().Get("approveAllPermissions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["approveAllPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproveAllPermissions(val)
        }
        return nil
    }
    res["packageIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPackageIds(res)
        }
        return nil
    }
    return res
}
// GetPackageIds gets the packageIds property value. The packageIds property
// returns a []string when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) GetPackageIds()([]string) {
    val, err := m.GetBackingStore().Get("packageIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("approveAllPermissions", m.GetApproveAllPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetPackageIds() != nil {
        err := writer.WriteCollectionOfStringValues("packageIds", m.GetPackageIds())
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
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetApproveAllPermissions sets the approveAllPermissions property value. The approveAllPermissions property
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) SetApproveAllPermissions(value *bool)() {
    err := m.GetBackingStore().Set("approveAllPermissions", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetPackageIds sets the packageIds property value. The packageIds property
func (m *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBody) SetPackageIds(value []string)() {
    err := m.GetBackingStore().Set("packageIds", value)
    if err != nil {
        panic(err)
    }
}
type AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApproveAllPermissions()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetPackageIds()([]string)
    SetApproveAllPermissions(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetPackageIds(value []string)()
}
