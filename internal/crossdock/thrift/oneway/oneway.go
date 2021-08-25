// Code generated by thriftrw v1.29.0. DO NOT EDIT.
// @generated

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package oneway

import (
	fmt "fmt"
	stream "go.uber.org/thriftrw/protocol/stream"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "oneway",
	Package:  "go.uber.org/yarpc/internal/crossdock/thrift/oneway",
	FilePath: "oneway.thrift",
	SHA1:     "2eb401ce231c98a633b0731ea9c0cdd24d40d3ca",
	Raw:      rawIDL,
}

const rawIDL = "service Oneway {\n    oneway void echo (1: string token)\n}\n"

// Oneway_Echo_Args represents the arguments for the Oneway.echo function.
//
// The arguments for echo are sent and received over the wire as this struct.
type Oneway_Echo_Args struct {
	Token *string `json:"token,omitempty"`
}

// ToWire translates a Oneway_Echo_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Oneway_Echo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Token != nil {
		w, err = wire.NewValueString(*(v.Token)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Oneway_Echo_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Oneway_Echo_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Oneway_Echo_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Oneway_Echo_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Token = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a Oneway_Echo_Args struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Oneway_Echo_Args struct could not be encoded.
func (v *Oneway_Echo_Args) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Token != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Token)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Oneway_Echo_Args struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Oneway_Echo_Args struct could not be generated from the wire
// representation.
func (v *Oneway_Echo_Args) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch fh.ID {
		case 1:
			if fh.Type == wire.TBinary {
				var x string
				x, err = sr.ReadString()
				v.Token = &x
				if err != nil {
					return err
				}

			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a Oneway_Echo_Args
// struct.
func (v *Oneway_Echo_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Token != nil {
		fields[i] = fmt.Sprintf("Token: %v", *(v.Token))
		i++
	}

	return fmt.Sprintf("Oneway_Echo_Args{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Oneway_Echo_Args match the
// provided Oneway_Echo_Args.
//
// This function performs a deep comparison.
func (v *Oneway_Echo_Args) Equals(rhs *Oneway_Echo_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Token, rhs.Token) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Oneway_Echo_Args.
func (v *Oneway_Echo_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Token != nil {
		enc.AddString("token", *v.Token)
	}
	return err
}

// GetToken returns the value of Token if it is set or its
// zero value if it is unset.
func (v *Oneway_Echo_Args) GetToken() (o string) {
	if v != nil && v.Token != nil {
		return *v.Token
	}

	return
}

// IsSetToken returns true if Token is not nil.
func (v *Oneway_Echo_Args) IsSetToken() bool {
	return v != nil && v.Token != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echo" for this struct.
func (v *Oneway_Echo_Args) MethodName() string {
	return "echo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be OneWay for this struct.
func (v *Oneway_Echo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

// Oneway_Echo_Helper provides functions that aid in handling the
// parameters and return values of the Oneway.echo
// function.
var Oneway_Echo_Helper = struct {
	// Args accepts the parameters of echo in-order and returns
	// the arguments struct for the function.
	Args func(
		token *string,
	) *Oneway_Echo_Args
}{}

func init() {
	Oneway_Echo_Helper.Args = func(
		token *string,
	) *Oneway_Echo_Args {
		return &Oneway_Echo_Args{
			Token: token,
		}
	}

}
