// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/contractmanager/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Failure message contains information about ACK failures and can be used to
// replay ACK in case of requirement.
// Note that Failure means that sudo handler to cosmwasm contract failed for some reason
type Failure struct {
	// Address of the failed contract
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Id of the failure under specific address
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Acknowledgement type
	AckType string `protobuf:"bytes,3,opt,name=ack_type,json=ackType,proto3" json:"ack_type,omitempty"`
	// IBC Packet
	Packet *types.Packet `protobuf:"bytes,4,opt,name=packet,proto3" json:"packet,omitempty"`
	// Acknowledgement
	Ack *types.Acknowledgement `protobuf:"bytes,5,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf4a1534315a7490, []int{0}
}
func (m *Failure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(m, src)
}
func (m *Failure) XXX_Size() int {
	return m.Size()
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Failure) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Failure) GetAckType() string {
	if m != nil {
		return m.AckType
	}
	return ""
}

func (m *Failure) GetPacket() *types.Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *Failure) GetAck() *types.Acknowledgement {
	if m != nil {
		return m.Ack
	}
	return nil
}

// GenesisState defines the contractmanager module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// List of the contract failures
	FailuresList []Failure `protobuf:"bytes,2,rep,name=failures_list,json=failuresList,proto3" json:"failures_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf4a1534315a7490, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFailuresList() []Failure {
	if m != nil {
		return m.FailuresList
	}
	return nil
}

// Deprecated. Used only for migration purposes.
type OldFailure struct {
	// ChannelId
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Address of the failed contract
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// id of the failure under specific address
	Id uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// ACK id to restore
	AckId uint64 `protobuf:"varint,4,opt,name=ack_id,json=ackId,proto3" json:"ack_id,omitempty"`
	// Acknowledgement type
	AckType string `protobuf:"bytes,5,opt,name=ack_type,json=ackType,proto3" json:"ack_type,omitempty"`
}

func (m *OldFailure) Reset()         { *m = OldFailure{} }
func (m *OldFailure) String() string { return proto.CompactTextString(m) }
func (*OldFailure) ProtoMessage()    {}
func (*OldFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf4a1534315a7490, []int{2}
}
func (m *OldFailure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OldFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OldFailure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OldFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldFailure.Merge(m, src)
}
func (m *OldFailure) XXX_Size() int {
	return m.Size()
}
func (m *OldFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_OldFailure.DiscardUnknown(m)
}

var xxx_messageInfo_OldFailure proto.InternalMessageInfo

func (m *OldFailure) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *OldFailure) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *OldFailure) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OldFailure) GetAckId() uint64 {
	if m != nil {
		return m.AckId
	}
	return 0
}

func (m *OldFailure) GetAckType() string {
	if m != nil {
		return m.AckType
	}
	return ""
}

func init() {
	proto.RegisterType((*Failure)(nil), "neutron.contractmanager.Failure")
	proto.RegisterType((*GenesisState)(nil), "neutron.contractmanager.GenesisState")
	proto.RegisterType((*OldFailure)(nil), "neutron.contractmanager.OldFailure")
}

func init() {
	proto.RegisterFile("neutron/contractmanager/genesis.proto", fileDescriptor_cf4a1534315a7490)
}

