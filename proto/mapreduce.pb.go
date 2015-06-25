// Code generated by protoc-gen-gogo.
// source: mapreduce.proto
// DO NOT EDIT!

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		mapreduce.proto

	It has these top-level messages:
		MapperRequest
		MapperResponse
		ReducerRequest
		ReducerResponse
		WorkRequest
		WorkConfigResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import io "io"
import fmt "fmt"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type MapperRequest struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MapperRequest) Reset()         { *m = MapperRequest{} }
func (m *MapperRequest) String() string { return proto1.CompactTextString(m) }
func (*MapperRequest) ProtoMessage()    {}

type MapperResponse struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MapperResponse) Reset()         { *m = MapperResponse{} }
func (m *MapperResponse) String() string { return proto1.CompactTextString(m) }
func (*MapperResponse) ProtoMessage()    {}

type ReducerRequest struct {
	Key   string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []string `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
}

func (m *ReducerRequest) Reset()         { *m = ReducerRequest{} }
func (m *ReducerRequest) String() string { return proto1.CompactTextString(m) }
func (*ReducerRequest) ProtoMessage()    {}

type ReducerResponse struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ReducerResponse) Reset()         { *m = ReducerResponse{} }
func (m *ReducerResponse) String() string { return proto1.CompactTextString(m) }
func (*ReducerResponse) ProtoMessage()    {}

type WorkRequest struct {
	TaskID uint64 `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
}

func (m *WorkRequest) Reset()         { *m = WorkRequest{} }
func (m *WorkRequest) String() string { return proto1.CompactTextString(m) }
func (*WorkRequest) ProtoMessage()    {}

