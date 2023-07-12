package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosiPadOSWebClip contains properties and inherited properties for iOS web apps.
type IosiPadOSWebClip struct {
    MobileApp
    // The OdataType property
    OdataType *string
}
// NewIosiPadOSWebClip instantiates a new iosiPadOSWebClip and sets the default values.
func NewIosiPadOSWebClip()(*IosiPadOSWebClip) {
    m := &IosiPadOSWebClip{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.iosiPadOSWebClip"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosiPadOSWebClipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosiPadOSWebClipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosiPadOSWebClip(), nil
}
// GetAppUrl gets the appUrl property value. Indicates iOS/iPadOS web clip app URL. Example: 'https://www.contoso.com'
func (m *IosiPadOSWebClip) GetAppUrl()(*string) {
    val, err := m.GetBackingStore().Get("appUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosiPadOSWebClip) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUrl(val)
        }
        return nil
    }
    res["fullScreenEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullScreenEnabled(val)
        }
        return nil
    }
    res["ignoreManifestScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreManifestScope(val)
        }
        return nil
    }
    res["preComposedIconEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreComposedIconEnabled(val)
        }
        return nil
    }
    res["targetApplicationBundleIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetApplicationBundleIdentifier(val)
        }
        return nil
    }
    res["useManagedBrowser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseManagedBrowser(val)
        }
        return nil
    }
    return res
}
// GetFullScreenEnabled gets the fullScreenEnabled property value. Whether or not to open the web clip as a full-screen web app. Defaults to false. If TRUE, opens the web clip as a full-screen web app. If FALSE, the web clip opens inside of another app, such as Safari or the app specified with targetApplicationBundleIdentifier.
func (m *IosiPadOSWebClip) GetFullScreenEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("fullScreenEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIgnoreManifestScope gets the ignoreManifestScope property value. Whether or not a full screen web clip can navigate to an external web site without showing the Safari UI. Defaults to false. If FALSE, the Safari UI appears when navigating away. If TRUE, the Safari UI will not be shown.
func (m *IosiPadOSWebClip) GetIgnoreManifestScope()(*bool) {
    val, err := m.GetBackingStore().Get("ignoreManifestScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPreComposedIconEnabled gets the preComposedIconEnabled property value. Whether or not the icon for the app is precomosed. Defaults to false. If TRUE, prevents SpringBoard from adding 'shine' to the icon. If FALSE, SpringBoard can add 'shine'.
func (m *IosiPadOSWebClip) GetPreComposedIconEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("preComposedIconEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTargetApplicationBundleIdentifier gets the targetApplicationBundleIdentifier property value. Specifies the application bundle identifier which opens the URL. Available in iOS 14 and later.
func (m *IosiPadOSWebClip) GetTargetApplicationBundleIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("targetApplicationBundleIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUseManagedBrowser gets the useManagedBrowser property value. Whether or not to use managed browser. When TRUE, the app will be required to be opened in Microsoft Edge. When FALSE, the app will not be required to be opened in Microsoft Edge. By default, this property is set to FALSE.
func (m *IosiPadOSWebClip) GetUseManagedBrowser()(*bool) {
    val, err := m.GetBackingStore().Get("useManagedBrowser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosiPadOSWebClip) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appUrl", m.GetAppUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fullScreenEnabled", m.GetFullScreenEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ignoreManifestScope", m.GetIgnoreManifestScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("preComposedIconEnabled", m.GetPreComposedIconEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetApplicationBundleIdentifier", m.GetTargetApplicationBundleIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useManagedBrowser", m.GetUseManagedBrowser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppUrl sets the appUrl property value. Indicates iOS/iPadOS web clip app URL. Example: 'https://www.contoso.com'
func (m *IosiPadOSWebClip) SetAppUrl(value *string)() {
    err := m.GetBackingStore().Set("appUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetFullScreenEnabled sets the fullScreenEnabled property value. Whether or not to open the web clip as a full-screen web app. Defaults to false. If TRUE, opens the web clip as a full-screen web app. If FALSE, the web clip opens inside of another app, such as Safari or the app specified with targetApplicationBundleIdentifier.
func (m *IosiPadOSWebClip) SetFullScreenEnabled(value *bool)() {
    err := m.GetBackingStore().Set("fullScreenEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIgnoreManifestScope sets the ignoreManifestScope property value. Whether or not a full screen web clip can navigate to an external web site without showing the Safari UI. Defaults to false. If FALSE, the Safari UI appears when navigating away. If TRUE, the Safari UI will not be shown.
func (m *IosiPadOSWebClip) SetIgnoreManifestScope(value *bool)() {
    err := m.GetBackingStore().Set("ignoreManifestScope", value)
    if err != nil {
        panic(err)
    }
}
// SetPreComposedIconEnabled sets the preComposedIconEnabled property value. Whether or not the icon for the app is precomosed. Defaults to false. If TRUE, prevents SpringBoard from adding 'shine' to the icon. If FALSE, SpringBoard can add 'shine'.
func (m *IosiPadOSWebClip) SetPreComposedIconEnabled(value *bool)() {
    err := m.GetBackingStore().Set("preComposedIconEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetApplicationBundleIdentifier sets the targetApplicationBundleIdentifier property value. Specifies the application bundle identifier which opens the URL. Available in iOS 14 and later.
func (m *IosiPadOSWebClip) SetTargetApplicationBundleIdentifier(value *string)() {
    err := m.GetBackingStore().Set("targetApplicationBundleIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetUseManagedBrowser sets the useManagedBrowser property value. Whether or not to use managed browser. When TRUE, the app will be required to be opened in Microsoft Edge. When FALSE, the app will not be required to be opened in Microsoft Edge. By default, this property is set to FALSE.
func (m *IosiPadOSWebClip) SetUseManagedBrowser(value *bool)() {
    err := m.GetBackingStore().Set("useManagedBrowser", value)
    if err != nil {
        panic(err)
    }
}
// IosiPadOSWebClipable 
type IosiPadOSWebClipable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppUrl()(*string)
    GetFullScreenEnabled()(*bool)
    GetIgnoreManifestScope()(*bool)
    GetPreComposedIconEnabled()(*bool)
    GetTargetApplicationBundleIdentifier()(*string)
    GetUseManagedBrowser()(*bool)
    SetAppUrl(value *string)()
    SetFullScreenEnabled(value *bool)()
    SetIgnoreManifestScope(value *bool)()
    SetPreComposedIconEnabled(value *bool)()
    SetTargetApplicationBundleIdentifier(value *string)()
    SetUseManagedBrowser(value *bool)()
}
