// Code generated by thriftrw v1.29.2. DO NOT EDIT.
// @generated

package NOSERVICES

import (
	fmt "fmt"
	stream "go.uber.org/thriftrw/protocol/stream"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type ExWithAnnotation struct {
	Foo *string `json:"foo,omitempty"`
}

// ToWire translates a ExWithAnnotation struct into a Thrift-level intermediate
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
func (v *ExWithAnnotation) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Foo != nil {
		w, err = wire.NewValueString(*(v.Foo)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ExWithAnnotation struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ExWithAnnotation struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ExWithAnnotation
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ExWithAnnotation) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Foo = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a ExWithAnnotation struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a ExWithAnnotation struct could not be encoded.
func (v *ExWithAnnotation) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Foo != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Foo)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a ExWithAnnotation struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a ExWithAnnotation struct could not be generated from the wire
// representation.
func (v *ExWithAnnotation) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Foo = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
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

// String returns a readable string representation of a ExWithAnnotation
// struct.
func (v *ExWithAnnotation) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Foo != nil {
		fields[i] = fmt.Sprintf("Foo: %v", *(v.Foo))
		i++
	}

	return fmt.Sprintf("ExWithAnnotation{%v}", strings.Join(fields[:i], ", "))
}

// ErrorName is the name of this type as defined in the Thrift
// file.
func (*ExWithAnnotation) ErrorName() string {
	return "ExWithAnnotation"
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this ExWithAnnotation match the
// provided ExWithAnnotation.
//
// This function performs a deep comparison.
func (v *ExWithAnnotation) Equals(rhs *ExWithAnnotation) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Foo, rhs.Foo) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ExWithAnnotation.
func (v *ExWithAnnotation) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Foo != nil {
		enc.AddString("foo", *v.Foo)
	}
	return err
}

// GetFoo returns the value of Foo if it is set or its
// zero value if it is unset.
func (v *ExWithAnnotation) GetFoo() (o string) {
	if v != nil && v.Foo != nil {
		return *v.Foo
	}

	return
}

// IsSetFoo returns true if Foo is not nil.
func (v *ExWithAnnotation) IsSetFoo() bool {
	return v != nil && v.Foo != nil
}

func (v *ExWithAnnotation) Error() string {
	return v.String()
}

type ExWithoutAnnotation struct {
	Bar *string `json:"bar,omitempty"`
}

// ToWire translates a ExWithoutAnnotation struct into a Thrift-level intermediate
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
func (v *ExWithoutAnnotation) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Bar != nil {
		w, err = wire.NewValueString(*(v.Bar)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ExWithoutAnnotation struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ExWithoutAnnotation struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ExWithoutAnnotation
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ExWithoutAnnotation) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Bar = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a ExWithoutAnnotation struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a ExWithoutAnnotation struct could not be encoded.
func (v *ExWithoutAnnotation) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Bar != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Bar)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a ExWithoutAnnotation struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a ExWithoutAnnotation struct could not be generated from the wire
// representation.
func (v *ExWithoutAnnotation) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Bar = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
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

// String returns a readable string representation of a ExWithoutAnnotation
// struct.
func (v *ExWithoutAnnotation) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Bar != nil {
		fields[i] = fmt.Sprintf("Bar: %v", *(v.Bar))
		i++
	}

	return fmt.Sprintf("ExWithoutAnnotation{%v}", strings.Join(fields[:i], ", "))
}

// ErrorName is the name of this type as defined in the Thrift
// file.
func (*ExWithoutAnnotation) ErrorName() string {
	return "ExWithoutAnnotation"
}

// Equals returns true if all the fields of this ExWithoutAnnotation match the
// provided ExWithoutAnnotation.
//
// This function performs a deep comparison.
func (v *ExWithoutAnnotation) Equals(rhs *ExWithoutAnnotation) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Bar, rhs.Bar) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ExWithoutAnnotation.
func (v *ExWithoutAnnotation) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Bar != nil {
		enc.AddString("bar", *v.Bar)
	}
	return err
}

// GetBar returns the value of Bar if it is set or its
// zero value if it is unset.
func (v *ExWithoutAnnotation) GetBar() (o string) {
	if v != nil && v.Bar != nil {
		return *v.Bar
	}

	return
}

// IsSetBar returns true if Bar is not nil.
func (v *ExWithoutAnnotation) IsSetBar() bool {
	return v != nil && v.Bar != nil
}

func (v *ExWithoutAnnotation) Error() string {
	return v.String()
}

type Struct struct {
	Baz *string `json:"baz,omitempty"`
}

// ToWire translates a Struct struct into a Thrift-level intermediate
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
func (v *Struct) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Baz != nil {
		w, err = wire.NewValueString(*(v.Baz)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Struct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Struct struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Struct
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Struct) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Baz = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a Struct struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Struct struct could not be encoded.
func (v *Struct) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Baz != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Baz)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Struct struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Struct struct could not be generated from the wire
// representation.
func (v *Struct) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Baz = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
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

// String returns a readable string representation of a Struct
// struct.
func (v *Struct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Baz != nil {
		fields[i] = fmt.Sprintf("Baz: %v", *(v.Baz))
		i++
	}

	return fmt.Sprintf("Struct{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Struct match the
// provided Struct.
//
// This function performs a deep comparison.
func (v *Struct) Equals(rhs *Struct) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Baz, rhs.Baz) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Struct.
func (v *Struct) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Baz != nil {
		enc.AddString("baz", *v.Baz)
	}
	return err
}

// GetBaz returns the value of Baz if it is set or its
// zero value if it is unset.
func (v *Struct) GetBaz() (o string) {
	if v != nil && v.Baz != nil {
		return *v.Baz
	}

	return
}

// IsSetBaz returns true if Baz is not nil.
func (v *Struct) IsSetBaz() bool {
	return v != nil && v.Baz != nil
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "NOSERVICES",
	Package:  "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/NOSERVICES",
	FilePath: "NOSERVICES.thrift",
	SHA1:     "bc28a223c8e87e320722bf77136e0eec1f1d18d6",
	Raw:      rawIDL,
}

const rawIDL = "// Thrift file with no service to ensure that types_yarpc.go is always\n// generated.\n\nexception ExWithAnnotation {\n    1: optional string foo\n} (\n    rpc.code = \"OUT_OF_RANGE\"\n)\n\nexception ExWithoutAnnotation {\n    1: optional string bar\n}\n\nstruct Struct {\n    1: optional string baz\n}\n"
