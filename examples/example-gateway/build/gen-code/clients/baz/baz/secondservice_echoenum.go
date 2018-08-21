// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// SecondService_EchoEnum_Args represents the arguments for the SecondService.echoEnum function.
//
// The arguments for echoEnum are sent and received over the wire as this struct.
type SecondService_EchoEnum_Args struct {
	Arg *Fruit `json:"arg,omitempty"`
}

func _Fruit_ptr(v Fruit) *Fruit {
	return &v
}

// ToWire translates a SecondService_EchoEnum_Args struct into a Thrift-level intermediate
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
func (v *SecondService_EchoEnum_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Arg == nil {
		v.Arg = _Fruit_ptr(FruitApple)
	}
	{
		w, err = v.Arg.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Fruit_Read(w wire.Value) (Fruit, error) {
	var v Fruit
	err := v.FromWire(w)
	return v, err
}

// FromWire deserializes a SecondService_EchoEnum_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SecondService_EchoEnum_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SecondService_EchoEnum_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SecondService_EchoEnum_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x Fruit
				x, err = _Fruit_Read(field.Value)
				v.Arg = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if v.Arg == nil {
		v.Arg = _Fruit_ptr(FruitApple)
	}

	return nil
}

// String returns a readable string representation of a SecondService_EchoEnum_Args
// struct.
func (v *SecondService_EchoEnum_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Arg != nil {
		fields[i] = fmt.Sprintf("Arg: %v", *(v.Arg))
		i++
	}

	return fmt.Sprintf("SecondService_EchoEnum_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Fruit_EqualsPtr(lhs, rhs *Fruit) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this SecondService_EchoEnum_Args match the
// provided SecondService_EchoEnum_Args.
//
// This function performs a deep comparison.
func (v *SecondService_EchoEnum_Args) Equals(rhs *SecondService_EchoEnum_Args) bool {
	if !_Fruit_EqualsPtr(v.Arg, rhs.Arg) {
		return false
	}

	return true
}

// GetArg returns the value of Arg if it is set or its
// default value if it is unset.
func (v *SecondService_EchoEnum_Args) GetArg() (o Fruit) {
	if v.Arg != nil {
		return *v.Arg
	}
	o = FruitApple
	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echoEnum" for this struct.
func (v *SecondService_EchoEnum_Args) MethodName() string {
	return "echoEnum"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SecondService_EchoEnum_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SecondService_EchoEnum_Helper provides functions that aid in handling the
// parameters and return values of the SecondService.echoEnum
// function.
var SecondService_EchoEnum_Helper = struct {
	// Args accepts the parameters of echoEnum in-order and returns
	// the arguments struct for the function.
	Args func(
		arg *Fruit,
	) *SecondService_EchoEnum_Args

	// IsException returns true if the given error can be thrown
	// by echoEnum.
	//
	// An error can be thrown by echoEnum only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echoEnum
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echoEnum into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echoEnum
	//
	//   value, err := echoEnum(args)
	//   result, err := SecondService_EchoEnum_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echoEnum: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(Fruit, error) (*SecondService_EchoEnum_Result, error)

	// UnwrapResponse takes the result struct for echoEnum
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echoEnum threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := SecondService_EchoEnum_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SecondService_EchoEnum_Result) (Fruit, error)
}{}

func init() {
	SecondService_EchoEnum_Helper.Args = func(
		arg *Fruit,
	) *SecondService_EchoEnum_Args {
		return &SecondService_EchoEnum_Args{
			Arg: arg,
		}
	}

	SecondService_EchoEnum_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	SecondService_EchoEnum_Helper.WrapResponse = func(success Fruit, err error) (*SecondService_EchoEnum_Result, error) {
		if err == nil {
			return &SecondService_EchoEnum_Result{Success: &success}, nil
		}

		return nil, err
	}
	SecondService_EchoEnum_Helper.UnwrapResponse = func(result *SecondService_EchoEnum_Result) (success Fruit, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// SecondService_EchoEnum_Result represents the result of a SecondService.echoEnum function call.
//
// The result of a echoEnum execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type SecondService_EchoEnum_Result struct {
	// Value returned by echoEnum after a successful execution.
	Success *Fruit `json:"success,omitempty"`
}

// ToWire translates a SecondService_EchoEnum_Result struct into a Thrift-level intermediate
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
func (v *SecondService_EchoEnum_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("SecondService_EchoEnum_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SecondService_EchoEnum_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SecondService_EchoEnum_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SecondService_EchoEnum_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SecondService_EchoEnum_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI32 {
				var x Fruit
				x, err = _Fruit_Read(field.Value)
				v.Success = &x
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
		return fmt.Errorf("SecondService_EchoEnum_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SecondService_EchoEnum_Result
// struct.
func (v *SecondService_EchoEnum_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("SecondService_EchoEnum_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SecondService_EchoEnum_Result match the
// provided SecondService_EchoEnum_Result.
//
// This function performs a deep comparison.
func (v *SecondService_EchoEnum_Result) Equals(rhs *SecondService_EchoEnum_Result) bool {
	if !_Fruit_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *SecondService_EchoEnum_Result) GetSuccess() (o Fruit) {
	if v.Success != nil {
		return *v.Success
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echoEnum" for this struct.
func (v *SecondService_EchoEnum_Result) MethodName() string {
	return "echoEnum"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SecondService_EchoEnum_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