var fileDescriptor_cf4a1534315a7490 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0xeb, 0xf4, 0xcf, 0x98, 0x3b, 0x38, 0x58, 0x20, 0x42, 0x11, 0x59, 0xa8, 0x86, 0xd4,
	0x0b, 0xb6, 0xd6, 0x49, 0xbb, 0x71, 0x60, 0x07, 0xd0, 0x04, 0x12, 0x55, 0xe0, 0xc4, 0xa5, 0x72,
	0x6c, 0x93, 0x59, 0x49, 0xec, 0xc8, 0x76, 0x07, 0x7b, 0x02, 0xae, 0x9c, 0x79, 0x15, 0x5e, 0x60,
	0xc7, 0x1d, 0x39, 0x21, 0xd4, 0xbe, 0x08, 0x4a, 0xe2, 0x20, 0xc6, 0xd6, 0xdb, 0xaf, 0xd5, 0xe7,
	0xfb, 0x75, 0x7e, 0x1f, 0x1b, 0x3e, 0x53, 0x62, 0xe5, 0x8c, 0x56, 0x84, 0x69, 0xe5, 0x0c, 0x65,
	0xae, 0xa4, 0x8a, 0x66, 0xc2, 0x90, 0x4c, 0x28, 0x61, 0xa5, 0xc5, 0x95, 0xd1, 0x4e, 0xa3, 0x87,
	0x1e, 0xc3, 0xff, 0x61, 0x93, 0xfb, 0x99, 0xce, 0x74, 0xc3, 0x90, 0x7a, 0x6a, 0xf1, 0xc9, 0xc1,
	0xb6, 0xd6, 0x8a, 0x1a, 0x5a, 0xfa, 0xd2, 0xc9, 0x53, 0x99, 0x32, 0xc2, 0xb4, 0x11, 0x84, 0x9d,
	0x51, 0xa5, 0x44, 0x41, 0xce, 0x0f, 0xbb, 0xb1, 0x45, 0xa6, 0x3f, 0x00, 0xdc, 0x79, 0x45, 0x65,
	0xb1, 0x32, 0x02, 0x85, 0x70, 0x87, 0x72, 0x6e, 0x84, 0xb5, 0x21, 0x88, 0xc1, 0x6c, 0x37, 0xe9,
	0x7e, 0xa2, 0x7b, 0x30, 0x90, 0x3c, 0x0c, 0x62, 0x30, 0x1b, 0x24, 0x81, 0xe4, 0xe8, 0x11, 0xbc,
	0x43, 0x59, 0xbe, 0x74, 0x17, 0x95, 0x08, 0xfb, 0x1e, 0x65, 0xf9, 0x87, 0x8b, 0x4a, 0xa0, 0x23,
	0x38, 0xaa, 0x28, 0xcb, 0x85, 0x0b, 0x07, 0x31, 0x98, 0x8d, 0xe7, 0x8f, 0xb1, 0x4c, 0x19, 0xae,
	0x3f, 0x02, 0x77, 0x27, 0x9f, 0x1f, 0xe2, 0x45, 0x83, 0x24, 0x1e, 0x45, 0xc7, 0xb0, 0x4f, 0x59,
	0x1e, 0x0e, 0x9b, 0xc4, 0xc1, 0xad, 0x89, 0x97, 0x2c, 0x57, 0xfa, 0x73, 0x21, 0x78, 0x26, 0x4a,
	0xa1, 0x5c, 0x52, 0x07, 0xa6, 0xdf, 0x01, 0xdc, 0x7b, 0xdd, 0x7a, 0x7c, 0xef, 0xa8, 0x13, 0xe8,
	0x45, 0x7d, 0x7a, 0x6d, 0xa0, 0xd9, 0x60, 0x3c, 0xdf, 0xc7, 0x5b, 0xbc, 0xe2, 0x45, 0x83, 0x9d,
	0x0c, 0x2e, 0x7f, 0xed, 0xf7, 0x12, 0x1f, 0x42, 0x6f, 0xe0, 0xdd, 0x4f, 0xad, 0x0c, 0xbb, 0x2c,
	0xa4, 0x75, 0x61, 0x10, 0xf7, 0x67, 0xe3, 0x79, 0xbc, 0xb5, 0xc5, 0xab, 0xf3, 0x35, 0x7b, 0x5d,
	0xf8, 0xad, 0xb4, 0x6e, 0xfa, 0x15, 0x40, 0xf8, 0xae, 0xe0, 0x9d, 0xdd, 0x27, 0x10, 0xfa, 0x75,
	0x96, 0x92, 0x7b, 0xc1, 0xbb, 0xfe, 0x9f, 0x53, 0xfe, 0xaf, 0xfc, 0xe0, 0x36, 0xf9, 0xfd, 0xbf,
	0xf2, 0x1f, 0xc0, 0x51, 0x2d, 0x5f, 0xf2, 0xc6, 0xf0, 0x20, 0x19, 0x52, 0x96, 0x9f, 0x5e, 0xbf,
	0x93, 0xe1, 0xb5, 0x3b, 0x39, 0x59, 0x5c, 0xae, 0x23, 0x70, 0xb5, 0x8e, 0xc0, 0xef, 0x75, 0x04,
	0xbe, 0x6d, 0xa2, 0xde, 0xd5, 0x26, 0xea, 0xfd, 0xdc, 0x44, 0xbd, 0x8f, 0xc7, 0x99, 0x74, 0x67,
	0xab, 0x14, 0x33, 0x5d, 0x12, 0xbf, 0xe3, 0x73, 0x6d, 0xb2, 0x6e, 0x26, 0x5f, 0x6e, 0x3c, 0xb0,
	0xba, 0xdf, 0xa6, 0xa3, 0xe6, 0xf5, 0x1c, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x59, 0xd9, 0x03,
	0xb9, 0xde, 0x02, 0x00, 0x00,
}

func (m *Failure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Failure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Failure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ack != nil {
		{
			size, err := m.Ack.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Packet != nil {
		{
			size, err := m.Packet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.AckType) > 0 {
		i -= len(m.AckType)
		copy(dAtA[i:], m.AckType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AckType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FailuresList) > 0 {
		for iNdEx := len(m.FailuresList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FailuresList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *OldFailure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OldFailure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OldFailure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AckType) > 0 {
		i -= len(m.AckType)
		copy(dAtA[i:], m.AckType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AckType)))
		i--
		dAtA[i] = 0x2a
	}
	if m.AckId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AckId))
		i--
		dAtA[i] = 0x20
	}
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Failure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	l = len(m.AckType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Packet != nil {
		l = m.Packet.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Ack != nil {
		l = m.Ack.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FailuresList) > 0 {
		for _, e := range m.FailuresList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *OldFailure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	if m.AckId != 0 {
		n += 1 + sovGenesis(uint64(m.AckId))
	}
	l = len(m.AckType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Failure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Failure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Failure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Packet == nil {
				m.Packet = &types.Packet{}
			}
			if err := m.Packet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ack", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ack == nil {
				m.Ack = &types.Acknowledgement{}
			}
			if err := m.Ack.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailuresList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FailuresList = append(m.FailuresList, Failure{})
			if err := m.FailuresList[len(m.FailuresList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OldFailure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OldFailure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OldFailure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckId", wireType)
			}
			m.AckId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AckId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
