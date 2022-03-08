package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IntuneBrandingProfileAssignment provides operations to manage the deviceManagement singleton.
type IntuneBrandingProfileAssignment struct {
    Entity
    // Assignment target that the branding profile is assigned to.
    target DeviceAndAppManagementAssignmentTargetable;
}
// NewIntuneBrandingProfileAssignment instantiates a new intuneBrandingProfileAssignment and sets the default values.
func NewIntuneBrandingProfileAssignment()(*IntuneBrandingProfileAssignment) {
    m := &IntuneBrandingProfileAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIntuneBrandingProfileAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntuneBrandingProfileAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewIntuneBrandingProfileAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntuneBrandingProfileAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. Assignment target that the branding profile is assigned to.
func (m *IntuneBrandingProfileAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *IntuneBrandingProfileAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IntuneBrandingProfileAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTarget sets the target property value. Assignment target that the branding profile is assigned to.
func (m *IntuneBrandingProfileAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
