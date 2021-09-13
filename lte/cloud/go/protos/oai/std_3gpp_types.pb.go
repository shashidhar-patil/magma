// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/oai/std_3gpp_types.proto

package oai

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// parameters_list
type ParametersList struct {
	Parameters           []*Parameter `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty"`
	NumParameters        uint32       `protobuf:"varint,2,opt,name=num_parameters,json=numParameters,proto3" json:"num_parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ParametersList) Reset()         { *m = ParametersList{} }
func (m *ParametersList) String() string { return proto.CompactTextString(m) }
func (*ParametersList) ProtoMessage()    {}
func (*ParametersList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{0}
}

func (m *ParametersList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParametersList.Unmarshal(m, b)
}
func (m *ParametersList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParametersList.Marshal(b, m, deterministic)
}
func (m *ParametersList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParametersList.Merge(m, src)
}
func (m *ParametersList) XXX_Size() int {
	return xxx_messageInfo_ParametersList.Size(m)
}
func (m *ParametersList) XXX_DiscardUnknown() {
	xxx_messageInfo_ParametersList.DiscardUnknown(m)
}

var xxx_messageInfo_ParametersList proto.InternalMessageInfo

func (m *ParametersList) GetParameters() []*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ParametersList) GetNumParameters() uint32 {
	if m != nil {
		return m.NumParameters
	}
	return 0
}

// parameter_t
type Parameter struct {
	ParameterIdentifier  uint32   `protobuf:"varint,1,opt,name=parameter_identifier,json=parameterIdentifier,proto3" json:"parameter_identifier,omitempty"`
	Length               uint32   `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Contents             []byte   `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{1}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetParameterIdentifier() uint32 {
	if m != nil {
		return m.ParameterIdentifier
	}
	return 0
}

func (m *Parameter) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Parameter) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

// protocol_configuration_options_t
type Pco struct {
	Ext                      uint32         `protobuf:"varint,1,opt,name=ext,proto3" json:"ext,omitempty"`
	Spare                    uint32         `protobuf:"varint,2,opt,name=spare,proto3" json:"spare,omitempty"`
	ConfigurationProtocol    uint32         `protobuf:"varint,3,opt,name=configuration_protocol,json=configurationProtocol,proto3" json:"configuration_protocol,omitempty"`
	NumProtocolOrContainerId uint32         `protobuf:"varint,4,opt,name=num_protocol_or_container_id,json=numProtocolOrContainerId,proto3" json:"num_protocol_or_container_id,omitempty"`
	PcoProtocol              []*PcoProtocol `protobuf:"bytes,5,rep,name=pco_protocol,json=pcoProtocol,proto3" json:"pco_protocol,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}       `json:"-"`
	XXX_unrecognized         []byte         `json:"-"`
	XXX_sizecache            int32          `json:"-"`
}

func (m *Pco) Reset()         { *m = Pco{} }
func (m *Pco) String() string { return proto.CompactTextString(m) }
func (*Pco) ProtoMessage()    {}
func (*Pco) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{2}
}

func (m *Pco) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pco.Unmarshal(m, b)
}
func (m *Pco) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pco.Marshal(b, m, deterministic)
}
func (m *Pco) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pco.Merge(m, src)
}
func (m *Pco) XXX_Size() int {
	return xxx_messageInfo_Pco.Size(m)
}
func (m *Pco) XXX_DiscardUnknown() {
	xxx_messageInfo_Pco.DiscardUnknown(m)
}

var xxx_messageInfo_Pco proto.InternalMessageInfo

func (m *Pco) GetExt() uint32 {
	if m != nil {
		return m.Ext
	}
	return 0
}

func (m *Pco) GetSpare() uint32 {
	if m != nil {
		return m.Spare
	}
	return 0
}

func (m *Pco) GetConfigurationProtocol() uint32 {
	if m != nil {
		return m.ConfigurationProtocol
	}
	return 0
}

