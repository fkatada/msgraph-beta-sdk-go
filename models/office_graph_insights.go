package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeGraphInsights 
type OfficeGraphInsights struct {
    Entity
}
// NewOfficeGraphInsights instantiates a new officeGraphInsights and sets the default values.
func NewOfficeGraphInsights()(*OfficeGraphInsights) {
    m := &OfficeGraphInsights{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOfficeGraphInsightsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeGraphInsightsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.itemInsights":
                        return NewItemInsights(), nil
                }
            }
        }
    }
    return NewOfficeGraphInsights(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeGraphInsights) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["shared"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharedInsightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedInsightable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SharedInsightable)
                }
            }
            m.SetShared(res)
        }
        return nil
    }
    res["trending"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrendingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Trendingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Trendingable)
                }
            }
            m.SetTrending(res)
        }
        return nil
    }
    res["used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUsedInsightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UsedInsightable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UsedInsightable)
                }
            }
            m.SetUsed(res)
        }
        return nil
    }
    return res
}
// GetShared gets the shared property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) GetShared()([]SharedInsightable) {
    val, err := m.GetBackingStore().Get("shared")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SharedInsightable)
    }
    return nil
}
// GetTrending gets the trending property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) GetTrending()([]Trendingable) {
    val, err := m.GetBackingStore().Get("trending")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Trendingable)
    }
    return nil
}
// GetUsed gets the used property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) GetUsed()([]UsedInsightable) {
    val, err := m.GetBackingStore().Get("used")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UsedInsightable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OfficeGraphInsights) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetShared() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShared()))
        for i, v := range m.GetShared() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("shared", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrending() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrending()))
        for i, v := range m.GetTrending() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("trending", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsed() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsed()))
        for i, v := range m.GetUsed() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("used", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetShared sets the shared property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) SetShared(value []SharedInsightable)() {
    err := m.GetBackingStore().Set("shared", value)
    if err != nil {
        panic(err)
    }
}
// SetTrending sets the trending property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) SetTrending(value []Trendingable)() {
    err := m.GetBackingStore().Set("trending", value)
    if err != nil {
        panic(err)
    }
}
// SetUsed sets the used property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) SetUsed(value []UsedInsightable)() {
    err := m.GetBackingStore().Set("used", value)
    if err != nil {
        panic(err)
    }
}
// OfficeGraphInsightsable 
type OfficeGraphInsightsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetShared()([]SharedInsightable)
    GetTrending()([]Trendingable)
    GetUsed()([]UsedInsightable)
    SetShared(value []SharedInsightable)()
    SetTrending(value []Trendingable)()
    SetUsed(value []UsedInsightable)()
}
