// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/rights.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Right is the enum that defines all the different rights to do something in
// the network.
type Right int32

const (
	RIGHT_APPLICATION_DELETE        Right = 0
	RIGHT_APPLICATION_SETTINGS      Right = 1
	RIGHT_APPLICATION_COLLABORATORS Right = 2
	RIGHT_GATEWAY_OWNER             Right = 10
)

var Right_name = map[int32]string{
	0:  "RIGHT_APPLICATION_DELETE",
	1:  "RIGHT_APPLICATION_SETTINGS",
	2:  "RIGHT_APPLICATION_COLLABORATORS",
	10: "RIGHT_GATEWAY_OWNER",
}
var Right_value = map[string]int32{
	"RIGHT_APPLICATION_DELETE":        0,
	"RIGHT_APPLICATION_SETTINGS":      1,
	"RIGHT_APPLICATION_COLLABORATORS": 2,
	"RIGHT_GATEWAY_OWNER":             10,
}

func (x Right) String() string {
	return proto.EnumName(Right_name, int32(x))
}
func (Right) EnumDescriptor() ([]byte, []int) { return fileDescriptorRights, []int{0} }

type APIKey struct {
	// name is the API key name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// key is the actual API key.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// rights are the rights this API key bears.
	Rights []Right `protobuf:"varint,3,rep,packed,name=rights,enum=ttn.v3.Right" json:"rights,omitempty"`
}

func (m *APIKey) Reset()                    { *m = APIKey{} }
func (*APIKey) ProtoMessage()               {}
func (*APIKey) Descriptor() ([]byte, []int) { return fileDescriptorRights, []int{0} }

func (m *APIKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *APIKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *APIKey) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func init() {
	proto.RegisterType((*APIKey)(nil), "ttn.v3.APIKey")
	proto.RegisterEnum("ttn.v3.Right", Right_name, Right_value)
}
func (m *APIKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *APIKey) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRights(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRights(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Rights) > 0 {
		dAtA2 := make([]byte, len(m.Rights)*10)
		var j1 int
		for _, num := range m.Rights {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRights(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func encodeFixed64Rights(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Rights(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRights(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *APIKey) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRights(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRights(uint64(l))
	}
	if len(m.Rights) > 0 {
		l = 0
		for _, e := range m.Rights {
			l += sovRights(uint64(e))
		}
		n += 1 + sovRights(uint64(l)) + l
	}
	return n
}

func sovRights(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRights(x uint64) (n int) {
	return sovRights(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *APIKey) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&APIKey{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Rights:` + fmt.Sprintf("%v", this.Rights) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRights(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *APIKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRights
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRights
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRights
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRights
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRights
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v Right
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRights
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (Right(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rights = append(m.Rights, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRights
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthRights
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v Right
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRights
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (Right(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rights = append(m.Rights, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rights", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRights(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRights
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
func skipRights(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRights
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
					return 0, ErrIntOverflowRights
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRights
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRights
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRights
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRights(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRights = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRights   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/rights.proto", fileDescriptorRights)
}

var fileDescriptorRights = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0xc9, 0x48, 0x0d, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f,
	0xf6, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x29, 0xc9, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4,
	0x2f, 0xca, 0x4c, 0xcf, 0x28, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x29,
	0xc9, 0xd3, 0x2b, 0x33, 0x96, 0xd2, 0x45, 0xd2, 0x99, 0x9e, 0x9f, 0x9e, 0xaf, 0x0f, 0x96, 0x4e,
	0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x4d, 0x29, 0x94, 0x8b, 0xcd, 0x31, 0xc0,
	0xd3, 0x3b, 0xb5, 0x52, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x02, 0x0b, 0x81,
	0x98, 0x42, 0xaa, 0x5c, 0x6c, 0x10, 0x6b, 0x25, 0x98, 0x15, 0x98, 0x35, 0xf8, 0x8c, 0x78, 0xf5,
	0x20, 0xf6, 0xea, 0x05, 0x81, 0x44, 0x83, 0xa0, 0x92, 0x5a, 0xbd, 0x8c, 0x5c, 0xac, 0x60, 0x11,
	0x21, 0x19, 0x2e, 0x89, 0x20, 0x4f, 0x77, 0x8f, 0x90, 0x78, 0xc7, 0x80, 0x00, 0x1f, 0x4f, 0x67,
	0xc7, 0x10, 0x4f, 0x7f, 0xbf, 0x78, 0x17, 0x57, 0x1f, 0xd7, 0x10, 0x57, 0x01, 0x06, 0x21, 0x39,
	0x2e, 0x29, 0x4c, 0xd9, 0x60, 0xd7, 0x90, 0x10, 0x4f, 0x3f, 0xf7, 0x60, 0x01, 0x46, 0x21, 0x65,
	0x2e, 0x79, 0x4c, 0x79, 0x67, 0x7f, 0x1f, 0x1f, 0x47, 0x27, 0xff, 0x20, 0xc7, 0x10, 0xff, 0xa0,
	0x60, 0x01, 0x26, 0x21, 0x71, 0x2e, 0x61, 0x88, 0x22, 0x77, 0xc7, 0x10, 0xd7, 0x70, 0xc7, 0xc8,
	0x78, 0xff, 0x70, 0x3f, 0xd7, 0x20, 0x01, 0x2e, 0x29, 0x8e, 0x8e, 0xc5, 0x72, 0x0c, 0x1b, 0x96,
	0xc8, 0x31, 0x38, 0xb9, 0xdf, 0x78, 0x28, 0xc7, 0xd0, 0xf0, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23,
	0x39, 0xc6, 0x28, 0x4d, 0x42, 0xa1, 0x5d, 0x90, 0x9d, 0x0e, 0xa2, 0x0b, 0x92, 0x92, 0xd8, 0xc0,
	0xc1, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x65, 0x12, 0x02, 0xa1, 0x01, 0x00, 0x00,
}
