package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody 
type AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The payloadIds property
    payloadIds []string
}
// NewAndroidManagedAppProtectionsHasPayloadLinksPostRequestBody instantiates a new AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody and sets the default values.
func NewAndroidManagedAppProtectionsHasPayloadLinksPostRequestBody()(*AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody) {
    m := &AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateAndroidManagedAppProtectionsHasPayloadLinksPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedAppProtectionsHasPayloadLinksPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedAppProtectionsHasPayloadLinksPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["payloadIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPayloadIds(res)
        }
        return nil
    }
    return res
}
// GetPayloadIds gets the payloadIds property value. The payloadIds property
func (m *AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody) GetPayloadIds()([]string) {
    return m.payloadIds
}
// Serialize serializes information the current object
func (m *AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPayloadIds() != nil {
        err := writer.WriteCollectionOfStringValues("payloadIds", m.GetPayloadIds())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPayloadIds sets the payloadIds property value. The payloadIds property
func (m *AndroidManagedAppProtectionsHasPayloadLinksPostRequestBody) SetPayloadIds(value []string)() {
    m.payloadIds = value
}
