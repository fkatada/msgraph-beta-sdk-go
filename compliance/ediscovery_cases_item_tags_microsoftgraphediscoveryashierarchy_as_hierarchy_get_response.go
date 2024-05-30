package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

type EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse instantiates a new EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse and sets the default values.
func NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse()(*EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse) {
    m := &EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []Tagable when successful
func (m *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse) GetValue()([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponse) SetValue(value []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
    SetValue(value []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)()
}
