// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Echo_EchoI32Map_Args represents the arguments for the Echo.echoI32Map function.
//
// The arguments for echoI32Map are sent and received over the wire as this struct.
type Echo_EchoI32Map_Args struct {
	Arg map[int32]*BarResponse `json:"arg,required"`
}

type _Map_I32_BarResponse_MapItemList map[int32]*BarResponse

func (m _Map_I32_BarResponse_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := wire.NewValueI32(k), error(nil)
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

func (m _Map_I32_BarResponse_MapItemList) Size() int {
	return len(m)
}

func (_Map_I32_BarResponse_MapItemList) KeyType() wire.Type {
	return wire.TI32
}

func (_Map_I32_BarResponse_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_I32_BarResponse_MapItemList) Close() {}

// ToWire translates a Echo_EchoI32Map_Args struct into a Thrift-level intermediate
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
func (v *Echo_EchoI32Map_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Arg == nil {
		return w, errors.New("field Arg of Echo_EchoI32Map_Args is required")
	}
	w, err = wire.NewValueMap(_Map_I32_BarResponse_MapItemList(v.Arg)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_I32_BarResponse_Read(m wire.MapItemList) (map[int32]*BarResponse, error) {
	if m.KeyType() != wire.TI32 {
		return nil, nil
	}

	if m.ValueType() != wire.TStruct {
		return nil, nil
	}

	o := make(map[int32]*BarResponse, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetI32(), error(nil)
		if err != nil {
			return err
		}

		v, err := _BarResponse_Read(x.Value)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

// FromWire deserializes a Echo_EchoI32Map_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_EchoI32Map_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Echo_EchoI32Map_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Echo_EchoI32Map_Args) FromWire(w wire.Value) error {
	var err error

	argIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TMap {
				v.Arg, err = _Map_I32_BarResponse_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}

	if !argIsSet {
		return errors.New("field Arg of Echo_EchoI32Map_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Echo_EchoI32Map_Args
// struct.
func (v *Echo_EchoI32Map_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++

	return fmt.Sprintf("Echo_EchoI32Map_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Map_I32_BarResponse_Equals(lhs, rhs map[int32]*BarResponse) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !lv.Equals(rv) {
			return false
		}
	}
	return true
}

// Equals returns true if all the fields of this Echo_EchoI32Map_Args match the
// provided Echo_EchoI32Map_Args.
//
// This function performs a deep comparison.
func (v *Echo_EchoI32Map_Args) Equals(rhs *Echo_EchoI32Map_Args) bool {
	if !_Map_I32_BarResponse_Equals(v.Arg, rhs.Arg) {
		return false
	}

	return true
}

// GetArg returns the value of Arg if it is set or its
// zero value if it is unset.
func (v *Echo_EchoI32Map_Args) GetArg() (o map[int32]*BarResponse) { return v.Arg }

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echoI32Map" for this struct.
func (v *Echo_EchoI32Map_Args) MethodName() string {
	return "echoI32Map"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Echo_EchoI32Map_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Echo_EchoI32Map_Helper provides functions that aid in handling the
// parameters and return values of the Echo.echoI32Map
// function.
var Echo_EchoI32Map_Helper = struct {
	// Args accepts the parameters of echoI32Map in-order and returns
	// the arguments struct for the function.
	Args func(
		arg map[int32]*BarResponse,
	) *Echo_EchoI32Map_Args

	// IsException returns true if the given error can be thrown
	// by echoI32Map.
	//
	// An error can be thrown by echoI32Map only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echoI32Map
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echoI32Map into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echoI32Map
	//
	//   value, err := echoI32Map(args)
	//   result, err := Echo_EchoI32Map_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echoI32Map: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(map[int32]*BarResponse, error) (*Echo_EchoI32Map_Result, error)

	// UnwrapResponse takes the result struct for echoI32Map
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echoI32Map threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Echo_EchoI32Map_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Echo_EchoI32Map_Result) (map[int32]*BarResponse, error)
}{}

func init() {
	Echo_EchoI32Map_Helper.Args = func(
		arg map[int32]*BarResponse,
	) *Echo_EchoI32Map_Args {
		return &Echo_EchoI32Map_Args{
			Arg: arg,
		}
	}

	Echo_EchoI32Map_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Echo_EchoI32Map_Helper.WrapResponse = func(success map[int32]*BarResponse, err error) (*Echo_EchoI32Map_Result, error) {
		if err == nil {
			return &Echo_EchoI32Map_Result{Success: success}, nil
		}

		return nil, err
	}
	Echo_EchoI32Map_Helper.UnwrapResponse = func(result *Echo_EchoI32Map_Result) (success map[int32]*BarResponse, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Echo_EchoI32Map_Result represents the result of a Echo.echoI32Map function call.
//
// The result of a echoI32Map execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Echo_EchoI32Map_Result struct {
	// Value returned by echoI32Map after a successful execution.
	Success map[int32]*BarResponse `json:"success,omitempty"`
}

// ToWire translates a Echo_EchoI32Map_Result struct into a Thrift-level intermediate
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
func (v *Echo_EchoI32Map_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueMap(_Map_I32_BarResponse_MapItemList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_EchoI32Map_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Echo_EchoI32Map_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_EchoI32Map_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Echo_EchoI32Map_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Echo_EchoI32Map_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TMap {
				v.Success, err = _Map_I32_BarResponse_Read(field.Value.GetMap())
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
		return fmt.Errorf("Echo_EchoI32Map_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Echo_EchoI32Map_Result
// struct.
func (v *Echo_EchoI32Map_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Echo_EchoI32Map_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Echo_EchoI32Map_Result match the
// provided Echo_EchoI32Map_Result.
//
// This function performs a deep comparison.
func (v *Echo_EchoI32Map_Result) Equals(rhs *Echo_EchoI32Map_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _Map_I32_BarResponse_Equals(v.Success, rhs.Success))) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Echo_EchoI32Map_Result) GetSuccess() (o map[int32]*BarResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echoI32Map" for this struct.
func (v *Echo_EchoI32Map_Result) MethodName() string {
	return "echoI32Map"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Echo_EchoI32Map_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
