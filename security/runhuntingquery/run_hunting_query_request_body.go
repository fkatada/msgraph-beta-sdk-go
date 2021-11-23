package runhuntingquery

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RunHuntingQueryRequestBody 
type RunHuntingQueryRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    query *string;
}
// NewRunHuntingQueryRequestBody instantiates a new runHuntingQueryRequestBody and sets the default values.
func NewRunHuntingQueryRequestBody()(*RunHuntingQueryRequestBody) {
    m := &RunHuntingQueryRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RunHuntingQueryRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetQuery gets the query property value. 
func (m *RunHuntingQueryRequestBody) GetQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RunHuntingQueryRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["query"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    return res
}
func (m *RunHuntingQueryRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RunHuntingQueryRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("query", m.GetQuery())
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
func (m *RunHuntingQueryRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetQuery sets the query property value. 
func (m *RunHuntingQueryRequestBody) SetQuery(value *string)() {
    m.query = value
}
