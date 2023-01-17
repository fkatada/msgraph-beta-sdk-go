package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendationsItemImpactedResourcesItemDismissPostRequestBody 
type RecommendationsItemImpactedResourcesItemDismissPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dismissReason property
    dismissReason *string
}
// NewRecommendationsItemImpactedResourcesItemDismissPostRequestBody instantiates a new RecommendationsItemImpactedResourcesItemDismissPostRequestBody and sets the default values.
func NewRecommendationsItemImpactedResourcesItemDismissPostRequestBody()(*RecommendationsItemImpactedResourcesItemDismissPostRequestBody) {
    m := &RecommendationsItemImpactedResourcesItemDismissPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateRecommendationsItemImpactedResourcesItemDismissPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecommendationsItemImpactedResourcesItemDismissPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecommendationsItemImpactedResourcesItemDismissPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecommendationsItemImpactedResourcesItemDismissPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDismissReason gets the dismissReason property value. The dismissReason property
func (m *RecommendationsItemImpactedResourcesItemDismissPostRequestBody) GetDismissReason()(*string) {
    return m.dismissReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecommendationsItemImpactedResourcesItemDismissPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dismissReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissReason(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RecommendationsItemImpactedResourcesItemDismissPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dismissReason", m.GetDismissReason())
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
func (m *RecommendationsItemImpactedResourcesItemDismissPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDismissReason sets the dismissReason property value. The dismissReason property
func (m *RecommendationsItemImpactedResourcesItemDismissPostRequestBody) SetDismissReason(value *string)() {
    m.dismissReason = value
}
