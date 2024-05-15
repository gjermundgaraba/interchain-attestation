// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package pessimist

import (
	binary "encoding/binary"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ClientState                     protoreflect.MessageDescriptor
	fd_ClientState_dependent_client_id protoreflect.FieldDescriptor
	fd_ClientState_latest_height       protoreflect.FieldDescriptor
)

func init() {
	file_hub_pessimist_pessimist_proto_init()
	md_ClientState = File_hub_pessimist_pessimist_proto.Messages().ByName("ClientState")
	fd_ClientState_dependent_client_id = md_ClientState.Fields().ByName("dependent_client_id")
	fd_ClientState_latest_height = md_ClientState.Fields().ByName("latest_height")
}

var _ protoreflect.Message = (*fastReflection_ClientState)(nil)

type fastReflection_ClientState ClientState

func (x *ClientState) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ClientState)(x)
}

func (x *ClientState) slowProtoReflect() protoreflect.Message {
	mi := &file_hub_pessimist_pessimist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ClientState_messageType fastReflection_ClientState_messageType
var _ protoreflect.MessageType = fastReflection_ClientState_messageType{}

type fastReflection_ClientState_messageType struct{}

func (x fastReflection_ClientState_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ClientState)(nil)
}
func (x fastReflection_ClientState_messageType) New() protoreflect.Message {
	return new(fastReflection_ClientState)
}
func (x fastReflection_ClientState_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ClientState
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ClientState) Descriptor() protoreflect.MessageDescriptor {
	return md_ClientState
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ClientState) Type() protoreflect.MessageType {
	return _fastReflection_ClientState_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ClientState) New() protoreflect.Message {
	return new(fastReflection_ClientState)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ClientState) Interface() protoreflect.ProtoMessage {
	return (*ClientState)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ClientState) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DependentClientId != "" {
		value := protoreflect.ValueOfString(x.DependentClientId)
		if !f(fd_ClientState_dependent_client_id, value) {
			return
		}
	}
	if x.LatestHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.LatestHeight)
		if !f(fd_ClientState_latest_height, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ClientState) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "hub.pessimist.ClientState.dependent_client_id":
		return x.DependentClientId != ""
	case "hub.pessimist.ClientState.latest_height":
		return x.LatestHeight != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: hub.pessimist.ClientState"))
		}
		panic(fmt.Errorf("message hub.pessimist.ClientState does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClientState) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "hub.pessimist.ClientState.dependent_client_id":
		x.DependentClientId = ""
	case "hub.pessimist.ClientState.latest_height":
		x.LatestHeight = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: hub.pessimist.ClientState"))
		}
		panic(fmt.Errorf("message hub.pessimist.ClientState does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ClientState) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "hub.pessimist.ClientState.dependent_client_id":
		value := x.DependentClientId
		return protoreflect.ValueOfString(value)
	case "hub.pessimist.ClientState.latest_height":
		value := x.LatestHeight
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: hub.pessimist.ClientState"))
		}
		panic(fmt.Errorf("message hub.pessimist.ClientState does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClientState) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "hub.pessimist.ClientState.dependent_client_id":
		x.DependentClientId = value.Interface().(string)
	case "hub.pessimist.ClientState.latest_height":
		x.LatestHeight = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: hub.pessimist.ClientState"))
		}
		panic(fmt.Errorf("message hub.pessimist.ClientState does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClientState) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "hub.pessimist.ClientState.dependent_client_id":
		panic(fmt.Errorf("field dependent_client_id of message hub.pessimist.ClientState is not mutable"))
	case "hub.pessimist.ClientState.latest_height":
		panic(fmt.Errorf("field latest_height of message hub.pessimist.ClientState is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: hub.pessimist.ClientState"))
		}
		panic(fmt.Errorf("message hub.pessimist.ClientState does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ClientState) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "hub.pessimist.ClientState.dependent_client_id":
		return protoreflect.ValueOfString("")
	case "hub.pessimist.ClientState.latest_height":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: hub.pessimist.ClientState"))
		}
		panic(fmt.Errorf("message hub.pessimist.ClientState does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ClientState) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in hub.pessimist.ClientState", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ClientState) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClientState) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ClientState) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ClientState) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ClientState)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DependentClientId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.LatestHeight != 0 {
			n += 9
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ClientState)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.LatestHeight != 0 {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.LatestHeight))
			i--
			dAtA[i] = 0x11
		}
		if len(x.DependentClientId) > 0 {
			i -= len(x.DependentClientId)
			copy(dAtA[i:], x.DependentClientId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DependentClientId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ClientState)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClientState: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DependentClientId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DependentClientId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
				}
				x.LatestHeight = 0
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LatestHeight = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: hub/pessimist/pessimist.proto

