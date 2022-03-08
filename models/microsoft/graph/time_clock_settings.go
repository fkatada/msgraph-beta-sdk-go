package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeClockSettings provides operations to manage the compliance singleton.
type TimeClockSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The aprroved location of the timeClock.
    approvedLocation GeoCoordinatesable;
}
// NewTimeClockSettings instantiates a new timeClockSettings and sets the default values.
func NewTimeClockSettings()(*TimeClockSettings) {
    m := &TimeClockSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTimeClockSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeClockSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTimeClockSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeClockSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApprovedLocation gets the approvedLocation property value. The aprroved location of the timeClock.
func (m *TimeClockSettings) GetApprovedLocation()(GeoCoordinatesable) {
    if m == nil {
        return nil
    } else {
        return m.approvedLocation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeClockSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approvedLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovedLocation(val.(GeoCoordinatesable))
        }
        return nil
    }
    return res
}
func (m *TimeClockSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeClockSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("approvedLocation", m.GetApprovedLocation())
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
func (m *TimeClockSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApprovedLocation sets the approvedLocation property value. The aprroved location of the timeClock.
func (m *TimeClockSettings) SetApprovedLocation(value GeoCoordinatesable)() {
    if m != nil {
        m.approvedLocation = value
    }
}
