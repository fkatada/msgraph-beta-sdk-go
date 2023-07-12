package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchDetectedSensitiveContent 
type ExactMatchDetectedSensitiveContent struct {
    DetectedSensitiveContentBase
}
// NewExactMatchDetectedSensitiveContent instantiates a new exactMatchDetectedSensitiveContent and sets the default values.
func NewExactMatchDetectedSensitiveContent()(*ExactMatchDetectedSensitiveContent) {
    m := &ExactMatchDetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    return m
}
// CreateExactMatchDetectedSensitiveContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchDetectedSensitiveContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchDetectedSensitiveContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDetectedSensitiveContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DetectedSensitiveContentBase.GetFieldDeserializers()
    res["matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitiveContentLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveContentLocationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SensitiveContentLocationable)
                }
            }
            m.SetMatches(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetMatches gets the matches property value. The matches property
func (m *ExactMatchDetectedSensitiveContent) GetMatches()([]SensitiveContentLocationable) {
    val, err := m.GetBackingStore().Get("matches")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SensitiveContentLocationable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExactMatchDetectedSensitiveContent) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExactMatchDetectedSensitiveContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DetectedSensitiveContentBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("matches", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMatches sets the matches property value. The matches property
func (m *ExactMatchDetectedSensitiveContent) SetMatches(value []SensitiveContentLocationable)() {
    err := m.GetBackingStore().Set("matches", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExactMatchDetectedSensitiveContent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// ExactMatchDetectedSensitiveContentable 
type ExactMatchDetectedSensitiveContentable interface {
    DetectedSensitiveContentBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMatches()([]SensitiveContentLocationable)
    GetOdataType()(*string)
    SetMatches(value []SensitiveContentLocationable)()
    SetOdataType(value *string)()
}
