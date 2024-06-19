// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/tokenfactory/params.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type WhitelistedHook struct {
	CodeID       uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	DenomCreator string `protobuf:"bytes,2,opt,name=denom_creator,json=denomCreator,proto3" json:"denom_creator,omitempty"`
}

func (m *WhitelistedHook) Reset()         { *m = WhitelistedHook{} }
func (m *WhitelistedHook) String() string { return proto.CompactTextString(m) }
func (*WhitelistedHook) ProtoMessage()    {}
func (*WhitelistedHook) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c297db7c49d1cf, []int{0}
}
func (m *WhitelistedHook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhitelistedHook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhitelistedHook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhitelistedHook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistedHook.Merge(m, src)
}
func (m *WhitelistedHook) XXX_Size() int {
	return m.Size()
}
func (m *WhitelistedHook) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistedHook.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistedHook proto.InternalMessageInfo

func (m *WhitelistedHook) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

func (m *WhitelistedHook) GetDenomCreator() string {
	if m != nil {
		return m.DenomCreator
	}
	return ""
}

// Params defines the parameters for the tokenfactory module.
type Params struct {
	// DenomCreationFee defines the fee to be charged on the creation of a new
	// denom. The fee is drawn from the MsgCreateDenom's sender account, and
	// transferred to the community pool.
	DenomCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=denom_creation_fee,json=denomCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"denom_creation_fee" yaml:"denom_creation_fee"`
	// DenomCreationGasConsume defines the gas cost for creating a new denom.
	// This is intended as a spam deterrence mechanism.
	//
	// See: https://github.com/CosmWasm/token-factory/issues/11
	DenomCreationGasConsume uint64 `protobuf:"varint,2,opt,name=denom_creation_gas_consume,json=denomCreationGasConsume,proto3" json:"denom_creation_gas_consume,omitempty" yaml:"denom_creation_gas_consume"`
	// FeeCollectorAddress is the address where fees collected from denom creation
	// are sent to
	FeeCollectorAddress string `protobuf:"bytes,3,opt,name=fee_collector_address,json=feeCollectorAddress,proto3" json:"fee_collector_address,omitempty"`
	// HookWhitelist is the list of hooks which are allowed to be added and executed
	WhitelistedHooks []*WhitelistedHook `protobuf:"bytes,4,rep,name=whitelisted_hooks,json=whitelistedHooks,proto3" json:"whitelisted_hooks,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c297db7c49d1cf, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDenomCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DenomCreationFee
	}
	return nil
}

func (m *Params) GetDenomCreationGasConsume() uint64 {
	if m != nil {
		return m.DenomCreationGasConsume
	}
	return 0
}

func (m *Params) GetFeeCollectorAddress() string {
	if m != nil {
		return m.FeeCollectorAddress
	}
	return ""
}

func (m *Params) GetWhitelistedHooks() []*WhitelistedHook {
	if m != nil {
		return m.WhitelistedHooks
	}
	return nil
}

func init() {
	proto.RegisterType((*WhitelistedHook)(nil), "osmosis.tokenfactory.WhitelistedHook")
	proto.RegisterType((*Params)(nil), "osmosis.tokenfactory.Params")
}

func init() { proto.RegisterFile("osmosis/tokenfactory/params.proto", fileDescriptor_09c297db7c49d1cf) }

