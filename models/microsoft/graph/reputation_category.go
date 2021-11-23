package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ReputationCategory 
type ReputationCategory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    description *string;
    // 
    name *string;
    // 
    vendor_escaped *string;
}
// NewReputationCategory instantiates a new reputationCategory and sets the default values.
func NewReputationCategory()(*ReputationCategory) {
    m := &ReputationCategory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReputationCategory) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. 
func (m *ReputationCategory) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetName gets the name property value. 
func (m *ReputationCategory) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetVendor_escaped gets the vendor_escaped property value. 
func (m *ReputationCategory) GetVendor_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReputationCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["vendor_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendor_escaped(val)
        }
        return nil
    }
    return res
}
func (m *ReputationCategory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ReputationCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendor_escaped", m.GetVendor_escaped())
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
func (m *ReputationCategory) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. 
func (m *ReputationCategory) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. 
func (m *ReputationCategory) SetName(value *string)() {
    m.name = value
}
// SetVendor_escaped sets the vendor_escaped property value. 
func (m *ReputationCategory) SetVendor_escaped(value *string)() {
    m.vendor_escaped = value
}