func (m *Pco) GetNumProtocolOrContainerId() uint32 {
	if m != nil {
		return m.NumProtocolOrContainerId
	}
	return 0
}

func (m *Pco) GetPcoProtocol() []*PcoProtocol {
	if m != nil {
		return m.PcoProtocol
	}
	return nil
}

// pco_protocol_or_container_id
type PcoProtocol struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Length               uint32   `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Contents             []byte   `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcoProtocol) Reset()         { *m = PcoProtocol{} }
func (m *PcoProtocol) String() string { return proto.CompactTextString(m) }
func (*PcoProtocol) ProtoMessage()    {}
func (*PcoProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{3}
}

func (m *PcoProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcoProtocol.Unmarshal(m, b)
}
func (m *PcoProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcoProtocol.Marshal(b, m, deterministic)
}
func (m *PcoProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcoProtocol.Merge(m, src)
}
func (m *PcoProtocol) XXX_Size() int {
	return xxx_messageInfo_PcoProtocol.Size(m)
}
func (m *PcoProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_PcoProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_PcoProtocol proto.InternalMessageInfo

func (m *PcoProtocol) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PcoProtocol) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *PcoProtocol) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

// fteid_t
type Fteid struct {
	Ipv4Address          uint32   `protobuf:"varint,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          []byte   `protobuf:"bytes,2,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	InterfaceType        uint32   `protobuf:"varint,3,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	Teid                 uint32   `protobuf:"varint,4,opt,name=teid,proto3" json:"teid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fteid) Reset()         { *m = Fteid{} }
func (m *Fteid) String() string { return proto.CompactTextString(m) }
func (*Fteid) ProtoMessage()    {}
func (*Fteid) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{4}
}

func (m *Fteid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fteid.Unmarshal(m, b)
}
func (m *Fteid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fteid.Marshal(b, m, deterministic)
}
func (m *Fteid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fteid.Merge(m, src)
}
func (m *Fteid) XXX_Size() int {
	return xxx_messageInfo_Fteid.Size(m)
}
func (m *Fteid) XXX_DiscardUnknown() {
	xxx_messageInfo_Fteid.DiscardUnknown(m)
}

var xxx_messageInfo_Fteid proto.InternalMessageInfo

func (m *Fteid) GetIpv4Address() uint32 {
	if m != nil {
		return m.Ipv4Address
	}
	return 0
}

func (m *Fteid) GetIpv6Address() []byte {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

func (m *Fteid) GetInterfaceType() uint32 {
	if m != nil {
		return m.InterfaceType
	}
	return 0
}

func (m *Fteid) GetTeid() uint32 {
	if m != nil {
		return m.Teid
	}
	return 0
}

// port_range
type PortRange struct {
	LowLimit             uint32   `protobuf:"varint,1,opt,name=low_limit,json=lowLimit,proto3" json:"low_limit,omitempty"`
	HighLimit            uint32   `protobuf:"varint,2,opt,name=high_limit,json=highLimit,proto3" json:"high_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortRange) Reset()         { *m = PortRange{} }
func (m *PortRange) String() string { return proto.CompactTextString(m) }
func (*PortRange) ProtoMessage()    {}
func (*PortRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{5}
}

func (m *PortRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortRange.Unmarshal(m, b)
}
func (m *PortRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortRange.Marshal(b, m, deterministic)
}
func (m *PortRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortRange.Merge(m, src)
}
func (m *PortRange) XXX_Size() int {
	return xxx_messageInfo_PortRange.Size(m)
}
func (m *PortRange) XXX_DiscardUnknown() {
	xxx_messageInfo_PortRange.DiscardUnknown(m)
}

var xxx_messageInfo_PortRange proto.InternalMessageInfo

func (m *PortRange) GetLowLimit() uint32 {
	if m != nil {
		return m.LowLimit
	}
	return 0
}

func (m *PortRange) GetHighLimit() uint32 {
	if m != nil {
		return m.HighLimit
	}
	return 0
}

// typdeofservice_trafficclass
type TypeOfServiceTrafficClass struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Mask                 uint32   `protobuf:"varint,2,opt,name=mask,proto3" json:"mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeOfServiceTrafficClass) Reset()         { *m = TypeOfServiceTrafficClass{} }
func (m *TypeOfServiceTrafficClass) String() string { return proto.CompactTextString(m) }
func (*TypeOfServiceTrafficClass) ProtoMessage()    {}
func (*TypeOfServiceTrafficClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{6}
}

func (m *TypeOfServiceTrafficClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeOfServiceTrafficClass.Unmarshal(m, b)
}
func (m *TypeOfServiceTrafficClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeOfServiceTrafficClass.Marshal(b, m, deterministic)
}
func (m *TypeOfServiceTrafficClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeOfServiceTrafficClass.Merge(m, src)
}
func (m *TypeOfServiceTrafficClass) XXX_Size() int {
	return xxx_messageInfo_TypeOfServiceTrafficClass.Size(m)
}
func (m *TypeOfServiceTrafficClass) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeOfServiceTrafficClass.DiscardUnknown(m)
}

var xxx_messageInfo_TypeOfServiceTrafficClass proto.InternalMessageInfo

func (m *TypeOfServiceTrafficClass) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TypeOfServiceTrafficClass) GetMask() uint32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

// ipvremoteaddr
type IpRemoteAddress struct {
	Addr                 uint32   `protobuf:"varint,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Mask                 uint32   `protobuf:"varint,2,opt,name=mask,proto3" json:"mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpRemoteAddress) Reset()         { *m = IpRemoteAddress{} }
