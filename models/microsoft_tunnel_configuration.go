package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelConfiguration entity that represents a collection of Microsoft Tunnel settings
type MicrosoftTunnelConfiguration struct {
    Entity
    // Additional settings that may be applied to the server
    advancedSettings []KeyValuePairable
    // The Default Domain appendix that will be used by the clients
    defaultDomainSuffix *string
    // The configuration's description (optional)
    description *string
    // When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to transfer data.
    disableUdpConnections *bool
    // The display name for the server configuration. This property is required when a server is created.
    displayName *string
    // The DNS servers that will be used by the clients
    dnsServers []string
    // When the configuration was last updated
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The port that both TCP and UPD will listen over on the server
    listenPort *int32
    // The subnet that will be used to allocate virtual address for the clients
    network *string
    // List of Scope Tags for this Entity instance
    roleScopeTagIds []string
    // Subsets of the routes that will not be routed by the server
    routeExcludes []string
    // The routes that will be routed by the server
    routeIncludes []string
    // Subsets of the routes that will not be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteExcludes'.
    routesExclude []string
    // The routes that will be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteIncludes'.
    routesInclude []string
    // The domains that will be resolved using the provided dns servers
    splitDNS []string
}
// NewMicrosoftTunnelConfiguration instantiates a new microsoftTunnelConfiguration and sets the default values.
func NewMicrosoftTunnelConfiguration()(*MicrosoftTunnelConfiguration) {
    m := &MicrosoftTunnelConfiguration{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.microsoftTunnelConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMicrosoftTunnelConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelConfiguration(), nil
}
// GetAdvancedSettings gets the advancedSettings property value. Additional settings that may be applied to the server
func (m *MicrosoftTunnelConfiguration) GetAdvancedSettings()([]KeyValuePairable) {
    return m.advancedSettings
}
// GetDefaultDomainSuffix gets the defaultDomainSuffix property value. The Default Domain appendix that will be used by the clients
func (m *MicrosoftTunnelConfiguration) GetDefaultDomainSuffix()(*string) {
    return m.defaultDomainSuffix
}
// GetDescription gets the description property value. The configuration's description (optional)
func (m *MicrosoftTunnelConfiguration) GetDescription()(*string) {
    return m.description
}
// GetDisableUdpConnections gets the disableUdpConnections property value. When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to transfer data.
func (m *MicrosoftTunnelConfiguration) GetDisableUdpConnections()(*bool) {
    return m.disableUdpConnections
}
// GetDisplayName gets the displayName property value. The display name for the server configuration. This property is required when a server is created.
func (m *MicrosoftTunnelConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetDnsServers gets the dnsServers property value. The DNS servers that will be used by the clients
func (m *MicrosoftTunnelConfiguration) GetDnsServers()([]string) {
    return m.dnsServers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetAdvancedSettings)
    res["defaultDomainSuffix"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultDomainSuffix)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["disableUdpConnections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDisableUdpConnections)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["dnsServers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDnsServers)
    res["lastUpdateDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastUpdateDateTime)
    res["listenPort"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetListenPort)
    res["network"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetwork)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["routeExcludes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRouteExcludes)
    res["routeIncludes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRouteIncludes)
    res["routesExclude"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoutesExclude)
    res["routesInclude"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoutesInclude)
    res["splitDNS"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSplitDNS)
    return res
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. When the configuration was last updated
func (m *MicrosoftTunnelConfiguration) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdateDateTime
}
// GetListenPort gets the listenPort property value. The port that both TCP and UPD will listen over on the server
func (m *MicrosoftTunnelConfiguration) GetListenPort()(*int32) {
    return m.listenPort
}
// GetNetwork gets the network property value. The subnet that will be used to allocate virtual address for the clients
func (m *MicrosoftTunnelConfiguration) GetNetwork()(*string) {
    return m.network
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance
func (m *MicrosoftTunnelConfiguration) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetRouteExcludes gets the routeExcludes property value. Subsets of the routes that will not be routed by the server
func (m *MicrosoftTunnelConfiguration) GetRouteExcludes()([]string) {
    return m.routeExcludes
}
// GetRouteIncludes gets the routeIncludes property value. The routes that will be routed by the server
func (m *MicrosoftTunnelConfiguration) GetRouteIncludes()([]string) {
    return m.routeIncludes
}
// GetRoutesExclude gets the routesExclude property value. Subsets of the routes that will not be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteExcludes'.
func (m *MicrosoftTunnelConfiguration) GetRoutesExclude()([]string) {
    return m.routesExclude
}
// GetRoutesInclude gets the routesInclude property value. The routes that will be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteIncludes'.
func (m *MicrosoftTunnelConfiguration) GetRoutesInclude()([]string) {
    return m.routesInclude
}
// GetSplitDNS gets the splitDNS property value. The domains that will be resolved using the provided dns servers
func (m *MicrosoftTunnelConfiguration) GetSplitDNS()([]string) {
    return m.splitDNS
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAdvancedSettings())
        err = writer.WriteCollectionOfObjectValues("advancedSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultDomainSuffix", m.GetDefaultDomainSuffix())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableUdpConnections", m.GetDisableUdpConnections())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDnsServers() != nil {
        err = writer.WriteCollectionOfStringValues("dnsServers", m.GetDnsServers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("listenPort", m.GetListenPort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetRouteExcludes() != nil {
        err = writer.WriteCollectionOfStringValues("routeExcludes", m.GetRouteExcludes())
        if err != nil {
            return err
        }
    }
    if m.GetRouteIncludes() != nil {
        err = writer.WriteCollectionOfStringValues("routeIncludes", m.GetRouteIncludes())
        if err != nil {
            return err
        }
    }
    if m.GetRoutesExclude() != nil {
        err = writer.WriteCollectionOfStringValues("routesExclude", m.GetRoutesExclude())
        if err != nil {
            return err
        }
    }
    if m.GetRoutesInclude() != nil {
        err = writer.WriteCollectionOfStringValues("routesInclude", m.GetRoutesInclude())
        if err != nil {
            return err
        }
    }
    if m.GetSplitDNS() != nil {
        err = writer.WriteCollectionOfStringValues("splitDNS", m.GetSplitDNS())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedSettings sets the advancedSettings property value. Additional settings that may be applied to the server
func (m *MicrosoftTunnelConfiguration) SetAdvancedSettings(value []KeyValuePairable)() {
    m.advancedSettings = value
}
// SetDefaultDomainSuffix sets the defaultDomainSuffix property value. The Default Domain appendix that will be used by the clients
func (m *MicrosoftTunnelConfiguration) SetDefaultDomainSuffix(value *string)() {
    m.defaultDomainSuffix = value
}
// SetDescription sets the description property value. The configuration's description (optional)
func (m *MicrosoftTunnelConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetDisableUdpConnections sets the disableUdpConnections property value. When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to transfer data.
func (m *MicrosoftTunnelConfiguration) SetDisableUdpConnections(value *bool)() {
    m.disableUdpConnections = value
}
// SetDisplayName sets the displayName property value. The display name for the server configuration. This property is required when a server is created.
func (m *MicrosoftTunnelConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDnsServers sets the dnsServers property value. The DNS servers that will be used by the clients
func (m *MicrosoftTunnelConfiguration) SetDnsServers(value []string)() {
    m.dnsServers = value
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. When the configuration was last updated
func (m *MicrosoftTunnelConfiguration) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
// SetListenPort sets the listenPort property value. The port that both TCP and UPD will listen over on the server
func (m *MicrosoftTunnelConfiguration) SetListenPort(value *int32)() {
    m.listenPort = value
}
// SetNetwork sets the network property value. The subnet that will be used to allocate virtual address for the clients
func (m *MicrosoftTunnelConfiguration) SetNetwork(value *string)() {
    m.network = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance
func (m *MicrosoftTunnelConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetRouteExcludes sets the routeExcludes property value. Subsets of the routes that will not be routed by the server
func (m *MicrosoftTunnelConfiguration) SetRouteExcludes(value []string)() {
    m.routeExcludes = value
}
// SetRouteIncludes sets the routeIncludes property value. The routes that will be routed by the server
func (m *MicrosoftTunnelConfiguration) SetRouteIncludes(value []string)() {
    m.routeIncludes = value
}
// SetRoutesExclude sets the routesExclude property value. Subsets of the routes that will not be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteExcludes'.
func (m *MicrosoftTunnelConfiguration) SetRoutesExclude(value []string)() {
    m.routesExclude = value
}
// SetRoutesInclude sets the routesInclude property value. The routes that will be routed by the server. This property is going to be deprecated with the option of using the new property, 'RouteIncludes'.
func (m *MicrosoftTunnelConfiguration) SetRoutesInclude(value []string)() {
    m.routesInclude = value
}
// SetSplitDNS sets the splitDNS property value. The domains that will be resolved using the provided dns servers
func (m *MicrosoftTunnelConfiguration) SetSplitDNS(value []string)() {
    m.splitDNS = value
}