var fileDescriptor_09c297db7c49d1cf = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x49, 0x14, 0xc4, 0x02, 0xa2, 0x98, 0x22, 0x9c, 0x1c, 0xec, 0xd4, 0x15, 0x52, 0x38,
	0xd4, 0xab, 0x16, 0xc4, 0x81, 0x1b, 0x31, 0x02, 0x7a, 0x40, 0xaa, 0x7c, 0x41, 0x82, 0x83, 0xb5,
	0xf1, 0x8e, 0x9d, 0x55, 0xe2, 0x9d, 0xc8, 0xbb, 0x69, 0xc9, 0x5f, 0xc0, 0x85, 0x8f, 0xe0, 0x4b,
	0x72, 0xec, 0x91, 0x93, 0x41, 0xc9, 0x1f, 0xf4, 0x0b, 0x50, 0xd6, 0xae, 0x9a, 0x94, 0x9c, 0xbc,
	0xf3, 0xde, 0xcc, 0x78, 0xde, 0x9b, 0x21, 0x07, 0xa8, 0x72, 0x54, 0x42, 0x51, 0x8d, 0x63, 0x90,
	0x29, 0x4b, 0x34, 0x16, 0x73, 0x3a, 0x65, 0x05, 0xcb, 0x55, 0x30, 0x2d, 0x50, 0xa3, 0xbd, 0x5f,
	0xa7, 0x04, 0x9b, 0x29, 0x5d, 0x37, 0x31, 0x30, 0x1d, 0x32, 0x05, 0xf4, 0xfc, 0x78, 0x08, 0x9a,
	0x1d, 0xd3, 0x04, 0x85, 0xac, 0xaa, 0xba, 0x9d, 0x8a, 0x8f, 0x4d, 0x44, 0xab, 0xa0, 0xa6, 0xf6,
	0x33, 0xcc, 0xb0, 0xc2, 0xd7, 0xaf, 0x0a, 0xf5, 0xbf, 0x92, 0x47, 0x9f, 0x47, 0x42, 0xc3, 0x44,
	0x28, 0x0d, 0xfc, 0x23, 0xe2, 0xd8, 0x3e, 0x24, 0x77, 0x13, 0xe4, 0x10, 0x0b, 0xee, 0x58, 0x3d,
	0xab, 0xdf, 0x1a, 0x90, 0x65, 0xe9, 0xb5, 0x43, 0xe4, 0x70, 0xfa, 0x2e, 0x6a, 0xaf, 0xa9, 0x53,
	0x6e, 0x1f, 0x92, 0x87, 0x1c, 0x24, 0xe6, 0x71, 0x52, 0x00, 0xd3, 0x58, 0x38, 0x77, 0x7a, 0x56,
	0xff, 0x5e, 0xf4, 0xc0, 0x80, 0x61, 0x85, 0xf9, 0x3f, 0x9a, 0xa4, 0x7d, 0x66, 0x44, 0xd9, 0x3f,
	0x2d, 0x62, 0x6f, 0x14, 0x08, 0x94, 0x71, 0x0a, 0xe0, 0x58, 0xbd, 0x66, 0xff, 0xfe, 0x49, 0x27,
	0xa8, 0x27, 0x5d, 0xcb, 0x0a, 0x6a, 0x59, 0x41, 0x88, 0x42, 0x0e, 0x3e, 0x2d, 0x4a, 0xaf, 0x71,
	0x55, 0x7a, 0x9d, 0x39, 0xcb, 0x27, 0x6f, 0xfc, 0xff, 0x5b, 0xf8, 0xbf, 0xfe, 0x78, 0xfd, 0x4c,
	0xe8, 0xd1, 0x6c, 0x18, 0x24, 0x98, 0xd7, 0x9a, 0xeb, 0xcf, 0x91, 0xe2, 0x63, 0xaa, 0xe7, 0x53,
	0x50, 0xa6, 0x9b, 0x8a, 0xf6, 0x6e, 0xe6, 0x13, 0x28, 0xdf, 0x03, 0xd8, 0x29, 0xe9, 0xde, 0x6a,
	0x9a, 0x31, 0x15, 0x27, 0x28, 0xd5, 0x2c, 0x07, 0xa3, 0xaa, 0x35, 0x78, 0xb1, 0x28, 0x3d, 0xeb,
	0xaa, 0xf4, 0x0e, 0x76, 0x0e, 0xb1, 0x91, 0xef, 0x47, 0xcf, 0xb6, 0x7e, 0xf0, 0x81, 0xa9, 0xb0,
	0x62, 0xec, 0x13, 0xf2, 0x34, 0x05, 0x88, 0x13, 0x9c, 0x4c, 0x60, 0xbd, 0xcb, 0x98, 0x71, 0x5e,
	0x80, 0x52, 0x4e, 0xd3, 0x18, 0xf7, 0x24, 0x05, 0x08, 0xaf, 0xb9, 0xb7, 0x15, 0x65, 0x47, 0xe4,
	0xf1, 0xc5, 0xcd, 0x72, 0xe2, 0x11, 0xe2, 0x58, 0x39, 0x2d, 0x63, 0xd9, 0xf3, 0x60, 0xd7, 0x7d,
	0x04, 0xb7, 0x76, 0x19, 0xed, 0x5d, 0x6c, 0x03, 0x6a, 0x70, 0xb6, 0x58, 0xba, 0xd6, 0xe5, 0xd2,
	0xb5, 0xfe, 0x2e, 0x5d, 0xeb, 0xfb, 0xca, 0x6d, 0x5c, 0xae, 0xdc, 0xc6, 0xef, 0x95, 0xdb, 0xf8,
	0xf2, 0x7a, 0xc3, 0x45, 0x09, 0x33, 0x5d, 0xa0, 0x3c, 0xc2, 0x22, 0xbb, 0x7e, 0xd3, 0xf3, 0x57,
	0xf4, 0xdb, 0xf6, 0xc1, 0x1a, 0x67, 0x87, 0x6d, 0x73, 0x49, 0x2f, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x0e, 0x4e, 0x8c, 0xba, 0xd5, 0x02, 0x00, 0x00,
}

func (m *WhitelistedHook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhitelistedHook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhitelistedHook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomCreator) > 0 {
		i -= len(m.DenomCreator)
		copy(dAtA[i:], m.DenomCreator)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DenomCreator)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeID != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WhitelistedHooks) > 0 {
		for iNdEx := len(m.WhitelistedHooks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhitelistedHooks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.FeeCollectorAddress) > 0 {
		i -= len(m.FeeCollectorAddress)
		copy(dAtA[i:], m.FeeCollectorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeCollectorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DenomCreationGasConsume != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DenomCreationGasConsume))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DenomCreationFee) > 0 {
		for iNdEx := len(m.DenomCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhitelistedHook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeID != 0 {
		n += 1 + sovParams(uint64(m.CodeID))
	}
	l = len(m.DenomCreator)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DenomCreationFee) > 0 {
		for _, e := range m.DenomCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.DenomCreationGasConsume != 0 {
		n += 1 + sovParams(uint64(m.DenomCreationGasConsume))
	}
	l = len(m.FeeCollectorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.WhitelistedHooks) > 0 {
		for _, e := range m.WhitelistedHooks {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhitelistedHook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: WhitelistedHook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhitelistedHook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomCreator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomCreationFee = append(m.DenomCreationFee, types.Coin{})
			if err := m.DenomCreationFee[len(m.DenomCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationGasConsume", wireType)
			}
			m.DenomCreationGasConsume = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DenomCreationGasConsume |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCollectorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeCollectorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedHooks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedHooks = append(m.WhitelistedHooks, &WhitelistedHook{})
			if err := m.WhitelistedHooks[len(m.WhitelistedHooks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
