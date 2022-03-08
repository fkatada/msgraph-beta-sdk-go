package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcBulkRemoteActionResult provides operations to call the bulkRestoreCloudPc method.
type CloudPcBulkRemoteActionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A list of all the Intune managed device IDs that completed the bulk action with a failure.
    failedDeviceIds []string;
    // A list of all the Intune managed device IDs that were not found when the bulk action was attempted.
    notFoundDeviceIds []string;
    // A list of all the Intune managed device IDs that were identified as unsupported for the bulk action.
    notSupportedDeviceIds []string;
    // A list of all the Intune managed device IDs that completed the bulk action successfully.
    successfulDeviceIds []string;
}
// NewCloudPcBulkRemoteActionResult instantiates a new cloudPcBulkRemoteActionResult and sets the default values.
func NewCloudPcBulkRemoteActionResult()(*CloudPcBulkRemoteActionResult) {
    m := &CloudPcBulkRemoteActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcBulkRemoteActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcBulkRemoteActionResultFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcBulkRemoteActionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcBulkRemoteActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFailedDeviceIds gets the failedDeviceIds property value. A list of all the Intune managed device IDs that completed the bulk action with a failure.
func (m *CloudPcBulkRemoteActionResult) GetFailedDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcBulkRemoteActionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["failedDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFailedDeviceIds(res)
        }
        return nil
    }
    res["notFoundDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotFoundDeviceIds(res)
        }
        return nil
    }
    res["notSupportedDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotSupportedDeviceIds(res)
        }
        return nil
    }
    res["successfulDeviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSuccessfulDeviceIds(res)
        }
        return nil
    }
    return res
}
// GetNotFoundDeviceIds gets the notFoundDeviceIds property value. A list of all the Intune managed device IDs that were not found when the bulk action was attempted.
func (m *CloudPcBulkRemoteActionResult) GetNotFoundDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notFoundDeviceIds
    }
}
// GetNotSupportedDeviceIds gets the notSupportedDeviceIds property value. A list of all the Intune managed device IDs that were identified as unsupported for the bulk action.
func (m *CloudPcBulkRemoteActionResult) GetNotSupportedDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notSupportedDeviceIds
    }
}
// GetSuccessfulDeviceIds gets the successfulDeviceIds property value. A list of all the Intune managed device IDs that completed the bulk action successfully.
func (m *CloudPcBulkRemoteActionResult) GetSuccessfulDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.successfulDeviceIds
    }
}
func (m *CloudPcBulkRemoteActionResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcBulkRemoteActionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetFailedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("failedDeviceIds", m.GetFailedDeviceIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotFoundDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("notFoundDeviceIds", m.GetNotFoundDeviceIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotSupportedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("notSupportedDeviceIds", m.GetNotSupportedDeviceIds())
        if err != nil {
            return err
        }
    }
    if m.GetSuccessfulDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("successfulDeviceIds", m.GetSuccessfulDeviceIds())
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
func (m *CloudPcBulkRemoteActionResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFailedDeviceIds sets the failedDeviceIds property value. A list of all the Intune managed device IDs that completed the bulk action with a failure.
func (m *CloudPcBulkRemoteActionResult) SetFailedDeviceIds(value []string)() {
    if m != nil {
        m.failedDeviceIds = value
    }
}
// SetNotFoundDeviceIds sets the notFoundDeviceIds property value. A list of all the Intune managed device IDs that were not found when the bulk action was attempted.
func (m *CloudPcBulkRemoteActionResult) SetNotFoundDeviceIds(value []string)() {
    if m != nil {
        m.notFoundDeviceIds = value
    }
}
// SetNotSupportedDeviceIds sets the notSupportedDeviceIds property value. A list of all the Intune managed device IDs that were identified as unsupported for the bulk action.
func (m *CloudPcBulkRemoteActionResult) SetNotSupportedDeviceIds(value []string)() {
    if m != nil {
        m.notSupportedDeviceIds = value
    }
}
// SetSuccessfulDeviceIds sets the successfulDeviceIds property value. A list of all the Intune managed device IDs that completed the bulk action successfully.
func (m *CloudPcBulkRemoteActionResult) SetSuccessfulDeviceIds(value []string)() {
    if m != nil {
        m.successfulDeviceIds = value
    }
}
