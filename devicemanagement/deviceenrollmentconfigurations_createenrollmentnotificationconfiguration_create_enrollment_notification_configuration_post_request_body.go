package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody instantiates a new DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody and sets the default values.
func NewDeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody()(*DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) {
    m := &DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceEnrollmentNotificationConfigurations gets the deviceEnrollmentNotificationConfigurations property value. The deviceEnrollmentNotificationConfigurations property
// returns a []DeviceEnrollmentConfigurationable when successful
func (m *DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) GetDeviceEnrollmentNotificationConfigurations()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable) {
    val, err := m.GetBackingStore().Get("deviceEnrollmentNotificationConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceEnrollmentNotificationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)
                }
            }
            m.SetDeviceEnrollmentNotificationConfigurations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceEnrollmentNotificationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceEnrollmentNotificationConfigurations()))
        for i, v := range m.GetDeviceEnrollmentNotificationConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("deviceEnrollmentNotificationConfigurations", cast)
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
func (m *DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceEnrollmentNotificationConfigurations sets the deviceEnrollmentNotificationConfigurations property value. The deviceEnrollmentNotificationConfigurations property
func (m *DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBody) SetDeviceEnrollmentNotificationConfigurations(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)() {
    err := m.GetBackingStore().Set("deviceEnrollmentNotificationConfigurations", value)
    if err != nil {
        panic(err)
    }
}
type DeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceEnrollmentNotificationConfigurations()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceEnrollmentNotificationConfigurations(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)()
}
