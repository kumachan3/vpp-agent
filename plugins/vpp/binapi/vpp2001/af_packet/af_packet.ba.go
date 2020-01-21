// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/af_packet.api.json

/*
Package af_packet is a generated VPP binary API for 'af_packet' module.

It consists of:
	  6 enums
	  2 aliases
	  8 messages
	  4 services
*/
package af_packet

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "af_packet"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xba745e20
)

// IfStatusFlags represents VPP binary API enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var IfStatusFlags_name = map[uint32]string{
	1: "IF_STATUS_API_FLAG_ADMIN_UP",
	2: "IF_STATUS_API_FLAG_LINK_UP",
}

var IfStatusFlags_value = map[string]uint32{
	"IF_STATUS_API_FLAG_ADMIN_UP": 1,
	"IF_STATUS_API_FLAG_LINK_UP":  2,
}

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IfType represents VPP binary API enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var IfType_name = map[uint32]string{
	1: "IF_API_TYPE_HARDWARE",
	2: "IF_API_TYPE_SUB",
	3: "IF_API_TYPE_P2P",
	4: "IF_API_TYPE_PIPE",
}

var IfType_value = map[string]uint32{
	"IF_API_TYPE_HARDWARE": 1,
	"IF_API_TYPE_SUB":      2,
	"IF_API_TYPE_P2P":      3,
	"IF_API_TYPE_PIPE":     4,
}

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LinkDuplex represents VPP binary API enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var LinkDuplex_name = map[uint32]string{
	0: "LINK_DUPLEX_API_UNKNOWN",
	1: "LINK_DUPLEX_API_HALF",
	2: "LINK_DUPLEX_API_FULL",
}

var LinkDuplex_value = map[string]uint32{
	"LINK_DUPLEX_API_UNKNOWN": 0,
	"LINK_DUPLEX_API_HALF":    1,
	"LINK_DUPLEX_API_FULL":    2,
}

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// MtuProto represents VPP binary API enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var MtuProto_name = map[uint32]string{
	1: "MTU_PROTO_API_L3",
	2: "MTU_PROTO_API_IP4",
	3: "MTU_PROTO_API_IP6",
	4: "MTU_PROTO_API_MPLS",
	5: "MTU_PROTO_API_N",
}

var MtuProto_value = map[string]uint32{
	"MTU_PROTO_API_L3":   1,
	"MTU_PROTO_API_IP4":  2,
	"MTU_PROTO_API_IP6":  3,
	"MTU_PROTO_API_MPLS": 4,
	"MTU_PROTO_API_N":    5,
}

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// RxMode represents VPP binary API enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var RxMode_name = map[uint32]string{
	0: "RX_MODE_API_UNKNOWN",
	1: "RX_MODE_API_POLLING",
	2: "RX_MODE_API_INTERRUPT",
	3: "RX_MODE_API_ADAPTIVE",
	4: "RX_MODE_API_DEFAULT",
}

var RxMode_value = map[string]uint32{
	"RX_MODE_API_UNKNOWN":   0,
	"RX_MODE_API_POLLING":   1,
	"RX_MODE_API_INTERRUPT": 2,
	"RX_MODE_API_ADAPTIVE":  3,
	"RX_MODE_API_DEFAULT":   4,
}

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SubIfFlags represents VPP binary API enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var SubIfFlags_name = map[uint32]string{
	1:   "SUB_IF_API_FLAG_NO_TAGS",
	2:   "SUB_IF_API_FLAG_ONE_TAG",
	4:   "SUB_IF_API_FLAG_TWO_TAGS",
	8:   "SUB_IF_API_FLAG_DOT1AD",
	16:  "SUB_IF_API_FLAG_EXACT_MATCH",
	32:  "SUB_IF_API_FLAG_DEFAULT",
	64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
	128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
	254: "SUB_IF_API_FLAG_MASK_VNET",
	256: "SUB_IF_API_FLAG_DOT1AH",
}

var SubIfFlags_value = map[string]uint32{
	"SUB_IF_API_FLAG_NO_TAGS":           1,
	"SUB_IF_API_FLAG_ONE_TAG":           2,
	"SUB_IF_API_FLAG_TWO_TAGS":          4,
	"SUB_IF_API_FLAG_DOT1AD":            8,
	"SUB_IF_API_FLAG_EXACT_MATCH":       16,
	"SUB_IF_API_FLAG_DEFAULT":           32,
	"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
	"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
	"SUB_IF_API_FLAG_MASK_VNET":         254,
	"SUB_IF_API_FLAG_DOT1AH":            256,
}

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// InterfaceIndex represents VPP binary API alias 'interface_index'.
type InterfaceIndex uint32

// MacAddress represents VPP binary API alias 'mac_address'.
type MacAddress [6]uint8

// AfPacketCreate represents VPP binary API message 'af_packet_create'.
type AfPacketCreate struct {
	HwAddr          MacAddress
	UseRandomHwAddr bool
	HostIfName      string `struc:"[64]byte"`
}

