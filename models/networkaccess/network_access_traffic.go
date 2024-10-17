package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type NetworkAccessTraffic struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewNetworkAccessTraffic instantiates a new NetworkAccessTraffic and sets the default values.
func NewNetworkAccessTraffic()(*NetworkAccessTraffic) {
    m := &NetworkAccessTraffic{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNetworkAccessTrafficFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkAccessTrafficFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkAccessTraffic(), nil
}
// GetAction gets the action property value. Indicates what action to take based on filtering policies. The possible values are: block, allow.
// returns a *FilteringPolicyAction when successful
func (m *NetworkAccessTraffic) GetAction()(*FilteringPolicyAction) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FilteringPolicyAction)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NetworkAccessTraffic) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAgentVersion gets the agentVersion property value. Represents the version of the Global Secure Access client agent software. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetAgentVersion()(*string) {
    val, err := m.GetBackingStore().Get("agentVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApplicationSnapshot gets the applicationSnapshot property value. The applicationSnapshot property
// returns a ApplicationSnapshotable when successful
func (m *NetworkAccessTraffic) GetApplicationSnapshot()(ApplicationSnapshotable) {
    val, err := m.GetBackingStore().Get("applicationSnapshot")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApplicationSnapshotable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *NetworkAccessTraffic) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetConnectionId gets the connectionId property value. Represents a unique identifier assigned to a connection. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetConnectionId()(*string) {
    val, err := m.GetBackingStore().Get("connectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Represents the date and time when a network access traffic log entry was created. Supports $filter (eq) and $orderby.
// returns a *Time when successful
func (m *NetworkAccessTraffic) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationFQDN gets the destinationFQDN property value. Represents the Fully Qualified Domain Name (FQDN) of the destination host or server in a network communication. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetDestinationFQDN()(*string) {
    val, err := m.GetBackingStore().Get("destinationFQDN")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationIp gets the destinationIp property value. Represents the IP address of the destination host or server in a network communication. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetDestinationIp()(*string) {
    val, err := m.GetBackingStore().Get("destinationIp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationPort gets the destinationPort property value. Represents the network port number on the destination host or server in a network communication. Supports $filter (eq) and $orderby.
// returns a *int32 when successful
func (m *NetworkAccessTraffic) GetDestinationPort()(*int32) {
    val, err := m.GetBackingStore().Get("destinationPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDestinationUrl gets the destinationUrl property value. The destinationUrl property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetDestinationUrl()(*string) {
    val, err := m.GetBackingStore().Get("destinationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationWebCategory gets the destinationWebCategory property value. The destinationWebCategory property
// returns a WebCategoryable when successful
func (m *NetworkAccessTraffic) GetDestinationWebCategory()(WebCategoryable) {
    val, err := m.GetBackingStore().Get("destinationWebCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WebCategoryable)
    }
    return nil
}
// GetDevice gets the device property value. The device property
// returns a Deviceable when successful
func (m *NetworkAccessTraffic) GetDevice()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable) {
    val, err := m.GetBackingStore().Get("device")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)
    }
    return nil
}
// GetDeviceCategory gets the deviceCategory property value. Represents the category classification of a device within a network infrastructure. The possible values are: client, branch, unknownFutureValue. Supports $filter (eq) and $orderby.
// returns a *DeviceCategory when successful
func (m *NetworkAccessTraffic) GetDeviceCategory()(*DeviceCategory) {
    val, err := m.GetBackingStore().Get("deviceCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceCategory)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. Represents a unique identifier assigned to a device within a network infrastructure. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceOperatingSystem gets the deviceOperatingSystem property value. Represents the operating system installed on a device within a network infrastructure. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetDeviceOperatingSystem()(*string) {
    val, err := m.GetBackingStore().Get("deviceOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceOperatingSystemVersion gets the deviceOperatingSystemVersion property value. Represents the version or release number of the operating system installed on a device within a network infrastructure. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetDeviceOperatingSystemVersion()(*string) {
    val, err := m.GetBackingStore().Get("deviceOperatingSystemVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkAccessTraffic) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFilteringPolicyAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*FilteringPolicyAction))
        }
        return nil
    }
    res["agentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgentVersion(val)
        }
        return nil
    }
    res["applicationSnapshot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationSnapshotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationSnapshot(val.(ApplicationSnapshotable))
        }
        return nil
    }
    res["connectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["destinationFQDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationFQDN(val)
        }
        return nil
    }
    res["destinationIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationIp(val)
        }
        return nil
    }
    res["destinationPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationPort(val)
        }
        return nil
    }
    res["destinationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationUrl(val)
        }
        return nil
    }
    res["destinationWebCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationWebCategory(val.(WebCategoryable))
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable))
        }
        return nil
    }
    res["deviceCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val.(*DeviceCategory))
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOperatingSystem(val)
        }
        return nil
    }
    res["deviceOperatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOperatingSystemVersion(val)
        }
        return nil
    }
    res["filteringProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilteringProfileId(val)
        }
        return nil
    }
    res["filteringProfileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilteringProfileName(val)
        }
        return nil
    }
    res["headers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHeadersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaders(val.(Headersable))
        }
        return nil
    }
    res["httpMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHttpMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpMethod(val.(*HttpMethod))
        }
        return nil
    }
    res["initiatingProcessName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatingProcessName(val)
        }
        return nil
    }
    res["networkProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkingProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkProtocol(val.(*NetworkingProtocol))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["policyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["policyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["policyRuleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyRuleId(val)
        }
        return nil
    }
    res["policyRuleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyRuleName(val)
        }
        return nil
    }
    res["privateAccessDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivateAccessDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateAccessDetails(val.(PrivateAccessDetailsable))
        }
        return nil
    }
    res["receivedBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedBytes(val)
        }
        return nil
    }
    res["remoteNetworkId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteNetworkId(val)
        }
        return nil
    }
    res["resourceTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceTenantId(val)
        }
        return nil
    }
    res["responseCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseCode(val)
        }
        return nil
    }
    res["sentBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentBytes(val)
        }
        return nil
    }
    res["sessionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionId(val)
        }
        return nil
    }
    res["sourceIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIp(val)
        }
        return nil
    }
    res["sourcePort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourcePort(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["threatType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatType(val)
        }
        return nil
    }
    res["trafficType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrafficType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrafficType(val.(*TrafficType))
        }
        return nil
    }
    res["transactionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionId(val)
        }
        return nil
    }
    res["transportProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkingProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransportProtocol(val.(*NetworkingProtocol))
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["vendorNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetVendorNames(res)
        }
        return nil
    }
    return res
}
// GetFilteringProfileId gets the filteringProfileId property value. The filteringProfileId property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetFilteringProfileId()(*string) {
    val, err := m.GetBackingStore().Get("filteringProfileId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFilteringProfileName gets the filteringProfileName property value. The filteringProfileName property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetFilteringProfileName()(*string) {
    val, err := m.GetBackingStore().Get("filteringProfileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHeaders gets the headers property value. Represents the headers included in a network request or response. Supports $filter (eq) and $orderby.
// returns a Headersable when successful
func (m *NetworkAccessTraffic) GetHeaders()(Headersable) {
    val, err := m.GetBackingStore().Get("headers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Headersable)
    }
    return nil
}
// GetHttpMethod gets the httpMethod property value. The httpMethod property
// returns a *HttpMethod when successful
func (m *NetworkAccessTraffic) GetHttpMethod()(*HttpMethod) {
    val, err := m.GetBackingStore().Get("httpMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*HttpMethod)
    }
    return nil
}
// GetInitiatingProcessName gets the initiatingProcessName property value. The initiatingProcessName property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetInitiatingProcessName()(*string) {
    val, err := m.GetBackingStore().Get("initiatingProcessName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkProtocol gets the networkProtocol property value. Represents the networking protocol used for communication.The possible values are: ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII, unknownFutureValue. Supports $filter (eq) and $orderby.
// returns a *NetworkingProtocol when successful
func (m *NetworkAccessTraffic) GetNetworkProtocol()(*NetworkingProtocol) {
    val, err := m.GetBackingStore().Get("networkProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NetworkingProtocol)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyId gets the policyId property value. Represents a unique identifier assigned to a policy. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetPolicyId()(*string) {
    val, err := m.GetBackingStore().Get("policyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyName gets the policyName property value. The policyName property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetPolicyName()(*string) {
    val, err := m.GetBackingStore().Get("policyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyRuleId gets the policyRuleId property value. Represents a unique identifier assigned to a policy rule. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetPolicyRuleId()(*string) {
    val, err := m.GetBackingStore().Get("policyRuleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyRuleName gets the policyRuleName property value. The policyRuleName property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetPolicyRuleName()(*string) {
    val, err := m.GetBackingStore().Get("policyRuleName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrivateAccessDetails gets the privateAccessDetails property value. The privateAccessDetails property
// returns a PrivateAccessDetailsable when successful
func (m *NetworkAccessTraffic) GetPrivateAccessDetails()(PrivateAccessDetailsable) {
    val, err := m.GetBackingStore().Get("privateAccessDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivateAccessDetailsable)
    }
    return nil
}
// GetReceivedBytes gets the receivedBytes property value. Represents the total number of bytes received in a network communication or data transfer. Supports $filter (eq) and $orderby.
// returns a *int64 when successful
func (m *NetworkAccessTraffic) GetReceivedBytes()(*int64) {
    val, err := m.GetBackingStore().Get("receivedBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRemoteNetworkId gets the remoteNetworkId property value. The remoteNetworkId property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetRemoteNetworkId()(*string) {
    val, err := m.GetBackingStore().Get("remoteNetworkId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceTenantId gets the resourceTenantId property value. The resourceTenantId property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetResourceTenantId()(*string) {
    val, err := m.GetBackingStore().Get("resourceTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResponseCode gets the responseCode property value. The responseCode property
// returns a *int32 when successful
func (m *NetworkAccessTraffic) GetResponseCode()(*int32) {
    val, err := m.GetBackingStore().Get("responseCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSentBytes gets the sentBytes property value. Represents the total number of bytes sent in a network communication or data transfer. Supports $filter (eq) and $orderby.
// returns a *int64 when successful
func (m *NetworkAccessTraffic) GetSentBytes()(*int64) {
    val, err := m.GetBackingStore().Get("sentBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSessionId gets the sessionId property value. Represents a unique identifier assigned to a session or connection within a network infrastructure. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetSessionId()(*string) {
    val, err := m.GetBackingStore().Get("sessionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSourceIp gets the sourceIp property value. Represents the source IP address in a network communication. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetSourceIp()(*string) {
    val, err := m.GetBackingStore().Get("sourceIp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSourcePort gets the sourcePort property value. Represents the network port number on the source host or device in a network communication. Supports $filter (eq) and $orderby.
// returns a *int32 when successful
func (m *NetworkAccessTraffic) GetSourcePort()(*int32) {
    val, err := m.GetBackingStore().Get("sourcePort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTenantId gets the tenantId property value. Represents a unique identifier assigned to a tenant within a network infrastructure. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThreatType gets the threatType property value. The threatType property
// returns a *string when successful
func (m *NetworkAccessTraffic) GetThreatType()(*string) {
    val, err := m.GetBackingStore().Get("threatType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTrafficType gets the trafficType property value. The trafficType property
// returns a *TrafficType when successful
func (m *NetworkAccessTraffic) GetTrafficType()(*TrafficType) {
    val, err := m.GetBackingStore().Get("trafficType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TrafficType)
    }
    return nil
}
// GetTransactionId gets the transactionId property value. Represents a unique identifier assigned to a specific transaction or operation. Key. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetTransactionId()(*string) {
    val, err := m.GetBackingStore().Get("transactionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTransportProtocol gets the transportProtocol property value. Represents the transport protocol used for communication.The possible values are: ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII, unknownFutureValue. Supports $filter (eq) and $orderby.
// returns a *NetworkingProtocol when successful
func (m *NetworkAccessTraffic) GetTransportProtocol()(*NetworkingProtocol) {
    val, err := m.GetBackingStore().Get("transportProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NetworkingProtocol)
    }
    return nil
}
// GetUser gets the user property value. The user property
// returns a Userable when successful
func (m *NetworkAccessTraffic) GetUser()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    val, err := m.GetBackingStore().Get("user")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)
    }
    return nil
}
// GetUserId gets the userId property value. Represents a unique identifier assigned to a user. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. Represents the user principal name (UPN) associated with a user. Supports $filter (eq) and $orderby.
// returns a *string when successful
func (m *NetworkAccessTraffic) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVendorNames gets the vendorNames property value. The vendorNames property
// returns a []string when successful
func (m *NetworkAccessTraffic) GetVendorNames()([]string) {
    val, err := m.GetBackingStore().Get("vendorNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *NetworkAccessTraffic) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("agentVersion", m.GetAgentVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("applicationSnapshot", m.GetApplicationSnapshot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectionId", m.GetConnectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationFQDN", m.GetDestinationFQDN())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationIp", m.GetDestinationIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("destinationPort", m.GetDestinationPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationUrl", m.GetDestinationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("destinationWebCategory", m.GetDestinationWebCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCategory() != nil {
        cast := (*m.GetDeviceCategory()).String()
        err := writer.WriteStringValue("deviceCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceOperatingSystem", m.GetDeviceOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceOperatingSystemVersion", m.GetDeviceOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filteringProfileId", m.GetFilteringProfileId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filteringProfileName", m.GetFilteringProfileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("headers", m.GetHeaders())
        if err != nil {
            return err
        }
    }
    if m.GetHttpMethod() != nil {
        cast := (*m.GetHttpMethod()).String()
        err := writer.WriteStringValue("httpMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("initiatingProcessName", m.GetInitiatingProcessName())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkProtocol() != nil {
        cast := (*m.GetNetworkProtocol()).String()
        err := writer.WriteStringValue("networkProtocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyRuleId", m.GetPolicyRuleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyRuleName", m.GetPolicyRuleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("privateAccessDetails", m.GetPrivateAccessDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("receivedBytes", m.GetReceivedBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remoteNetworkId", m.GetRemoteNetworkId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceTenantId", m.GetResourceTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("responseCode", m.GetResponseCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("sentBytes", m.GetSentBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionId", m.GetSessionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceIp", m.GetSourceIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sourcePort", m.GetSourcePort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("threatType", m.GetThreatType())
        if err != nil {
            return err
        }
    }
    if m.GetTrafficType() != nil {
        cast := (*m.GetTrafficType()).String()
        err := writer.WriteStringValue("trafficType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("transactionId", m.GetTransactionId())
        if err != nil {
            return err
        }
    }
    if m.GetTransportProtocol() != nil {
        cast := (*m.GetTransportProtocol()).String()
        err := writer.WriteStringValue("transportProtocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetVendorNames() != nil {
        err := writer.WriteCollectionOfStringValues("vendorNames", m.GetVendorNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. Indicates what action to take based on filtering policies. The possible values are: block, allow.
func (m *NetworkAccessTraffic) SetAction(value *FilteringPolicyAction)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkAccessTraffic) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAgentVersion sets the agentVersion property value. Represents the version of the Global Secure Access client agent software. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetAgentVersion(value *string)() {
    err := m.GetBackingStore().Set("agentVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationSnapshot sets the applicationSnapshot property value. The applicationSnapshot property
func (m *NetworkAccessTraffic) SetApplicationSnapshot(value ApplicationSnapshotable)() {
    err := m.GetBackingStore().Set("applicationSnapshot", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *NetworkAccessTraffic) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetConnectionId sets the connectionId property value. Represents a unique identifier assigned to a connection. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetConnectionId(value *string)() {
    err := m.GetBackingStore().Set("connectionId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Represents the date and time when a network access traffic log entry was created. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *NetworkAccessTraffic) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationFQDN sets the destinationFQDN property value. Represents the Fully Qualified Domain Name (FQDN) of the destination host or server in a network communication. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetDestinationFQDN(value *string)() {
    err := m.GetBackingStore().Set("destinationFQDN", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationIp sets the destinationIp property value. Represents the IP address of the destination host or server in a network communication. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetDestinationIp(value *string)() {
    err := m.GetBackingStore().Set("destinationIp", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationPort sets the destinationPort property value. Represents the network port number on the destination host or server in a network communication. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetDestinationPort(value *int32)() {
    err := m.GetBackingStore().Set("destinationPort", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationUrl sets the destinationUrl property value. The destinationUrl property
func (m *NetworkAccessTraffic) SetDestinationUrl(value *string)() {
    err := m.GetBackingStore().Set("destinationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationWebCategory sets the destinationWebCategory property value. The destinationWebCategory property
func (m *NetworkAccessTraffic) SetDestinationWebCategory(value WebCategoryable)() {
    err := m.GetBackingStore().Set("destinationWebCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetDevice sets the device property value. The device property
func (m *NetworkAccessTraffic) SetDevice(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)() {
    err := m.GetBackingStore().Set("device", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCategory sets the deviceCategory property value. Represents the category classification of a device within a network infrastructure. The possible values are: client, branch, unknownFutureValue. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetDeviceCategory(value *DeviceCategory)() {
    err := m.GetBackingStore().Set("deviceCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. Represents a unique identifier assigned to a device within a network infrastructure. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceOperatingSystem sets the deviceOperatingSystem property value. Represents the operating system installed on a device within a network infrastructure. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetDeviceOperatingSystem(value *string)() {
    err := m.GetBackingStore().Set("deviceOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceOperatingSystemVersion sets the deviceOperatingSystemVersion property value. Represents the version or release number of the operating system installed on a device within a network infrastructure. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetDeviceOperatingSystemVersion(value *string)() {
    err := m.GetBackingStore().Set("deviceOperatingSystemVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetFilteringProfileId sets the filteringProfileId property value. The filteringProfileId property
func (m *NetworkAccessTraffic) SetFilteringProfileId(value *string)() {
    err := m.GetBackingStore().Set("filteringProfileId", value)
    if err != nil {
        panic(err)
    }
}
// SetFilteringProfileName sets the filteringProfileName property value. The filteringProfileName property
func (m *NetworkAccessTraffic) SetFilteringProfileName(value *string)() {
    err := m.GetBackingStore().Set("filteringProfileName", value)
    if err != nil {
        panic(err)
    }
}
// SetHeaders sets the headers property value. Represents the headers included in a network request or response. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetHeaders(value Headersable)() {
    err := m.GetBackingStore().Set("headers", value)
    if err != nil {
        panic(err)
    }
}
// SetHttpMethod sets the httpMethod property value. The httpMethod property
func (m *NetworkAccessTraffic) SetHttpMethod(value *HttpMethod)() {
    err := m.GetBackingStore().Set("httpMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetInitiatingProcessName sets the initiatingProcessName property value. The initiatingProcessName property
func (m *NetworkAccessTraffic) SetInitiatingProcessName(value *string)() {
    err := m.GetBackingStore().Set("initiatingProcessName", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkProtocol sets the networkProtocol property value. Represents the networking protocol used for communication.The possible values are: ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII, unknownFutureValue. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetNetworkProtocol(value *NetworkingProtocol)() {
    err := m.GetBackingStore().Set("networkProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NetworkAccessTraffic) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyId sets the policyId property value. Represents a unique identifier assigned to a policy. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetPolicyId(value *string)() {
    err := m.GetBackingStore().Set("policyId", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyName sets the policyName property value. The policyName property
func (m *NetworkAccessTraffic) SetPolicyName(value *string)() {
    err := m.GetBackingStore().Set("policyName", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyRuleId sets the policyRuleId property value. Represents a unique identifier assigned to a policy rule. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetPolicyRuleId(value *string)() {
    err := m.GetBackingStore().Set("policyRuleId", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyRuleName sets the policyRuleName property value. The policyRuleName property
func (m *NetworkAccessTraffic) SetPolicyRuleName(value *string)() {
    err := m.GetBackingStore().Set("policyRuleName", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivateAccessDetails sets the privateAccessDetails property value. The privateAccessDetails property
func (m *NetworkAccessTraffic) SetPrivateAccessDetails(value PrivateAccessDetailsable)() {
    err := m.GetBackingStore().Set("privateAccessDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivedBytes sets the receivedBytes property value. Represents the total number of bytes received in a network communication or data transfer. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetReceivedBytes(value *int64)() {
    err := m.GetBackingStore().Set("receivedBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteNetworkId sets the remoteNetworkId property value. The remoteNetworkId property
func (m *NetworkAccessTraffic) SetRemoteNetworkId(value *string)() {
    err := m.GetBackingStore().Set("remoteNetworkId", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceTenantId sets the resourceTenantId property value. The resourceTenantId property
func (m *NetworkAccessTraffic) SetResourceTenantId(value *string)() {
    err := m.GetBackingStore().Set("resourceTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetResponseCode sets the responseCode property value. The responseCode property
func (m *NetworkAccessTraffic) SetResponseCode(value *int32)() {
    err := m.GetBackingStore().Set("responseCode", value)
    if err != nil {
        panic(err)
    }
}
// SetSentBytes sets the sentBytes property value. Represents the total number of bytes sent in a network communication or data transfer. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetSentBytes(value *int64)() {
    err := m.GetBackingStore().Set("sentBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetSessionId sets the sessionId property value. Represents a unique identifier assigned to a session or connection within a network infrastructure. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetSessionId(value *string)() {
    err := m.GetBackingStore().Set("sessionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceIp sets the sourceIp property value. Represents the source IP address in a network communication. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetSourceIp(value *string)() {
    err := m.GetBackingStore().Set("sourceIp", value)
    if err != nil {
        panic(err)
    }
}
// SetSourcePort sets the sourcePort property value. Represents the network port number on the source host or device in a network communication. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetSourcePort(value *int32)() {
    err := m.GetBackingStore().Set("sourcePort", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. Represents a unique identifier assigned to a tenant within a network infrastructure. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetThreatType sets the threatType property value. The threatType property
func (m *NetworkAccessTraffic) SetThreatType(value *string)() {
    err := m.GetBackingStore().Set("threatType", value)
    if err != nil {
        panic(err)
    }
}
// SetTrafficType sets the trafficType property value. The trafficType property
func (m *NetworkAccessTraffic) SetTrafficType(value *TrafficType)() {
    err := m.GetBackingStore().Set("trafficType", value)
    if err != nil {
        panic(err)
    }
}
// SetTransactionId sets the transactionId property value. Represents a unique identifier assigned to a specific transaction or operation. Key. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetTransactionId(value *string)() {
    err := m.GetBackingStore().Set("transactionId", value)
    if err != nil {
        panic(err)
    }
}
// SetTransportProtocol sets the transportProtocol property value. Represents the transport protocol used for communication.The possible values are: ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII, unknownFutureValue. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetTransportProtocol(value *NetworkingProtocol)() {
    err := m.GetBackingStore().Set("transportProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetUser sets the user property value. The user property
func (m *NetworkAccessTraffic) SetUser(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    err := m.GetBackingStore().Set("user", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. Represents a unique identifier assigned to a user. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Represents the user principal name (UPN) associated with a user. Supports $filter (eq) and $orderby.
func (m *NetworkAccessTraffic) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorNames sets the vendorNames property value. The vendorNames property
func (m *NetworkAccessTraffic) SetVendorNames(value []string)() {
    err := m.GetBackingStore().Set("vendorNames", value)
    if err != nil {
        panic(err)
    }
}
type NetworkAccessTrafficable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*FilteringPolicyAction)
    GetAgentVersion()(*string)
    GetApplicationSnapshot()(ApplicationSnapshotable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetConnectionId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDestinationFQDN()(*string)
    GetDestinationIp()(*string)
    GetDestinationPort()(*int32)
    GetDestinationUrl()(*string)
    GetDestinationWebCategory()(WebCategoryable)
    GetDevice()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)
    GetDeviceCategory()(*DeviceCategory)
    GetDeviceId()(*string)
    GetDeviceOperatingSystem()(*string)
    GetDeviceOperatingSystemVersion()(*string)
    GetFilteringProfileId()(*string)
    GetFilteringProfileName()(*string)
    GetHeaders()(Headersable)
    GetHttpMethod()(*HttpMethod)
    GetInitiatingProcessName()(*string)
    GetNetworkProtocol()(*NetworkingProtocol)
    GetOdataType()(*string)
    GetPolicyId()(*string)
    GetPolicyName()(*string)
    GetPolicyRuleId()(*string)
    GetPolicyRuleName()(*string)
    GetPrivateAccessDetails()(PrivateAccessDetailsable)
    GetReceivedBytes()(*int64)
    GetRemoteNetworkId()(*string)
    GetResourceTenantId()(*string)
    GetResponseCode()(*int32)
    GetSentBytes()(*int64)
    GetSessionId()(*string)
    GetSourceIp()(*string)
    GetSourcePort()(*int32)
    GetTenantId()(*string)
    GetThreatType()(*string)
    GetTrafficType()(*TrafficType)
    GetTransactionId()(*string)
    GetTransportProtocol()(*NetworkingProtocol)
    GetUser()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetVendorNames()([]string)
    SetAction(value *FilteringPolicyAction)()
    SetAgentVersion(value *string)()
    SetApplicationSnapshot(value ApplicationSnapshotable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetConnectionId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDestinationFQDN(value *string)()
    SetDestinationIp(value *string)()
    SetDestinationPort(value *int32)()
    SetDestinationUrl(value *string)()
    SetDestinationWebCategory(value WebCategoryable)()
    SetDevice(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable)()
    SetDeviceCategory(value *DeviceCategory)()
    SetDeviceId(value *string)()
    SetDeviceOperatingSystem(value *string)()
    SetDeviceOperatingSystemVersion(value *string)()
    SetFilteringProfileId(value *string)()
    SetFilteringProfileName(value *string)()
    SetHeaders(value Headersable)()
    SetHttpMethod(value *HttpMethod)()
    SetInitiatingProcessName(value *string)()
    SetNetworkProtocol(value *NetworkingProtocol)()
    SetOdataType(value *string)()
    SetPolicyId(value *string)()
    SetPolicyName(value *string)()
    SetPolicyRuleId(value *string)()
    SetPolicyRuleName(value *string)()
    SetPrivateAccessDetails(value PrivateAccessDetailsable)()
    SetReceivedBytes(value *int64)()
    SetRemoteNetworkId(value *string)()
    SetResourceTenantId(value *string)()
    SetResponseCode(value *int32)()
    SetSentBytes(value *int64)()
    SetSessionId(value *string)()
    SetSourceIp(value *string)()
    SetSourcePort(value *int32)()
    SetTenantId(value *string)()
    SetThreatType(value *string)()
    SetTrafficType(value *TrafficType)()
    SetTransactionId(value *string)()
    SetTransportProtocol(value *NetworkingProtocol)()
    SetUser(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetVendorNames(value []string)()
}
