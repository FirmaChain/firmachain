// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firmachain/contract/contract_log.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ContractLog struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	ContractHash string `protobuf:"bytes,3,opt,name=contractHash,proto3" json:"contractHash,omitempty"`
	TimeStamp    uint64 `protobuf:"varint,4,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	EventName    string `protobuf:"bytes,5,opt,name=eventName,proto3" json:"eventName,omitempty"`
	OwnerAddress string `protobuf:"bytes,6,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	JsonString   string `protobuf:"bytes,7,opt,name=jsonString,proto3" json:"jsonString,omitempty"`
}

func (m *ContractLog) Reset()         { *m = ContractLog{} }
func (m *ContractLog) String() string { return proto.CompactTextString(m) }
func (*ContractLog) ProtoMessage()    {}
func (*ContractLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_55035304f5ef7b23, []int{0}
}
func (m *ContractLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractLog.Merge(m, src)
}
func (m *ContractLog) XXX_Size() int {
	return m.Size()
}
func (m *ContractLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractLog.DiscardUnknown(m)
}

var xxx_messageInfo_ContractLog proto.InternalMessageInfo

func (m *ContractLog) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ContractLog) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContractLog) GetContractHash() string {
	if m != nil {
		return m.ContractHash
	}
	return ""
}

func (m *ContractLog) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *ContractLog) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *ContractLog) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *ContractLog) GetJsonString() string {
	if m != nil {
		return m.JsonString
	}
	return ""
}

func init() {
	proto.RegisterType((*ContractLog)(nil), "firmachain.firmachain.contract.ContractLog")
}

func init() {
	proto.RegisterFile("firmachain/contract/contract_log.proto", fileDescriptor_55035304f5ef7b23)
}

var fileDescriptor_55035304f5ef7b23 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe3, 0x7c, 0xfd, 0x5a, 0xd5, 0x20, 0x06, 0x8b, 0xc1, 0x42, 0xc8, 0xaa, 0x3a, 0xa0,
	0x4e, 0x89, 0x10, 0x4f, 0x00, 0x2c, 0x48, 0x20, 0x86, 0x76, 0x63, 0x41, 0x6e, 0x62, 0x1c, 0x23,
	0xe2, 0x1b, 0x39, 0x97, 0x7f, 0x6f, 0xc1, 0x63, 0x31, 0x76, 0x44, 0x62, 0x41, 0xc9, 0x8b, 0xa0,
	0xb8, 0x84, 0x98, 0xed, 0xf8, 0xe7, 0xdf, 0x3d, 0xc3, 0xa1, 0x47, 0x77, 0xc6, 0x95, 0x32, 0x2b,
	0xa4, 0xb1, 0x69, 0x06, 0x16, 0x9d, 0xcc, 0xf0, 0x37, 0xdc, 0x3e, 0x80, 0x4e, 0x2a, 0x07, 0x08,
	0x4c, 0x0c, 0x5e, 0x12, 0xc4, 0xde, 0x3c, 0xd8, 0xd7, 0xa0, 0xc1, 0xab, 0x69, 0x97, 0xb6, 0x57,
	0xf3, 0x4f, 0x42, 0x77, 0xce, 0x7f, 0x94, 0x2b, 0xd0, 0x8c, 0xd3, 0x49, 0xe6, 0x94, 0x44, 0x70,
	0x9c, 0xcc, 0xc8, 0x62, 0xba, 0xec, 0x9f, 0x6c, 0x8f, 0xc6, 0x26, 0xe7, 0xf1, 0x8c, 0x2c, 0x46,
	0xcb, 0xd8, 0xe4, 0x6c, 0x4e, 0x77, 0xfb, 0xee, 0x0b, 0x59, 0x17, 0xfc, 0x9f, 0xd7, 0xff, 0x30,
	0x76, 0x48, 0xa7, 0x68, 0x4a, 0xb5, 0x42, 0x59, 0x56, 0x7c, 0xe4, 0x4f, 0x07, 0xd0, 0xfd, 0xaa,
	0x27, 0x65, 0xf1, 0x5a, 0x96, 0x8a, 0xff, 0xf7, 0xe7, 0x03, 0xe8, 0xfa, 0xe1, 0xd9, 0x2a, 0x77,
	0x9a, 0xe7, 0x4e, 0xd5, 0x35, 0x1f, 0x6f, 0xfb, 0x43, 0xc6, 0x04, 0xa5, 0xf7, 0x35, 0xd8, 0x15,
	0x3a, 0x63, 0x35, 0x9f, 0x78, 0x23, 0x20, 0x67, 0x97, 0xef, 0x8d, 0x20, 0x9b, 0x46, 0x90, 0xaf,
	0x46, 0x90, 0xb7, 0x56, 0x44, 0x9b, 0x56, 0x44, 0x1f, 0xad, 0x88, 0x6e, 0x8e, 0xb5, 0xc1, 0xe2,
	0x71, 0x9d, 0x64, 0x50, 0xa6, 0xc1, 0xc0, 0x41, 0x7c, 0x19, 0xd6, 0xc6, 0xd7, 0x4a, 0xd5, 0xeb,
	0xb1, 0x5f, 0xec, 0xe4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x3e, 0xad, 0xc4, 0x91, 0x01, 0x00,
	0x00,
}

func (m *ContractLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.JsonString) > 0 {
		i -= len(m.JsonString)
		copy(dAtA[i:], m.JsonString)
		i = encodeVarintContractLog(dAtA, i, uint64(len(m.JsonString)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintContractLog(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.EventName) > 0 {
		i -= len(m.EventName)
		copy(dAtA[i:], m.EventName)
		i = encodeVarintContractLog(dAtA, i, uint64(len(m.EventName)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeStamp != 0 {
		i = encodeVarintContractLog(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ContractHash) > 0 {
		i -= len(m.ContractHash)
		copy(dAtA[i:], m.ContractHash)
		i = encodeVarintContractLog(dAtA, i, uint64(len(m.ContractHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintContractLog(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintContractLog(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContractLog(dAtA []byte, offset int, v uint64) int {
	offset -= sovContractLog(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovContractLog(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovContractLog(uint64(m.Id))
	}
	l = len(m.ContractHash)
	if l > 0 {
		n += 1 + l + sovContractLog(uint64(l))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovContractLog(uint64(m.TimeStamp))
	}
	l = len(m.EventName)
	if l > 0 {
		n += 1 + l + sovContractLog(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovContractLog(uint64(l))
	}
	l = len(m.JsonString)
	if l > 0 {
		n += 1 + l + sovContractLog(uint64(l))
	}
	return n
}

func sovContractLog(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContractLog(x uint64) (n int) {
	return sovContractLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContractLog
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
			return fmt.Errorf("proto: ContractLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractLog
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
				return ErrInvalidLengthContractLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractLog
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
				return fmt.Errorf("proto: wrong wireType = %d for field ContractHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractLog
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
				return ErrInvalidLengthContractLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractLog
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
				return ErrInvalidLengthContractLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractLog
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
				return ErrInvalidLengthContractLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractLog
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
				return ErrInvalidLengthContractLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractLog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JsonString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContractLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContractLog
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
func skipContractLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContractLog
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
					return 0, ErrIntOverflowContractLog
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
					return 0, ErrIntOverflowContractLog
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
				return 0, ErrInvalidLengthContractLog
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContractLog
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContractLog
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContractLog        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContractLog          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContractLog = fmt.Errorf("proto: unexpected end of group")
)