func (*AfPacketCreate) GetMessageName() string {
	return "af_packet_create"
}
func (*AfPacketCreate) GetCrcString() string {
	return "a190415f"
}
func (*AfPacketCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketCreateReply represents VPP binary API message 'af_packet_create_reply'.
type AfPacketCreateReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (*AfPacketCreateReply) GetMessageName() string {
	return "af_packet_create_reply"
}
func (*AfPacketCreateReply) GetCrcString() string {
	return "5383d31f"
}
func (*AfPacketCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketDelete represents VPP binary API message 'af_packet_delete'.
type AfPacketDelete struct {
	HostIfName string `struc:"[64]byte"`
}

func (*AfPacketDelete) GetMessageName() string {
	return "af_packet_delete"
}
func (*AfPacketDelete) GetCrcString() string {
	return "863fa648"
}
func (*AfPacketDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketDeleteReply represents VPP binary API message 'af_packet_delete_reply'.
type AfPacketDeleteReply struct {
	Retval int32
}

func (*AfPacketDeleteReply) GetMessageName() string {
	return "af_packet_delete_reply"
}
func (*AfPacketDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AfPacketDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketDetails represents VPP binary API message 'af_packet_details'.
type AfPacketDetails struct {
	SwIfIndex  InterfaceIndex
	HostIfName string `struc:"[64]byte"`
}

func (*AfPacketDetails) GetMessageName() string {
	return "af_packet_details"
}
func (*AfPacketDetails) GetCrcString() string {
	return "58c7c042"
}
func (*AfPacketDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketDump represents VPP binary API message 'af_packet_dump'.
type AfPacketDump struct{}

func (*AfPacketDump) GetMessageName() string {
	return "af_packet_dump"
}
func (*AfPacketDump) GetCrcString() string {
	return "51077d14"
}
func (*AfPacketDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketSetL4CksumOffload represents VPP binary API message 'af_packet_set_l4_cksum_offload'.
type AfPacketSetL4CksumOffload struct {
	SwIfIndex InterfaceIndex
	Set       bool
}

func (*AfPacketSetL4CksumOffload) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload"
}
func (*AfPacketSetL4CksumOffload) GetCrcString() string {
	return "319cd5c8"
}
func (*AfPacketSetL4CksumOffload) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketSetL4CksumOffloadReply represents VPP binary API message 'af_packet_set_l4_cksum_offload_reply'.
type AfPacketSetL4CksumOffloadReply struct {
	Retval int32
}

func (*AfPacketSetL4CksumOffloadReply) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload_reply"
}
func (*AfPacketSetL4CksumOffloadReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AfPacketSetL4CksumOffloadReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*AfPacketCreate)(nil), "af_packet.AfPacketCreate")
	api.RegisterMessage((*AfPacketCreateReply)(nil), "af_packet.AfPacketCreateReply")
	api.RegisterMessage((*AfPacketDelete)(nil), "af_packet.AfPacketDelete")
	api.RegisterMessage((*AfPacketDeleteReply)(nil), "af_packet.AfPacketDeleteReply")
	api.RegisterMessage((*AfPacketDetails)(nil), "af_packet.AfPacketDetails")
	api.RegisterMessage((*AfPacketDump)(nil), "af_packet.AfPacketDump")
	api.RegisterMessage((*AfPacketSetL4CksumOffload)(nil), "af_packet.AfPacketSetL4CksumOffload")
	api.RegisterMessage((*AfPacketSetL4CksumOffloadReply)(nil), "af_packet.AfPacketSetL4CksumOffloadReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AfPacketCreate)(nil),
		(*AfPacketCreateReply)(nil),
		(*AfPacketDelete)(nil),
		(*AfPacketDeleteReply)(nil),
		(*AfPacketDetails)(nil),
		(*AfPacketDump)(nil),
		(*AfPacketSetL4CksumOffload)(nil),
		(*AfPacketSetL4CksumOffloadReply)(nil),
	}
}

// RPCService represents RPC service API for af_packet module.
type RPCService interface {
	DumpAfPacket(ctx context.Context, in *AfPacketDump) (RPCService_DumpAfPacketClient, error)
	AfPacketCreate(ctx context.Context, in *AfPacketCreate) (*AfPacketCreateReply, error)
	AfPacketDelete(ctx context.Context, in *AfPacketDelete) (*AfPacketDeleteReply, error)
	AfPacketSetL4CksumOffload(ctx context.Context, in *AfPacketSetL4CksumOffload) (*AfPacketSetL4CksumOffloadReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpAfPacket(ctx context.Context, in *AfPacketDump) (RPCService_DumpAfPacketClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpAfPacketClient{stream}
	return x, nil
}

type RPCService_DumpAfPacketClient interface {
	Recv() (*AfPacketDetails, error)
}

type serviceClient_DumpAfPacketClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpAfPacketClient) Recv() (*AfPacketDetails, error) {
	m := new(AfPacketDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) AfPacketCreate(ctx context.Context, in *AfPacketCreate) (*AfPacketCreateReply, error) {
	out := new(AfPacketCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AfPacketDelete(ctx context.Context, in *AfPacketDelete) (*AfPacketDeleteReply, error) {
	out := new(AfPacketDeleteReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AfPacketSetL4CksumOffload(ctx context.Context, in *AfPacketSetL4CksumOffload) (*AfPacketSetL4CksumOffloadReply, error) {
	out := new(AfPacketSetL4CksumOffloadReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack