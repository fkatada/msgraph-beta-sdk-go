package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Schema 
type Schema struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewSchema instantiates a new schema and sets the default values.
func NewSchema()(*Schema) {
    m := &Schema{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchema(), nil
}
// GetBaseType gets the baseType property value. Must be set to microsoft.graph.externalItem. Required.
func (m *Schema) GetBaseType()(*string) {
    val, err := m.GetBackingStore().Get("baseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Schema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseType(val)
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Propertyable, len(val))
            for i, v := range val {
                res[i] = v.(Propertyable)
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
func (m *Schema) GetProperties()([]Propertyable) {
    val, err := m.GetBackingStore().Get("properties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Propertyable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Schema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("baseType", m.GetBaseType())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBaseType sets the baseType property value. Must be set to microsoft.graph.externalItem. Required.
func (m *Schema) SetBaseType(value *string)() {
    err := m.GetBackingStore().Set("baseType", value)
    if err != nil {
        panic(err)
    }
}
// SetProperties sets the properties property value. The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
func (m *Schema) SetProperties(value []Propertyable)() {
    err := m.GetBackingStore().Set("properties", value)
    if err != nil {
        panic(err)
    }
}
// Schemaable 
type Schemaable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseType()(*string)
    GetProperties()([]Propertyable)
    SetBaseType(value *string)()
    SetProperties(value []Propertyable)()
}