func (m *IpRemoteAddress) String() string { return proto.CompactTextString(m) }
func (*IpRemoteAddress) ProtoMessage()    {}
func (*IpRemoteAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{7}
}

func (m *IpRemoteAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpRemoteAddress.Unmarshal(m, b)
}
func (m *IpRemoteAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpRemoteAddress.Marshal(b, m, deterministic)
}
func (m *IpRemoteAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpRemoteAddress.Merge(m, src)
}
func (m *IpRemoteAddress) XXX_Size() int {
	return xxx_messageInfo_IpRemoteAddress.Size(m)
}
func (m *IpRemoteAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_IpRemoteAddress.DiscardUnknown(m)
}

var xxx_messageInfo_IpRemoteAddress proto.InternalMessageInfo

func (m *IpRemoteAddress) GetAddr() uint32 {
	if m != nil {
		return m.Addr
	}
	return 0
}

func (m *IpRemoteAddress) GetMask() uint32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

// packet_filter_contents
type PacketFilterContents struct {
	Flags                        uint32                     `protobuf:"varint,1,opt,name=flags,proto3" json:"flags,omitempty"`
	Ipv4RemoteAddresses          []*IpRemoteAddress         `protobuf:"bytes,2,rep,name=ipv4_remote_addresses,json=ipv4RemoteAddresses,proto3" json:"ipv4_remote_addresses,omitempty"`
	Ipv6RemoteAddresses          []*IpRemoteAddress         `protobuf:"bytes,3,rep,name=ipv6_remote_addresses,json=ipv6RemoteAddresses,proto3" json:"ipv6_remote_addresses,omitempty"`
	ProtocolIdentifierNextheader uint32                     `protobuf:"varint,4,opt,name=protocol_identifier_nextheader,json=protocolIdentifierNextheader,proto3" json:"protocol_identifier_nextheader,omitempty"`
	SingleLocalPort              uint32                     `protobuf:"varint,10,opt,name=single_local_port,json=singleLocalPort,proto3" json:"single_local_port,omitempty"`
	LocalPortRange               *PortRange                 `protobuf:"bytes,11,opt,name=local_port_range,json=localPortRange,proto3" json:"local_port_range,omitempty"`
	SingleRemotePort             uint32                     `protobuf:"varint,12,opt,name=single_remote_port,json=singleRemotePort,proto3" json:"single_remote_port,omitempty"`
	RemotePortRange              *PortRange                 `protobuf:"bytes,13,opt,name=remote_port_range,json=remotePortRange,proto3" json:"remote_port_range,omitempty"`
	SecurityParameterIndex       uint32                     `protobuf:"varint,20,opt,name=security_parameter_index,json=securityParameterIndex,proto3" json:"security_parameter_index,omitempty"`
	TypeOfServiceTrafficClass    *TypeOfServiceTrafficClass `protobuf:"bytes,21,opt,name=type_of_service_traffic_class,json=typeOfServiceTrafficClass,proto3" json:"type_of_service_traffic_class,omitempty"`
	FlowLabel                    uint32                     `protobuf:"varint,22,opt,name=flow_label,json=flowLabel,proto3" json:"flow_label,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                   `json:"-"`
	XXX_unrecognized             []byte                     `json:"-"`
	XXX_sizecache                int32                      `json:"-"`
}

func (m *PacketFilterContents) Reset()         { *m = PacketFilterContents{} }
func (m *PacketFilterContents) String() string { return proto.CompactTextString(m) }
func (*PacketFilterContents) ProtoMessage()    {}
func (*PacketFilterContents) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e2a41045ec6d643, []int{8}
}

func (m *PacketFilterContents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketFilterContents.Unmarshal(m, b)
}
func (m *PacketFilterContents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketFilterContents.Marshal(b, m, deterministic)
}
func (m *PacketFilterContents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketFilterContents.Merge(m, src)
}
func (m *PacketFilterContents) XXX_Size() int {
	return xxx_messageInfo_PacketFilterContents.Size(m)
}
func (m *PacketFilterContents) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketFilterContents.DiscardUnknown(m)
}

var xxx_messageInfo_PacketFilterContents proto.InternalMessageInfo

func (m *PacketFilterContents) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *PacketFilterContents) GetIpv4RemoteAddresses() []*IpRemoteAddress {
	if m != nil {
		return m.Ipv4RemoteAddresses
	}
	return nil
}

func (m *PacketFilterContents) GetIpv6RemoteAddresses() []*IpRemoteAddress {
	if m != nil {
		return m.Ipv6RemoteAddresses
	}
	return nil
}

func (m *PacketFilterContents) GetProtocolIdentifierNextheader() uint32 {
	if m != nil {
		return m.ProtocolIdentifierNextheader
	}
	return 0
}

func (m *PacketFilterContents) GetSingleLocalPort() uint32 {
	if m != nil {
		return m.SingleLocalPort
	}
	return 0
}

func (m *PacketFilterContents) GetLocalPortRange() *PortRange {
	if m != nil {
		return m.LocalPortRange
	}
	return nil
}

func (m *PacketFilterContents) GetSingleRemotePort() uint32 {
	if m != nil {
		return m.SingleRemotePort
	}
	return 0
}

func (m *PacketFilterContents) GetRemotePortRange() *PortRange {
	if m != nil {
		return m.RemotePortRange
	}
	return nil
}

func (m *PacketFilterContents) GetSecurityParameterIndex() uint32 {
	if m != nil {
		return m.SecurityParameterIndex
	}
	return 0
}

func (m *PacketFilterContents) GetTypeOfServiceTrafficClass() *TypeOfServiceTrafficClass {
	if m != nil {
		return m.TypeOfServiceTrafficClass
	}
	return nil
}

func (m *PacketFilterContents) GetFlowLabel() uint32 {
	if m != nil {
		return m.FlowLabel
	}
	return 0
}

func init() {
	proto.RegisterType((*ParametersList)(nil), "magma.lte.oai.ParametersList")
	proto.RegisterType((*Parameter)(nil), "magma.lte.oai.Parameter")
	proto.RegisterType((*Pco)(nil), "magma.lte.oai.Pco")
	proto.RegisterType((*PcoProtocol)(nil), "magma.lte.oai.PcoProtocol")
	proto.RegisterType((*Fteid)(nil), "magma.lte.oai.Fteid")
	proto.RegisterType((*PortRange)(nil), "magma.lte.oai.PortRange")
	proto.RegisterType((*TypeOfServiceTrafficClass)(nil), "magma.lte.oai.TypeOfServiceTrafficClass")
	proto.RegisterType((*IpRemoteAddress)(nil), "magma.lte.oai.IpRemoteAddress")
	proto.RegisterType((*PacketFilterContents)(nil), "magma.lte.oai.PacketFilterContents")
}

func init() {
	proto.RegisterFile("lte/protos/oai/std_3gpp_types.proto", fileDescriptor_7e2a41045ec6d643)
}

var fileDescriptor_7e2a41045ec6d643 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0x23, 0x35,
	0x18, 0x55, 0x92, 0x66, 0xb5, 0xf9, 0xf2, 0xd3, 0xae, 0x37, 0x8d, 0x66, 0x97, 0xcd, 0x12, 0x82,
	0x56, 0x8a, 0x10, 0x4a, 0x04, 0x0b, 0xd1, 0x72, 0x01, 0x12, 0x6d, 0x29, 0x8a, 0x14, 0xd1, 0x30,
	0xf4, 0x8a, 0x1b, 0xcb, 0x9d, 0xf9, 0x66, 0x62, 0xea, 0x19, 0x0f, 0x1e, 0xa7, 0x3f, 0x6f, 0xc0,
	0x7b, 0xf1, 0x28, 0xbc, 0x08, 0xb2, 0xc7, 0x33, 0x49, 0x43, 0x2b, 0x10, 0x77, 0xfe, 0xfe, 0xce,
	0xb1, 0xcf, 0x1c, 0x7b, 0xe0, 0x53, 0xa1, 0x71, 0x96, 0x29, 0xa9, 0x65, 0x3e, 0x93, 0x8c, 0xcf,
	0x72, 0x1d, 0xd2, 0xf7, 0x71, 0x96, 0x51, 0x7d, 0x9f, 0x61, 0x3e, 0xb5, 0x15, 0xd2, 0x4d, 0x58,
	0x9c, 0xb0, 0xa9, 0xd0, 0x38, 0x95, 0x8c, 0x8f, 0x7f, 0x87, 0xde, 0x8a, 0x29, 0x96, 0xa0, 0x46,
	0x95, 0x2f, 0x79, 0xae, 0xc9, 0x07, 0x80, 0xac, 0xca, 0x78, 0xb5, 0x51, 0x63, 0xd2, 0xfe, 0xd2,
	0x9b, 0x3e, 0x98, 0x9a, 0x56, 0x23, 0xfe, 0x4e, 0x2f, 0x79, 0x07, 0xbd, 0x74, 0x93, 0xd0, 0x9d,
	0xe9, 0xfa, 0xa8, 0x36, 0xe9, 0xfa, 0xdd, 0x74, 0x93, 0x6c, 0x49, 0xc6, 0x0a, 0x5a, 0x55, 0x44,
	0xbe, 0x80, 0x7e, 0xd5, 0x4f, 0x79, 0x88, 0xa9, 0xe6, 0x11, 0x47, 0xe5, 0xd5, 0xec, 0xe4, 0xcb,
	0xaa, 0xb6, 0xa8, 0x4a, 0x64, 0x00, 0xcf, 0x04, 0xa6, 0xb1, 0x5e, 0x3b, 0x78, 0x17, 0x91, 0xd7,
	0xf0, 0x3c, 0x90, 0xa9, 0xc6, 0x54, 0xe7, 0x5e, 0x63, 0x54, 0x9b, 0x74, 0xfc, 0x2a, 0x1e, 0xff,
	0x55, 0x83, 0xc6, 0x2a, 0x90, 0xe4, 0x08, 0x1a, 0x78, 0xa7, 0x1d, 0xba, 0x59, 0x92, 0x3e, 0x34,
	0xf3, 0x8c, 0x29, 0x74, 0x60, 0x45, 0x40, 0xbe, 0x86, 0x41, 0x20, 0xd3, 0x88, 0xc7, 0x1b, 0xc5,
	0x34, 0x97, 0x29, 0xb5, 0xe2, 0x05, 0x52, 0x58, 0xe4, 0xae, 0x7f, 0xfc, 0xa0, 0xba, 0x72, 0x45,
	0xf2, 0x1d, 0xbc, 0xb1, 0x0a, 0xb8, 0x98, 0x4a, 0x45, 0xcd, 0x16, 0x18, 0x4f, 0xed, 0xe9, 0xbc,
	0x03, 0x3b, 0xec, 0x19, 0x3d, 0x5c, 0xcb, 0x85, 0x3a, 0x2d, 0x1b, 0x16, 0x21, 0xf9, 0x16, 0x3a,
	0x59, 0x20, 0xb7, 0x64, 0x4d, 0xab, 0xfe, 0xeb, 0x7d, 0xf5, 0x03, 0x59, 0x8e, 0xfb, 0xed, 0x6c,
	0x1b, 0x8c, 0x7f, 0x86, 0xf6, 0x4e, 0x8d, 0xf4, 0xa0, 0xce, 0x43, 0x77, 0xd6, 0x3a, 0x0f, 0xff,
	0x97, 0x70, 0x7f, 0xd4, 0xa0, 0x79, 0xae, 0x91, 0x87, 0xe4, 0x13, 0xe8, 0xf0, 0xec, 0xe6, 0x2b,
	0xca, 0xc2, 0x50, 0x61, 0x9e, 0x3b, 0xdc, 0xb6, 0xc9, 0x7d, 0x5f, 0xa4, 0x5c, 0xcb, 0xbc, 0x6a,
	0xa9, 0x5b, 0x30, 0xd3, 0x32, 0x2f, 0x5b, 0xde, 0x41, 0x8f, 0xa7, 0x1a, 0x55, 0xc4, 0x02, 0xb4,
	0xbe, 0x74, 0x82, 0x76, 0xab, 0xec, 0xe5, 0x7d, 0x86, 0x84, 0xc0, 0x81, 0x21, 0x75, 0x82, 0xd9,
	0xf5, 0xf8, 0x47, 0x68, 0xad, 0xa4, 0xd2, 0x3e, 0x4b, 0x63, 0x24, 0x1f, 0x41, 0x4b, 0xc8, 0x5b,
	0x2a, 0x78, 0xc2, 0xcb, 0xcf, 0xf9, 0x5c, 0xc8, 0xdb, 0xa5, 0x89, 0xc9, 0x10, 0x60, 0xcd, 0xe3,
	0xb5, 0xab, 0x16, 0x87, 0x6d, 0x99, 0x8c, 0x2d, 0x8f, 0x7f, 0x80, 0x57, 0x86, 0xe4, 0x22, 0xfa,
	0x05, 0xd5, 0x0d, 0x0f, 0xf0, 0x52, 0xb1, 0x28, 0xe2, 0xc1, 0xa9, 0x60, 0x79, 0x6e, 0xfc, 0x70,
	0xc3, 0xc4, 0x06, 0x1d, 0x68, 0x11, 0x98, 0xfd, 0x24, 0x2c, 0xbf, 0x76, 0x58, 0x76, 0x3d, 0xfe,
	0x06, 0x0e, 0x17, 0x99, 0x8f, 0x89, 0xd4, 0x58, 0x9e, 0x8e, 0xc0, 0x81, 0x39, 0xbb, 0x9b, 0xb5,
	0xeb, 0x47, 0x47, 0xff, 0x6c, 0x42, 0x7f, 0xc5, 0x82, 0x6b, 0xd4, 0xe7, 0x5c, 0x68, 0xb4, 0x1e,
	0x30, 0x72, 0x1b, 0xf6, 0x48, 0xb0, 0xb8, 0x54, 0xb7, 0x08, 0x88, 0x0f, 0xc7, 0x56, 0x7a, 0x65,
	0xc9, 0x4a, 0x79, 0xd1, 0x08, 0x6c, 0xfc, 0xf1, 0x76, 0xcf, 0x1f, 0x7b, 0xbb, 0xf2, 0x5f, 0x9a,
	0xe1, 0x07, 0x29, 0x2c, 0x31, 0xe7, 0xff, 0xc4, 0x6c, 0xfc, 0x67, 0xcc, 0xf9, 0x3e, 0xe6, 0x19,
	0xbc, 0xad, 0xac, 0xbf, 0xbd, 0xcb, 0x34, 0xc5, 0x3b, 0xbd, 0x46, 0x16, 0xa2, 0x72, 0xdf, 0xf3,
	0x4d, 0xd9, 0xb5, 0xbd, 0xd5, 0x3f, 0x55, 0x3d, 0xe4, 0x33, 0x78, 0x91, 0xf3, 0x34, 0x16, 0x48,
	0x85, 0x0c, 0x98, 0xa0, 0x99, 0x54, 0xda, 0x03, 0x3b, 0x78, 0x58, 0x14, 0x96, 0x26, 0x6f, 0xbc,
	0x40, 0x4e, 0xe0, 0x68, 0xdb, 0x44, 0x95, 0xb1, 0x86, 0xd7, 0x1e, 0xd5, 0x1e, 0x7b, 0xb2, 0x4a,
	0xeb, 0xf8, 0x3d, 0x51, 0x8e, 0x17, 0x56, 0xfa, 0x1c, 0x88, 0xe3, 0x73, 0x5a, 0x58, 0xc2, 0x8e,
	0x25, 0x3c, 0x2a, 0x2a, 0xc5, 0x41, 0x2d, 0xe3, 0x19, 0xbc, 0xd8, 0x69, 0x73, 0x94, 0xdd, 0x7f,
	0xa1, 0x3c, 0x54, 0x15, 0x40, 0xc1, 0xf9, 0x01, 0xbc, 0x1c, 0x83, 0x8d, 0xe2, 0xfa, 0x9e, 0xee,
	0xbc, 0x7f, 0x69, 0x88, 0x77, 0x5e, 0xdf, 0x32, 0x0f, 0xca, 0x7a, 0xf5, 0x56, 0x2e, 0x4c, 0x95,
	0xfc, 0x06, 0x43, 0x73, 0x6d, 0xa8, 0x8c, 0x68, 0x5e, 0xd8, 0x97, 0xea, 0xc2, 0xbf, 0x34, 0x30,
	0x06, 0xf6, 0x8e, 0xed, 0x5e, 0x26, 0x7b, 0x7b, 0x79, 0xd2, 0xf0, 0xfe, 0x2b, 0xfd, 0xe4, 0x5d,
	0x18, 0x02, 0x44, 0xf6, 0x96, 0xb1, 0x2b, 0x14, 0xde, 0xa0, 0xb8, 0x47, 0x26, 0xb3, 0x34, 0x89,
	0x93, 0x8f, 0x7f, 0x1d, 0x5a, 0x92, 0x99, 0xf9, 0xef, 0x04, 0x42, 0x6e, 0xc2, 0x59, 0x2c, 0x77,
	0x7e, 0x40, 0x57, 0xcf, 0xec, 0xfa, 0xfd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x22, 0xab, 0x1d,
	0x07, 0x99, 0x06, 0x00, 0x00,
}