// Code generated by protoc-gen-gogo.
// source: subpkg/subproto.proto
// DO NOT EDIT!

/*
Package subpkg is a generated protocol buffer package.

It is generated from these files:
	subpkg/subproto.proto

It has these top-level messages:
	SubObject
*/
package subpkg

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SubObject struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SubObject) Reset()         { *m = SubObject{} }
func (m *SubObject) String() string { return proto.CompactTextString(m) }
func (*SubObject) ProtoMessage()    {}

func init() {
	proto.RegisterType((*SubObject)(nil), "subpkg.SubObject")
}
