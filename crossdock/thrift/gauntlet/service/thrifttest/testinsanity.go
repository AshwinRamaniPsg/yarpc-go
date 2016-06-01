// Code generated by thriftrw

// Copyright (c) 2016 Uber Technologies, Inc.
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

package thrifttest

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"github.com/yarpc/yarpc-go/crossdock/thrift/gauntlet"
	"strings"
)

type TestInsanityArgs struct {
	Argument *gauntlet.Insanity `json:"argument,omitempty"`
}

func (v *TestInsanityArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Argument != nil {
		w, err = v.Argument.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Insanity_Read(w wire.Value) (*gauntlet.Insanity, error) {
	var v gauntlet.Insanity
	err := v.FromWire(w)
	return &v, err
}

func (v *TestInsanityArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Argument, err = _Insanity_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *TestInsanityArgs) String() string {
	var fields [1]string
	i := 0
	if v.Argument != nil {
		fields[i] = fmt.Sprintf("Argument: %v", v.Argument)
		i++
	}
	return fmt.Sprintf("TestInsanityArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestInsanityArgs) MethodName() string {
	return "testInsanity"
}

func (v *TestInsanityArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type TestInsanityResult struct {
	Success map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity `json:"success"`
}

type _Map_Numberz_Insanity_MapItemList map[gauntlet.Numberz]*gauntlet.Insanity

func (m _Map_Numberz_Insanity_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := k.ToWire()
		if err != nil {
			return err
		}
		vw, err := v.ToWire()
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_Numberz_Insanity_MapItemList) Size() int {
	return len(m)
}

func (_Map_Numberz_Insanity_MapItemList) KeyType() wire.Type {
	return wire.TI32
}

func (_Map_Numberz_Insanity_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_Numberz_Insanity_MapItemList) Close() {
}

type _Map_UserId_Map_Numberz_Insanity_MapItemList map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity

func (m _Map_UserId_Map_Numberz_Insanity_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := k.ToWire()
		if err != nil {
			return err
		}
		vw, err := wire.NewValueMap(_Map_Numberz_Insanity_MapItemList(v)), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_UserId_Map_Numberz_Insanity_MapItemList) Size() int {
	return len(m)
}

func (_Map_UserId_Map_Numberz_Insanity_MapItemList) KeyType() wire.Type {
	return wire.TI64
}

func (_Map_UserId_Map_Numberz_Insanity_MapItemList) ValueType() wire.Type {
	return wire.TMap
}

func (_Map_UserId_Map_Numberz_Insanity_MapItemList) Close() {
}

func (v *TestInsanityResult) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueMap(_Map_UserId_Map_Numberz_Insanity_MapItemList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("TestInsanityResult should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UserId_Read(w wire.Value) (gauntlet.UserId, error) {
	var x gauntlet.UserId
	err := x.FromWire(w)
	return x, err
}

func _Map_Numberz_Insanity_Read(m wire.MapItemList) (map[gauntlet.Numberz]*gauntlet.Insanity, error) {
	if m.KeyType() != wire.TI32 {
		return nil, nil
	}
	if m.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make(map[gauntlet.Numberz]*gauntlet.Insanity, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _Numberz_Read(x.Key)
		if err != nil {
			return err
		}
		v, err := _Insanity_Read(x.Value)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func _Map_UserId_Map_Numberz_Insanity_Read(m wire.MapItemList) (map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, error) {
	if m.KeyType() != wire.TI64 {
		return nil, nil
	}
	if m.ValueType() != wire.TMap {
		return nil, nil
	}
	o := make(map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _UserId_Read(x.Key)
		if err != nil {
			return err
		}
		v, err := _Map_Numberz_Insanity_Read(x.Value.GetMap())
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func (v *TestInsanityResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TMap {
				v.Success, err = _Map_UserId_Map_Numberz_Insanity_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("TestInsanityResult should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *TestInsanityResult) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("TestInsanityResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestInsanityResult) MethodName() string {
	return "testInsanity"
}

func (v *TestInsanityResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var TestInsanityHelper = struct {
	IsException    func(error) bool
	Args           func(argument *gauntlet.Insanity) *TestInsanityArgs
	WrapResponse   func(map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, error) (*TestInsanityResult, error)
	UnwrapResponse func(*TestInsanityResult) (map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, error)
}{}

func init() {
	TestInsanityHelper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	TestInsanityHelper.Args = func(argument *gauntlet.Insanity) *TestInsanityArgs {
		return &TestInsanityArgs{Argument: argument}
	}
	TestInsanityHelper.WrapResponse = func(success map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, err error) (*TestInsanityResult, error) {
		if err == nil {
			return &TestInsanityResult{Success: success}, nil
		}
		return nil, err
	}
	TestInsanityHelper.UnwrapResponse = func(result *TestInsanityResult) (success map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}