package getskypeforbusinesspeertopeeractivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod 
type GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    appSharing *int64;
    // 
    audio *int64;
    // 
    fileTransfer *int64;
    // 
    im *int64;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    video *int64;
}
// NewGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod instantiates a new getSkypeForBusinessPeerToPeerActivityCountsWithPeriod and sets the default values.
func NewGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod()(*GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) {
    m := &GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAppSharing gets the appSharing property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetAppSharing()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.appSharing
    }
}
// GetAudio gets the audio property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetAudio()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
// GetFileTransfer gets the fileTransfer property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetFileTransfer()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileTransfer
    }
}
// GetIm gets the im property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetIm()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.im
    }
}
// GetReportDate gets the reportDate property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetVideo gets the video property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetVideo()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appSharing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppSharing(val)
        }
        return nil
    }
    res["audio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudio(val)
        }
        return nil
    }
    res["fileTransfer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileTransfer(val)
        }
        return nil
    }
    res["im"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIm(val)
        }
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["video"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideo(val)
        }
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("appSharing", m.GetAppSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("audio", m.GetAudio())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("fileTransfer", m.GetFileTransfer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("im", m.GetIm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("video", m.GetVideo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppSharing sets the appSharing property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetAppSharing(value *int64)() {
    m.appSharing = value
}
// SetAudio sets the audio property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetAudio(value *int64)() {
    m.audio = value
}
// SetFileTransfer sets the fileTransfer property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetFileTransfer(value *int64)() {
    m.fileTransfer = value
}
// SetIm sets the im property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetIm(value *int64)() {
    m.im = value
}
// SetReportDate sets the reportDate property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetVideo sets the video property value. 
func (m *GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod) SetVideo(value *int64)() {
    m.video = value
}
