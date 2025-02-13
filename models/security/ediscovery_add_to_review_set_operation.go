package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EdiscoveryAddToReviewSetOperation struct {
    CaseOperation
}
// NewEdiscoveryAddToReviewSetOperation instantiates a new EdiscoveryAddToReviewSetOperation and sets the default values.
func NewEdiscoveryAddToReviewSetOperation()(*EdiscoveryAddToReviewSetOperation) {
    m := &EdiscoveryAddToReviewSetOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoveryAddToReviewSetOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoveryAddToReviewSetOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryAddToReviewSetOperation(), nil
}
// GetAdditionalDataOptions gets the additionalDataOptions property value. The additionalDataOptions property
// returns a *AdditionalDataOptions when successful
func (m *EdiscoveryAddToReviewSetOperation) GetAdditionalDataOptions()(*AdditionalDataOptions) {
    val, err := m.GetBackingStore().Get("additionalDataOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AdditionalDataOptions)
    }
    return nil
}
// GetCloudAttachmentVersion gets the cloudAttachmentVersion property value. The cloudAttachmentVersion property
// returns a *CloudAttachmentVersion when successful
func (m *EdiscoveryAddToReviewSetOperation) GetCloudAttachmentVersion()(*CloudAttachmentVersion) {
    val, err := m.GetBackingStore().Get("cloudAttachmentVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAttachmentVersion)
    }
    return nil
}
// GetDocumentVersion gets the documentVersion property value. The documentVersion property
// returns a *DocumentVersion when successful
func (m *EdiscoveryAddToReviewSetOperation) GetDocumentVersion()(*DocumentVersion) {
    val, err := m.GetBackingStore().Get("documentVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DocumentVersion)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdiscoveryAddToReviewSetOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["additionalDataOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdditionalDataOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDataOptions(val.(*AdditionalDataOptions))
        }
        return nil
    }
    res["cloudAttachmentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAttachmentVersion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudAttachmentVersion(val.(*CloudAttachmentVersion))
        }
        return nil
    }
    res["documentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDocumentVersion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentVersion(val.(*DocumentVersion))
        }
        return nil
    }
    res["itemsToInclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseItemsToInclude)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemsToInclude(val.(*ItemsToInclude))
        }
        return nil
    }
    res["reviewSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(EdiscoveryReviewSetable))
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(EdiscoverySearchable))
        }
        return nil
    }
    return res
}
// GetItemsToInclude gets the itemsToInclude property value. The itemsToInclude property
// returns a *ItemsToInclude when successful
func (m *EdiscoveryAddToReviewSetOperation) GetItemsToInclude()(*ItemsToInclude) {
    val, err := m.GetBackingStore().Get("itemsToInclude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ItemsToInclude)
    }
    return nil
}
// GetReviewSet gets the reviewSet property value. eDiscovery review set to which items matching source collection query gets added.
// returns a EdiscoveryReviewSetable when successful
func (m *EdiscoveryAddToReviewSetOperation) GetReviewSet()(EdiscoveryReviewSetable) {
    val, err := m.GetBackingStore().Get("reviewSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EdiscoveryReviewSetable)
    }
    return nil
}
// GetSearch gets the search property value. eDiscovery search that gets added to review set.
// returns a EdiscoverySearchable when successful
func (m *EdiscoveryAddToReviewSetOperation) GetSearch()(EdiscoverySearchable) {
    val, err := m.GetBackingStore().Get("search")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EdiscoverySearchable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoveryAddToReviewSetOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalDataOptions() != nil {
        cast := (*m.GetAdditionalDataOptions()).String()
        err = writer.WriteStringValue("additionalDataOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudAttachmentVersion() != nil {
        cast := (*m.GetCloudAttachmentVersion()).String()
        err = writer.WriteStringValue("cloudAttachmentVersion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDocumentVersion() != nil {
        cast := (*m.GetDocumentVersion()).String()
        err = writer.WriteStringValue("documentVersion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetItemsToInclude() != nil {
        cast := (*m.GetItemsToInclude()).String()
        err = writer.WriteStringValue("itemsToInclude", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewSet", m.GetReviewSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalDataOptions sets the additionalDataOptions property value. The additionalDataOptions property
func (m *EdiscoveryAddToReviewSetOperation) SetAdditionalDataOptions(value *AdditionalDataOptions)() {
    err := m.GetBackingStore().Set("additionalDataOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudAttachmentVersion sets the cloudAttachmentVersion property value. The cloudAttachmentVersion property
func (m *EdiscoveryAddToReviewSetOperation) SetCloudAttachmentVersion(value *CloudAttachmentVersion)() {
    err := m.GetBackingStore().Set("cloudAttachmentVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentVersion sets the documentVersion property value. The documentVersion property
func (m *EdiscoveryAddToReviewSetOperation) SetDocumentVersion(value *DocumentVersion)() {
    err := m.GetBackingStore().Set("documentVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetItemsToInclude sets the itemsToInclude property value. The itemsToInclude property
func (m *EdiscoveryAddToReviewSetOperation) SetItemsToInclude(value *ItemsToInclude)() {
    err := m.GetBackingStore().Set("itemsToInclude", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewSet sets the reviewSet property value. eDiscovery review set to which items matching source collection query gets added.
func (m *EdiscoveryAddToReviewSetOperation) SetReviewSet(value EdiscoveryReviewSetable)() {
    err := m.GetBackingStore().Set("reviewSet", value)
    if err != nil {
        panic(err)
    }
}
// SetSearch sets the search property value. eDiscovery search that gets added to review set.
func (m *EdiscoveryAddToReviewSetOperation) SetSearch(value EdiscoverySearchable)() {
    err := m.GetBackingStore().Set("search", value)
    if err != nil {
        panic(err)
    }
}
type EdiscoveryAddToReviewSetOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalDataOptions()(*AdditionalDataOptions)
    GetCloudAttachmentVersion()(*CloudAttachmentVersion)
    GetDocumentVersion()(*DocumentVersion)
    GetItemsToInclude()(*ItemsToInclude)
    GetReviewSet()(EdiscoveryReviewSetable)
    GetSearch()(EdiscoverySearchable)
    SetAdditionalDataOptions(value *AdditionalDataOptions)()
    SetCloudAttachmentVersion(value *CloudAttachmentVersion)()
    SetDocumentVersion(value *DocumentVersion)()
    SetItemsToInclude(value *ItemsToInclude)()
    SetReviewSet(value EdiscoveryReviewSetable)()
    SetSearch(value EdiscoverySearchable)()
}
