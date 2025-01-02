package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FallbackToMicrosoftProviderOnError struct {
    CustomExtensionBehaviorOnError
}
// NewFallbackToMicrosoftProviderOnError instantiates a new FallbackToMicrosoftProviderOnError and sets the default values.
func NewFallbackToMicrosoftProviderOnError()(*FallbackToMicrosoftProviderOnError) {
    m := &FallbackToMicrosoftProviderOnError{
        CustomExtensionBehaviorOnError: *NewCustomExtensionBehaviorOnError(),
    }
    odataTypeValue := "#microsoft.graph.fallbackToMicrosoftProviderOnError"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFallbackToMicrosoftProviderOnErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFallbackToMicrosoftProviderOnErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFallbackToMicrosoftProviderOnError(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FallbackToMicrosoftProviderOnError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionBehaviorOnError.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *FallbackToMicrosoftProviderOnError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionBehaviorOnError.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type FallbackToMicrosoftProviderOnErrorable interface {
    CustomExtensionBehaviorOnErrorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
