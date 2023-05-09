package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
)

// CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse 
type CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewCallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse instantiates a new CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse and sets the default values.
func NewCallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse()(*CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse) {
    m := &CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreatePstnOnlineMeetingDialoutReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnOnlineMeetingDialoutReportable, len(val))
            for i, v := range val {
                res[i] = v.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnOnlineMeetingDialoutReportable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse) GetValue()([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnOnlineMeetingDialoutReportable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnOnlineMeetingDialoutReportable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponse) SetValue(value []iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnOnlineMeetingDialoutReportable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponseable 
type CallRecordsMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnOnlineMeetingDialoutReportable)
    SetValue(value []iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnOnlineMeetingDialoutReportable)()
}