// buf:lint:ignore PACKAGE_VERSION_SUFFIX

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DependentClientId string `protobuf:"bytes,1,opt,name=dependent_client_id,json=dependentClientId,proto3" json:"dependent_client_id,omitempty"`
	LatestHeight      int64  `protobuf:"fixed64,2,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height,omitempty"`
}

func (x *ClientState) Reset() {
	*x = ClientState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_pessimist_pessimist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientState) ProtoMessage() {}

// Deprecated: Use ClientState.ProtoReflect.Descriptor instead.
func (*ClientState) Descriptor() ([]byte, []int) {
	return file_hub_pessimist_pessimist_proto_rawDescGZIP(), []int{0}
}

func (x *ClientState) GetDependentClientId() string {
	if x != nil {
		return x.DependentClientId
	}
	return ""
}

func (x *ClientState) GetLatestHeight() int64 {
	if x != nil {
		return x.LatestHeight
	}
	return 0
}

var File_hub_pessimist_pessimist_proto protoreflect.FileDescriptor

var file_hub_pessimist_pessimist_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x2f,
	0x70, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0c, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x42,
	0x98, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x65, 0x73, 0x73,
	0x69, 0x6d, 0x69, 0x73, 0x74, 0x42, 0x0e, 0x50, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73,
	0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x65,
	0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0xa2, 0x02, 0x03, 0x48, 0x50, 0x58, 0xaa, 0x02, 0x0d,
	0x48, 0x75, 0x62, 0x2e, 0x50, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0xca, 0x02, 0x0d,
	0x48, 0x75, 0x62, 0x5c, 0x50, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0xe2, 0x02, 0x19,
	0x48, 0x75, 0x62, 0x5c, 0x50, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x48, 0x75, 0x62, 0x3a,
	0x3a, 0x50, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hub_pessimist_pessimist_proto_rawDescOnce sync.Once
	file_hub_pessimist_pessimist_proto_rawDescData = file_hub_pessimist_pessimist_proto_rawDesc
)

func file_hub_pessimist_pessimist_proto_rawDescGZIP() []byte {
	file_hub_pessimist_pessimist_proto_rawDescOnce.Do(func() {
		file_hub_pessimist_pessimist_proto_rawDescData = protoimpl.X.CompressGZIP(file_hub_pessimist_pessimist_proto_rawDescData)
	})
	return file_hub_pessimist_pessimist_proto_rawDescData
}

var file_hub_pessimist_pessimist_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hub_pessimist_pessimist_proto_goTypes = []interface{}{
	(*ClientState)(nil), // 0: hub.pessimist.ClientState
}
var file_hub_pessimist_pessimist_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hub_pessimist_pessimist_proto_init() }
func file_hub_pessimist_pessimist_proto_init() {
	if File_hub_pessimist_pessimist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hub_pessimist_pessimist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hub_pessimist_pessimist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hub_pessimist_pessimist_proto_goTypes,
		DependencyIndexes: file_hub_pessimist_pessimist_proto_depIdxs,
		MessageInfos:      file_hub_pessimist_pessimist_proto_msgTypes,
	}.Build()
	File_hub_pessimist_pessimist_proto = out.File
	file_hub_pessimist_pessimist_proto_rawDesc = nil
	file_hub_pessimist_pessimist_proto_goTypes = nil
	file_hub_pessimist_pessimist_proto_depIdxs = nil
}
