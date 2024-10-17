package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SiteProtectionUnitsBulkAdditionJob struct {
    ProtectionUnitsBulkJobBase
}
// NewSiteProtectionUnitsBulkAdditionJob instantiates a new SiteProtectionUnitsBulkAdditionJob and sets the default values.
func NewSiteProtectionUnitsBulkAdditionJob()(*SiteProtectionUnitsBulkAdditionJob) {
    m := &SiteProtectionUnitsBulkAdditionJob{
        ProtectionUnitsBulkJobBase: *NewProtectionUnitsBulkJobBase(),
    }
    odataTypeValue := "#microsoft.graph.siteProtectionUnitsBulkAdditionJob"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSiteProtectionUnitsBulkAdditionJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSiteProtectionUnitsBulkAdditionJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSiteProtectionUnitsBulkAdditionJob(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SiteProtectionUnitsBulkAdditionJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectionUnitsBulkJobBase.GetFieldDeserializers()
    res["siteIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSiteIds(res)
        }
        return nil
    }
    res["siteWebUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSiteWebUrls(res)
        }
        return nil
    }
    return res
}
// GetSiteIds gets the siteIds property value. The list of SharePoint site IDs to add to the SharePoint protection policy.
// returns a []string when successful
func (m *SiteProtectionUnitsBulkAdditionJob) GetSiteIds()([]string) {
    val, err := m.GetBackingStore().Get("siteIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSiteWebUrls gets the siteWebUrls property value. The list of SharePoint site URLs to add to the SharePoint protection policy.
// returns a []string when successful
func (m *SiteProtectionUnitsBulkAdditionJob) GetSiteWebUrls()([]string) {
    val, err := m.GetBackingStore().Get("siteWebUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SiteProtectionUnitsBulkAdditionJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectionUnitsBulkJobBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSiteIds() != nil {
        err = writer.WriteCollectionOfStringValues("siteIds", m.GetSiteIds())
        if err != nil {
            return err
        }
    }
    if m.GetSiteWebUrls() != nil {
        err = writer.WriteCollectionOfStringValues("siteWebUrls", m.GetSiteWebUrls())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSiteIds sets the siteIds property value. The list of SharePoint site IDs to add to the SharePoint protection policy.
func (m *SiteProtectionUnitsBulkAdditionJob) SetSiteIds(value []string)() {
    err := m.GetBackingStore().Set("siteIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteWebUrls sets the siteWebUrls property value. The list of SharePoint site URLs to add to the SharePoint protection policy.
func (m *SiteProtectionUnitsBulkAdditionJob) SetSiteWebUrls(value []string)() {
    err := m.GetBackingStore().Set("siteWebUrls", value)
    if err != nil {
        panic(err)
    }
}
type SiteProtectionUnitsBulkAdditionJobable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectionUnitsBulkJobBaseable
    GetSiteIds()([]string)
    GetSiteWebUrls()([]string)
    SetSiteIds(value []string)()
    SetSiteWebUrls(value []string)()
}
