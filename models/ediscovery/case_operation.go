package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CaseOperation provides operations to manage the collection of accessReviewDecision entities.
type CaseOperation struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
    action *CaseAction
    // The date and time the operation was completed.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user that created the operation.
    createdBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The date and time the operation was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The progress of the operation.
    percentProgress *int32
    // Contains success and failure-specific result information.
    resultInfo ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResultInfoable
    // The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
    status *CaseOperationStatus
}
// NewCaseOperation instantiates a new caseOperation and sets the default values.
func NewCaseOperation()(*CaseOperation) {
    m := &CaseOperation{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.caseOperation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCaseOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCaseOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.ediscovery.addToReviewSetOperation":
                        return NewAddToReviewSetOperation(), nil
                    case "#microsoft.graph.ediscovery.caseExportOperation":
                        return NewCaseExportOperation(), nil
                    case "#microsoft.graph.ediscovery.caseHoldOperation":
                        return NewCaseHoldOperation(), nil
                    case "#microsoft.graph.ediscovery.caseIndexOperation":
                        return NewCaseIndexOperation(), nil
                    case "#microsoft.graph.ediscovery.estimateStatisticsOperation":
                        return NewEstimateStatisticsOperation(), nil
                    case "#microsoft.graph.ediscovery.purgeDataOperation":
                        return NewPurgeDataOperation(), nil
                    case "#microsoft.graph.ediscovery.tagOperation":
                        return NewTagOperation(), nil
                }
            }
        }
    }
    return NewCaseOperation(), nil
}
// GetAction gets the action property value. The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
func (m *CaseOperation) GetAction()(*CaseAction) {
    return m.action
}
// GetCompletedDateTime gets the completedDateTime property value. The date and time the operation was completed.
func (m *CaseOperation) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedBy gets the createdBy property value. The user that created the operation.
func (m *CaseOperation) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the operation was created.
func (m *CaseOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CaseOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCaseAction , m.SetAction)
    res["completedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCompletedDateTime)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["percentProgress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPercentProgress)
    res["resultInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResultInfoFromDiscriminatorValue , m.SetResultInfo)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCaseOperationStatus , m.SetStatus)
    return res
}
// GetPercentProgress gets the percentProgress property value. The progress of the operation.
func (m *CaseOperation) GetPercentProgress()(*int32) {
    return m.percentProgress
}
// GetResultInfo gets the resultInfo property value. Contains success and failure-specific result information.
func (m *CaseOperation) GetResultInfo()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResultInfoable) {
    return m.resultInfo
}
// GetStatus gets the status property value. The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
func (m *CaseOperation) GetStatus()(*CaseOperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CaseOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentProgress", m.GetPercentProgress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resultInfo", m.GetResultInfo())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
func (m *CaseOperation) SetAction(value *CaseAction)() {
    m.action = value
}
// SetCompletedDateTime sets the completedDateTime property value. The date and time the operation was completed.
func (m *CaseOperation) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedBy sets the createdBy property value. The user that created the operation.
func (m *CaseOperation) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the operation was created.
func (m *CaseOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetPercentProgress sets the percentProgress property value. The progress of the operation.
func (m *CaseOperation) SetPercentProgress(value *int32)() {
    m.percentProgress = value
}
// SetResultInfo sets the resultInfo property value. Contains success and failure-specific result information.
func (m *CaseOperation) SetResultInfo(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResultInfoable)() {
    m.resultInfo = value
}
// SetStatus sets the status property value. The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
func (m *CaseOperation) SetStatus(value *CaseOperationStatus)() {
    m.status = value
}