type WorkConfigResponse struct {
	Key   []string `protobuf:"bytes,1,rep,name=key" json:"key,omitempty"`
	Value []string `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
}

func (m *WorkConfigResponse) Reset()         { *m = WorkConfigResponse{} }
func (m *WorkConfigResponse) String() string { return proto1.CompactTextString(m) }
func (*WorkConfigResponse) ProtoMessage()    {}

func init() {
}
func (m *MapperRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMapreduce(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *MapperResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMapreduce(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ReducerRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMapreduce(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ReducerResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMapreduce(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *WorkRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskID", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TaskID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMapreduce(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *WorkConfigResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMapreduce(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipMapreduce(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			return iNdEx, nil
		case 3:
			for {
				var wire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				wireType := int(wire & 0x7)
				if wireType == 4 {
					break
				}
				next, err := skipMapreduce(data[start:])
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
func (m *MapperRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMapreduce(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMapreduce(uint64(l))
	}
	return n
}

func (m *MapperResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMapreduce(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMapreduce(uint64(l))
	}
	return n
}

func (m *ReducerRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMapreduce(uint64(l))
	}
	if len(m.Value) > 0 {
		for _, s := range m.Value {
			l = len(s)
			n += 1 + l + sovMapreduce(uint64(l))
		}
	}
	return n
}

func (m *ReducerResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMapreduce(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMapreduce(uint64(l))
	}
	return n
}

func (m *WorkRequest) Size() (n int) {
	var l int
	_ = l
	if m.TaskID != 0 {
		n += 1 + sovMapreduce(uint64(m.TaskID))
	}
	return n
}

func (m *WorkConfigResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Key) > 0 {
		for _, s := range m.Key {
			l = len(s)
			n += 1 + l + sovMapreduce(uint64(l))
		}
	}
	if len(m.Value) > 0 {
		for _, s := range m.Value {
			l = len(s)
			n += 1 + l + sovMapreduce(uint64(l))
		}
	}
	return n
}

func sovMapreduce(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMapreduce(x uint64) (n int) {
	return sovMapreduce(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MapperRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MapperRequest) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMapreduce(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if len(m.Value) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintMapreduce(data, i, uint64(len(m.Value)))
		i += copy(data[i:], m.Value)
	}
	return i, nil
}

func (m *MapperResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MapperResponse) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMapreduce(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if len(m.Value) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintMapreduce(data, i, uint64(len(m.Value)))
		i += copy(data[i:], m.Value)
	}
	return i, nil
}

func (m *ReducerRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReducerRequest) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMapreduce(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if len(m.Value) > 0 {
		for _, s := range m.Value {
			data[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *ReducerResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReducerResponse) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMapreduce(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if len(m.Value) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintMapreduce(data, i, uint64(len(m.Value)))
		i += copy(data[i:], m.Value)
	}
	return i, nil
}

func (m *WorkRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WorkRequest) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TaskID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintMapreduce(data, i, uint64(m.TaskID))
	}
	return i, nil
}

func (m *WorkConfigResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WorkConfigResponse) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		for _, s := range m.Key {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Value) > 0 {
		for _, s := range m.Value {
			data[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Mapreduce(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Mapreduce(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMapreduce(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}

// Client API for Mapper service

type MapperClient interface {
	GetEmitResult(ctx context.Context, in *MapperRequest, opts ...grpc.CallOption) (Mapper_GetEmitResultClient, error)
}

type mapperClient struct {
	cc *grpc.ClientConn
}

func NewMapperClient(cc *grpc.ClientConn) MapperClient {
	return &mapperClient{cc}
}

func (c *mapperClient) GetEmitResult(ctx context.Context, in *MapperRequest, opts ...grpc.CallOption) (Mapper_GetEmitResultClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Mapper_serviceDesc.Streams[0], c.cc, "/proto.Mapper/GetEmitResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapperGetEmitResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mapper_GetEmitResultClient interface {
	Recv() (*MapperResponse, error)
	grpc.ClientStream
}

type mapperGetEmitResultClient struct {
	grpc.ClientStream
}

func (x *mapperGetEmitResultClient) Recv() (*MapperResponse, error) {
	m := new(MapperResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Mapper service

type MapperServer interface {
	GetEmitResult(*MapperRequest, Mapper_GetEmitResultServer) error
}

func RegisterMapperServer(s *grpc.Server, srv MapperServer) {
	s.RegisterService(&_Mapper_serviceDesc, srv)
}

func _Mapper_GetEmitResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MapperRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapperServer).GetEmitResult(m, &mapperGetEmitResultServer{stream})
}

type Mapper_GetEmitResultServer interface {
	Send(*MapperResponse) error
	grpc.ServerStream
}

type mapperGetEmitResultServer struct {
	grpc.ServerStream
}

func (x *mapperGetEmitResultServer) Send(m *MapperResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Mapper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Mapper",
	HandlerType: (*MapperServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEmitResult",
			Handler:       _Mapper_GetEmitResult_Handler,
			ServerStreams: true,
		},
	},
}

// Client API for MapperStream service

type MapperStreamClient interface {
	GetStreamEmitResult(ctx context.Context, opts ...grpc.CallOption) (MapperStream_GetStreamEmitResultClient, error)
}

type mapperStreamClient struct {
	cc *grpc.ClientConn
}

func NewMapperStreamClient(cc *grpc.ClientConn) MapperStreamClient {
	return &mapperStreamClient{cc}
}

func (c *mapperStreamClient) GetStreamEmitResult(ctx context.Context, opts ...grpc.CallOption) (MapperStream_GetStreamEmitResultClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MapperStream_serviceDesc.Streams[0], c.cc, "/proto.MapperStream/GetStreamEmitResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapperStreamGetStreamEmitResultClient{stream}
	return x, nil
}

type MapperStream_GetStreamEmitResultClient interface {
	Send(*MapperRequest) error
	Recv() (*MapperResponse, error)
	grpc.ClientStream
}

type mapperStreamGetStreamEmitResultClient struct {
	grpc.ClientStream
}

func (x *mapperStreamGetStreamEmitResultClient) Send(m *MapperRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mapperStreamGetStreamEmitResultClient) Recv() (*MapperResponse, error) {
	m := new(MapperResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MapperStream service

type MapperStreamServer interface {
	GetStreamEmitResult(MapperStream_GetStreamEmitResultServer) error
}

func RegisterMapperStreamServer(s *grpc.Server, srv MapperStreamServer) {
	s.RegisterService(&_MapperStream_serviceDesc, srv)
}

func _MapperStream_GetStreamEmitResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MapperStreamServer).GetStreamEmitResult(&mapperStreamGetStreamEmitResultServer{stream})
}

type MapperStream_GetStreamEmitResultServer interface {
	Send(*MapperResponse) error
	Recv() (*MapperRequest, error)
	grpc.ServerStream
}

type mapperStreamGetStreamEmitResultServer struct {
	grpc.ServerStream
}

func (x *mapperStreamGetStreamEmitResultServer) Send(m *MapperResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mapperStreamGetStreamEmitResultServer) Recv() (*MapperRequest, error) {
	m := new(MapperRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MapperStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MapperStream",
	HandlerType: (*MapperStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStreamEmitResult",
			Handler:       _MapperStream_GetStreamEmitResult_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

// Client API for Reducer service

type ReducerClient interface {
	GetCollectResult(ctx context.Context, in *ReducerRequest, opts ...grpc.CallOption) (Reducer_GetCollectResultClient, error)
}

type reducerClient struct {
	cc *grpc.ClientConn
}

func NewReducerClient(cc *grpc.ClientConn) ReducerClient {
	return &reducerClient{cc}
}

func (c *reducerClient) GetCollectResult(ctx context.Context, in *ReducerRequest, opts ...grpc.CallOption) (Reducer_GetCollectResultClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Reducer_serviceDesc.Streams[0], c.cc, "/proto.Reducer/GetCollectResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &reducerGetCollectResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Reducer_GetCollectResultClient interface {
	Recv() (*ReducerResponse, error)
	grpc.ClientStream
}

type reducerGetCollectResultClient struct {
	grpc.ClientStream
}

func (x *reducerGetCollectResultClient) Recv() (*ReducerResponse, error) {
	m := new(ReducerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Reducer service

type ReducerServer interface {
	GetCollectResult(*ReducerRequest, Reducer_GetCollectResultServer) error
}

func RegisterReducerServer(s *grpc.Server, srv ReducerServer) {
	s.RegisterService(&_Reducer_serviceDesc, srv)
}

func _Reducer_GetCollectResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReducerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReducerServer).GetCollectResult(m, &reducerGetCollectResultServer{stream})
}

type Reducer_GetCollectResultServer interface {
	Send(*ReducerResponse) error
	grpc.ServerStream
}

type reducerGetCollectResultServer struct {
	grpc.ServerStream
}

func (x *reducerGetCollectResultServer) Send(m *ReducerResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Reducer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Reducer",
	HandlerType: (*ReducerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCollectResult",
			Handler:       _Reducer_GetCollectResult_Handler,
			ServerStreams: true,
		},
	},
}

// Client API for ReducerStream service

type ReducerStreamClient interface {
	GetStreamCollectResult(ctx context.Context, opts ...grpc.CallOption) (ReducerStream_GetStreamCollectResultClient, error)
}

type reducerStreamClient struct {
	cc *grpc.ClientConn
}

func NewReducerStreamClient(cc *grpc.ClientConn) ReducerStreamClient {
	return &reducerStreamClient{cc}
}

func (c *reducerStreamClient) GetStreamCollectResult(ctx context.Context, opts ...grpc.CallOption) (ReducerStream_GetStreamCollectResultClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ReducerStream_serviceDesc.Streams[0], c.cc, "/proto.ReducerStream/GetStreamCollectResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &reducerStreamGetStreamCollectResultClient{stream}
	return x, nil
}

type ReducerStream_GetStreamCollectResultClient interface {
	Send(*ReducerRequest) error
	Recv() (*ReducerResponse, error)
	grpc.ClientStream
}

type reducerStreamGetStreamCollectResultClient struct {
	grpc.ClientStream
}

func (x *reducerStreamGetStreamCollectResultClient) Send(m *ReducerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reducerStreamGetStreamCollectResultClient) Recv() (*ReducerResponse, error) {
	m := new(ReducerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ReducerStream service

type ReducerStreamServer interface {
	GetStreamCollectResult(ReducerStream_GetStreamCollectResultServer) error
}

func RegisterReducerStreamServer(s *grpc.Server, srv ReducerStreamServer) {
	s.RegisterService(&_ReducerStream_serviceDesc, srv)
}

func _ReducerStream_GetStreamCollectResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReducerStreamServer).GetStreamCollectResult(&reducerStreamGetStreamCollectResultServer{stream})
}

type ReducerStream_GetStreamCollectResultServer interface {
	Send(*ReducerResponse) error
	Recv() (*ReducerRequest, error)
	grpc.ServerStream
}

type reducerStreamGetStreamCollectResultServer struct {
	grpc.ServerStream
}

func (x *reducerStreamGetStreamCollectResultServer) Send(m *ReducerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reducerStreamGetStreamCollectResultServer) Recv() (*ReducerRequest, error) {
	m := new(ReducerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ReducerStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ReducerStream",
	HandlerType: (*ReducerStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStreamCollectResult",
			Handler:       _ReducerStream_GetStreamCollectResult_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

// Client API for Master service

type MasterClient interface {
	GetWork(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (Master_GetWorkClient, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) GetWork(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (Master_GetWorkClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Master_serviceDesc.Streams[0], c.cc, "/proto.Master/GetWork", opts...)
	if err != nil {
		return nil, err
	}
	x := &masterGetWorkClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Master_GetWorkClient interface {
	Recv() (*WorkConfigResponse, error)
	grpc.ClientStream
}

type masterGetWorkClient struct {
	grpc.ClientStream
}

func (x *masterGetWorkClient) Recv() (*WorkConfigResponse, error) {
	m := new(WorkConfigResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Master service

type MasterServer interface {
	GetWork(*WorkRequest, Master_GetWorkServer) error
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_GetWork_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WorkRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MasterServer).GetWork(m, &masterGetWorkServer{stream})
}

type Master_GetWorkServer interface {
	Send(*WorkConfigResponse) error
	grpc.ServerStream
}

type masterGetWorkServer struct {
	grpc.ServerStream
}

func (x *masterGetWorkServer) Send(m *WorkConfigResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Master",
	HandlerType: (*MasterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWork",
			Handler:       _Master_GetWork_Handler,
			ServerStreams: true,
		},
	},
}

// Client API for Worker service

type WorkerClient interface {
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

// Server API for Worker service

type WorkerServer interface {
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
}
