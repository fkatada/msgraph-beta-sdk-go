package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IpApplicationSegment struct {
    ApplicationSegment
}
// NewIpApplicationSegment instantiates a new IpApplicationSegment and sets the default values.
func NewIpApplicationSegment()(*IpApplicationSegment) {
    m := &IpApplicationSegment{
        ApplicationSegment: *NewApplicationSegment(),
    }
    odataTypeValue := "#microsoft.graph.ipApplicationSegment"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIpApplicationSegmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIpApplicationSegmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpApplicationSegment(), nil
}
// GetApplication gets the application property value. The on-premises nonweb application published through Microsoft Entra application proxy. Expanded by default and supports $expand.
// returns a Applicationable when successful
func (m *IpApplicationSegment) GetApplication()(Applicationable) {
    val, err := m.GetBackingStore().Get("application")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Applicationable)
    }
    return nil
}
// GetDestinationHost gets the destinationHost property value. Either the IP address, IP range, or FQDN of the applicationSegment, with or without wildcards.
// returns a *string when successful
func (m *IpApplicationSegment) GetDestinationHost()(*string) {
    val, err := m.GetBackingStore().Get("destinationHost")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationType gets the destinationType property value. The possible values are: ipAddress, ipRange, ipRangeCidr, fqdn, dnsSuffix, unknownFutureValue.
// returns a *PrivateNetworkDestinationType when successful
func (m *IpApplicationSegment) GetDestinationType()(*PrivateNetworkDestinationType) {
    val, err := m.GetBackingStore().Get("destinationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrivateNetworkDestinationType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IpApplicationSegment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApplicationSegment.GetFieldDeserializers()
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(Applicationable))
        }
        return nil
    }
    res["destinationHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationHost(val)
        }
        return nil
    }
    res["destinationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivateNetworkDestinationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationType(val.(*PrivateNetworkDestinationType))
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["ports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPorts(res)
        }
        return nil
    }
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivateNetworkProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val.(*PrivateNetworkProtocol))
        }
        return nil
    }
    return res
}
// GetPort gets the port property value. Port supported for the application segment. DO NOT USE.
// returns a *int32 when successful
func (m *IpApplicationSegment) GetPort()(*int32) {
    val, err := m.GetBackingStore().Get("port")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPorts gets the ports property value. List of ports supported for the application segment.
// returns a []string when successful
func (m *IpApplicationSegment) GetPorts()([]string) {
    val, err := m.GetBackingStore().Get("ports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetProtocol gets the protocol property value. Indicates the protocol of the network traffic acquired for the application segment. The possible values are: tcp, udp, unknownFutureValue.
// returns a *PrivateNetworkProtocol when successful
func (m *IpApplicationSegment) GetProtocol()(*PrivateNetworkProtocol) {
    val, err := m.GetBackingStore().Get("protocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PrivateNetworkProtocol)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IpApplicationSegment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ApplicationSegment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("destinationHost", m.GetDestinationHost())
        if err != nil {
            return err
        }
    }
    if m.GetDestinationType() != nil {
        cast := (*m.GetDestinationType()).String()
        err = writer.WriteStringValue("destinationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    if m.GetPorts() != nil {
        err = writer.WriteCollectionOfStringValues("ports", m.GetPorts())
        if err != nil {
            return err
        }
    }
    if m.GetProtocol() != nil {
        cast := (*m.GetProtocol()).String()
        err = writer.WriteStringValue("protocol", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The on-premises nonweb application published through Microsoft Entra application proxy. Expanded by default and supports $expand.
func (m *IpApplicationSegment) SetApplication(value Applicationable)() {
    err := m.GetBackingStore().Set("application", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationHost sets the destinationHost property value. Either the IP address, IP range, or FQDN of the applicationSegment, with or without wildcards.
func (m *IpApplicationSegment) SetDestinationHost(value *string)() {
    err := m.GetBackingStore().Set("destinationHost", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationType sets the destinationType property value. The possible values are: ipAddress, ipRange, ipRangeCidr, fqdn, dnsSuffix, unknownFutureValue.
func (m *IpApplicationSegment) SetDestinationType(value *PrivateNetworkDestinationType)() {
    err := m.GetBackingStore().Set("destinationType", value)
    if err != nil {
        panic(err)
    }
}
// SetPort sets the port property value. Port supported for the application segment. DO NOT USE.
func (m *IpApplicationSegment) SetPort(value *int32)() {
    err := m.GetBackingStore().Set("port", value)
    if err != nil {
        panic(err)
    }
}
// SetPorts sets the ports property value. List of ports supported for the application segment.
func (m *IpApplicationSegment) SetPorts(value []string)() {
    err := m.GetBackingStore().Set("ports", value)
    if err != nil {
        panic(err)
    }
}
// SetProtocol sets the protocol property value. Indicates the protocol of the network traffic acquired for the application segment. The possible values are: tcp, udp, unknownFutureValue.
func (m *IpApplicationSegment) SetProtocol(value *PrivateNetworkProtocol)() {
    err := m.GetBackingStore().Set("protocol", value)
    if err != nil {
        panic(err)
    }
}
type IpApplicationSegmentable interface {
    ApplicationSegmentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(Applicationable)
    GetDestinationHost()(*string)
    GetDestinationType()(*PrivateNetworkDestinationType)
    GetPort()(*int32)
    GetPorts()([]string)
    GetProtocol()(*PrivateNetworkProtocol)
    SetApplication(value Applicationable)()
    SetDestinationHost(value *string)()
    SetDestinationType(value *PrivateNetworkDestinationType)()
    SetPort(value *int32)()
    SetPorts(value []string)()
    SetProtocol(value *PrivateNetworkProtocol)()
}
